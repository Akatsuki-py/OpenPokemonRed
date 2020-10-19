package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap/blk"
)

var RedsHouse1F = &Header{
	Tileset:     tileset.RedsHouse,
	Height:      4,
	Width:       4,
	blk:         blk.RedsHouse1F[:],
	Connections: Connections{},
}
