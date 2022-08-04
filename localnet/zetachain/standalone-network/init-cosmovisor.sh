#!/usr/bin/env bash


export DAEMON_HOME=$HOME/.zetacore
export DAEMON_NAME=zetacored
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true
export CLIENT_DAEMON_NAME=zetaclientd
export CLIENT_DAEMON_ARGS="-enable-chains,GOERLI,-val zeta"
#export DAEMON_DATA_BACKUP_DIR=$DAEMON_HOME

make clean
make install
# Genesis
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/zetacored $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/zetaclientd $DAEMON_HOME/cosmovisor/genesis/bin


chmod +x $DAEMON_HOME/cosmovisor/genesis/bin/zetacored
chmod +x $DAEMON_HOME/cosmovisor/genesis/bin/zetaclientd

zetacored init test --chain-id=localnet -o
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | zetacored keys add zeta --recover --keyring-backend=test
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | zetacored keys add mario --recover --keyring-backend=test
zetacored add-genesis-account $(zetacored keys show zeta -a --keyring-backend=test) 500000000000000000000000000000000stake --keyring-backend=test
zetacored add-genesis-account $(zetacored keys show mario -a --keyring-backend=test)  500000000000000000000000000000000stake --keyring-backend=test
zetacored gentx zeta 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test
zetacored collect-gentxs
zetacored validate-genesis

cosmovisor start --home ~/.zetacore/ --p2p.laddr 0.0.0.0:27655  --grpc.address 0.0.0.0:9096 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:26657
