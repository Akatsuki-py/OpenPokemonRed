package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap/blk"
)

var BluesHouse = &Header{
	Tileset: tileset.House,
	Height:  4,
	Width:   4,
	blk:     blk.BluesHouse[:],
	Text: []string{
		txt.BluesHouseText1,
		txt.BluesHouseText2,
		txt.BluesHouseText3,
	},
	Connections: Connections{},
}
