package keeper_test

import (
	"context"
	"testing"

	keepertest "chemotronix/testutil/keeper"
	"chemotronix/x/chemotronix/keeper"
	"chemotronix/x/chemotronix/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChemotronixKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
