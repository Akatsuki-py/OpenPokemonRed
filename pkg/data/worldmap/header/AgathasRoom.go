package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap/blk"
)

var AgathasRoom = &Header{
	Tileset: tileset.Cemetery,
	Height:  6,
	Width:   5,
	blk:     blk.AgathasRoom[:],
}
