package header

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap/blk"
)

var PalletTown = &Header{
	Tileset: tileset.Overworld,
	Height:  9,
	Width:   10,
	blk:     blk.PalletTown[:],
}
