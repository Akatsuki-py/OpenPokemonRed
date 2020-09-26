package text

import (
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

func PlaceChar(target *ebiten.Image, char string, x, y int) {
	charcode, ok := charmap[char]
	if !ok {
		return
	}

	switch charcode {
	case 0x52:
		place0x52(target, x, y)
	default:
		charImage := chardata[charcode]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(util.TileToFPixel(x, y))
		target.DrawImage(charImage, op)
	}
}

func place0x52(target *ebiten.Image, x, y int) {}
