package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/widget"
)

func widgetStartMenu() {
	SetID(WidgetStartMenu2)
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
			SetID(Halt)
		case "ITEM":
			SetID(WidgetBag)
			menu.NewListMenuID(menu.ItemListMenu, store.BagItems)
		case "RED":
			m.Close()
			SetID(WidgetTrainerCard)
			widget.DrawTrainerCard()
		case "SAVE":
			m.Close()
			SetID(WidgetNamingScreen)
			widget.DrawNameScreen(widget.PlayerName)
		}
	case pressed.B:
		m.Close()
		SetID(Halt)
	}
}

func widgetBag() {
	pressed := menu.DisplayListMenuIDLoop()
	switch {
	case pressed.A:
		switch menu.CurListMenu.Item() {
		case menu.Cancel:
			menu.CurListMenu.Close()
			SetID(WidgetStartMenu2)
		}
	case pressed.B:
		menu.CurListMenu.Close()
		SetID(WidgetStartMenu2)
	}
}

func widgetTrainerCard() {
	if joypad.ABButtonPress() {
		widget.CloseTrainerCard()
		SetID(WidgetStartMenu)
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
