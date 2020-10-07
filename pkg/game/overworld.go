package game

import (
	"pokered/pkg/joypad"
	"pokered/pkg/script"
)

func execOverworld() {
	joypad.Joypad()
	if joypad.JoyHeld.Start {
		script.ScriptID = script.WidgetStartMenu
	}
}
