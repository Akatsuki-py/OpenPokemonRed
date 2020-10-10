package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/widget"
)

func widgetStartMenu() {
	SetScriptID(WidgetStartMenu2)
	widget.DrawStartMenu()
}

func widgetStartMenu2() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()
	switch {
	case pressed.A:
		switch m.Item() {
		case "EXIT":
			m.Close()
			SetScriptID(Halt)
		case "ITEM":
			SetScriptID(WidgetBag)
			menu.NewListMenuID(menu.ItemListMenu, store.BagItems)
		case "RED":
			SetScriptID(WidgetTrainerCard)
			widget.InitTrainerCard()
		}
	case pressed.B:
		m.Close()
		SetScriptID(Halt)
	}
}

func widgetBag() {
	pressed := menu.DisplayListMenuIDLoop()
	switch {
	case pressed.A:
		switch menu.CurListMenu.Item() {
		case menu.Cancel:
			menu.CurListMenu.Close()
			SetScriptID(WidgetStartMenu2)
		}
	case pressed.B:
		menu.CurListMenu.Close()
		SetScriptID(WidgetStartMenu2)
	}
}

func widgetTrainerCard() {
	widget.DrawStartMenu()
	if joypad.ABButtonPress() {
		widget.ExitTrainerCard()
		SetScriptID(WidgetStartMenu2)
	}
}
