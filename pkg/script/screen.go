package script

import (
	"pokered/pkg/screen"
	"pokered/pkg/store"
)

var screenCounter uint

// DoWhiteScreen make screen white for frame
func DoWhiteScreen(frame uint, doPush bool) {
	screenCounter = frame
	if doPush {
		store.PushScriptID(store.WhiteScreen)
		return
	}
	store.SetScriptID(store.WhiteScreen)
}

func whiteScreen() {
	if screenCounter == 0 {
		nextScript()
		return
	}
	screen.FillWhite()
	screenCounter--
	if screenCounter == 0 {
		nextScript()
		return
	}
}
