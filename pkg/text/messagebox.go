package text

import "github.com/hajimehoshi/ebiten"

// MessageBox text box window
type MessageBox struct {
	X, Y, Z       int
	Width, Height int
	Cache         *ebiten.Image
}

// XYZ return x, y, z
func (m *MessageBox) XYZ() (int, int, int) {
	return m.X, m.Y, m.Z
}
