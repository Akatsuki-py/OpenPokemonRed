package text

import "github.com/hajimehoshi/ebiten"

// BuySell how many buy/sell items?
type BuySell struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (b *BuySell) area() boxArea {
	return boxArea{
		X: b.X,
		Y: b.Y,
		Z: b.Z,
		W: b.W,
		H: b.H,
	}
}
