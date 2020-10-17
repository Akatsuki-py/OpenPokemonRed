package world

import (
	"pokered/pkg/data/blockset"
	"pokered/pkg/data/tileset"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// Blockset cur map blockset
type Blockset struct {
	TilesetID uint
	Bytes     [][16]byte
	Data      []*ebiten.Image
}

// CurBlockset current block & tileset data
var CurBlockset Blockset

// loadBlockset load block set
func loadBlockset(tilesetID uint) {
	if CurBlockset.TilesetID == tilesetID {
		return
	}

	bs := blockset.Get(tilesetID)
	length := len(bs) / 16
	result := make([]*ebiten.Image, length)
	tiles := make([][16]byte, length)

	for i := 0; i < length; i++ {
		b := [16]byte{}
		block, _ := ebiten.NewImage(8*4, 8*4, ebiten.FilterDefault)

		for j := 0; j < 16; j++ {
			tileID := bs[i*16+j]
			b[j] = tileID
			tile := tileset.Tile(tilesetID, uint(tileID))
			x, y := 8*(j%4), 8*(j/4)
			util.DrawImagePixel(block, tile, x, y)
		}

		tiles[i] = b
		result[i] = block
	}

	CurBlockset = Blockset{
		TilesetID: tilesetID,
		Bytes:     tiles,
		Data:      result,
	}
}
