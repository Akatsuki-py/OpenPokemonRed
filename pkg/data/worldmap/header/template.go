package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/data/worldmap/blk"
)

var template = &Header{
	Tileset: tileset.Overworld,
	Height:  9,
	Width:   10,
	blk:     blk.PalletTown[:],
	Text:    []string{},
	Connections: Connections{
		North: Connection{
			OK:        true,
			DestMapID: worldmap.ROUTE_1,
			Coords:    []uint{10, 11},
		},
	},
}
