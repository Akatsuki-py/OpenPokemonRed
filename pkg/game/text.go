package game

import (
	"pokered/pkg/joypad"
	"pokered/pkg/text"
)

func execText() {
	if text.InDelay() {
		joypad.Joypad()
		if joypad.JoyHeld.A || joypad.JoyHeld.B {
			text.FrameCounter = 0
			return
		}
		text.FrameCounter--
		return
	}
	text.PrintText()
}
