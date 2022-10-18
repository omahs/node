//go:generate sh -c "cat ConnectorEth.json | jq .abi | abigen --abi - --pkg evm --type Connector --out ConnectorEth.go"
//go:generate sh -c "cat ZetaEth.json | jq .abi | abigen --abi - --pkg evm --type ZetaEth --out ZetaEth.go"
//go:generate sh -c "cat ZetaNonEth.json | jq .abi | abigen --abi - --pkg evm --type ZetaNonEth --out ZetaNonEth.go"

package evm

import (
	_ "embed"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

type CompiledContract struct {
	ABI abi.ABI
	Bin evmtypes.HexString
}

var _ = Connector{}
var _ = ZetaEth{}
var _ = ZetaNonEth{}

var (
	//go:embed ConnectorEth.json
	ConnectorEthJSON []byte // nolint: golint

	//go:embed ConnectorNonEth.json
	ConnectorNonEthJSON []byte // nolint: golint

	//go:embed ZetaEth.json
	ZetaEthJSON []byte // nolint: golint

	//go:embed ZetaNonEth.json
	ZetaNonEthJSON []byte // nolint: golint

	ConnectorEthContract    CompiledContract
	ConnectorNonEthContract CompiledContract
	ZetaEthContract         CompiledContract
	ZetaNonEthContract      CompiledContract
)

func init() {
	err := json.Unmarshal(ConnectorEthJSON, &ConnectorEthContract)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ConnectorNonEthJSON, &ConnectorNonEthContract)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ZetaEthJSON, &ZetaEthContract)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ZetaNonEthJSON, &ZetaNonEthContract)
	if err != nil {
		panic(err)
	}

	if len(ConnectorEthContract.Bin) == 0 {
		panic("load contract failed")
	}
	if len(ConnectorNonEthContract.Bin) == 0 {
		panic("load contract failed")
	}
	if len(ZetaEthContract.Bin) == 0 {
		panic("load contract failed")
	}
	if len(ZetaNonEthContract.Bin) == 0 {
		panic("load contract failed")
	}
}
