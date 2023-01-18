package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/TessorNetwork/dredger/testutil/keeper"
	"github.com/TessorNetwork/dredger/testutil/nullify"
	"github.com/TessorNetwork/dredger/x/stakeibc/keeper"
	"github.com/TessorNetwork/dredger/x/stakeibc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEpochTracker(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EpochTracker {
	items := make([]types.EpochTracker, n)
	for i := range items {
		items[i].EpochIdentifier = strconv.Itoa(i)
		keeper.SetEpochTracker(ctx, items[i])
	}
	return items
}

func TestEpochTrackerGet(t *testing.T) {
	keeper, ctx := keepertest.StakeibcKeeper(t)
	items := createNEpochTracker(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEpochTracker(ctx,
			item.EpochIdentifier,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestEpochTrackerRemove(t *testing.T) {
	keeper, ctx := keepertest.StakeibcKeeper(t)
	items := createNEpochTracker(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEpochTracker(ctx,
			item.EpochIdentifier,
		)
		_, found := keeper.GetEpochTracker(ctx,
			item.EpochIdentifier,
		)
		require.False(t, found)
	}
}

func TestEpochTrackerGetAll(t *testing.T) {
	keeper, ctx := keepertest.StakeibcKeeper(t)
	items := createNEpochTracker(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEpochTracker(ctx)),
	)
}
