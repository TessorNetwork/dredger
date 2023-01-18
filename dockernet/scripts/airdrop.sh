### AIRDROP TESTING FLOW
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/../config.sh

# First, start the network with `make start-docker`
# Then, run this script with `bash dockernet/scripts/airdrop.sh`

# NOTE: First, store the keys using the following mnemonics
# distributor address: dredger1z835j3j65nqr6ng257q0xkkc9gta72gf48txwl
# distributor mnemonic: barrel salmon half click confirm crunch sense defy salute process cart fiscal sport clump weasel render private manage picture spell wreck hill frozen before
echo "barrel salmon half click confirm crunch sense defy salute process cart fiscal sport clump weasel render private manage picture spell wreck hill frozen before" | \
    $DREDGER_MAIN_CMD keys add distributor-test --recover

# airdrop-test address: dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr
# airdrop claimer mnemonic: royal auction state december october hip monster hotel south help bulk supreme history give deliver pigeon license gold carpet rabbit raw wool fatigue donate
echo "royal auction state december october hip monster hotel south help bulk supreme history give deliver pigeon license gold carpet rabbit raw wool fatigue donate" | \
    $DREDGER_MAIN_CMD keys add airdrop-test --recover

## AIRDROP SETUP
echo "Funding accounts..."
# Transfer uatom from gaia to dredger, so that we can liquid stake later
$GAIA_MAIN_CMD tx ibc-transfer transfer transfer channel-0 dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr 1000000uatom --from ${GAIA_VAL_PREFIX}1 -y 
sleep 5
# Fund the distributor account
$DREDGER_MAIN_CMD tx bank send val1 dredger1z835j3j65nqr6ng257q0xkkc9gta72gf48txwl 600000udredger --from val1 -y
sleep 5
# Fund the airdrop account
$DREDGER_MAIN_CMD tx bank send val1 dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr 1000000000udredger --from val1 -y
sleep 5
# Create the airdrop, so that the airdrop account can claim tokens
$DREDGER_MAIN_CMD tx claim create-airdrop dredger 1666792900 40000000 udredger --from distributor-test -y
sleep 5
# Set airdrop allocations
$DREDGER_MAIN_CMD tx claim set-airdrop-allocations dredger dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr 1 --from distributor-test -y
sleep 5

# AIRDROP CLAIMS
# Check balances before claims
echo "Initial balance before claim:"
$DREDGER_MAIN_CMD query bank balances dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr
# NOTE: You can claim here using the CLI, or from the frontend!
# Claim 20% of the free tokens
echo "Claiming fee amount..."
$DREDGER_MAIN_CMD tx claim claim-free-amount --from airdrop-test --gas 400000
sleep 5
echo "Balance after claim:" 
$DREDGER_MAIN_CMD query bank balances dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr
# Stake, to claim another 20%
echo "Staking..."
$DREDGER_MAIN_CMD tx staking delegate dredgervaloper1nnurja9zt97huqvsfuartetyjx63tc5zrj5x9f 100udredger --from airdrop-test --gas 400000
sleep 5
echo "Balance after stake:" 
$DREDGER_MAIN_CMD query bank balances dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr
# Liquid stake, to claim the final 60% of tokens
echo "Liquid staking..."
$DREDGER_MAIN_CMD tx stakeibc liquid-stake 1000 uatom --from airdrop-test --gas 400000
sleep 5
echo "Balance after liquid stake:" 
$DREDGER_MAIN_CMD query bank balances dredger1nf6v2paty9m22l3ecm7dpakq2c92ueyununayr
