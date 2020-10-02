package menu

import (
	"pokered/pkg/joypad"
	"pokered/pkg/util"
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
}

func CurMenu() Menu {
	z := MaxZIndex()
	for _, s := range CurSelectMenus {
		if s.z == z {
			return &s
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
// - - -
// INPUT: [wMenuWatchedKeys] = 反応する対象のキー入力 上下ボタンは必ず反応して選択オフセットを上下に移動させる
//
// OUTPUT:
// a = キー入力 [↓, ↑, ←, →, Start, Select, B, A]
// [wCurrentMenuItem] = 選択されたメニューアイテム
// [wMenuCursorLocation] = カーソルのあるタイルのアドレス
func HandleMenuInput() {
	PlaceCursor()
	// TODO: AnimatePartyMon

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return // TODO: blink
	}

	m := CurMenu()
	switch m := m.(type) {
	case *SelectMenu:
		handleSelectMenuInput(m)
	case *ListMenu:
		handleListMenuInput(m)
	}
}

func handleSelectMenuInput(m *SelectMenu) {
	switch {
	case joypad.Joy5.Up:
	case joypad.Joy5.Down:
	}
}

func handleListMenuInput(m *ListMenu) {
	switch {
	case joypad.Joy5.Up:
	case joypad.Joy5.Down:
	}
}
