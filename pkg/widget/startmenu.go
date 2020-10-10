package widget

import (
	"pokered/pkg/event"
	"pokered/pkg/menu"
	"pokered/pkg/util"
)

// DrawStartMenu draw start menu
// ref: DrawStartMenu
func DrawStartMenu() {
	height := 12
	elm := []string{
		util.Pokemon,
		"ITEM",
		"RED",
		"SAVE",
		"OPTION",
		"EXIT",
	}
	if event.CheckEvent(event.EVENT_GOT_POKEDEX) {
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
