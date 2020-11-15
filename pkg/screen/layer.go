package screen

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// preset layer's z value
const (
	World  = 1
	Sprite = 2
	Widget = 100
	Mask   = 1000
)

// Layer struct
type Layer struct {
	name           string
	z              int
	image          *ebiten.Image
	pixelX, pixelY int
}

// Layers holds each screen layer
type Layers []Layer

var layers = Layers{}

// sort interface
func (ls Layers) Len() int           { return len(ls) }
func (ls Layers) Swap(i, j int)      { ls[i], ls[j] = ls[j], ls[i] }
func (ls Layers) Less(i, j int) bool { return ls[i].z < ls[j].z }

// MaxZ returns max z value in layers
func MaxZ() int {
	result := 0
	for _, l := range layers {
		if l.z > result {
			result = l.z
		}
	}
	return result
}

// IsOK returns if layer holds layer data
func (l Layer) IsOK() bool {
	return len(l.name) > 0 && l.z > 0 && l.image != nil
}

func (l Layer) Name() string {
	return l.name
}

// AddLayer adds layer into screen
func AddLayer(name string, z int, image *ebiten.Image, x, y int) bool {
	dup := false
	for _, l := range layers {
		if l.z == z {
			dup = true
			break
		}
	}

	if !dup {
		layers = append(layers, Layer{
			name:   name,
			z:      z,
			image:  image,
			pixelX: x,
			pixelY: y,
		})
		return true
	}
	return false
}

func AddLayerOnTop(name string, image *ebiten.Image, x, y int) bool {
	z := MaxZ() + 1
	return AddLayer(name, z, image, x, y)
}

func FillWhite() {
	m := newScreen()
	m.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
	AddLayer("whitemask", Mask, m, 0, 0)
}
