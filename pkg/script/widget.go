package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/widget"
)

func widgetStartMenu() {
	store.SetScriptID(store.WidgetStartMenu2)
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
			store.SetScriptID(store.Overworld)
		case "ITEM":
			store.SetScriptID(store.WidgetBag)
			menu.NewListMenuID(menu.ItemListMenu, store.BagItems)
		case "RED":
			m.Close()
			store.SetScriptID(store.WidgetTrainerCard)
			widget.DrawTrainerCard()
		case util.Pokemon:
			if store.PartyMonLen() > 0 {
				m.Close()
				store.SetScriptID(store.WidgetPartyMenu)
				widget.DrawPartyMenu()
			}
		}
	case pressed.B:
		m.Close()
		store.SetScriptID(store.Overworld)
	}
}

func widgetBag() {
	pressed := menu.DisplayListMenuIDLoop()
	switch {
	case pressed.A:
		switch menu.CurListMenu.Item() {
		case menu.Cancel:
			menu.CurListMenu.Close()
			store.SetScriptID(store.WidgetStartMenu2)
		}
	case pressed.B:
		menu.CurListMenu.Close()
		store.SetScriptID(store.WidgetStartMenu2)
	}
}

func widgetTrainerCard() {
	if joypad.ABButtonPress() {
		widget.CloseTrainerCard()
		store.SetScriptID(store.WidgetStartMenu)
	}
}

func widgetNamingScreen() {
	widget.UpdateNameScreen()

	joypad.JoypadLowSensitivity()
	switch {
	case joypad.Joy5.Up:
		widget.SetNameCursor(0, -1)
	case joypad.Joy5.Down:
		widget.SetNameCursor(0, 1)
	case joypad.Joy5.Left:
		widget.SetNameCursor(-1, 0)
	case joypad.Joy5.Right:
		widget.SetNameCursor(1, 0)
	case joypad.Joy5.Select:
		widget.ToggleCase()
	case joypad.Joy5.A:
		widget.NextChar()
	case joypad.Joy5.B:
		widget.EraseChar()
	}
}

func widgetPartyMenu() {
	pressed := widget.HandlePartyMenuInput()
	switch {
	case pressed.B:
		widget.ClosePartyMenu()
		store.SetScriptID(store.WidgetStartMenu)
	}
}
