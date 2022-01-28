package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog"
	"github.com/zeta-chain/zetacore/cmd"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/common/cosmos"
	mc "github.com/zeta-chain/zetacore/zetaclient"

	"github.com/cosmos/cosmos-sdk/types"
	mcconfig "github.com/zeta-chain/zetacore/zetaclient/config"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-peerstore/addr"
	maddr "github.com/multiformats/go-multiaddr"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const ()

func main() {
	var mockFlag = flag.String("mock", "", "[mocknet|tsstest]")
	var validatorName = flag.String("val", "alice", "validator name")
	var tssTestFlag = flag.Bool("tss", false, "2 node TSS test mode")
	var peer = flag.String("peer", "", "peer address, e.g. /dns/tss1/tcp/6668/ipfs/16Uiu2HAmACG5DtqmQsHtXg4G2sLS65ttv84e7MrL4kapkjfmhxAp")
	var chainid = flag.String("chain-id", "testing", "chain-id; default testing")
	flag.Parse()
	//BOOTSTRAP_PEER := "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"

	var peers addr.AddrList
	fmt.Println("peer", *peer)
	if *peer != "" {
		address, err := maddr.NewMultiaddr(*peer)
		if err != nil {
			log.Error().Err(err).Msg("NewMultiaddr error")
			return
		}
		peers = append(peers, address)
	}

	eth_end := os.Getenv("ETH_ENDPOINT")
	if eth_end != "" {
		fmt.Println("ETH_ENDPOINT: ", eth_end)
		mcconfig.ETH_ENDPOINT = eth_end
	}
	poly_end := os.Getenv("POLY_ENDPOINT")
	if poly_end != "" {
		fmt.Println("POLY_ENDPOINT: ", poly_end)
		mcconfig.POLY_ENDPOINT = poly_end
	}
	bsc_end := os.Getenv("BSC_ENDPOINT")
	if bsc_end != "" {
		fmt.Println("BSC_ENDPOINT: ", bsc_end)
		mcconfig.BSC_ENDPOINT = bsc_end
	}
	eth_addr := os.Getenv("ETH_ZETALOCK_ADDRESS")
	if eth_addr != "" {
		fmt.Println("ETH_ZETALOCK_ADDRESS: ", eth_addr)
		mcconfig.ETH_ZETALOCK_ADDRESS = eth_addr
	}
	poly_addr := os.Getenv("POLYGON_TOKEN_ADDRESS")
	if poly_addr != "" {
		fmt.Println("POLYGON_TOKEN_ADDRESS: ", poly_addr)
		mcconfig.POLYGON_TOKEN_ADDRESS = poly_addr
	}
	bsc_addr := os.Getenv("BSC_TOKEN_ADDRESS")
	if bsc_addr != "" {
		fmt.Println("BSC_TOKEN_ADDRESS: ", bsc_addr)
		mcconfig.BSC_TOKEN_ADDRESS = bsc_addr
	}

	if *chainid != "" {
		cmd.CHAINID = *chainid
	}
	if *tssTestFlag {
		SetupConfigForTest()
		fmt.Println("testing TSS signing")

		tssServer, _, err := mc.SetupTSSServer(peers, "")
		if err != nil {
			log.Error().Err(err).Msg("setup TSS server error")
			return
		}

		time.Sleep(5 * time.Second)
		kgRes := mc.TestKeygen(tssServer)
		log.Debug().Msgf("keygen succeeds! TSS pubkey: %s", kgRes.PubKey)

		log.Debug().Msgf("Keysign test begins...")

		mc.TestKeysign(kgRes.PubKey, tssServer)

		// wait....
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Info().Msg("stop signal received")
		return
	}

	if *mockFlag == "tsstest" {
		fmt.Println("single node multiple clients tests")
		mock_integration_test() // single node testing environment; mocking multiple clients
		return
	} else if *mockFlag == "mocknet" {
		fmt.Println("single node single clients mocknet")
		mocknet_test(*validatorName)
		return
	} else {
		fmt.Println("multi-node client")
		integration_test(*validatorName, peers)
		return
	}

}

func SetupConfigForTest() {
	config := cosmos.GetConfig()
	config.SetBech32PrefixForAccount(cmd.Bech32PrefixAccAddr, cmd.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(cmd.Bech32PrefixValAddr, cmd.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(cmd.Bech32PrefixConsAddr, cmd.Bech32PrefixConsPub)
	//config.SetCoinType(cmd.MetaChainCoinType)
	config.SetFullFundraiserPath(cmd.ZetaChainHDPath)
	types.SetCoinDenomRegex(func() string {
		return cmd.DenomRegex
	})

	rand.Seed(time.Now().UnixNano())

}

func integration_test(validatorName string, peers addr.AddrList) {
	SetupConfigForTest() // setup meta-prefix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// wait until zetacore is up
	log.Info().Msg("Waiting for ZetaCore to open 9090 port...")
	for {
		_, err := grpc.Dial(
			fmt.Sprintf("127.0.0.1:9090"),
			grpc.WithInsecure(),
		)
		if err != nil {
			log.Warn().Err(err).Msg("grpc dial fail")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	log.Info().Msgf("ZetaCore to open 9090 port...")

	// setup 2 metabridges
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("UserHomeDir error")
		return
	}
	chainHomeFoler := filepath.Join(homeDir, ".zetacore")

	// first signer & bridge
	signerName := validatorName
	signerPass := "password"
	bridge1, done := CreateMetaBridge(chainHomeFoler, signerName, signerPass)
	if done {
		return
	}

	// setup mock TSS signers:
	// The following privkey has address 0xE80B6467863EbF8865092544f441da8fD3cF6074
	//privateKey, err := crypto.HexToECDSA(mcconfig.TSS_TEST_PRIVKEY)
	//if err != nil {
	//	log.Err(err).Msg("TEST private key error")
	//	return
	//}
	//tss := mc.TestSigner{
	//	PrivKey: privateKey,
	//}

	tss, err := mc.NewTSS(peers)
	if err != nil {
		log.Error().Err(err).Msg("NewTSS error")
		return
	}

	signerMap1, err := CreateSignerMap(tss)
	if err != nil {
		log.Err(err).Msg("CreateSignerMap")
		return
	}

	userDir, _ := os.UserHomeDir()
	dbpath := filepath.Join(userDir, ".zetaclient/chainobserver")
	chainClientMap1, err := CreateChainClientMap(bridge1, tss, dbpath)
	if err != nil {
		log.Err(err).Msg("CreateSignerMap")
		return
	}

	fmt.Print("Press 'Enter' to start...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	httpServer := mc.NewHTTPServer()

	log.Info().Msg("starting zetacore observer...")
	mo1 := mc.NewCoreObserver(bridge1, signerMap1, *chainClientMap1, httpServer)

	mo1.MonitorCore()

	// printout debug info from SIGUSR1
	// trigger by $ kill -SIGUSR1 <PID of zetaclient>
	usr := make(chan os.Signal, 1)
	signal.Notify(usr, syscall.SIGUSR1)
	go func() {
		for {
			<-usr
			fmt.Printf("Last blocks:\n")
			fmt.Printf("ETH     %d:\n", (*chainClientMap1)[common.ETHChain].LastBlock)
			fmt.Printf("BSC     %d:\n", (*chainClientMap1)[common.BSCChain].LastBlock)
			fmt.Printf("POLYGON %d:\n", (*chainClientMap1)[common.POLYGONChain].LastBlock)

		}
	}()

	// wait....
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Info().Msg("stop signal received")

	(*chainClientMap1)[common.ETHChain].Stop()
	(*chainClientMap1)[common.BSCChain].Stop()
	(*chainClientMap1)[common.POLYGONChain].Stop()
}

func mocknet_test(validatorName string) {
	SetupConfigForTest() // setup meta-prefix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// wait until zetacore is up
	log.Info().Msg("Waiting for ZetaCore to open 9090 port...")
	for {
		_, err := grpc.Dial(
			fmt.Sprintf("127.0.0.1:9090"),
			grpc.WithInsecure(),
		)
		if err != nil {
			log.Warn().Err(err).Msg("grpc dial fail")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	log.Info().Msgf("ZetaCore to open 9090 port...")

	// setup 2 metabridges
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("UserHomeDir error")
		return
	}
	chainHomeFoler := filepath.Join(homeDir, ".zetacore")

	// first signer & bridge
	signerName := validatorName
	signerPass := "password"
	bridge1, done := CreateMetaBridge(chainHomeFoler, signerName, signerPass)
	if done {
		return
	}

	//setup mock TSS signers:
	//The following privkey has address 0xE80B6467863EbF8865092544f441da8fD3cF6074
	privateKey, err := crypto.HexToECDSA(mcconfig.TSS_TEST_PRIVKEY)
	if err != nil {
		log.Err(err).Msg("TEST private key error")
		return
	}
	tss := mc.TestSigner{
		PrivKey: privateKey,
	}

	signerMap1, err := CreateSignerMap(tss)
	if err != nil {
		log.Err(err).Msg("CreateSignerMap")
		return
	}

	userDir, _ := os.UserHomeDir()
	dbpath := filepath.Join(userDir, ".zetaclient/chainobserver")
	chainClientMap1, err := CreateChainClientMap(bridge1, tss, dbpath)
	if err != nil {
		log.Err(err).Msg("CreateSignerMap")
		return
	}

	fmt.Print("Press 'Enter' to start...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	httpServer := mc.NewHTTPServer()

	log.Info().Msg("starting zetacore observer...")
	mo1 := mc.NewCoreObserver(bridge1, signerMap1, *chainClientMap1, httpServer)

	mo1.MonitorCore()

	// printout debug info from SIGUSR1
	// trigger by $ kill -SIGUSR1 <PID of zetaclient>
	usr := make(chan os.Signal, 1)
	signal.Notify(usr, syscall.SIGUSR1)
	go func() {
		for {
			<-usr
			fmt.Printf("Last blocks:\n")
			fmt.Printf("ETH     %d:\n", (*chainClientMap1)[common.ETHChain].LastBlock)
			fmt.Printf("BSC     %d:\n", (*chainClientMap1)[common.BSCChain].LastBlock)
			fmt.Printf("POLYGON %d:\n", (*chainClientMap1)[common.POLYGONChain].LastBlock)

		}
	}()

	// wait....
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Info().Msg("stop signal received")

	(*chainClientMap1)[common.ETHChain].Stop()
	(*chainClientMap1)[common.BSCChain].Stop()
	(*chainClientMap1)[common.POLYGONChain].Stop()
}
