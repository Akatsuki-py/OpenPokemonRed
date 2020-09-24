package text

import "github.com/hajimehoshi/ebiten"

// MessageBox text box window
type MessageBox struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

func (m *MessageBox) area() boxArea {
	return boxArea{
		X: m.X,
		Y: m.Y,
		Z: m.Z,
		W: m.W,
		H: m.H,
	}
}
