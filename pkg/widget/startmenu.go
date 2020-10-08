package widget

import (
	"pokered/pkg/audio"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// DrawStartMenu draw start menu
// ref: DrawStartMenu
func DrawStartMenu() {
	audio.PlaySound(audio.SFX_START_MENU)
	height := 12
	elm := []string{
		util.Pokemon,
		"ITEM",
		"RED",
		"SAVE",
		"OPTION",
		"EXIT",
	}
	if store.CheckEvent(store.EVENT_GOT_POKEDEX) {
		height = 15
		elm = []string{
			util.Pokedex,
			util.Pokemon,
			"ITEM",
			"RED",
			"SAVE",
			"OPTION",
			"EXIT",
		}
	}
	menu.NewSelectMenu(elm, 10, 0, 8, height, true, true)
}
