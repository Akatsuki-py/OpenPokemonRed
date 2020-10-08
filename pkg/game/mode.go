package game

import (
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

func execScript() {
	script.Current()()
}
