package main

import (
	"github.com/nanmu42/etherscan-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEtherscanAPI(t *testing.T) {
	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     "You key here",
		BaseURL: "https://api-testnet.polygonscan.com/api?",
		Verbose: false,
	})
	//balance, err := client.AccountBalance("0x6dA30bFA65E85a16b05bCE3846339ed2BC746316")
	//t.Logf("balance: %s, err: %v", balance.Int().String(), err)

	startTime := time.Now()
	startBlock := 29444807
	txs, err := client.NormalTxByAddress("0x7c125C1d515b8945841b3d5144a060115C58725F", &startBlock, nil, 1, 100, false)
	t.Logf("NormalTxByAddress takes %s", time.Since(startTime))
	assert.Equal(t, nil, err)
	t.Logf("txs: %d", len(txs))
	for _, tx := range txs {
		t.Logf("tx nonce: %d hash %s", tx.Nonce, tx.Hash)
	}
}
