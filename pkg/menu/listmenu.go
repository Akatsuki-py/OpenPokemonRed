package menu

import (
	"pokered/pkg/textbox"

	"github.com/hajimehoshi/ebiten"
)

// ListMenu list menu
type ListMenu struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (l *ListMenu) area() textbox.Area {
	return textbox.Area{
		X: l.X,
		Y: l.Y,
		Z: l.Z,
		W: l.W,
		H: l.H,
	}
}
