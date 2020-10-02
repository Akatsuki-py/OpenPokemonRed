package game

import "pokered/pkg/menu"

func execMenu() {
	z := menu.MaxZIndex()
	if menu.CurListMenu.Z() == z {
		menu.DisplayListMenuIDLoop()
		return
	}
}
