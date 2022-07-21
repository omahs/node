package app

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	m "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const releaseVersion = "0.2.1"

func SetupHandlers(app *App) {
	app.UpgradeKeeper.SetUpgradeHandler(releaseVersion, func(ctx sdk.Context, plan types.Plan, vm m.VersionMap) (m.VersionMap, error) {
		app.Logger().Info("Running upgrade handler for " + releaseVersion)
		fmt.Println("Version map :", vm)
		for moduleName, _ := range vm {
			vm[moduleName] = app.mm.Modules[moduleName].ConsensusVersion()
			fmt.Println("Setting Consensus Version : ", moduleName, vm[moduleName])
		}
		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})

	//upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	//if err != nil {
	//	panic(err)
	//}
	//if upgradeInfo.Name == releaseVersion && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
	//	storeUpgrades := storetypes.StoreUpgrades{}
	//	// Use upgrade store loader for the initial loading of all stores when app starts,
	//	// it checks if version == upgradeHeight and applies store upgrades before loading the stores,
	//	// so that new stores start with the correct version (the current height of chain),
	//	// instead the default which is the latest version that store last committed i.e 0 for new stores.
	//	app.SetStoreLoader(types.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	//}
}
