package game

import "pokered/pkg/text"

func execText() {
	if text.InDelay() {
		return
	}
	text.PrintText()
}
