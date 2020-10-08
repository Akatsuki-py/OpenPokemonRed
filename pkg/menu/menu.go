package menu

import (
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

// Cancelled menu cancel
const Cancelled = "CANCELLED"

type Menu interface {
	Z() uint
	Hide()
	Top() (util.Tile, util.Tile)
	Len() int
	Wrap() bool
	Current() uint
	SetCurrent(uint)
	Item() string
	Image() *ebiten.Image
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
	if !done {
		util.DrawImage(store.TileMap, CurListMenu.image, 0, 0)
	}
	CurSelectMenus = newCurSelectMenus
}
