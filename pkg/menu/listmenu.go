package menu

import (
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
)

type ListMenuID = uint

const (
	PCPokemonListMenu ListMenuID = iota
	MovesListMenu
	PricedItemListMenu
	ItemListMenu
	SpecialListMenu
)

// ListMenuElm list menu element
type ListMenuElm struct {
	ID  uint // pokemonID or itemID
	Num uint // ITEMLISTMENU only
}

// ListMenu list menu
// ref: https://github.com/Akatsuki-py/understanding-pokemon-red
type ListMenu struct {
	ID         ListMenuID // wListMenuID
	Elm        []ListMenuElm
	Z          uint      // zindex 0:hide
	TopX, TopY util.Tile // wTopMenuItemX,Y
	Swap       uint      // wMenuItemToSwap
	Wrap       bool      // !wMenuWatchMovingOutOfBounds
	Select     uint      // wCurrentMenuItem
}

func defaultListMenu() ListMenu {
	return ListMenu{
		Z: 0,
	}
}

// CurListMenu list menu displayed now
var CurListMenu = defaultListMenu()

// LastListMenu list menu where player select item or exit
var LastListMenu = defaultListMenu()

// InitListMenuID initialize list menu
func InitListMenuID(id ListMenuID, elm []ListMenuElm) {
	util.SetBit(store.D730, 6)
	text.DisplayTextBoxID(text.LIST_MENU_BOX)

	CurListMenu = ListMenu{
		ID:     id,
		Elm:    elm,
		Z:      maxZIndex(),
		TopX:   5,
		TopY:   4,
		Swap:   0,
		Wrap:   false,
		Select: 0,
	}
}

// ExitListMenu exit list menu if player cancel list menu
func ExitListMenu() {
	LastListMenu = CurListMenu
	CurListMenu = defaultListMenu()
	MenuExitMethod = CancelledMenu
	util.ResBit(store.D730, 6)
}
