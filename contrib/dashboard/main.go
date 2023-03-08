package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeta-chain/zetacore/contracts/zevm"
	"math/big"
	"net/http"
	"time"
)

var (
	gETHPair   = ethcommon.HexToAddress("0x719E237DaE24fe0b7A81DDcF4d54c69F48232406")
	tBNBPair   = ethcommon.HexToAddress("0x0C0B35C5eF00d9caD8D2ce65b147ea2A27d526Bc")
	tMATICPair = ethcommon.HexToAddress("0x16Ef1b018026E389FDA93c1e993E987CF6E852E7")

	gETHZRC20   = ethcommon.HexToAddress("0x91d18e54DAf4F677cB28167158d6dd21F6aB3921")
	tBNBZRC20   = ethcommon.HexToAddress("0x13A0c5930C028511Dc02665E7285134B6d11A5f4")
	tMATICZRC20 = ethcommon.HexToAddress("0xd97B1de3619ed2c6BEb3860147E30cA8A7dC9891")

	WZETAAddress = ethcommon.HexToAddress("0x5F0b1a82749cb4E2278EC87F8BF6B618dC71a8bf")

	zevmRPCEndpoint = "https://api.athens2.zetachain.com/evm"
)

type Server struct {
	zevmClient *ethclient.Client
	ethPair    *zevm.UniswapV2Pair
	bnbPair    *zevm.UniswapV2Pair
	maticPair  *zevm.UniswapV2Pair
}

func main() {
	server := &Server{}
	zevmClient, err := ethclient.Dial(zevmRPCEndpoint)
	if err != nil {
		panic(err)
	}
	server.zevmClient = zevmClient

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
	server.ethPair = ethPair
	server.bnbPair = bnbPair
	server.maticPair = maticPair

	http.HandleFunc("/", server.reserveHandler)
	http.HandleFunc("/ethpool", server.ethEventsHandler)
	http.HandleFunc("/bnbpool", server.bnbEventsHandler)
	http.HandleFunc("/maticpool", server.maticEventsHandler)

	http.ListenAndServe(":8088", nil)

	bnbToken0, err := bnbPair.Token0(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	bnbToken1, err := bnbPair.Token1(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	ethToken0, err := ethPair.Token0(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	ethToken1, err := ethPair.Token1(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	maticToken0, err := maticPair.Token0(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	maticToken1, err := maticPair.Token1(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("bnb token0: %v\n", bnbToken0.Hex())
	fmt.Printf("bnb token1: %v\n", bnbToken1.Hex())
	fmt.Printf("eth token0: %v\n", ethToken0.Hex())
	fmt.Printf("eth token1: %v\n", ethToken1.Hex())
	fmt.Printf("matic token0: %v\n", maticToken0.Hex())
	fmt.Printf("matic token1: %v\n", maticToken1.Hex())

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

func (s *Server) reserveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<!doctype html><html><body> ")
	fmt.Fprintf(w, "<div><a href=\"/ethpool\">ETH swap events</a> | <a href=\"/bnbpool\">BNB swap events</a> | <a href=\"/maticpool\">MATIC swap events</a></div>")

	fmt.Fprintf(w, "<h2>zEVM system pool reserves</h2>")

	fmt.Fprintf(w, "<pre>")
	fmt.Fprintf(w, "gETH ZRC20: %v\n", gETHZRC20.Hex())
	fmt.Fprintf(w, "tBNB ZRC20: %v\n", tBNBZRC20.Hex())
	fmt.Fprintf(w, "tMATIC ZRC20: %v\n", tMATICZRC20.Hex())
	fmt.Fprintf(w, "WZETA: %v\n", WZETAAddress.Hex())
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "gETH/ZETA pair: %v\n", gETHPair.Hex())
	fmt.Fprintf(w, "tBNB/ZETA pair: %v\n", tBNBPair.Hex())
	fmt.Fprintf(w, "tMATIC/ZETA pair: %v\n", tMATICPair.Hex())
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "(*) marks gETH or tBNB or tMATIC, as opposed to WZETA\n")

	res, err := s.bnbPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "tBNB/ZETA pool reserves:\n")
	fmt.Fprintf(w, "  Reserve0(*): %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  Reserve1: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))

	res, err = s.maticPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "tMATIC/ZETA pool reserves:\n")
	fmt.Fprintf(w, "  Reserve0: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  Reserve1(*): %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))

	res, err = s.ethPair.GetReserves(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "gETH/ZETA pool reserves:\n")
	fmt.Fprintf(w, "  Reserve0: %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve0), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  Reserve1(*): %v\n", big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.Reserve1), big.NewFloat(1e18)))
	fmt.Fprintf(w, "  BlockTimestampLast: %v\n", time.Unix(int64(res.BlockTimestampLast), 0))

	fmt.Fprintf(w, "</pre>")
}

func (s *Server) bnbEventsHandler(w http.ResponseWriter, r *http.Request) {
	s.getEventsByPair(w, s.bnbPair)
}

func (s *Server) ethEventsHandler(w http.ResponseWriter, r *http.Request) {
	s.getEventsByPair(w, s.ethPair)
}

func (s *Server) maticEventsHandler(w http.ResponseWriter, r *http.Request) {
	s.getEventsByPair(w, s.maticPair)
}

func (s *Server) getEventsByPair(w http.ResponseWriter, pair *zevm.UniswapV2Pair) {
	bn, err := s.zevmClient.BlockNumber(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	iter, err := pair.FilterSwap(&bind.FilterOpts{
		Start:   bn - 300,
		End:     nil,
		Context: context.Background(),
	}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	token0, err := pair.Token0(&bind.CallOpts{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	token1, err := pair.Token1(&bind.CallOpts{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	var logs []*zevm.UniswapV2PairSwap
	for iter.Next() {
		logs = append(logs, iter.Event)
	}

	fmt.Fprintf(w, "gETH/ZETA swap event (latest 100 events of last 1000 blocks)\n")
	cnt := 0
	for i := len(logs) - 1; i >= 0 && cnt <= 100; i-- {
		fmt.Fprintf(w, "%3d ", cnt)
		log := logs[i]
		if log.Amount0In.Cmp(big.NewInt(0)) > 0 {
			fmt.Fprintf(w, " %5s: %9f => ", addrToTokenName(token0), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(log.Amount0In), big.NewFloat(1e18)))
		}
		if log.Amount1In.Cmp(big.NewInt(0)) > 0 {
			fmt.Fprintf(w, " %5s: %9f => ", addrToTokenName(token1), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(log.Amount1In), big.NewFloat(1e18)))
		}
		if log.Amount0Out.Cmp(big.NewInt(0)) > 0 {
			fmt.Fprintf(w, " %5s: %9f", addrToTokenName(token0), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(log.Amount0Out), big.NewFloat(1e18)))
		}
		if log.Amount1Out.Cmp(big.NewInt(0)) > 0 {
			fmt.Fprintf(w, " %5s: %9f", addrToTokenName(token1), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(log.Amount1Out), big.NewFloat(1e18)))
		}
		fmt.Fprintf(w, "  \tto: %v\t", log.To)
		fmt.Fprintf(w, "  block: %v\n", log.Raw.BlockNumber)
		//fmt.Fprintf(w, "  tx: %v\n", log.Raw.TxHash)
		//fmt.Fprintf(w, "\n")
		cnt++
	}
}

func addrToTokenName(addr ethcommon.Address) string {
	switch addr {
	case tBNBZRC20:
		return "tBNB"
	case tMATICZRC20:
		return "tMATIC"
	case gETHZRC20:
		return "gETH"
	case WZETAAddress:
		return "WZETA"
	}
	return "unknown"
}
