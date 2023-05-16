
CHAINID="localnet_101-1"
KEYRING="test"
signer=zeta
PID=1
HOSTNAME=$(hostname)

zetacored tx gov submit-legacy-proposal param-change ../athens3/update_blockgas.json --from $signer --gas=auto --gas-adjustment=1.5 --gas-prices=0.001azeta --chain-id=$CHAINID --keyring-backend=$KEYRING -y --broadcast-mode=block
zetacored tx gov vote "$PID" yes --from $signer --keyring-backend $KEYRING --chain-id $CHAINID --yes --fees=40azeta --broadcast-mode=block
