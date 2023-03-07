package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeta-chain/zetacore/contracts/zevm"
	"math/big"
	"time"
)

var (
	gETHPair   = ethcommon.HexToAddress("0x719E237DaE24fe0b7A81DDcF4d54c69F48232406")
	tBNBPair   = ethcommon.HexToAddress("0x0C0B35C5eF00d9caD8D2ce65b147ea2A27d526Bc")
	tMATICPair = ethcommon.HexToAddress("0x16Ef1b018026E389FDA93c1e993E987CF6E852E7")

	gETHZRC20   = ethcommon.HexToAddress("0x91d18e54DAf4F677cB28167158d6dd21F6aB3921")
	tBNBZRC20   = ethcommon.HexToAddress("0x13A0c5930C028511Dc02665E7285134B6d11A5f4")
	tMATICZRC20 = ethcommon.HexToAddress("0xd97B1de3619ed2c6BEb3860147E30cA8A7dC9891")

	zevmRPCEndpoint = "https://api.athens2.zetachain.com/evm"
)

func main() {
	zevmClient, err := ethclient.Dial(zevmRPCEndpoint)
	if err != nil {
		panic(err)
	}
	ethPair, err := zevm.NewUniswapV2Pair(gETHPair, zevmClient)
	if err != nil {
		panic(err)
	}
	bnbPair, err := zevm.NewUniswapV2Pair(tBNBPair, zevmClient)
	if err != nil {
		panic(err)
	}
	maticPair, err := zevm.NewUniswapV2Pair(tMATICPair, zevmClient)
	if err != nil {
		panic(err)
	}

	res, err := bnbPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("BN/ZETA pool reserves:\n")
	fmt.Printf("  Reserve0: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Printf("  Reserve1: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Printf("  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))

	res, err = maticPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("MATIC/ZETA pool reserves:\n")
	fmt.Printf("  Reserve0: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Printf("  Reserve1: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Printf("  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))

	res, err = ethPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ETH/ZETA pool reserves:\n")
	fmt.Printf("  Reserve0: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Printf("  Reserve1: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Printf("  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))
}
