package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetValidatorConsAddress(goCtx context.Context, req *types.QueryGetValidatorConsAddressRequest) (*types.QueryGetValidatorConsAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	valOperator, err := sdk.ValAddressFromBech32(req.ValidatorOperatorAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid validator operator address")
	}
	validator, found := k.StakingKeeper.GetValidator(ctx, valOperator)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "validator not found")
	}
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid validator consensus address")
	}

	return &types.QueryGetValidatorConsAddressResponse{
		ValidatorConsAddress: consAddr.String(),
	}, nil
}
