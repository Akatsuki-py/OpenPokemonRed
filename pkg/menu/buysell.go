package menu

import (
	"github.com/hajimehoshi/ebiten"
)

// BuySell how many buy/sell items?
type BuySell struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}
