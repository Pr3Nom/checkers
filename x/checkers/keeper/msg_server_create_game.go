package keeper

import (
	"context"

	"github.com/Pr3Nom/checkers/x/checkers/types"
	"github.com/Pr3Nom/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	newGame := rules.New()
	storedGame := types.StoredGame{
	    Index: newIndex,
	    Board: newGame.String(),
	    Turn:  rules.PieceStrings[newGame.Turn],
	    Black: msg.Black,
	    Red:   msg.Red,
	}

	err := storedGame.Validate()
	if err != nil {
	    return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateGameResponse{
	    GameIndex: newIndex
	}, nil
}
