//go:generate sh -c "cat ConnectorEth.json | jq .abi | abigen --abi - --pkg evm --type Connector --out Connector.go"
//go:generate sh -c "cat ZetaEth.json | jq .abi | abigen --abi - --pkg evm --type ZetaEth --out ZetaEth.go"
//go:generate sh -c "cat ZetaNonEth.json | jq .abi | abigen --abi - --pkg evm --type ZetaNonEth --out ZetaNonEth.go"

package evm

var _ = Connector{}
var _ = ZetaEth{}
var _ = ZetaNonEth{}
