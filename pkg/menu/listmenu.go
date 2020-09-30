package menu

import (
	"github.com/hajimehoshi/ebiten"
)

// ListMenu list menu
type ListMenu struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

// DisplayListMenuIDLoop list menu でユーザーの入力を待って対応する処理
func DisplayListMenuIDLoop() bool {
	selected := true
	return selected
}
