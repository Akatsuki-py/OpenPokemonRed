package script

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/menu"
	"pokered/pkg/text"
	"pokered/pkg/widget"
)

func widgetStartMenu() {
	SetScriptID(WidgetStartMenu2)
	widget.DisplayStartMenu()
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
		case "RED":
			SetScriptID(ExecText)
			text.PrintText(text.Image, txt.AgathaBeforeBattleText)
		}
	case pressed.B:
		m.Close()
		SetScriptID(Halt)
	}
}
