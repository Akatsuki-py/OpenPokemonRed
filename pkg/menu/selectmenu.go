package menu

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"
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

// CurSelectMenu get current handled select menu
func CurSelectMenu() *SelectMenu {
	z := MaxZIndex()
	for _, s := range CurSelectMenus {
		if s.z == z {
			return s
		}
	}
	return nil
}

func (s *SelectMenu) Close() {
	s.z = 0
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

// HandleSelectMenuInput メニューでのキー入力に対処するハンドラ
func HandleSelectMenuInput() joypad.Input {
	s := CurSelectMenu()
	PlaceCursor(s.image, s)
	store.DelayFrames = 3

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	maxItem := uint(len(s.Elm) - 1)
	s.current = HandleMenuInput(s.current, maxItem, s.wrap)
	return joypad.Joy5
}
