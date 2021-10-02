package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/tokenregistry/test"
	"github.com/Sifchain/sifnode/x/tokenregistry/types"
	"github.com/stretchr/testify/assert"
)

func TestKeeper_CheckDenomPermissions(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "rowan",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t1",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	})
	// Duplicate permission is interpreted correctly.
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t2",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	})
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	entry := app.TokenRegistryKeeper.GetEntry(registry, "rowan")
	entry2 := app.TokenRegistryKeeper.GetEntry(registry, "t2")
	assert.True(t, app.TokenRegistryKeeper.CheckDenomPermissions(entry, []types.Permission{types.Permission_CLP}))
	assert.False(t, app.TokenRegistryKeeper.CheckDenomPermissions(entry, []types.Permission{types.Permission_IBCIMPORT}))
	assert.False(t, app.TokenRegistryKeeper.CheckDenomPermissions(entry, []types.Permission{types.Permission_CLP, types.Permission_IBCIMPORT}))
	assert.False(t, app.TokenRegistryKeeper.CheckDenomPermissions(entry2, []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCIMPORT}))
	assert.True(t, app.TokenRegistryKeeper.CheckDenomPermissions(entry, []types.Permission{}))
}
