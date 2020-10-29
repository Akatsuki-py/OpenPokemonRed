package game

import (
	"pokered/pkg/script"
	"pokered/pkg/store"
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
	return store.ScriptID() == store.Overworld
}

func execScript() {
	script.Current()()
}
