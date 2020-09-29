package menu

import (
	"github.com/hajimehoshi/ebiten"
)

// Toss how many toss items?
type Toss struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}
