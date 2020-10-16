package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/data/worldmap/blk"
)

// Route1_h:
// 	db OVERWORLD ; tileset
// 	db ROUTE_1_HEIGHT, ROUTE_1_WIDTH ; dimensions (y, x)
// 	dw Route1_Blocks ; blocks
// 	dw Route1_TextPointers ; texts
// 	dw Route1_Script ; scripts
// 	db NORTH | SOUTH ; connections
// 	NORTH_MAP_CONNECTION ROUTE_1, VIRIDIAN_CITY, -3, 2, ViridianCity_Blocks
// 	SOUTH_MAP_CONNECTION ROUTE_1, PALLET_TOWN, 0, 0, PalletTown_Blocks, 1
// 	dw Route1_Object ; objects

var Route1 = &Header{
	Tileset: tileset.Overworld,
	Height:  18,
	Width:   10,
	blk:     blk.Route1[:],
	Text: []string{
		txt.Route1Text1,
		txt.Route1Text2,
		txt.Route1Text3,
	},
	Connections: Connections{
		South: Connection{
			OK:        true,
			DestMapID: worldmap.PALLET_TOWN,
			Coords:    []uint{10, 11},
		},
	},
}
