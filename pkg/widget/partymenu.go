package widget

import (
	"pokered/pkg/store"
	"pokered/pkg/text"

	"github.com/hajimehoshi/ebiten"
)

var partyMenu *ebiten.Image

// 0: no swap 1: first selected offset for swap(starting from 1)
var partyMenuSwapID uint

// DrawPartyMenu draw party menu
// this func is always used when party menu is needed.
// e.g. Pokemon, item target, ...
// ref: RedrawPartyMenu_
func DrawPartyMenu() {
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

	text.PlaceStringAtOnce(partyMenu, mon.Nick, 3, y)

	if partyMenuSwapID > 0 {
		drawWhitePartyCursor()
	}

	// status condition
	hp, status := mon.HP, mon.Status
	printStatusCondition(offset, hp, status)

	// hp
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
