package menu

import "pokered/pkg/util"

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

var CurSelectMenus = []SelectMenu{}
