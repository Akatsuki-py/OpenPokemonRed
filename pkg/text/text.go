package text

import "github.com/hajimehoshi/ebiten"

type Box interface {
	ID() string
}

// MessageBox text box window
type MessageBox struct {
	X, Y, Z       int
	Width, Height int
	Cache         *ebiten.Image
}

func (m *MessageBox) ID() string {
	return "Message Box"
}

// ListMenu list menu
type ListMenu struct {
	X, Y, Z       int
	Width, Height int
	Cache         *ebiten.Image
}

func (l *ListMenu) ID() string {
	return "List Menu"
}
