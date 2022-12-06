package indexer

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nanmu42/etherscan-api"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"

	"testing"
)

type EthTx struct {
	TxHash  ethcommon.Hash
	Nonce   uint64            `gorm:"primaryKey"`
	ChainID uint64            `gorm:"primaryKey"`
	From    ethcommon.Address `gorm:"primaryKey"`
	Height  uint64
}

func Test1(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&EthTx{})

	tx := EthTx{
		Nonce:   1,
		TxHash:  ethcommon.HexToHash("0x123"),
		ChainID: 80001,
		From:    ethcommon.HexToAddress("0x1337"),
		Height:  1234,
	}
	db.Create(&tx)

	tx2 := EthTx{
		Nonce:   1,
		ChainID: 80001,
		From:    ethcommon.HexToAddress("0x1337"),
	}
	db.First(&tx2)
	t.Logf("tx2: nonce %d from %s txhash %s", tx2.Nonce, tx2.From.Hex(), tx2.TxHash.Hex())

	tx3 := EthTx{
		Nonce:   2,
		TxHash:  ethcommon.HexToHash("0x132"),
		ChainID: 80001,
		From:    ethcommon.HexToAddress("0x1337"),
		Height:  4567,
	}
	db.Create(&tx3)

	tx4 := EthTx{
		Nonce:   2,
		ChainID: 80001,
		From:    ethcommon.HexToAddress("0x1337"),
	}
	db.First(&tx4)
	t.Logf("tx4: nonce %d from %s height %d", tx4.Nonce, tx4.From.Hex(), tx4.Height)

	latestBlock := uint64(0)
	db.Table("eth_txes").Where("chain_id = ?", 80001).Select("MAX(height)").Row().Scan(&latestBlock)
	t.Logf("latest block: %d", latestBlock)
	assert.Equal(t, uint64(4567), latestBlock)
}

func Test2(t *testing.T) {
	MUMBAI_ETHERSCAN_ENDPOINT := "https://api-testnet.polygonscan.com/api?"
	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     "You key here",
		BaseURL: MUMBAI_ETHERSCAN_ENDPOINT,
		Verbose: false,
	})

	startTime := time.Now()
	startBlock := 29139311 // nov14, 2022
	txs, err := client.NormalTxByAddress("0x7c125C1d515b8945841b3d5144a060115C58725F", &startBlock, nil, 1, 100, false)
	t.Logf("NormalTxByAddress takes %s", time.Since(startTime))
	assert.Equal(t, nil, err)
	t.Logf("txs: %d", len(txs))
	for _, tx := range txs {
		t.Logf("tx nonce: %d hash %s", tx.Nonce, tx.Hash)
	}
}
