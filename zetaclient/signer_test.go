package zetaclient

import (
	"bytes"
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/contracts/evm"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	. "gopkg.in/check.v1"
	"math/big"
	"os"
)

type SignerSuite struct {
	signer *Signer
}

var _ = Suite(&SignerSuite{})

func (s *SignerSuite) SetUpTest(c *C) {
	privkey := os.Getenv(("PRIVKEY"))
	c.Assert(privkey, Not(Equals), "")
	tss, err := NewTestSigner(privkey)
	c.Assert(err, IsNil)
	metaContractAddress := ethcommon.HexToAddress(config.Chains[common.BSCTestnetChain.String()].ConnectorContractAddress)
	nodeEnv := os.Getenv("GOERLI_ENDPOINT")
	c.Logf("nodeEnv: %s", nodeEnv)
	c.Assert(nodeEnv, Not(Equals), "")
	signer, err := NewSigner(common.GoerliChain, nodeEnv, tss, config.ConnectorAbiString, metaContractAddress)
	c.Assert(err, IsNil)
	c.Logf("TSS Address %s", tss.Address().Hex())
	c.Logf("Contract on chain %s: %s", signer.chain, metaContractAddress.Hex())
	s.signer = signer
}

func (s *SignerSuite) TestSign(c *C) {
	data := []byte("1234")
	tx, sig, hash, err := s.signer.Sign(data, s.signer.TssSigner.Address(), 109, big.NewInt(2), 23)
	_ = tx
	c.Assert(err, IsNil)
	pubkey, err := crypto.Ecrecover(hash, sig)
	c.Assert(err, IsNil)
	c.Assert(bytes.Equal(pubkey, s.signer.TssSigner.Pubkey()), Equals, true)
}

func (s *SignerSuite) TestDeployZetaEth(c *C) {
	nonce, err := s.signer.Client.NonceAt(context.Background(), s.signer.TssSigner.Address(), nil)
	c.Assert(err, IsNil)
	gasprice, err := s.signer.Client.SuggestGasPrice(context.Background())
	c.Assert(err, IsNil)
	cABI, err := evm.ZetaEthMetaData.GetAbi()
	c.Assert(err, IsNil)

	tx, err := s.signer.SignContractDeployTx(nonce, gasprice, 1_000_000, cABI, evm.ZetaEthContract.Bin, big.NewInt(1e18))
	c.Assert(err, IsNil)
	err = s.signer.Broadcast(tx)
	c.Assert(err, IsNil)

	contractAddr := crypto.CreateAddress(s.signer.TssSigner.Address(), nonce)

	c.Logf("Deploy ZetaEth Tx: %s; contract address: %s", tx.Hash().Hex(), contractAddr.Hex())
}

func (s *SignerSuite) TestDeployZetaAndConnector(c *C) {
	nonce, err := s.signer.Client.PendingNonceAt(context.Background(), s.signer.TssSigner.Address())
	c.Assert(err, IsNil)

	tx, zeta, err := s.signer.DeployZetaEth(big.NewInt(1000), nonce)
	c.Assert(err, IsNil)
	c.Logf("Deploy ZetaEth Tx: %s; contract address: %s", tx.Hex(), zeta.Hex())
	nonce++
	tx, conn, err := s.signer.DeployConnectorEth(zeta, nonce)
	c.Assert(err, IsNil)
	c.Logf("Deploy ConnectorEth Tx: %s; contract address: %s", tx.Hex(), conn.Hex())
	nonce++

	tx, zeta, err = s.signer.DeployZetaNonEth(nonce)
	c.Assert(err, IsNil)
	c.Logf("Deploy ZetaNonEth Tx: %s; contract address: %s", tx.Hex(), zeta.Hex())
	nonce++

	tx, conn, err = s.signer.DeployConnectorNonEth(zeta, nonce)
	c.Assert(err, IsNil)
	c.Logf("Deploy ConnectorNonEth Tx: %s; contract address: %s", tx.Hex(), conn.Hex())
	nonce++

}
