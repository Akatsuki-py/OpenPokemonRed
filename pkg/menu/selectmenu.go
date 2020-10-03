package menu

import (
	"pokered/pkg/text"
	"pokered/pkg/util"
)

type SelectMenu struct {
	Elm        []string
	z          uint // zindex 0:hide
	topX, topY util.Tile
	wrap       bool
	current    uint
}

// Z return z index
func (s *SelectMenu) Z() uint {
	return s.z
}

// Top return top tiles
func (s *SelectMenu) Top() (util.Tile, util.Tile) {
	return s.topX, s.topY
}

// Len return a number of items
func (s *SelectMenu) Len() int {
	return len(s.Elm)
}

// Wrap return menu wrap is enabled
func (s *SelectMenu) Wrap() bool {
	return s.wrap
}

// Current return current selected
func (s *SelectMenu) Current() uint {
	return s.current
}

// SetCurrent set current
func (s *SelectMenu) SetCurrent(c uint) {
	s.current = c
}

// CurSelectMenus current menus
var CurSelectMenus = []SelectMenu{}

// NewSelectMenu create new select menu
func NewSelectMenu(elm []string, x0, y0, width, height util.Tile, space, wrap bool) {
	topX, topY := x0+1, y0+1
	if space {
		topY++
	}
	text.DrawTextBox(x0, y0, width+1, height+1)
	newSelectMenu := SelectMenu{
		Elm:  elm,
		z:    MaxZIndex() + 1,
		topX: topX,
		topY: topY,
		wrap: wrap,
	}
	CurSelectMenus = append(CurSelectMenus, newSelectMenu)
	for i, elm := range newSelectMenu.Elm {
		text.PlaceStringAtOnce(elm, topX+1, topY+2*i)
	}
}
