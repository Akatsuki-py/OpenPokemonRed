package screen

import (
	"image/color"
	"pokered/pkg/util"
	"sort"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// constants
const (
	PixelWidth  = 160
	PixelHeight = 144
	TileWidth   = 20
	TileHeight  = 18
)

var tileMap = ebiten.NewImage(8*20, 8*18)
var background = newBackground()

func newBackground() *ebiten.Image {
	bg := ebiten.NewImage(8*20, 8*18)
	bg.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
	return bg
}

func newScreen() *ebiten.Image {
	return ebiten.NewImage(8*20, 8*18)
}

func TileMap() *ebiten.Image {
	return tileMap
}

func VBlank() {
	util.DrawImage(tileMap, background, 0, 0)
	sort.Sort(layers)
	for _, layer := range layers {
		if layer.IsOK() {
			util.DrawImagePixel(tileMap, layer.image, layer.pixelX, layer.pixelY)
		}
	}
	layers = Layers{}
}
