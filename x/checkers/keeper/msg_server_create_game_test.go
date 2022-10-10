package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Pr3Nom/checkers/testutil/keeper"
	"github.com/Pr3Nom/checkers/x/checkers"
	"github.com/Pr3Nom/checkers/x/checkers/keeper"
	"github.com/Pr3Nom/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	alice = "cosmos1pzcza54ljnklvayxl9q4g8w38x3yz9jd56n3a7"
	bob   = "cosmos1he3t9mgecp5r8vgmcsf2jsgqmlrzpuanqmlcgw"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}


func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", 
	}, *createResponse)
}

