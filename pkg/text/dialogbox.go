package text

import (
	"pokered/pkg/textbox"

	"github.com/hajimehoshi/ebiten"
)

// DialogBox text box window
type DialogBox struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (d *DialogBox) area() textbox.Area {
	return textbox.Area{
		X: d.X,
		Y: d.Y,
		Z: d.Z,
		W: d.W,
		H: d.H,
	}
}
