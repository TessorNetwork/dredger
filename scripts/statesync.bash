#!/bin/bash
# microtick and bitcanna contributed significantly here.
# Pebbledb state sync script.
# invoke like: bash scripts/ss.bash

## USAGE RUNDOWN
# Not for use on live nodes
# For use when testing.
# Assumes that ~/.evmosd doesn't exist
# can be modified to suit your purposes if ~/.evmosd does already exist

set -uxe

# Set Golang environment variables.
export GOPATH=~/go
export PATH=$PATH:~/go/bin

# Install with pebbledb
go mod edit -replace github.com/tendermint/tm-db=github.com/baabeetaa/tm-db@pebble
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb -X github.com/tendermint/tm-db.ForceSync=1' -tags pebbledb ./...

# go install ./...

# NOTE: ABOVE YOU CAN USE ALTERNATIVE DATABASES, HERE ARE THE EXACT COMMANDS
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags rocksdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags badgerdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb' -tags boltdb ./...

# Initialize chain.
dredger init test

# Get Genesis
wget https://raw.githubusercontent.com/Stride-Labs/testnet/main/mainnet/genesis.json
mv genesis.json ~/.dredger/config/genesis.json

# Get "trust_hash" and "trust_height".
INTERVAL=1000
LATEST_HEIGHT=$(curl -s https://dredger-rpc.polkachu.com/block | jq -r .result.block.header.height)
BLOCK_HEIGHT=$(($LATEST_HEIGHT - $INTERVAL))
TRUST_HASH=$(curl -s "https://dredger-rpc.polkachu.com/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables.
export DREDGERD_STATESYNC_ENABLE=true
export DREDGERD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export DREDGERD_STATESYNC_RPC_SERVERS="https://dredger-rpc.polkachu.com:443,https://dredger-rpc.polkachu.com:443"
export DREDGERD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export DREDGERD_STATESYNC_TRUST_HASH=$TRUST_HASH

# Fetch and set list of seeds from chain registry.
DREDGERD_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/dredger/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')
export DREDGERD_P2P_SEEDS

# Start chain.
dredger start --x-crisis-skip-assert-invariants --db_backend pebbledb
