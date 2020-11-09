package widget

import (
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/pkmn"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var animationCounter uint

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

	h := pkmnd.Header(mon.ID)
	ico := pkmn.IconGen1[h.IconGen1][0]
	util.DrawImage(partyMenu, ico, 1, y)

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
	animationCounter++

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	partyMenuPrev := partyMenuCurrent
	partyMenuCurrent = menu.HandleMenuInput(partyMenuCurrent, uint(store.PartyMonLen()-1), true)
	if partyMenuPrev != partyMenuCurrent {
		clearPartyMonAnimation()
	}
	return joypad.Joy5
}

func ClosePartyMenu() {
	partyMenu = nil
}

func AnimatePartyMon() {
	offset := partyMenuCurrent
	y := int(offset * 2)

	mon := store.PartyMons[offset]
	if !mon.Initialized {
		return
	}

	spd := pkmn.PartyMonSpeeds(mon.HP, mon.MaxHP)
	prevIconIndex := ((animationCounter - 1) / spd) % 2
	curIconIndex := (animationCounter / spd) % 2

	if curIconIndex != prevIconIndex {
		h := pkmnd.Header(mon.ID)
		icon := h.IconGen1
		drawPartyMonGen1(icon, curIconIndex, y)
	}
}

func clearPartyMonAnimation() {
	animationCounter = 0

	for i, mon := range store.PartyMons {
		y := i * 2
		if !mon.Initialized {
			break
		}

		h := pkmnd.Header(mon.ID)
		icon := h.IconGen1
		drawPartyMonGen1(icon, 0, y)
	}
}

func drawPartyMonGen1(icon, index uint, y util.Tile) {
	util.ClearScreenArea(partyMenu, 1, y, 2, 2)
	iconImage := pkmn.IconGen1[icon][index]
	util.DrawImage(partyMenu, iconImage, 1, y)
}
