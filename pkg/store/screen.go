package store

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// TileMap c3a0
var TileMap, _ = ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)

// FilteredTileMap get palette filtered TileMap
func FilteredTileMap() *ebiten.Image {
	return defaultFilter(TileMap)
}

func defaultFilter(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)

	sheet, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{95, 125, 100, 0xff})
	result.DrawImage(sheet, nil)

	src0, _ := ebiten.NewImageFromImage(target, ebiten.FilterDefault)
	op0 := &ebiten.DrawImageOptions{}
	op0.ColorM.Scale(0.65, 0.85, 0.65, 0.95)
	result.DrawImage(src0, op0)

	return result
}
