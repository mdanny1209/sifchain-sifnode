package keeper

import (
	"context"
	"fmt"

	tokenregistrytypes "github.com/Sifchain/sifnode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	bankKeeper          bankkeeper.Keeper
	tokenRegistryKeeper tokenregistrytypes.Keeper
	sdkMsgServer        types.MsgServer
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(sdkMsgServer types.MsgServer, bankKeeper bankkeeper.Keeper, tokenRegistryKeeper tokenregistrytypes.Keeper) types.MsgServer {
	return &msgServer{
		sdkMsgServer:        sdkMsgServer,
		bankKeeper:          bankKeeper,
		tokenRegistryKeeper: tokenRegistryKeeper,
	}
}

// Transfer defines a rpc handler method for MsgTransfer.
func (srv msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get token registry entry for sent token
	registryEntry := srv.tokenRegistryKeeper.GetDenom(ctx, msg.Token.Denom)
	// check if registry entry has an IBC decimal field
	if registryEntry.IbcDenom != "" && registryEntry.Decimals > registryEntry.IbcDecimals {
		sender, err := sdk.AccAddressFromBech32(msg.Sender)
		if err != nil {
			return nil, err
		}
		// calculate the conversion difference and reduce precision
		po := registryEntry.Decimals - registryEntry.IbcDecimals
		decAmount := sdk.NewDecFromInt(msg.Token.Amount)
		convAmountDec := ReducePrecision(decAmount, po)

		convAmount := sdk.NewIntFromBigInt(convAmountDec.TruncateInt().BigInt())
		// create converted and sifchain tokens with corresponding denoms and amounts
		convToken := sdk.NewCoin(registryEntry.IbcDenom, convAmount)
		// increase convAmount precision to ensure amount deducted from address is the same that gets sent
		tokenAmountDec := IncreasePrecision(convAmountDec, po)
		tokenAmount := sdk.NewIntFromBigInt(tokenAmountDec.TruncateInt().BigInt())
		token := sdk.NewCoin(msg.Token.Denom, tokenAmount)
		// send coins from account to module
		err = srv.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(token))
		if err != nil {
			return nil, err
		}
		// mint ibcdenom coins
		err = srv.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(convToken))
		if err != nil {
			return nil, err
		}
		// send coins from module account to address
		err = srv.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(convToken))
		if err != nil {
			return nil, err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				tokenregistrytypes.EventTypeCovertTransfer,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(tokenregistrytypes.AttributeKeySentAmount, fmt.Sprintf("%d", token.Amount)),
				sdk.NewAttribute(tokenregistrytypes.AttributeKeySentDenom, token.Denom),
				sdk.NewAttribute(tokenregistrytypes.AttributeKeyConvertAmount, fmt.Sprintf("%d", convToken.Amount)),
				sdk.NewAttribute(tokenregistrytypes.AttributeKeyConvertDenom, convToken.Denom),
			),
		)

		convMsg := types.NewMsgTransfer(
			msg.SourcePort, msg.SourceChannel, convToken, sender, msg.Receiver, msg.TimeoutHeight, msg.TimeoutTimestamp,
		)
		convCtx := sdk.WrapSDKContext(ctx)
		return srv.sdkMsgServer.Transfer(convCtx, convMsg)
	}
	return srv.sdkMsgServer.Transfer(goCtx, msg)
}

func IncreasePrecision(dec sdk.Dec, po int64) sdk.Dec {
	p := sdk.NewDec(10).Power(uint64(po))
	return dec.Mul(p)
}

func ReducePrecision(dec sdk.Dec, po int64) sdk.Dec {
	p := sdk.NewDec(10).Power(uint64(po))
	return dec.Quo(p)
}
