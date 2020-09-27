package menu

import (
	"pokered/pkg/textbox"

	"github.com/hajimehoshi/ebiten"
)

// Toss how many toss items?
type Toss struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (t *Toss) area() textbox.Area {
	return textbox.Area{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
		W: t.W,
		H: t.H,
	}
}
