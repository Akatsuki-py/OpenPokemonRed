package menu

import (
	"github.com/hajimehoshi/ebiten"
)

// ListMenu list menu
type ListMenu struct {
	X, Y, Z, W, H int
	Cache         *ebiten.Image
}

// DisplayListMenuID list menu でユーザーの入力を待って対応する処理
func DisplayListMenuID() bool {
	selected := true
	return selected
}
