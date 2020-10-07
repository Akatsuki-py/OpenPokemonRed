package menu

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"sort"

	"github.com/hajimehoshi/ebiten"
)

const (
	// ChoseMenuItem Aボタンでアイテムを選択した or 2択menuの上
	ChoseMenuItem uint = iota + 1

	// CancelledMenu Bボタンでキャンセルした or 2択menuの下
	CancelledMenu
)

type Menu interface {
	Z() uint
	Top() (util.Tile, util.Tile)
	Len() int
	Wrap() bool
	Current() uint
	SetCurrent(uint)
	Image() *ebiten.Image
}

// CurMenu get current handled menu
func CurMenu() Menu {
	z := MaxZIndex()
	for _, s := range CurSelectMenus {
		if s.z == z {
			return s
		}
	}
	if CurListMenu.z == z {
		return &CurListMenu
	}
	return nil
}

var downArrowBlinkCnt = 6 * 10

// MenuExitMethod プレイヤーが menu からどのように抜けたかを記録している
var MenuExitMethod uint

// MaxZIndex get max z index
func MaxZIndex() uint {
	selectZ := uint(0)
	for _, s := range CurSelectMenus {
		if s.z > selectZ {
			selectZ = s.z
		}
	}
	if CurListMenu.z > selectZ {
		return CurListMenu.z
	}
	return selectZ
}

// HandleMenuInput メニューでのキー入力に対処するハンドラ
func HandleMenuInput(target *ebiten.Image) joypad.Input {
	PlaceCursor(target)
	store.DelayFrames = 3
	// TODO: AnimatePartyMon

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	m := CurMenu()
	return handleMenuInput(m)
}

func handleMenuInput(m Menu) joypad.Input {
	maxItem := uint(m.Len() - 1)
	switch m.(type) {
	case *ListMenu:
		if maxItem > 2 {
			maxItem = 2
		} else {
			maxItem++
		}
	}

	switch {
	case joypad.Joy5.Up:
		if m.Current() > 0 {
			m.SetCurrent(m.Current() - 1)
		} else if m.Wrap() {
			m.SetCurrent(maxItem)
		}
	case joypad.Joy5.Down:
		if m.Current() < maxItem {
			m.SetCurrent(m.Current() + 1)
		} else if m.Wrap() {
			m.SetCurrent(0)
		}
	}

	if joypad.Joy5.A || joypad.Joy5.B {
		if !util.ReadBit(store.CD60, 5) {
			audio.PlaySound(audio.SFX_PRESS_AB)
		}
	}
	return joypad.Joy5
}

func VBlank() {
	listZ, done := CurListMenu.z, false
	sort.Sort(CurSelectMenus)
	for _, m := range CurSelectMenus {
		if m.z == 0 {
			continue
		}

		if listZ > 0 && listZ < m.z {
			util.DrawImage(store.TileMap, CurListMenu.image, 0, 0)
			done = true
		}
		util.DrawImage(store.TileMap, m.Image(), 0, 0)
	}
	if !done {
		util.DrawImage(store.TileMap, CurListMenu.image, 0, 0)
	}
}
