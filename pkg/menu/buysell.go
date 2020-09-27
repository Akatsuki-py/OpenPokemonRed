package menu

import (
	"pokered/pkg/textbox"

	"github.com/hajimehoshi/ebiten"
)

// BuySell how many buy/sell items?
type BuySell struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (b *BuySell) area() textbox.Area {
	return textbox.Area{
		X: b.X,
		Y: b.Y,
		Z: b.Z,
		W: b.W,
		H: b.H,
	}
}
