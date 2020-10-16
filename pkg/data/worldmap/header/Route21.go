package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/data/worldmap/blk"
)

// Route21_h:
// 	db OVERWORLD ; tileset
// 	db ROUTE_21_HEIGHT, ROUTE_21_WIDTH ; dimensions (y, x)
// 	dw Route21_Blocks ; blocks
// 	dw Route21_TextPointers ; texts
// 	dw Route21_Script ; scripts
// 	db NORTH | SOUTH ; connections
// 	NORTH_MAP_CONNECTION ROUTE_21, PALLET_TOWN, 0, 0, PalletTown_Blocks
// 	SOUTH_MAP_CONNECTION ROUTE_21, CINNABAR_ISLAND, 0, 0, CinnabarIsland_Blocks, 1
// 	dw Route21_Object ; objects

var Route21 = &Header{
	Tileset: tileset.Overworld,
	Height:  45,
	Width:   10,
	blk:     blk.Route21[:],
	Text:    []string{},
	Connections: Connections{
		North: Connection{
			OK:        true,
			DestMapID: worldmap.PALLET_TOWN,
			Coords:    []uint{4, 5, 6, 7},
		},
	},
}
