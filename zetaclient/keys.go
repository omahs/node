package zetaclient

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	ckeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/common/cosmos"
	"io"
	"os"
)

// Keys manages all the keys used by zeta client
type Keys struct {
	signerName      string
	password        string // TODO this is a bad way , need to fix it
	kb              ckeys.Keyring
	operatorAddress sdk.AccAddress
}

// NewKeysWithKeybase create a new instance of Keys
func NewKeysWithKeybase(kb ckeys.Keyring, granterAddress sdk.AccAddress, name, password string) *Keys {
	return &Keys{
		signerName:      name,
		password:        password,
		kb:              kb,
		operatorAddress: granterAddress,
	}
}

func GetGranteeKeyName(_ common.KeyType, signerName string) string {
	return fmt.Sprintf("%s", signerName)
}

// GetKeyringKeybase return keyring and key info
func GetKeyringKeybase(hotkeyName, chainHomeFolder, password string) (ckeys.Keyring, error) {
	logger := log.Logger.With().Str("module", "GetKeyringKeybase").Logger()
	if len(hotkeyName) == 0 {
		return nil, fmt.Errorf("signer name is empty")
	}
	if len(password) == 0 {
		return nil, fmt.Errorf("password is empty")
	}

	buf := bytes.NewBufferString(password)
	// the library used by keyring is using ReadLine , which expect a new line
	buf.WriteByte('\n')
	buf.WriteString(password)
	buf.WriteByte('\n')
	kb, err := getKeybase(chainHomeFolder, buf)
	if err != nil {
		return nil, fmt.Errorf("fail to get keybase,err:%w", err)
	}
	oldStdIn := os.Stdin
	defer func() {
		os.Stdin = oldStdIn
	}()
	os.Stdin = nil

	logger.Debug().Msgf("Checking for Key: %s  \n , Folder %s \n,Backend %s \n", hotkeyName, chainHomeFolder, kb.Backend())
	_, err = kb.Key(hotkeyName)
	if err != nil {
		return nil, fmt.Errorf("key not presnt with name (%s): %w", hotkeyName, err)
	}

	return kb, nil
}

// getKeybase will create an instance of Keybase
func getKeybase(zetaCoreHome string, reader io.Reader) (ckeys.Keyring, error) {
	cliDir := zetaCoreHome
	if len(zetaCoreHome) == 0 {
		return nil, fmt.Errorf("zetaCoreHome is empty")
	}
	//FIXME: BackendTest is used for convenient testing with Starport generated accouts.
	// Change to BackendFile with password!
	//op := func(options *ckeys.Options) {
	//	options.SupportedAlgos = append(options.SupportedAlgos, hd.EthSecp256k1)
	//	options.SupportedAlgosLedger = append(options.SupportedAlgosLedger, hd.EthSecp256k1)
	//}
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
	return ckeys.New(sdk.KeyringServiceName(), ckeys.BackendTest, cliDir, reader, cdc)
}

// GetSignerInfo return signer info
func (k *Keys) GetSignerInfo(keyType common.KeyType) *ckeys.Record {
	signer := GetGranteeKeyName(keyType, k.signerName)
	info, err := k.kb.Key(signer)
	if err != nil {
		panic(err)
	}
	return info
}

func (k *Keys) GetOperatorAddress() sdk.AccAddress {
	return k.operatorAddress
}

func (k *Keys) GetAddress(keyType common.KeyType) sdk.AccAddress {
	signer := GetGranteeKeyName(keyType, k.signerName)
	info, err := k.kb.Key(signer)
	if err != nil {
		panic(err)
	}
	addr, _ := info.GetAddress()
	return addr
}

// GetPrivateKey return the private key
func (k *Keys) GetPrivateKey(keyType common.KeyType) (cryptotypes.PrivKey, error) {
	// return k.kb.ExportPrivateKeyObject(k.signerName)
	signer := GetGranteeKeyName(keyType, k.signerName)
	privKeyArmor, err := k.kb.ExportPrivKeyArmor(signer, k.password)
	if err != nil {
		return nil, err
	}
	priKey, _, err := crypto.UnarmorDecryptPrivKey(privKeyArmor, k.password)
	if err != nil {
		return nil, fmt.Errorf("fail to unarmor private key: %w", err)
	}
	return priKey, nil
}

// GetKeybase return the keybase
func (k *Keys) GetKeybase() ckeys.Keyring {
	return k.kb
}

func (k *Keys) GetPubKeySet(keyType common.KeyType) (common.PubKeySet, error) {
	pubkeySet := common.PubKeySet{
		Secp256k1: "",
		Ed25519:   "",
	}

	pK, err := k.GetPrivateKey(keyType)
	if err != nil {
		return pubkeySet, err
	}

	s, err := cosmos.Bech32ifyPubKey(cosmos.Bech32PubKeyTypeAccPub, pK.PubKey())
	if err != nil {
		return pubkeySet, ErrBech32ifyPubKey
	}
	pubkey, err := common.NewPubKey(s)
	if err != nil {
		return pubkeySet, ErrNewPubKey
	}
	pubkeySet.Secp256k1 = pubkey
	return pubkeySet, nil
}
