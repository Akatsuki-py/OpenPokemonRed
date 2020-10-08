package menu

import (
	"pokered/pkg/audio"
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

func (s *SelectMenu) Hide() {
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
	m := CurSelectMenu()
	PlaceCursor(m.image, m)
	store.DelayFrames = 3
	// TODO: AnimatePartyMon

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	return handleSelectMenuInput(m)
}

func handleSelectMenuInput(s *SelectMenu) joypad.Input {
	maxItem := uint(len(s.Elm) - 1)

	switch {
	case joypad.Joy5.Up:
		if s.current > 0 {
			s.current--
		} else if s.wrap {
			s.current = maxItem
		}
	case joypad.Joy5.Down:
		if s.current < maxItem {
			s.current++
		} else if s.wrap {
			s.current = 0
		}
	}

	if joypad.Joy5.A || joypad.Joy5.B {
		if !util.ReadBit(store.CD60, 5) {
			audio.PlaySound(audio.SFX_PRESS_AB)
		}
	}
	return joypad.Joy5
}
