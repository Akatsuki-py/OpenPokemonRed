package menu

import (
	"pokered/pkg/text"
	"pokered/pkg/util"
)

var TopMenuItemX, TopMenuItemY util.Tile = 0, 0

func setTopMenuItem(x, y util.Tile) {
	TopMenuItemX, TopMenuItemY = x, y
}

var PreviousMenuItem, CurrentMenuItem uint = 0, 0

// Cursor cursor location in tileMap
type Cursor struct {
	X, Y util.Tile
}

// CursorLocation current cursor tile location in tileMap
var CursorLocation = Cursor{}

// PlaceCursor place "▶︎"
// 1. erase previous cursor
// 2. calc current CursorLocation
// 3. place "▶︎"
// 4. set CursorLocation
// ref: PlaceMenuCursor
func PlaceCursor() {
	text.PlaceChar(" ", TopMenuItemX, TopMenuItemY+util.Tile(PreviousMenuItem))
	X, Y := TopMenuItemX, TopMenuItemY+util.Tile(CurrentMenuItem)
	text.PlaceChar("▶︎", X, Y)
	CursorLocation = Cursor{X, Y}
	PreviousMenuItem = CurrentMenuItem
}

// PlaceUnfilledArrowCursor replace current cursor with "▷"
// ref: PlaceUnfilledArrowMenuCursor
func PlaceUnfilledArrowCursor() {
	x, y := CursorLocation.X, CursorLocation.Y
	text.PlaceChar("▷", x, y)
}

// EraseCursor erase cursor
// ref: EraseMenuCursor
func EraseCursor() {
	x, y := CursorLocation.X, CursorLocation.Y
	text.PlaceChar(" ", x, y)
}
