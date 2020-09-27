package text

import (
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

func PlaceChar(target *ebiten.Image, char string) {
	charcode, ok := charmap[char]
	if !ok {
		return
	}

	switch charcode {
	case 0x52:
		place0x52(target)
	default:
		placeChar(target, charcode)
	}
}

func placeChar(target *ebiten.Image, charcode CharCode) {
	font := fontmap[charcode]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(util.TileToFPixel(Caret()))
	target.DrawImage(font, op)
	Next()
}

func place0x52(target *ebiten.Image) {
	name := store.PlayerName
	finishDTE(target, name)
}

func finishDTE(target *ebiten.Image, str string) {
	for _, char := range str {
		PlaceChar(target, string(char))
	}
}
