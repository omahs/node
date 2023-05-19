#!/usr/bin/env bash

CHAINID="localnet_101-1"
KEYRING="test"

rm -rf ~/.zetacored
kill -9 $(lsof -ti:26657)
export DAEMON_HOME=$HOME/.zetacored
export DAEMON_NAME=zetacored
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true
export CLIENT_DAEMON_NAME=zetaclientd
export CLIENT_DAEMON_ARGS="start"
export DAEMON_DATA_BACKUP_DIR=$DAEMON_HOME
export CLIENT_SKIP_UPGRADE=false
export CLIENT_START_PROCESS=false
export UNSAFE_SKIP_BACKUP=true



make install
# Genesis
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin


cp $GOPATH/bin/zetacored $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/zetaclientd $DAEMON_HOME/cosmovisor/genesis/bin


chmod +x $DAEMON_HOME/cosmovisor/genesis/bin/zetacored
chmod +x $DAEMON_HOME/cosmovisor/genesis/bin/zetaclientd

zetacored init test --chain-id=$CHAINID -o
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | zetacored keys add zeta --recover --keyring-backend=test --algo=secp256k1
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | zetacored keys add mario --recover --keyring-backend=test --algo=secp256k1
zetacored add-genesis-account $(zetacored keys show zeta -a --keyring-backend=test) 500000000000000000000000000000000azeta --keyring-backend=test
zetacored add-genesis-account $(zetacored keys show mario -a --keyring-backend=test)  500000000000000000000000000000000azeta --keyring-backend=test

cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="azeta"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="azeta"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="azeta"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="azeta"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="azeta"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="500000000"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json
cat "$DAEMON_HOME"/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="10s"' > "$DAEMON_HOME"/config/tmp_genesis.json && mv "$DAEMON_HOME"/config/tmp_genesis.json "$DAEMON_HOME"/config/genesis.json

zetacored gentx zeta 1000000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING
zetacored collect-gentxs
zetacored validate-genesis


zetacored config keyring-backend $KEYRING --home ~/.zetacored
zetaclientd init test --chain-id=$CHAINID --operator zeta1syavy2npfyt9tcncdtsdzf7kny9lh777heefxk --hotkey=zeta --public-ip=0.0.0.0

#cosmovisor start --log_format json
cosmovisor start --log_format json >> ~/.zetacored/cosmovisor.json 2>&1  &

sleep 7
printf "Raising the governance proposal:\n"
zetacored tx gov submit-legacy-proposal software-upgrade v1.0.0 \
  --from zeta \
  --deposit 10000000000000000000azeta \
  --upgrade-height 6 \
  --upgrade-info '{"binaries":{"zetaclientd-darwin/arm64":"https://athens3-testnet.s3.amazonaws.com/binaries/zetaclientd-mac?checksum=sha256:9a10a058f33bb55ce6af420716bf6df37f43596dfe7ccdc86f90d2a89efdc027","zetacored-darwin/arm64":"https://athens3-testnet.s3.amazonaws.com/binaries/zetacored-mac?checksum=sha256:b014b086c0de4ddce46087a8f24940e68cf115b3406f3a9792464a41247fb59e"}}' \
  --description "test" \
  --title "test-upgrade" \
  --from zeta \
  --keyring-backend $KEYRING \
  --chain-id $CHAINID \
  --gas=auto --gas-adjustment=1.5 --gas-prices=0.001azeta \
  --yes
sleep 7
zetacored tx gov vote 1 yes --from zeta --keyring-backend $KEYRING --chain-id $CHAINID --yes
clear
sleep 10
zetacored query gov proposal 1

tail -f ~/.zetacored/cosmovisor.json

#'{"binaries":{"zetaclientd-darwin/arm64":"https://filebin.net/4awhitgraq8eenpd/zetaclientd","zetacored-darwin/arm64":"https://filebin.net/4awhitgraq8eenpd/zetacored"}}'
#'{"binaries":{"zetacored-linux/amd64": "https://zetachain-external-files.s3.amazonaws.com/binaries/athens3/v1.0.0/zetacored-ubuntu-22-amd64?checksum=sha256:8106290054b0297a2d00adaeea0257d1e968fcf71ed50885effd1ffb2edc8c64","zetacored-linux/arm64": "https://zetachain-external-files.s3.amazonaws.com/binaries/athens3/v1.0.0/zetacored-ubuntu-22-arm64?checksum=sha256:533b8938e1f1e8b0e48f6dbbaf509e15cda3fc4f81a971f5534e349d23d7d3da","zetaclientd-linux/amd64": "https://zetachain-external-files.s3.amazonaws.com/binaries/athens3/v1.0.0/zetaclientd-ubuntu-22-arm64?checksum=sha256:d188cdd2dbdad4d1e46d96c8abac9e60039f7bbed5c76539f17568329d391232"}}'