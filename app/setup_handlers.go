package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	m "github.com/cosmos/cosmos-sdk/types/module"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	types3 "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	"strings"
)

const releaseVersion = "0.11.0"

var (
	rc1 = GetUpgradeName(releaseVersion, "rc1")
)

func SetupHandlers(app *SifchainApp) {
	app.UpgradeKeeper.SetUpgradeHandler(rc1, func(ctx sdk.Context, plan types.Plan, vm m.VersionMap) (m.VersionMap, error) {
		app.Logger().Info("Running upgrade handler for " + rc1)
		app.IBCKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
		minter := types3.Minter{
			Inflation:        sdk.MustNewDecFromStr("0.417800000000000000"),
			AnnualProvisions: sdk.ZeroDec(),
		}
		app.MintKeeper.SetMinter(ctx, minter)
		params := app.MintKeeper.GetParams(ctx)
		params.InflationMax = sdk.MustNewDecFromStr("0.500000000000000000")
		params.InflationMin = sdk.MustNewDecFromStr("0.370000000000000000")
		app.MintKeeper.SetParams(ctx, params)
		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}
	if upgradeInfo.Name == releaseVersion && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{feegrant.ModuleName, crisistypes.ModuleName},
		}
		// Use upgrade store loader for the initial loading of all stores when app starts,
		// it checks if version == upgradeHeight and applies store upgrades before loading the stores,
		// so that new stores start with the correct version (the current height of chain),
		// instead the default which is the latest version that store last committed i.e 0 for new stores.
		app.SetStoreLoader(types.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func GetUpgradeName(rv, rc string) string {
	return strings.Join([]string{rv, rc}, "-")
}
