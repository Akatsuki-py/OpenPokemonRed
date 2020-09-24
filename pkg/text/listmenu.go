package text

import "github.com/hajimehoshi/ebiten"

// ListMenu list menu
type ListMenu struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (l *ListMenu) area() boxArea {
	return boxArea{
		X: l.X,
		Y: l.Y,
		Z: l.Z,
		W: l.W,
		H: l.H,
	}
}
