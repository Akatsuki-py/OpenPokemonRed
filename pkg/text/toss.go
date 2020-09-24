package text

import "github.com/hajimehoshi/ebiten"

// Toss how many toss items?
type Toss struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (t *Toss) area() boxArea {
	return boxArea{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
		W: t.W,
		H: t.H,
	}
}
