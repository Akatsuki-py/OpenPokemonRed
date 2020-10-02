package menu

import (
	"pokered/pkg/data/constant"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
)

const (
	ListMenuTopX, ListMenuTopY util.Tile = 5, 4
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
	ID      ListMenuID // wListMenuID
	Elm     []ListMenuElm
	z       uint // zindex 0:hide
	Swap    uint // wMenuItemToSwap
	Wrap    bool // !wMenuWatchMovingOutOfBounds
	Current uint // wCurrentMenuItem
}

// CurListMenu list menu displayed now
var CurListMenu = defaultListMenu()

// LastListMenu list menu where player select item or exit
var LastListMenu = defaultListMenu()

func defaultListMenu() ListMenu {
	return ListMenu{
		z: 0,
	}
}

// Z return zindex
func (l *ListMenu) Z() uint {
	return l.z
}

// Top return top tiles
func (l *ListMenu) Top() (util.Tile, util.Tile) {
	return ListMenuTopX, ListMenuTopY
}

// InitListMenuID initialize list menu
func InitListMenuID(id ListMenuID, elm []ListMenuElm) {
	util.SetBit(store.D730, 6)
	text.DisplayTextBoxID(text.LIST_MENU_BOX)

	CurListMenu = ListMenu{
		ID:  id,
		Elm: elm,
		z:   MaxZIndex() + 1,
	}
}

// DisplayListMenuIDLoop wait for a player's action
func DisplayListMenuIDLoop() {
	// TODO: old man battle
	HandleMenuInput()
	PlaceCursor()
}

// ExitListMenu exit list menu if player cancel list menu
func ExitListMenu() {
	LastListMenu = CurListMenu
	CurListMenu = defaultListMenu()
	MenuExitMethod = CancelledMenu
	util.ResBit(store.D730, 6)
}

// PrintEntries print list menu entries in text box
// ref: PrintListMenuEntries
func (l *ListMenu) PrintEntries() {
	util.ClearScreenArea(5, 3, 9, 14)
	for i, e := range l.Elm {
		nameAtX, nameAtY := ListMenuTopX+1, ListMenuTopY+i*2

		// if a number of entries is more than 4, blink ▼
		if i == 4 {
			text.PlaceChar("▼", nameAtX+13, nameAtY+1)
			break
		}

		switch l.ID {
		case PCPokemonListMenu:
			name := constant.PokemonNameMap[e.ID]
			text.PlaceStringAtOnce(name, nameAtX, nameAtY)
		case MovesListMenu:
			name := constant.MoveNameMap[e.ID]
			text.PlaceStringAtOnce(name, nameAtX, nameAtY)
		case PricedItemListMenu:
			name := constant.ItemNameMap[e.ID]
			text.PlaceStringAtOnce(name, nameAtX, nameAtY)
			price := constant.ItemPriceMap[e.ID]
			text.PlaceChar("¥", nameAtX+9, nameAtY+1)
			text.PlaceUintAtOnce(price, nameAtX+10, nameAtY+1)
		case ItemListMenu:
			name := constant.ItemNameMap[e.ID]
			text.PlaceStringAtOnce(name, nameAtX, nameAtY)
		}

		// print cancel
		if i == len(l.Elm)-1 && i <= 2 {
			text.PlaceStringAtOnce("CANCEL", nameAtX, nameAtY+2)
		}
	}
}
