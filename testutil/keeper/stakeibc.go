package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	strideapp "github.com/TessorNetwork/dredger/app"
	"github.com/TessorNetwork/dredger/x/stakeibc/keeper"
)

func StakeibcKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := strideapp.InitStrideTestApp(true)
	stakeibcKeeper := app.StakeibcKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "dredger-1", Time: time.Now().UTC()})

	return &stakeibcKeeper, ctx
}
