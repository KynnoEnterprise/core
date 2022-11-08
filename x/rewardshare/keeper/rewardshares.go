package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

// GetRewardshares returns all registered rewardshares.
func (k Keeper) GetRewardshares(ctx sdk.Context) []types.Rewardshare {
	rewardshares := []types.Rewardshare{}

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixRewardshare)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var rewardshare types.Rewardshare
		k.cdc.MustUnmarshal(iterator.Value(), &rewardshare)

		rewardshares = append(rewardshares, rewardshare)
	}

	return rewardshares
}

// IterateRewardshares iterates over all registered contracts and performs a
// callback with the corresponding Rewardshare.
func (k Keeper) IterateRewardshares(
	ctx sdk.Context,
	handlerFn func(fee types.Rewardshare) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixRewardshare)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var rewardshare types.Rewardshare
		k.cdc.MustUnmarshal(iterator.Value(), &rewardshare)

		if handlerFn(rewardshare) {
			break
		}
	}
}

// GetRewardshare returns the Rewardshare for a registered contract
func (k Keeper) GetRewardshare(
	ctx sdk.Context,
	contract common.Address,
) (types.Rewardshare, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRewardshare)
	bz := store.Get(contract.Bytes())
	if len(bz) == 0 {
		return types.Rewardshare{}, false
	}

	var rewardshare types.Rewardshare
	k.cdc.MustUnmarshal(bz, &rewardshare)
	return rewardshare, true
}

// SetRewardshare stores the Rewardshare for a registered contract.
func (k Keeper) SetRewardshare(ctx sdk.Context, rewardshare types.Rewardshare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRewardshare)
	key := rewardshare.GetContractAddr()
	bz := k.cdc.MustMarshal(&rewardshare)
	store.Set(key.Bytes(), bz)
}

// DeleteRewardshare deletes a Rewardshare of a registered contract.
func (k Keeper) DeleteRewardshare(ctx sdk.Context, fee types.Rewardshare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRewardshare)
	key := fee.GetContractAddr()
	store.Delete(key.Bytes())
}

// SetDeployerMap stores a contract-by-deployer mapping
func (k Keeper) SetDeployerMap(
	ctx sdk.Context,
	deployer sdk.AccAddress,
	contract common.Address,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDeployer)
	key := append(deployer.Bytes(), contract.Bytes()...)
	store.Set(key, []byte{1})
}

// DeleteDeployerMap deletes a contract-by-deployer mapping
func (k Keeper) DeleteDeployerMap(
	ctx sdk.Context,
	deployer sdk.AccAddress,
	contract common.Address,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDeployer)
	key := append(deployer.Bytes(), contract.Bytes()...)
	store.Delete(key)
}

// SetWithdrawerMap stores a contract-by-withdrawer mapping
func (k Keeper) SetWithdrawerMap(
	ctx sdk.Context,
	withdrawer sdk.AccAddress,
	contract common.Address,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWithdrawer)
	key := append(withdrawer.Bytes(), contract.Bytes()...)
	store.Set(key, []byte{1})
}

// DeleteWithdrawMap deletes a contract-by-withdrawer mapping
func (k Keeper) DeleteWithdrawerMap(
	ctx sdk.Context,
	withdrawer sdk.AccAddress,
	contract common.Address,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWithdrawer)
	key := append(withdrawer.Bytes(), contract.Bytes()...)
	store.Delete(key)
}

// IsRewardshareRegistered checks if a contract was registered for receiving
// transaction fees
func (k Keeper) IsRewardshareRegistered(
	ctx sdk.Context,
	contract common.Address,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRewardshare)
	return store.Has(contract.Bytes())
}

// IsDeployerMapSet checks if a given contract-by-withdrawer mapping is set in
// store
func (k Keeper) IsDeployerMapSet(
	ctx sdk.Context,
	deployer sdk.AccAddress,
	contract common.Address,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDeployer)
	key := append(deployer.Bytes(), contract.Bytes()...)
	return store.Has(key)
}

// IsWithdrawerMapSet checks if a giveb contract-by-withdrawer mapping is set in
// store
func (k Keeper) IsWithdrawerMapSet(
	ctx sdk.Context,
	withdrawer sdk.AccAddress,
	contract common.Address,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWithdrawer)
	key := append(withdrawer.Bytes(), contract.Bytes()...)
	return store.Has(key)
}
