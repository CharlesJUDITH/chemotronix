package keeper_test

import (
	"testing"

	testkeeper "chemotronix/testutil/keeper"
	"chemotronix/x/chemotronix/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChemotronixKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
