package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/screen"
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
		case store.Player.Name:
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

func handleNamingScreen() (string, bool) {
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
	case joypad.Joy5.Start:
		name := widget.CloseNameScreen()
		return name, true
	}
	return "", false
}

func widgetPartyMenu() {
	pressed := widget.HandlePartyMenuInput()
	widget.AnimatePartyMon()

	switch {
	case pressed.A:
		store.SetScriptID(store.WidgetPartyMenuSelect)
		width, height := 7, 5
		elm := []string{
			"STATS",
			"SWITCH",
			menu.Cancel,
		}
		menu.NewSelectMenu(elm, 11, 11, width, height, false, false, screen.Widget)
	case pressed.B:
		widget.ClosePartyMenu()
		store.SetScriptID(store.WidgetStartMenu)
	}
}

func widgetPartyMenuSelect() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()
	switch {
	case pressed.A:
		switch m.Item() {
		case "STATS":
		case "SWITCH":
		case menu.Cancel:
			m.Close()
			store.SetScriptID(store.WidgetPartyMenu)
		}
	case pressed.B:
		m.Close()
		store.SetScriptID(store.WidgetPartyMenu)
	}
}
