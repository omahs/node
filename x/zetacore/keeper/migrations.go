package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) MigrateToVer2(ctx sdk.Context) error {
	ctx.Logger().Info("Migrating to ZetaCORE ver 2")
	fmt.Println("Migrating to ZetaCORE ver 2")
	return nil
}
