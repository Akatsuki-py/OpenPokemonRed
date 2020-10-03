package menu

import (
	"pokered/pkg/text"
	"pokered/pkg/util"
)

var TopMenuItemX, TopMenuItemY util.Tile = 0, 0

func setTopMenuItem(x, y util.Tile) {
	TopMenuItemX, TopMenuItemY = x, y
}

// Cursor cursor location in tileMap
type Cursor struct {
	X, Y util.Tile
}

// CursorLocation current cursor tile location in tileMap
var CursorLocation = Cursor{}

// PlaceCursor set "▶︎" into current menu
// ref: PlaceMenuCursor
func PlaceCursor() {
	m := CurMenu()

	// erase old cursor
	switch m := m.(type) {
	case *SelectMenu:
		for i := 0; i < len(m.Elm); i++ {
			text.PlaceChar(" ", m.topX, m.topY+2*i)
		}
	case *ListMenu:
		for i := 0; i < 4; i++ {
			text.PlaceChar(" ", ListMenuTopX, ListMenuTopY+2*i)
		}
	}

	// place current cursor
	switch m := m.(type) {
	case *SelectMenu:
		text.PlaceChar("▶︎", m.topX, m.topY+util.Tile(2*m.current))
	case *ListMenu:
		text.PlaceChar("▶︎", ListMenuTopX, ListMenuTopY+util.Tile(2*m.current))
	}
}

// PlaceUnfilledArrowCursor replace current cursor with "▷"
// ref: PlaceUnfilledArrowMenuCursor
func PlaceUnfilledArrowCursor() {
	m := CurMenu()
	switch m := m.(type) {
	case *SelectMenu:
		text.PlaceChar("▷", m.topX, m.topY+util.Tile(2*m.current))
	case *ListMenu:
		text.PlaceChar("▷", ListMenuTopX, ListMenuTopY+util.Tile(2*m.current))
	}
}

// EraseCursor erase cursor
// ref: EraseMenuCursor
func EraseCursor() {
	m := CurMenu()
	switch m := m.(type) {
	case *SelectMenu:
		text.PlaceChar(" ", m.topX, m.topY+util.Tile(2*m.current))
	case *ListMenu:
		text.PlaceChar(" ", ListMenuTopX, ListMenuTopY+util.Tile(2*m.current))
	}
}
