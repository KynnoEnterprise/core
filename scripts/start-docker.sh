#!/bin/bash

KEY="mykey"
CHAINID="kynno_9700-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t kynno-datadir.XXXXX)

echo "create and add new keys"
./kynnod keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init kynno with moniker=$MONIKER and chain-id=$CHAINID"
./kynnod init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./kynnod add-genesis-account \
"$(./kynnod keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000akynno,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./kynnod gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./kynnod collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./kynnod validate-genesis --home $DATA_DIR

echo "starting kynno node $i in background ..."
./kynnod start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started kynno node"
tail -f /dev/null