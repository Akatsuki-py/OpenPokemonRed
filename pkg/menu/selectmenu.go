package menu

import "pokered/pkg/util"

type SelectMenu struct {
	Elm        []string
	z          uint // zindex 0:hide
	topX, topY util.Tile
	Current    uint
}

// Z return z index
func (s *SelectMenu) Z() uint {
	return s.z
}

// Top return top tiles
func (s *SelectMenu) Top() (util.Tile, util.Tile) {
	return s.topX, s.topY
}

var CurSelectMenus = []SelectMenu{}
