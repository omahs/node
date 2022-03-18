package main

import (
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/zeta-chain/zetacore/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	//authtx "github.com/cosmos/cosmos-sdk/x/auth/tx" in later cosmos versions
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

type ChainCosmos struct {
	name    common.Chain
	nodeURI string
	chainID string

	MPI_CONTRACT                 string
	DEFAULT_DESTINATION_CONTRACT string
}

func (c *ChainCosmos) newContext() (client.Context, error) {
	httpclient, err := rpchttp.New(c.nodeURI, "/websocket")
	if err != nil {
		panic(err)
	}
	return client.Context{}.
		WithChainID(c.chainID).
		WithOutput(os.Stdout).
		WithOutputFormat("text").
		WithNodeURI(c.nodeURI).
		WithClient(httpclient), nil
}

func (c *ChainCosmos) Start() {
	ctx, err := c.newContext()
	if err != nil {
		log.Fatalf("failed to create context: %s", err)
	}

	// where mpi_contract is the tx hash
	output, err := authclient.QueryTx(ctx, c.MPI_CONTRACT)
	if err != nil {
		log.Fatalf("failed to query tx: %s", err)
	}

	if output.Empty() {
		log.Fatalf("no such transaction found: %s", c.MPI_CONTRACT)
	}

	if err := ctx.PrintProto(output); err != nil {
		log.Fatalf("error during tx output")
	}
}
