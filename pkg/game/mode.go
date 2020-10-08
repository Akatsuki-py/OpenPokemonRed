package game

import (
	"pokered/pkg/joypad"
	"pokered/pkg/script"
)

const (
	Script uint = iota
	Overworld
)

func mode() uint {
	if isOverworld() {
		return Overworld
	}
	return Script
}

func isOverworld() bool {
	return script.ScriptID() == script.Halt
}

func execOverworld() {
	joypad.Joypad()
	if joypad.JoyHeld.Start {
		script.SetScriptID(script.WidgetStartMenu)
	}
}

func execScript() {
	script.Current()()
}
