package menu

import (
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// Cursor cursor location in tileMap
type Cursor struct {
	X, Y util.Tile
}

// CursorLocation current cursor tile location in tileMap
var CursorLocation = Cursor{}

// EraseAllCursors erase old cursor
// ref: PlaceMenuCursor
func EraseAllCursors(target *ebiten.Image, topX, topY util.Tile, length int, space util.Tile) {
	for i := 0; i < length; i++ {
		text.PlaceChar(target, " ", topX, topY+space*i)
	}
}

// PlaceMenuCursor set "▶︎" into current menu
// ref: PlaceMenuCursor
func PlaceMenuCursor(target *ebiten.Image, topX, topY util.Tile, current int, space util.Tile) {
	text.PlaceChar(target, "▶︎", topX, topY+space*current)
}

// PlaceUnfilledArrowCursor replace current cursor with "▷"
// ref: PlaceUnfilledArrowMenuCursor
func PlaceUnfilledArrowCursor(target *ebiten.Image, topX, topY util.Tile, current int, space util.Tile) {
	text.PlaceChar(target, "▷", topX, topY+space*current)
}

// EraseCursor erase cursor
// ref: EraseMenuCursor
func EraseCursor(target *ebiten.Image, m interface{}) {
	switch m := m.(type) {
	case *SelectMenu:
		text.PlaceChar(target, " ", m.topX, m.topY+util.Tile(2*m.current))
	case *ListMenu:
		text.PlaceChar(target, " ", ListMenuTopX, ListMenuTopY+util.Tile(2*m.current))
	}
}
