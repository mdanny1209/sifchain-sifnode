package keeper

import (
	"fmt"
	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the clp store
type Keeper struct {
	storeKey     sdk.StoreKey
	cdc          codec.BinaryMarshaler
	bankKeeper   types.BankKeeper
	supplyKeeper types.SupplyKeeper
}

// NewKeeper creates a dispensation keeper
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, bankkeeper types.BankKeeper, supplyKeeper types.SupplyKeeper, ps paramtypes.Subspace) Keeper {
	keeper := Keeper{
		storeKey:     key,
		cdc:          cdc,
		bankKeeper:   bankkeeper,
		supplyKeeper: supplyKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Codec() codec.BinaryMarshaler {
	return k.cdc
}

func (k Keeper) GetBankKeeper() types.BankKeeper {
	return k.bankKeeper
}

func (k Keeper) GetSupplyKeeper() types.SupplyKeeper {
	return k.supplyKeeper
}

func (k Keeper) Exists(ctx sdk.Context, key []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(key)
}

func (k Keeper) SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) error {
	return k.bankKeeper.SendCoins(ctx, from, to, coins)
}

func (k Keeper) HasCoins(ctx sdk.Context, user sdk.AccAddress, coins sdk.Coins) bool {
	return k.supplyKeeper.HasCoins(ctx, user, coins)
}
