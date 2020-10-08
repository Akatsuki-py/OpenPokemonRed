package script

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/widget"
)

const (
	Halt uint = iota
	WidgetStartMenu
	WidgetStartMenu2
	ExecText
)

// ScriptID current script ID
var ScriptID = Halt

// ScriptMap script ID -> script
var scriptMap = newScriptMap()

func newScriptMap() map[uint]func() {
	result := map[uint]func(){}
	result[Halt] = halt
	result[WidgetStartMenu] = widgetStartMenu
	result[WidgetStartMenu2] = widgetStartMenu2
	result[ExecText] = execText
	return result
}

func Current() func() {
	s, ok := scriptMap[ScriptID]
	if !ok {
		return halt
	}
	return s
}

func halt() {}

func execText() {
	if text.InScroll {
		text.ScrollTextUpOneLine(text.Image)
		return
	}
	if store.FrameCounter > 0 {
		joypad.Joypad()
		if joypad.JoyHeld.A || joypad.JoyHeld.B {
			store.FrameCounter = 0
			return
		}
		store.FrameCounter--
		if store.FrameCounter > 0 {
			store.DelayFrames = 1
			return
		}
		return
	}
	text.CurText = text.PlaceStringOneByOne(text.Image, text.CurText)
	if len([]rune(text.CurText)) == 0 {
		ScriptID = Halt
	}
}

func widgetStartMenu() {
	ScriptID = WidgetStartMenu2
	widget.DisplayStartMenu()
}

func widgetStartMenu2() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()
	switch {
	case pressed.A:
		switch m.Item() {
		case "EXIT":
			m.Hide()
			ScriptID = Halt
		case "RED":
			ScriptID = ExecText
			text.PrintText(text.Image, txt.AgathaBeforeBattleText)
		}
	case pressed.B:
		m.Hide()
		ScriptID = Halt
	}
}
