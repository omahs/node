package zetaclient

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/zetaclient/types"
)

type TestSigner struct {
	PrivKey *ecdsa.PrivateKey
}

var _ types.TSSSignerI = TestSigner{} // type check

func NewTestSigner(privkeyHex string) (*TestSigner, error) {
	privKey, err := crypto.HexToECDSA(privkeyHex)
	if err != nil {
		return nil, err
	}
	return &TestSigner{
		PrivKey: privKey,
	}, nil
}

func (s TestSigner) Keygen(_ []string) error {
	return nil
}

func (s TestSigner) Sign(digest []byte) ([65]byte, error) {
	sig, err := crypto.Sign(digest, s.PrivKey)
	if err != nil {
		return [65]byte{}, err
	}
	var sigbyte [65]byte
	copy(sigbyte[:], sig[:65])
	return sigbyte, nil
}

func (s TestSigner) Pubkey() []byte {
	publicKeyBytes := crypto.FromECDSAPub(&s.PrivKey.PublicKey)
	return publicKeyBytes
}

func (s TestSigner) PubkeyString() string {
	return sdk.AccAddress(s.Pubkey()).String()
}

func (s TestSigner) Address() ethcommon.Address {
	return crypto.PubkeyToAddress(s.PrivKey.PublicKey)
}

func (s TestSigner) TestKeysign() bool {
	log.Info().Msg("trying keysign...")
	data := []byte("hello meta")
	H := crypto.Keccak256Hash(data)
	log.Info().Msgf("hash of data (hello meta) is %s", H)

	sig, err := s.Sign(H.Bytes())
	if err != nil {
		log.Error().Err(err).Msg("signing failed")
		return false
	}
	sigPublicKey, err := crypto.SigToPub(H.Bytes(), sig[:])
	if err != nil {
		log.Error().Err(err).Msg("SigToPub failed")
		return false
	}
	//compressedPubkey := crypto.CompressPubkey(sigPublicKey)
	uncompressedPubkey := crypto.FromECDSAPub(sigPublicKey)
	log.Info().Msgf("compressed pubkey %s", hex.EncodeToString(uncompressedPubkey))
	log.Info().Msgf("pubkey            %s", hex.EncodeToString(s.Pubkey()))

	return bytes.Compare(s.Pubkey(), uncompressedPubkey) == 0

}
