package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kynnoenterprise/core/x/nft/types"
)

// GetOwner gets all the ID collections owned by an address and denom ID
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress, denom string) types.Owner {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, denom, ""))
	defer iterator.Close()

	owner := types.Owner{
		Address:       address.String(),
		IDCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]string)

	for ; iterator.Valid(); iterator.Next() {
		_, denomID, tokenID, _ := types.SplitKeyOwner(iterator.Key())
		if ids, ok := idsMap[denomID]; ok {
			idsMap[denomID] = append(ids, tokenID)
		} else {
			idsMap[denomID] = []string{tokenID}
			owner.IDCollections = append(
				owner.IDCollections,
				types.IDCollection{DenomId: denomID},
			)
		}
	}

	for i := 0; i < len(owner.IDCollections); i++ {
		owner.IDCollections[i].TokenIds = idsMap[owner.IDCollections[i].DenomId]
	}

	return owner
}

// GetOwners gets all the ID collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners types.Owners) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.KeyOwner(nil, "", ""))
	defer iterator.Close()

	idcsMap := make(map[string]types.IDCollections)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, denom, id, _ := types.SplitKeyOwner(key)
		if _, ok := idcsMap[address.String()]; !ok {
			idcsMap[address.String()] = types.IDCollections{}
			owners = append(
				owners,
				types.Owner{Address: address.String()},
			)
		}
		idcs := idcsMap[address.String()]
		idcs = idcs.Add(denom, id)
		idcsMap[address.String()] = idcs
	}
	for i, owner := range owners {
		owners[i].IDCollections = idcsMap[owner.Address]
	}

	return owners
}

func (k Keeper) deleteOwner(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denomID, tokenID))
}

func (k Keeper) setOwner(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := types.MustMarshalTokenID(k.cdc, tokenID)
	store.Set(types.KeyOwner(owner, denomID, tokenID), bz)
}

func (k Keeper) swapOwner(ctx sdk.Context, denomID, tokenID string, srcOwner, dstOwner sdk.AccAddress) {

	// delete old owner key
	k.deleteOwner(ctx, denomID, tokenID, srcOwner)

	// set new owner key
	k.setOwner(ctx, denomID, tokenID, dstOwner)
}
