package menu

import (
	"pokered/pkg/data/constant"
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
	z          uint      // zindex 0:hide
	TopX, TopY util.Tile // wTopMenuItemX,Y
	Swap       uint      // wMenuItemToSwap
	Wrap       bool      // !wMenuWatchMovingOutOfBounds
	Select     uint      // wCurrentMenuItem
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

// InitListMenuID initialize list menu
func InitListMenuID(id ListMenuID, elm []ListMenuElm) {
	util.SetBit(store.D730, 6)
	text.DisplayTextBoxID(text.LIST_MENU_BOX)
	util.ClearScreenArea(5, 3, 9, 14)

	CurListMenu = ListMenu{
		ID:     id,
		Elm:    elm,
		z:      maxZIndex(),
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

// PrintEntries PrintListMenuEntries
func (l *ListMenu) PrintEntries() {
	for i, e := range l.Elm {
		nameAtX, nameAtY := 6, 4+i*2
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
	}
}
