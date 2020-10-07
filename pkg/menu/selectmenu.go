package menu

import (
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

type SelectMenu struct {
	Elm        []string
	z          uint // zindex 0:hide
	topX, topY util.Tile
	wrap       bool
	current    uint
	image      *ebiten.Image
}

// Z return z index
func (s *SelectMenu) Z() uint {
	return s.z
}

func (s *SelectMenu) Hide() {
	s.z = 0
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

func (s *SelectMenu) Image() *ebiten.Image {
	return s.image
}

func (s *SelectMenu) Item() string {
	if s.current >= uint(len(s.Elm)) {
		return ""
	}
	return s.Elm[s.current]
}

// SetCurrent set current
func (s *SelectMenu) SetCurrent(c uint) {
	s.current = c
}

type SelectMenus []*SelectMenu

// CurSelectMenus current menus
var CurSelectMenus = SelectMenus{}

// sort interface
func (sm SelectMenus) Len() int           { return len(sm) }
func (sm SelectMenus) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm SelectMenus) Less(i, j int) bool { return sm[i].z < sm[j].z }

// NewSelectMenu create new select menu
func NewSelectMenu(elm []string, x0, y0, width, height util.Tile, space, wrap bool) {
	topX, topY := x0+1, y0+1
	if space {
		topY++
	}
	newSelectMenu := &SelectMenu{
		Elm:   elm,
		z:     MaxZIndex() + 1,
		topX:  topX,
		topY:  topY,
		wrap:  wrap,
		image: util.NewImage(),
	}
	text.DrawTextBoxWH(newSelectMenu.image, x0, y0, width, height)
	CurSelectMenus = append(CurSelectMenus, newSelectMenu)
	for i, elm := range newSelectMenu.Elm {
		text.PlaceStringAtOnce(newSelectMenu.image, elm, topX+1, topY+2*i)
	}
}
