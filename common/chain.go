package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zeta-chain/zetacore/packages/zetaclient/pkg/model"
)

var (
	// mainnets
	EmptyChain   = Chain("")
	BSCChain     = Chain("BSC")
	ETHChain     = Chain("ETH")
	POLYGONChain = Chain("POLYGON")
	ZETAChain    = Chain("ZETA")
	BitcoinChain = Chain("Bitcoin")

	SigningAlgoSecp256k1 = SigninAlgo("secp256k1")
	SigningAlgoEd25519   = SigninAlgo("ed25519")

	// testnets
	BSCTestnetChain     = Chain("BSCTestnet")
	GoerliChain         = Chain("Goerli")
	RopstenChain        = Chain("ROPSTEN")
	MumbaiChain         = Chain("Mumbai")
	BaobabChain         = Chain("Baobab")
	BitcoinTestnetChain = Chain("BitcoinTestnet")
)

type SigninAlgo string

// Chain is an alias of string , represent a block chain
type Chain string

// Chains represent a slice of Chain
type Chains []Chain

// Validate validates chain format, should consist only of uppercase letters
func (chain Chain) Validate() error {
	if len(chain) < 3 {
		return errors.New("chain id len is less than 3")
	}
	if len(chain) > 10 {
		return errors.New("chain id len is more than 10")
	}
	for _, ch := range string(chain) {
		if ch < 'A' || ch > 'Z' {
			return errors.New("chain id can consist only of uppercase letters")
		}
	}
	return nil
}

// NewChain create a new Chain and default the siging_algo to Secp256k1
func NewChain(chainID string) (Chain, error) {
	chain := Chain(strings.ToUpper(chainID))
	if err := chain.Validate(); err != nil {
		return chain, err
	}
	return chain, nil
}

func ParseChain(chainName string) (Chain, error) {
	switch chainName {
	case "ETH":
		return ETHChain, nil
	case "BSC":
		return BSCChain, nil
	case "POLYGON":
		return POLYGONChain, nil
	case "ROPSTEN":
		return RopstenChain, nil
	case "MUMBAI":
		return MumbaiChain, nil
	case "BSCTESTNET":
		return BSCTestnetChain, nil
	case "GOERLI":
		return GoerliChain, nil
	case "BAOBAB":
		return BaobabChain, nil
	case "BITCOIN":
		return BitcoinChain, nil
	default:
		return EmptyChain, fmt.Errorf("unsupported chain %s", chainName)
	}
}

func (chain Chain) GetNativeTokenSymbol() string {
	switch chain {
	case ETHChain:
		return "ETH"
	case BSCChain:
		return "BNB"
	case POLYGONChain:
		return "MATIC"
	case RopstenChain:
		return "rETH"
	case GoerliChain:
		return "gETH"
	case MumbaiChain:
		return "tMATIC"
	case BSCTestnetChain:
		return "tBNB"
	default:
		return "" // should not happen
	}
}

// Equals compare two chain to see whether they represent the same chain
func (chain Chain) Equals(c2 Chain) bool {
	return strings.EqualFold(chain.String(), c2.String())
}

func (chain Chain) IsZetaChain() bool {
	return chain.Equals(ZETAChain)
}

// IsEmpty is to determinate whether the chain is empty
func (chain Chain) IsEmpty() bool {
	return strings.TrimSpace(chain.String()) == ""
}

// String implement fmt.Stringer
func (chain Chain) String() string {
	// convert it to upper case again just in case someone created a ticker via Chain("rune")
	return strings.ToUpper(string(chain))
}

// GetSigningAlgo get the signing algorithm for the given chain
func (chain Chain) GetSigningAlgo() SigninAlgo {
	switch chain {
	case ETHChain, POLYGONChain, BSCChain, RopstenChain:
		return SigningAlgoSecp256k1
	default:
		return SigningAlgoSecp256k1
	}
}

func (chain Chain) Type() string {
	switch chain {
	case BitcoinChain, BitcoinTestnetChain:
		return model.ChainTypeBTC
	default:
		return model.ChainTypeETH
	}
}

// Has check whether chain c is in the list
func (chains Chains) Has(c Chain) bool {
	for _, ch := range chains {
		if ch.Equals(c) {
			return true
		}
	}
	return false
}

// Distinct return a distinct set of chains, no duplicates
func (chains Chains) Distinct() Chains {
	var newChains Chains
	for _, chain := range chains {
		if !newChains.Has(chain) {
			newChains = append(newChains, chain)
		}
	}
	return newChains
}

func (chains Chains) Strings() []string {
	strings := make([]string, len(chains))
	for i, c := range chains {
		strings[i] = c.String()
	}
	return strings
}
