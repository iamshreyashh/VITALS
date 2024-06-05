package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"insurance/x/insurance/types"
)

// GetinsuranceCount get the total number of insurance
func (k Keeper) GetinsuranceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.insuranceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetinsuranceCount set the total number of insurance
func (k Keeper) SetinsuranceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.insuranceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// Appendinsurance appends a insurance in the store with a new id and update the count
func (k Keeper) Appendinsurance(
	ctx sdk.Context,
	insurance types.insurance,
) uint64 {
	// Create the insurance
	count := k.GetinsuranceCount(ctx)

	// Set the ID of the appended value
	insurance.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.insuranceKey))
	appendedValue := k.cdc.MustMarshal(&insurance)
	store.Set(GetinsuranceIDBytes(insurance.Id), appendedValue)

	// Update insurance count
	k.SetinsuranceCount(ctx, count+1)

	return count
}

// Setinsurance set a specific insurance in the store
func (k Keeper) Setinsurance(ctx sdk.Context, insurance types.insurance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.insuranceKey))
	b := k.cdc.MustMarshal(&insurance)
	store.Set(GetinsuranceIDBytes(insurance.Id), b)
}

// Getinsurance returns a insurance from its id
func (k Keeper) Getinsurance(ctx sdk.Context, id uint64) (val types.insurance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.insuranceKey))
	b := store.Get(GetinsuranceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Removeinsurance removes a insurance from the store
func (k Keeper) Removeinsurance(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.insuranceKey))
	store.Delete(GetinsuranceIDBytes(id))
}

// GetAllinsurance returns all insurance
func (k Keeper) GetAllinsurance(ctx sdk.Context) (list []types.insurance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.insuranceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.insurance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetinsuranceIDBytes returns the byte representation of the ID
func GetinsuranceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetinsuranceIDFromBytes returns ID in uint64 format from a byte array
func GetinsuranceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
