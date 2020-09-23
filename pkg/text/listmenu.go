package text

import "github.com/hajimehoshi/ebiten"

// ListMenu list menu
type ListMenu struct {
	X, Y, Z       int
	Width, Height int
	Cache         *ebiten.Image
}

// XYZ return x, y, z
func (l *ListMenu) XYZ() (int, int, int) {
	return l.X, l.Y, l.Z
}
