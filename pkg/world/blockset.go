package world

import (
	"pokered/pkg/data/blockset"
	"pokered/pkg/data/tileset"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// Blockset cur map blockset
var Blockset = []*ebiten.Image{}

// LoadBlockset load block set
func LoadBlockset(tilesetID uint) {
	bs := blockset.Get(tilesetID)
	length := len(bs) / 16
	result := make([]*ebiten.Image, length)
	for i := 0; i < length; i++ {
		block, _ := ebiten.NewImage(8*4, 8*4, ebiten.FilterDefault)
		for j := 0; j < 16; j++ {
			tileID := bs[i*16+j]
			tile := tileset.Tile(tilesetID, uint(tileID))
			x, y := 8*(j%4), 8*(j/4)
			util.DrawImagePixel(block, tile, x, y)
		}
		result[i] = block
	}
	Blockset = result
}
