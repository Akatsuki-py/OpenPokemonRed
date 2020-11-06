package widget

import (
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

var partyMenu *ebiten.Image

// 0: no swap 1: first selected offset for swap(starting from 1)
var partyMenuSwapID uint
var partyMenuCurrent uint

// DrawPartyMenu draw party menu
// this func is always used when party menu is needed.
// e.g. Pokemon, item target, ...
// ref: RedrawPartyMenu_
func DrawPartyMenu() {
	partyMenu = util.NewImage()
	util.WhiteScreen(partyMenu)
	length := store.PartyMonLen()
	for i := 0; i < length; i++ {
		if i >= 6 {
			break
		}
		drawPartyPokemon(i)
	}
}

func drawPartyPokemon(offset int) {
	y := offset * 2
	mon := store.PartyMons[offset]
	if !mon.Initialized {
		return
	}

	text.PlaceStringAtOnce(partyMenu, mon.Nick, 3, y)

	if partyMenuSwapID > 0 {
		drawWhitePartyCursor()
	}

	// status condition
	hp, status := mon.HP, mon.Status
	printStatusCondition(offset, hp, status)

	// hp
	DrawHP(partyMenu, hp, mon.MaxHP, 4, y+1, true)

	// ABLE or NOT ABLE

	// level
	PrintLevel(partyMenu, mon.Level, 13, y)
}

func drawPartyCursor()      {}
func drawWhitePartyCursor() {}

func printStatusCondition(offset int, hp uint, status store.NonVolatileStatus) {
	x, y := 17, offset*2
	if hp == 0 {
		text.PlaceStringAtOnce(partyMenu, "FNT", x, y)
		return
	}

	if status != store.OK {
		text.PlaceStringAtOnce(partyMenu, status.String(), x, y)
	}
}

// HandlePartyMenuInput handle input on party menu
// ref: HandlePartyMenuInput
func HandlePartyMenuInput() joypad.Input {
	length := store.PartyMonLen()
	menu.EraseAllCursors(partyMenu, 0, 1, length, 2)
	menu.PlaceMenuCursor(partyMenu, 0, 1, int(partyMenuCurrent), 2)
	store.DelayFrames = 3
	// TODO: AnimatePartyMon

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	partyMenuCurrent = menu.HandleMenuInput(partyMenuCurrent, uint(store.PartyMonLen()), true)
	return joypad.Joy5
}

func ClosePartyMenu() {
	partyMenu = nil
}
