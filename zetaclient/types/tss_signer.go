package types

import ethcommon "github.com/ethereum/go-ethereum/common"

type TSSSignerI interface {
	Pubkey() []byte
	Sign(data []byte) ([65]byte, error)
	Address() ethcommon.Address
	PubkeyString() string
	Keygen(pubkeys []string) error
	TestKeysign() bool
}
