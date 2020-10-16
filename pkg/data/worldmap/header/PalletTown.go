package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap/blk"
)

// PalletTown_h:
// 	db OVERWORLD ; tileset
// 	db PALLET_TOWN_HEIGHT, PALLET_TOWN_WIDTH ; dimensions (y, x)
// 	dw PalletTown_Blocks ; blocks
// 	dw PalletTown_TextPointers ; texts
// 	dw PalletTown_Script ; scripts
// 	db NORTH | SOUTH ; connections
// 	NORTH_MAP_CONNECTION PALLET_TOWN, ROUTE_1, 0, 0, Route1_Blocks
// 	SOUTH_MAP_CONNECTION PALLET_TOWN, ROUTE_21, 0, 0, Route21_Blocks, 1
// 	dw PalletTown_Object ; objects

var PalletTown = &Header{
	Tileset: tileset.Overworld,
	Height:  9,
	Width:   10,
	blk:     blk.PalletTown[:],
}
