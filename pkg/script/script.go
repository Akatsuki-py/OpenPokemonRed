package script

import (
	"fmt"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/text"
)

const (
	Halt uint = iota
	ExecText
	WidgetStartMenu
	WidgetStartMenu2
	WidgetBag
)

// ScriptID current script ID
var scriptID = Halt

func ScriptID() uint      { return scriptID }
func SetScriptID(id uint) { scriptID = id }

// ScriptMap script ID -> script
var scriptMap = newScriptMap()

func newScriptMap() map[uint]func() {
	result := map[uint]func(){}
	result[Halt] = halt
	result[ExecText] = execText
	result[WidgetStartMenu] = widgetStartMenu
	result[WidgetStartMenu2] = widgetStartMenu2
	result[WidgetBag] = widgetBag
	return result
}

func Current() func() {
	s, ok := scriptMap[scriptID]
	if !ok {
		fmt.Println("this scriptID is not registered")
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
		SetScriptID(Halt)
	}
}
