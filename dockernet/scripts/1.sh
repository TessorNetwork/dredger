### IBC TRANSFER
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/../config.sh

## IBC ATOM from GAIA to DREDGER
$GAIA_MAIN_CMD tx ibc-transfer transfer transfer channel-0 $(DREDGER_ADDRESS) 1000000uatom --from ${GAIA_VAL_PREFIX}1 -y 
sleep 10
$DREDGER_MAIN_CMD q bank balances $(DREDGER_ADDRESS)