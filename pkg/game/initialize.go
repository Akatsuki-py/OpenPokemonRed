package game

import (
	"pokered/pkg/script"
	"pokered/pkg/store"
)

func initialize() {
	store.SetScriptID(store.TitleCopyright)
}

func initializeWorld() {
	script.InitializeOverworld()
}
