package zetaclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/zeta-chain/zetacore/contracts/evm"
	mcconfig "github.com/zeta-chain/zetacore/zetaclient/config"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/zetaclient/types"
	"math/big"
	"strings"
	"time"
)

var (
	nullHash     = ethcommon.Hash{}
	nullAddress  = ethcommon.Address{}
	deployerAddr = ethcommon.HexToAddress(mcconfig.DeployerAddress)
)

type Signer struct {
	Client              *ethclient.Client
	chain               common.Chain
	chainID             *big.Int
	TssSigner           types.TSSSignerI
	ethSigner           ethtypes.Signer
	abi                 abi.ABI
	metaContractAddress ethcommon.Address
	logger              zerolog.Logger
}

func NewSigner(chain common.Chain, endpoint string, tssSigner types.TSSSignerI, abiString string, metaContract ethcommon.Address) (*Signer, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}
	ethSigner := ethtypes.LatestSignerForChainID(chainID)
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}

	return &Signer{
		Client:              client,
		chain:               chain,
		TssSigner:           tssSigner,
		chainID:             chainID,
		ethSigner:           ethSigner,
		abi:                 abi,
		metaContractAddress: metaContract,
		logger:              log.With().Str("module", "Signer").Logger(),
	}, nil
}

// given data, and metadata (gas, nonce, etc)
// returns a signed transaction, sig bytes, hash bytes, and error
func (signer *Signer) Sign(data []byte, to ethcommon.Address, gasLimit uint64, gasPrice *big.Int, nonce uint64) (*ethtypes.Transaction, []byte, []byte, error) {
	tx := ethtypes.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, data)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.TssSigner.Sign(hashBytes)
	if err != nil {
		return nil, nil, nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, nil, nil, err
	}
	return signedTX, sig[:], hashBytes[:], nil
}

// takes in signed tx, broadcast to external chain node
func (signer *Signer) Broadcast(tx *ethtypes.Transaction) error {
	ctxt, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return signer.Client.SendTransaction(ctxt, tx)
}

//    function onReceive(
//        bytes calldata originSenderAddress,
//        uint256 originChainId,
//        address destinationAddress,
//        uint zetaAmount,
//        bytes calldata message,
//        bytes32 internalSendHash
//    ) external virtual {}
func (signer *Signer) SignOutboundTx(sender ethcommon.Address, srcChainID *big.Int, to ethcommon.Address, amount *big.Int, gasLimit uint64, message []byte, sendHash [32]byte, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	if len(sendHash) < 32 {
		return nil, fmt.Errorf("sendHash len %d must be 32", len(sendHash))
	}
	var data []byte
	var err error

	data, err = signer.abi.Pack("onReceive", sender.Bytes(), srcChainID, to, amount, message, sendHash)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.metaContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}

//function onRevert(
//address originSenderAddress,
//uint256 originChainId,
//bytes calldata destinationAddress,
//uint256 destinationChainId,
//uint256 zetaAmount,
//bytes calldata message,
//bytes32 internalSendHash
//) external override whenNotPaused onlyTssAddress
func (signer *Signer) SignRevertTx(sender ethcommon.Address, srcChainID *big.Int, to []byte, toChainID *big.Int, amount *big.Int, gasLimit uint64, message []byte, sendHash [32]byte, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	var data []byte
	var err error

	data, err = signer.abi.Pack("onRevert", sender, srcChainID, to, toChainID, amount, message, sendHash)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.metaContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}

func (signer *Signer) SignCancelTx(nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	tx := ethtypes.NewTransaction(nonce, signer.TssSigner.Address(), big.NewInt(0), 21000, gasPrice, nil)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.TssSigner.Sign(hashBytes)
	if err != nil {
		return nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, err
	}

	return signedTX, nil
}

// bytecode is hex string, withouth 0x prefix
func (signer *Signer) SignContractDeployTx(nonce uint64, gasPrice *big.Int, gasLimit uint64, cABI *abi.ABI, bytecode evmtypes.HexString, args ...interface{}) (*ethtypes.Transaction, error) {
	ctorArgs, err := cABI.Pack(
		"", // function--empty string for constructor
		args...,
	)
	if err != nil {
		return nil, err
	}
	data := make([]byte, len(bytecode)+len(ctorArgs))
	copy(data[:len(bytecode)], bytecode)
	copy(data[len(bytecode):], ctorArgs)
	tx := ethtypes.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, data)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.TssSigner.Sign(hashBytes)
	if err != nil {
		return nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, err
	}

	return signedTX, nil
}

//TODO: move contract creation directive to zetacore; and register contract address with zetacore
func (signer *Signer) DeployZetaEth(initialSupply *big.Int, nonce uint64) (ethcommon.Hash, ethcommon.Address, error) {

	gasPrice, err := signer.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nullHash, nullAddress, err
	}
	cABI, err := evm.ZetaEthMetaData.GetAbi()
	if err != nil {
		return nullHash, nullAddress, err
	}

	tx, err := signer.SignContractDeployTx(nonce, gasPrice, 1_000_000, cABI, evm.ZetaEthContract.Bin, initialSupply)
	if err != nil {
		return nullHash, nullAddress, err
	}

	err = signer.Broadcast(tx)
	if err != nil {
		return nullHash, nullAddress, err
	}

	contractAddr := crypto.CreateAddress(signer.TssSigner.Address(), nonce)

	return tx.Hash(), contractAddr, nil
}

func (signer *Signer) DeployZetaNonEth(nonce uint64) (ethcommon.Hash, ethcommon.Address, error) {
	gasPrice, err := signer.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nullHash, nullAddress, err
	}
	cABI, err := evm.ZetaNonEthMetaData.GetAbi()
	if err != nil {
		return nullHash, nullAddress, err
	}

	tx, err := signer.SignContractDeployTx(nonce, gasPrice, 1_000_000, cABI, evm.ZetaNonEthContract.Bin,
		signer.TssSigner.Address(), deployerAddr)
	if err != nil {
		return nullHash, nullAddress, err
	}

	err = signer.Broadcast(tx)
	if err != nil {
		return nullHash, nullAddress, err
	}

	contractAddr := crypto.CreateAddress(signer.TssSigner.Address(), nonce)

	return tx.Hash(), contractAddr, nil
}

func (signer *Signer) DeployConnectorEth(zetaTokenAddress ethcommon.Address, nonce uint64) (ethcommon.Hash, ethcommon.Address, error) {
	gasPrice, err := signer.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nullHash, nullAddress, err
	}
	cABI, err := evm.ConnectorMetaData.GetAbi()
	if err != nil {
		return nullHash, nullAddress, err
	}

	tx, err := signer.SignContractDeployTx(nonce, gasPrice, 2_000_000, cABI, evm.ConnectorEthContract.Bin,
		// zetaTokenAddress, tss, tssUpdater, pauser
		zetaTokenAddress, signer.TssSigner.Address(), deployerAddr, deployerAddr,
	)
	if err != nil {
		return nullHash, nullAddress, err
	}

	err = signer.Broadcast(tx)
	if err != nil {
		return nullHash, nullAddress, err
	}

	contractAddr := crypto.CreateAddress(signer.TssSigner.Address(), nonce)
	return tx.Hash(), contractAddr, nil
}

func (signer *Signer) DeployConnectorNonEth(zetaTokenAddress ethcommon.Address, nonce uint64) (ethcommon.Hash, ethcommon.Address, error) {
	gasPrice, err := signer.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nullHash, nullAddress, err
	}
	cABI, err := evm.ConnectorMetaData.GetAbi()
	if err != nil {
		return nullHash, nullAddress, err
	}

	tx, err := signer.SignContractDeployTx(nonce, gasPrice, 2_000_000, cABI, evm.ConnectorNonEthContract.Bin,
		// zetaTokenAddress, tss, tssUpdater, pauser
		zetaTokenAddress, signer.TssSigner.Address(), deployerAddr, deployerAddr,
	)
	if err != nil {
		return nullHash, nullAddress, err
	}

	err = signer.Broadcast(tx)
	if err != nil {
		return nullHash, nullAddress, err
	}

	contractAddr := crypto.CreateAddress(signer.TssSigner.Address(), nonce)
	return tx.Hash(), contractAddr, nil
}
