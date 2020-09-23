package text

import "github.com/hajimehoshi/ebiten"

type Box interface {
	ID() string
}

// MessageBox text box window
type MessageBox struct {
	Z             int
	X, Y          uint
	Width, Height int
	data          *ebiten.Image
}

func (m *MessageBox) ID() string {
	return "Message Box"
}

// ListMenu list menu
type ListMenu struct {
	Z             int
	X, Y          uint
	Width, Height int
	data          *ebiten.Image
}

func (l *ListMenu) ID() string {
	return "List Menu"
}
