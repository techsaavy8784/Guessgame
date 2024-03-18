package keeper

import (
	"encoding/binary"

	"game/x/game/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPlayerCount get the total number of player
func (k Keeper) GetPlayerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PlayerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPlayerCount set the total number of player
func (k Keeper) SetPlayerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PlayerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPlayer appends a player in the store with a new id and update the count
func (k Keeper) AppendPlayer(
	ctx sdk.Context,
	player types.Player,
) uint64 {
	// Create the player
	count := k.GetPlayerCount(ctx)

	// Set the ID of the appended value
	player.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKey))
	appendedValue := k.cdc.MustMarshal(&player)
	store.Set(GetPlayerIDBytes(player.Id), appendedValue)

	// Update player count
	k.SetPlayerCount(ctx, count+1)

	return count
}

// SetPlayer set a specific player in the store
func (k Keeper) SetPlayer(ctx sdk.Context, player types.Player) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKey))
	b := k.cdc.MustMarshal(&player)
	store.Set(GetPlayerIDBytes(player.Id), b)
}

// GetPlayer returns a player from its id
func (k Keeper) GetPlayer(ctx sdk.Context, id uint64) (val types.Player, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKey))
	b := store.Get(GetPlayerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlayer removes a player from the store
func (k Keeper) RemovePlayer(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKey))
	store.Delete(GetPlayerIDBytes(id))
}

// GetAllPlayer returns all player
func (k Keeper) GetAllPlayer(ctx sdk.Context) (list []types.Player) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Player
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPlayerIDBytes returns the byte representation of the ID
func GetPlayerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPlayerIDFromBytes returns ID in uint64 format from a byte array
func GetPlayerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
