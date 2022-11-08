package rewardshare

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kynnoenterprise/core/x/rewardshare/keeper"
	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)

	for _, rewardshare := range data.Rewardshares {
		contract := rewardshare.GetContractAddr()
		deployer := rewardshare.GetDeployerAddr()
		withdrawer := rewardshare.GetWithdrawerAddr()

		// Set initial contracts receiving transaction fees
		k.SetRewardshare(ctx, rewardshare)
		k.SetDeployerMap(ctx, deployer, contract)

		if len(withdrawer) != 0 {
			k.SetWithdrawerMap(ctx, withdrawer, contract)
		}
	}
}

// ExportGenesis export module state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:   k.GetParams(ctx),
		Rewardshares: k.GetRewardshares(ctx),
	}
}
