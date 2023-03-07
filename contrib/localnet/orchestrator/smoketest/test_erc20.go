//go:build PRIVNET
// +build PRIVNET

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	contracts "github.com/zeta-chain/zetacore/contracts/zevm"
)

func (sm *SmokeTest) TestERC20Deposit() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("test finishes in %s\n", time.Since(startTime))
	}()
	LoudPrintf("Deposit USDT ERC20 into ZEVM\n")
	txhash := sm.DepositERC20(big.NewInt(1e9), []byte{})
	WaitCctxMinedByInTxHash(txhash.Hex(), sm.cctxClient)
	usdtZRC20, err := contracts.NewZRC20(ethcommon.HexToAddress(USDTZRC20Addr), sm.zevmClient)
	if err != nil {
		panic(err)
	}
	bal, err := usdtZRC20.BalanceOf(&bind.CallOpts{}, DeployerAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of deployer on USDT ZRC20: %d\n", bal)
	supply, err := usdtZRC20.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("supply of USDT ZRC20: %d\n", supply)
	if bal.Int64() != 1e9 {
		panic("balance is not correct")
	}
}

func (sm *SmokeTest) DepositERC20(amount *big.Int, msg []byte) ethcommon.Hash {
	USDT := sm.USDTERC20
	tx, err := USDT.Mint(sm.goerliAuth, big.NewInt(1e10))
	if err != nil {
		panic(err)
	}
	receipt := MustWaitForTxReceipt(sm.goerliClient, tx)
	fmt.Printf("Mint receipt tx hash: %s\n", tx.Hash().Hex())

	tx, err = USDT.Approve(sm.goerliAuth, sm.ERC20CustodyAddr, big.NewInt(1e10))
	if err != nil {
		panic(err)
	}
	receipt = MustWaitForTxReceipt(sm.goerliClient, tx)
	fmt.Printf("USDT Approve receipt tx hash: %s\n", tx.Hash().Hex())

	tx, err = sm.ERC20Custody.Deposit(sm.goerliAuth, DeployerAddress.Bytes(), sm.USDTERC20Addr, amount, msg)
	if err != nil {
		panic(err)
	}
	receipt = MustWaitForTxReceipt(sm.goerliClient, tx)
	fmt.Printf("Deposit receipt tx hash: %s, status %d\n", receipt.TxHash.Hex(), receipt.Status)
	for _, log := range receipt.Logs {
		event, err := sm.ERC20Custody.ParseDeposited(*log)
		if err != nil {
			continue
		}
		fmt.Printf("Deposited event: \n")
		fmt.Printf("  Recipient address: %x, \n", event.Recipient)
		fmt.Printf("  ERC20 address: %s, \n", event.Asset.Hex())
		fmt.Printf("  Amount: %d, \n", event.Amount)
		fmt.Printf("  Message: %x, \n", event.Message)
	}
	fmt.Printf("gas limit %d\n", sm.zevmAuth.GasLimit)
	return tx.Hash()
	//WaitCctxMinedByInTxHash(tx.Hash().Hex(), sm.cctxClient)
}

func (sm *SmokeTest) TestERC20Withdraw() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("test finishes in %s\n", time.Since(startTime))
	}()
	LoudPrintf("Withdraw USDT ZRC20\n")

	sm.WithdrawERC20()
}

func (sm *SmokeTest) WithdrawERC20() {
	zevmClient := sm.zevmClient
	//goerliClient := sm.goerliClient
	//cctxClient := sm.cctxClient

	usdtZRC20, err := contracts.NewZRC20(ethcommon.HexToAddress(USDTZRC20Addr), zevmClient)
	if err != nil {
		panic(err)
	}
	bal, err := usdtZRC20.BalanceOf(&bind.CallOpts{}, DeployerAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of deployer on USDT ZRC20: %d\n", bal)
	supply, err := usdtZRC20.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("supply of USDT ZRC20: %d\n", supply)
	if bal.Int64() != 1e9 {
		panic("balance is not correct")
	}

	gasZRC20, gasFee, err := usdtZRC20.WithdrawGasFee(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("gasZRC20: %s, gasFee: %d\n", gasZRC20.Hex(), gasFee)

	ethZRC20, err := contracts.NewZRC20(gasZRC20, zevmClient)
	if err != nil {
		panic(err)
	}
	bal, err = ethZRC20.BalanceOf(&bind.CallOpts{}, DeployerAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of deployer on ETH ZRC20: %d\n", bal)
	if bal.Int64() <= 0 {
		panic("not enough ETH ZRC20 balance!")
	}

	// Approve
	tx, err := ethZRC20.Approve(sm.zevmAuth, ethcommon.HexToAddress(USDTZRC20Addr), big.NewInt(1e18))
	if err != nil {
		panic(err)
	}
	receipt := MustWaitForTxReceipt(zevmClient, tx)
	fmt.Printf("eth zrc20 approve receipt: status %d\n", receipt.Status)

	//get different clients
	client1 := getClient("http://zetacore1:8545")
	client2 := getClient("http://zetacore2:8545")
	client3 := getClient("http://zetacore3:8545")

	// Withdraw
	var wg sync.WaitGroup
	wg.Add(4)
	go sm.Withdraw(zevmClient, &wg)
	go sm.Withdraw(client1, &wg)
	go sm.Withdraw(client2, &wg)
	go sm.Withdraw(client3, &wg)

	wg.Wait()
	//for {
	//	tx, err = usdtZRC20.Withdraw(sm.zevmAuth, DeployerAddress.Bytes(), big.NewInt(100))
	//	if err != nil {
	//		panic(err)
	//	}
	//	receipt = MustWaitForTxReceipt(zevmClient, tx)
	//	fmt.Printf("Receipt txhash %s status %d\n", receipt.TxHash, receipt.Status)
	//}
	//for _, log := range receipt.Logs {
	//	event, err := usdtZRC20.ParseWithdrawal(*log)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Printf("  logs: from %s, to %x, value %d, gasfee %d\n", event.From.Hex(), event.To, event.Value, event.Gasfee)
	//}

	//sm.wg.Add(1)
	//go func() {
	//	defer sm.wg.Done()
	//	cctx := WaitCctxMinedByInTxHash(receipt.TxHash.Hex(), cctxClient)
	//	fmt.Printf("outTx hash %s\n", cctx.GetCurrentOutTxParam().OutboundTxHash)
	//
	//	USDTERC20, err := erc20.NewUSDT(ethcommon.HexToAddress(USDTERC20Addr), goerliClient)
	//	if err != nil {
	//		panic(err)
	//	}
	//	bal, err = USDTERC20.BalanceOf(&bind.CallOpts{}, DeployerAddress)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("USDT ERC20 bal: %d\n", bal)
	//
	//	receipt, err := sm.goerliClient.TransactionReceipt(context.Background(), ethcommon.HexToHash(cctx.GetCurrentOutTxParam().OutboundTxHash))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("Receipt txhash %s status %d\n", receipt.TxHash, receipt.Status)
	//	for _, log := range receipt.Logs {
	//		event, err := USDTERC20.ParseTransfer(*log)
	//		if err != nil {
	//			continue
	//		}
	//		fmt.Printf("  logs: from %s, to %s, value %d\n", event.From.Hex(), event.To.Hex(), event.Value)
	//		if event.Value.Int64() != 100 {
	//			panic("value is not correct")
	//		}
	//	}
	//}()
	//sm.wg.Wait()
}

func getClient(url string) *ethclient.Client {
	var client *ethclient.Client
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("dialing zevm client: \n", url)
		var err error
		client, err = ethclient.Dial(url)
		if err != nil {
			continue
		}
		break
	}
	return client
}

func (sm *SmokeTest) Withdraw(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	usdtZRC20, _ := contracts.NewZRC20(ethcommon.HexToAddress(USDTZRC20Addr), client)
	for {
		tx, err := usdtZRC20.Withdraw(sm.zevmAuth, DeployerAddress.Bytes(), big.NewInt(100))
		if err != nil {
			panic(err)
		}
		receipt := MustWaitForTxReceipt(client, tx)
		fmt.Printf("Receipt txhash %s status %d\n", receipt.TxHash, receipt.Status)
	}
}
