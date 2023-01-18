#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

DREDGER_CHAIN_ID=local-test-1
HOST_CHAIN_ID=osmosis-1
HOST_ENDPOINT=osmo-fleet-direct.main.dredgernet.co
HOST_ACCOUNT_PREFIX=osmo
HOST_DENOM=uosmo
HOST_BINARY=build/osmosisd
HOST_VAL_NAME_1=imperator
HOST_VAL_ADDRESS_1=osmovaloper1t8qckan2yrygq7kl9apwhzfalwzgc2429p8f0s
HOST_VAL_NAME_2=notional
HOST_VAL_ADDRESS_2=osmovaloper1083svrca4t350mphfv9x45wq9asrs60c6rv0j5
HOT_WALLET_ADDRESS=osmo1c37n9aywapx2v0s6vk2yedydkkhq65zz38jfnc

STATE=$SCRIPT_DIR/../state
LOGS=$SCRIPT_DIR/../logs
DREDGER_LOGS=$LOGS/dredger.log
DREDGER_HOME=$STATE/dred1
DOCKER_COMPOSE="docker-compose -f $SCRIPT_DIR/docker-compose.yml"

HERMES_DREDGER_MNEMONIC="alter old invest friend relief slot swear pioneer syrup economy vendor tray focus hedgehog artist legend antenna hair almost donkey spice protect sustain increase"
RELAYER_DREDGER_MNEMONIC="pride narrow breeze fitness sign bounce dose smart squirrel spell length federal replace coral lunar thunder vital push nuclear crouch fun accident hood need"

# cleanup any stale state
make stop-docker
rm -rf $STATE $LOGS
mkdir -p $STATE
mkdir -p $LOGS

# Start dredger
bash ${SCRIPT_DIR}/init_dredger.sh $DREDGER_CHAIN_ID

$DOCKER_COMPOSE up -d dred1
$DOCKER_COMPOSE logs -f dred1 | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" > $DREDGER_LOGS 2>&1 &

printf "Waiting for Dredger to start..."
( tail -f -n0 $DREDGER_LOGS & ) | grep -q "finalizing commit of block"
echo "Done"

# Setup relayers
mkdir -p $STATE/hermes
mkdir -p $STATE/relayer/config
HERMES_CONFIG_FILE="$STATE/hermes/config.toml"
RELAYER_CONFIG_FILE="$STATE/relayer/config/config.yaml"
cp ${SCRIPT_DIR}/templates/hermes_config.toml $HERMES_CONFIG_FILE
cp ${SCRIPT_DIR}/templates/relayer_config.yaml $RELAYER_CONFIG_FILE

# Update relayer templates
sed -i -E "s|DREDGER_CHAIN_ID|$DREDGER_CHAIN_ID|g" $HERMES_CONFIG_FILE
sed -i -E "s|HOST_CHAIN_ID|$HOST_CHAIN_ID|g" $HERMES_CONFIG_FILE
sed -i -E "s|HOST_ENDPOINT|$HOST_ENDPOINT|g" $HERMES_CONFIG_FILE
sed -i -E "s|HOST_ACCOUNT_PREFIX|$HOST_ACCOUNT_PREFIX|g" $HERMES_CONFIG_FILE
sed -i -E "s|HOST_DENOM|$HOST_DENOM|g" $HERMES_CONFIG_FILE

sed -i -E "s|DREDGER_CHAIN_ID|$DREDGER_CHAIN_ID|g" $RELAYER_CONFIG_FILE
sed -i -E "s|HOST_CHAIN_ID|$HOST_CHAIN_ID|g" $RELAYER_CONFIG_FILE
sed -i -E "s|HOST_ENDPOINT|$HOST_ENDPOINT|g" $RELAYER_CONFIG_FILE
sed -i -E "s|HOST_ACCOUNT_PREFIX|$HOST_ACCOUNT_PREFIX|g" $RELAYER_CONFIG_FILE
sed -i -E "s|HOST_DENOM|$HOST_DENOM|g" $RELAYER_CONFIG_FILE

echo "Adding Hermes keys"
HERMES_CMD="$SCRIPT_DIR/../../build/hermes/release/hermes --config $STATE/hermes/config.toml"
TMP_MNEMONICS=$STATE/mnemonic.txt 
echo "$HERMES_DREDGER_MNEMONIC" > $TMP_MNEMONICS
$HERMES_CMD keys add --key-name hrly1 --chain $DREDGER_CHAIN_ID --mnemonic-file $TMP_MNEMONICS --overwrite
echo "$HOT_WALLET_2_MNEMONIC" > $TMP_MNEMONICS
$HERMES_CMD keys add --key-name hrly2 --chain $HOST_CHAIN_ID --mnemonic-file $TMP_MNEMONICS --overwrite
rm -f $TMP_MNEMONICS

echo "Adding Relayer keys"
RELAYER_CMD="$SCRIPT_DIR/../../build/relayer --home $STATE/relayer"
$RELAYER_CMD keys restore dredger rly1 "$RELAYER_DREDGER_MNEMONIC" 
$RELAYER_CMD keys restore host rly2 "$HOT_WALLET_3_MNEMONIC" 

# Update commands template
COMMANDS_FILE=${SCRIPT_DIR}/commands.sh
cp ${SCRIPT_DIR}/templates/commands.sh $COMMANDS_FILE
DOCKER_COMPOSE_RELATIVE="docker-compose -f scripts/local-to-mainnet/docker-compose.yml"
STATE_RELATIVE=scripts/state
LOGS_RELATIVE=scripts/logs
sed -i -E '1s/^/############################################\n### WARNING: THIS FILE IS AUTOGENERATED. ###\n###   ANY CHANGES WILL BE OVERWRITTEN.   ###\n############################################\n/' $COMMANDS_FILE
sed -i -E "s|DOCKER_COMPOSE|$DOCKER_COMPOSE_RELATIVE|g" $COMMANDS_FILE
sed -i -E "s|STATE|$STATE_RELATIVE|g" $COMMANDS_FILE
sed -i -E "s|LOGS|$LOGS_RELATIVE|g" $COMMANDS_FILE
sed -i -E "s|DREDGER_HOME|s|g" $COMMANDS_FILE
sed -i -E "s|DREDGER_CHAIN_ID|$DREDGER_CHAIN_ID|g" $COMMANDS_FILE
sed -i -E "s|HOST_CHAIN_ID|$HOST_CHAIN_ID|g" $COMMANDS_FILE
sed -i -E "s|HOST_BINARY|$HOST_BINARY|g" $COMMANDS_FILE
sed -i -E "s|HOST_DENOM|$HOST_DENOM|g" $COMMANDS_FILE
sed -i -E "s|HOST_ACCOUNT_PREFIX|$HOST_ACCOUNT_PREFIX|g" $COMMANDS_FILE
sed -i -E "s|HOST_ENDPOINT|$HOST_ENDPOINT|g" $COMMANDS_FILE
sed -i -E "s|HOST_VAL_NAME_1|$HOST_VAL_NAME_1|g" $COMMANDS_FILE
sed -i -E "s|HOST_VAL_NAME_2|$HOST_VAL_NAME_2|g" $COMMANDS_FILE
sed -i -E "s|HOST_VAL_ADDRESS_1|$HOST_VAL_ADDRESS_1|g" $COMMANDS_FILE
sed -i -E "s|HOST_VAL_ADDRESS_2|$HOST_VAL_ADDRESS_2|g" $COMMANDS_FILE
sed -i -E "s|HOT_WALLET_ADDRESS|$HOT_WALLET_ADDRESS|g" $COMMANDS_FILE

rm -f $COMMANDS_FILE-E
echo "Done"