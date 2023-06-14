package main

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"net/http"
)

type OutTxDigestRequest struct {
	SendHash       string `json:"sendHash"`
	OutTxHash      string `json:"outTxHash"`
	OutBlockHeight uint64 `json:"outBlockHeight"`
	Amount         string `json:"amount"`
	ChainID        int64  `json:"chainID"`
	Nonce          uint64 `json:"nonce"`
	CoinType       string `json:"cointType"`
}

type OutTxDigestResponse struct {
	Digest string `json:"digest"`
}

type HashService struct{}

func (h *HashService) OutTxDigest(_ *http.Request, outtx *OutTxDigestRequest, result *OutTxDigestResponse) error {
	coinType, found := common.CoinType_value[outtx.CoinType]
	if !found {
		return fmt.Errorf("invalid coin type: %s", outtx.CoinType)
	}
	amount := math.NewUintFromString(outtx.Amount)

	msg := types.NewMsgReceiveConfirmation("", outtx.SendHash, outtx.OutTxHash, outtx.OutBlockHeight, amount, common.ReceiveStatus_Failed, outtx.ChainID, outtx.Nonce, common.CoinType(coinType))
	*result = OutTxDigestResponse{
		Digest: msg.Digest(),
	}
	return nil
}

// to test, create a file with content:
//
//	{
//	 "method": "hash.OutTxDigest",
//	 "params": [{
//	     "sendHash": "0x598fdd00ef3e0c62f65b388095c7e1f87795908da002962e0a27a2113e10ce32",
//	     "outTxHash": "0xb8c8707dc8e90673dcde2c4799a2c0d35acc1b43a8b3f93c9d9661b715cac193",
//	     "outBlockHeight": 30635730,
//	     "amount": "0",
//	     "chainId": 97,
//	     "nonce": 0,
//	     "cointType": "Zeta"
//	 }],
//	 "id": 1
//	}
//
// and use CURL to post it:
// curl -X POST -H "Content-Type: application/json" --data @payload.json http://localhost:9001/out_tx_digest
// you should get a response like:
// {"result":{"digest":"0x3bc920a14e8fa9885a0940c084dd5c921c191d81bf018e7e7ed9cf342e0008bc"},"error":null,"id":1}
func main() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	hashService := new(HashService)
	rpcServer.RegisterService(hashService, "hash")

	router := mux.NewRouter()
	router.Handle("/out_tx_digest", rpcServer)
	port := 9001
	fmt.Printf("Listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
