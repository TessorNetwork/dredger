#!/bin/sh
set -e
set -o pipefail

DREDGER_HOME=$HOME/.dredger
CONFIG_FOLDER=$DREDGER_HOME/config

DEFAULT_MNEMONIC="deer gaze swear marine one perfect hero twice turkey symbol mushroom hub escape accident prevent rifle horse arena secret endless panel equal rely payment"
DEFAULT_CHAIN_ID="localdredger"
DEFAULT_MONIKER="val"

# Override default values with environment variables
MNEMONIC=${MNEMONIC:-$DEFAULT_MNEMONIC}
CHAIN_ID=${CHAIN_ID:-$DEFAULT_CHAIN_ID}
MONIKER=${MONIKER:-$DEFAULT_MONIKER}

install_prerequisites () {
    sudo apk add -q --no-cache \
        python3 \
        py3-pip
}

edit_config () {

    # Remove seeds
    dasel put string -f $CONFIG_FOLDER/config.toml '.p2p.seeds' ''

    # Disable fast_sync
    dasel put bool -f $CONFIG_FOLDER/config.toml '.fast_sync' 'false'

    # Expose the rpc
    dasel put string -f $CONFIG_FOLDER/config.toml '.rpc.laddr' "tcp://0.0.0.0:26657"
}

if [[ ! -d $CONFIG_FOLDER ]]
then

    install_prerequisites

    echo "Chain ID: $CHAIN_ID"
    echo "Moniker:  $MONIKER"
    echo "MNEMONIC: $MNEMONIC"
    echo "DREDGER_HOME: $DREDGER_HOME"

    echo $MNEMONIC | dredger init localdredger -o --chain-id=$CHAIN_ID --home $DREDGER_HOME
    echo $MNEMONIC | dredger keys add val --recover --keyring-backend test

    ACCOUNT_PUBKEY=$(dredger keys show --keyring-backend test val --pubkey | dasel -r json '.key' --plain)
    ACCOUNT_ADDRESS=$(dredger keys show -a --keyring-backend test val --bech acc)

    VALIDATOR_PUBKEY_JSON=$(dredger tendermint show-validator --home $DREDGER_HOME)
    VALIDATOR_PUBKEY=$(echo $VALIDATOR_PUBKEY_JSON | dasel -r json '.key' --plain)
    VALIDATOR_HEX_ADDRESS=$(dredger debug pubkey $VALIDATOR_PUBKEY_JSON 2>&1 --home $DREDGER_HOME | grep Address | cut -d " " -f 2)
    VALIDATOR_ACCOUNT_ADDRESS=$(dredger debug addr $VALIDATOR_HEX_ADDRESS 2>&1  --home $DREDGER_HOME | grep Acc | cut -d " " -f 3)
    VALIDATOR_OPERATOR_ADDRESS=$(dredger debug addr $VALIDATOR_HEX_ADDRESS 2>&1  --home $DREDGER_HOME | grep Val | cut -d " " -f 3)
    VALIDATOR_CONSENSUS_ADDRESS=$(dredger tendermint show-address --home $DREDGER_HOME)

    python3 -u testnetify.py \
    -i /home/dredger/state_export.json \
    -o $CONFIG_FOLDER/genesis.json \
    -c $CHAIN_ID \
    --validator-hex-address $VALIDATOR_HEX_ADDRESS \
    --validator-operator-address $VALIDATOR_OPERATOR_ADDRESS \
    --validator-consensus-address $VALIDATOR_CONSENSUS_ADDRESS \
    --validator-pubkey $VALIDATOR_PUBKEY \
    --account-pubkey $ACCOUNT_PUBKEY \
    --account-address $ACCOUNT_ADDRESS

    edit_config
fi

dredger start --home $DREDGER_HOME --x-crisis-skip-assert-invariants
