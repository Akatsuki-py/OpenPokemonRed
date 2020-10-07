package widget

import (
	"pokered/pkg/audio"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/text"
)

// DrawStartMenu draw start menu
// ref: DrawStartMenu
func DrawStartMenu() {
	height := 12
	elm := []string{
		"POKéMON",
		"ITEM",
		"RED",
		"SAVE",
		"OPTION",
		"EXIT",
	}
	if store.CheckEvent(store.EVENT_GOT_POKEDEX) {
		height = 15
		elm = []string{
			"POKéDEX",
			"POKéMON",
			"ITEM",
			"RED",
			"SAVE",
			"OPTION",
			"EXIT",
		}
	}
	text.DrawTextBoxWH(10, 0, 8, height)

	menu.NewSelectMenu(elm, 10, 0, 8, height, true, true)
}

func DisplayStartMenu() {
	audio.PlaySound(audio.SFX_START_MENU)
	DrawStartMenu()
}
