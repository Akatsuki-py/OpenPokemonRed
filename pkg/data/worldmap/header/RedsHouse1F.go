package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap/blk"
)

var RedsHouse1F = &Header{
	Tileset: tileset.RedsHouse,
	Height:  4,
	Width:   4,
	blk:     blk.RedsHouse1F[:],
	Text: []string{
		txt.RedsHouse1FText1,
		txt.RedsHouse1FText2,
	},
	Connections: Connections{},
}
