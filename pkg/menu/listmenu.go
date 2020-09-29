package menu

import (
	"github.com/hajimehoshi/ebiten"
)

// ListMenu list menu
type ListMenu struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}
