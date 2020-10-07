package game

import "pokered/pkg/menu"

func execMenu() {
	z := menu.MaxZIndex()
	if menu.Quantity.Quantity > 0 {
		menu.DisplayChooseQuantityMenu()
		return
	}
	if menu.CurListMenu.Z() == z {
		menu.DisplayListMenuIDLoop()
		return
	}
	menu.HandleSelectMenuInput()
}
