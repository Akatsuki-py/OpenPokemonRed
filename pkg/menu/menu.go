package menu

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"sort"
)

const (
	// ChoseMenuItem Aボタンでアイテムを選択した or 2択menuの上
	ChoseMenuItem uint = iota + 1

	// CancelledMenu Bボタンでキャンセルした or 2択menuの下
	CancelledMenu
)

// Cancel menu cancel
const Cancel = "CANCEL"

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

func VBlank() {
	listZ, done := CurListMenu.z, false
	sort.Sort(CurSelectMenus)

	newCurSelectMenus := []*SelectMenu{}
	for _, m := range CurSelectMenus {
		if m.z == 0 {
			continue
		}

		if listZ > 0 && listZ < m.z {
			util.DrawImage(store.TileMap, CurListMenu.image, 0, 0)
			done = true
		}
		util.DrawImage(store.TileMap, m.image, 0, 0)
		newCurSelectMenus = append(newCurSelectMenus, m)
	}
	if !done && CurListMenu.z > 0 {
		util.DrawImage(store.TileMap, CurListMenu.image, 0, 0)
	}
	CurSelectMenus = newCurSelectMenus
}

func HandleMenuInput(current, maxItem uint, wrap bool) uint {
	switch {
	case joypad.Joy5.Up:
		if current > 0 {
			return current - 1
		} else if wrap {
			return maxItem
		}
	case joypad.Joy5.Down:
		if current < maxItem {
			return current + 1
		} else if wrap {
			return 0
		}
	}

	if joypad.Joy5.A || joypad.Joy5.B {
		if !util.ReadBit(store.CD60, 5) {
			audio.PlaySound(audio.SFX_PRESS_AB)
		}
	}
	return current
}

// func InMenu() bool {
// 	inSelect := CurSelectMenus.Len() > 0
// 	inListMenu := CurListMenu.z > 0
// 	return inSelect || inListMenu
// }
