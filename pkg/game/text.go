package game

import "pokered/pkg/text"

func execText() {
	if text.IsDelay() {
		return
	}
	text.PrintText()
}
