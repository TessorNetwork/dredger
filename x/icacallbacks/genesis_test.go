package icacallbacks_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/TessorNetwork/dredger/testutil/keeper"
	"github.com/TessorNetwork/dredger/testutil/nullify"
	"github.com/TessorNetwork/dredger/x/icacallbacks"
	"github.com/TessorNetwork/dredger/x/icacallbacks/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		CallbackDataList: []types.CallbackData{
			{
				CallbackKey: "0",
			},
			{
				CallbackKey: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IcacallbacksKeeper(t)
	icacallbacks.InitGenesis(ctx, *k, genesisState)
	got := icacallbacks.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.CallbackDataList, got.CallbackDataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
