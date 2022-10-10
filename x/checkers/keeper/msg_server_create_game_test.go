package keeper_test

import (
	"testing"

	"github.com/Pr3Nom/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1pzcza54ljnklvayxl9q4g8w38x3yz9jd56n3a7"
	bob   = "cosmos1he3t9mgecp5r8vgmcsf2jsgqmlrzpuanqmlcgw"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
