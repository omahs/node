zetacored tx gov submit-legacy-proposal param-change standalone-network/gov/observer-params.json --from zeta --fees=40azeta --chain-id=localnet_101-1 --keyring-backend=test -y --broadcast-mode=block
zetacored tx gov vote 1 yes --from zeta --keyring-backend test --chain-id localnet_101-1 --yes --fees=40azeta --broadcast-mode=block

