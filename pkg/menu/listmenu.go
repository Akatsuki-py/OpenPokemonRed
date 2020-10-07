package menu

import (
	"pokered/pkg/data/constant"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
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

func ParseListMenuElm(src string) (uint, uint) {
	s := strings.Split(src, "@")
	if len(s) == 1 {
		num := uint(0)
		id, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			return 0, num
		}
		return uint(id), num
	}

	id, err := strconv.ParseUint(s[0], 10, 64)
	if err != nil {
		return 0, 0
	}
	num, err := strconv.ParseUint(s[1], 10, 64)
	if err != nil {
		return 0, 0
	}
	return uint(id), uint(num)
}

// ListMenu list menu
// ref: https://github.com/Akatsuki-py/understanding-pokemon-red
type ListMenu struct {
	ID      ListMenuID // wListMenuID
	Elm     []string   // // "A@B" A: pokemonID or itemID, B: Num
	z       uint       // zindex 0:hide
	Swap    uint       // wMenuItemToSwap
	wrap    bool       // !wMenuWatchMovingOutOfBounds
	offset  uint       // wListScrollOffset
	current uint       // wCurrentMenuItem
	image   *ebiten.Image
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

func (l *ListMenu) Hide() {
	l.z = 0
}

// Top return top tiles
func (l *ListMenu) Top() (util.Tile, util.Tile) {
	return ListMenuTopX, ListMenuTopY
}

// Len return a number of items
func (l *ListMenu) Len() int {
	return len(l.Elm)
}

// Wrap return menu wrap is enabled
func (l *ListMenu) Wrap() bool {
	return l.wrap
}

// Current return current selected
func (l *ListMenu) Current() uint {
	return l.current
}

// SetCurrent set current
func (l *ListMenu) SetCurrent(c uint) {
	l.current = c
}

func (l *ListMenu) Image() *ebiten.Image {
	return l.image
}

func (l *ListMenu) Item() string {
	if l.current >= uint(len(l.Elm)) {
		return ""
	}
	return l.Elm[l.current]
}

// NewListMenuID initialize list menu
func NewListMenuID(id ListMenuID, elm []string) {
	image := util.NewImage()
	util.SetBit(store.D730, 6)
	text.DisplayTextBoxID(image, text.LIST_MENU_BOX)
	CurListMenu = ListMenu{
		ID:    id,
		Elm:   elm,
		z:     MaxZIndex() + 1,
		image: image,
	}
}

// DisplayListMenuIDLoop wait for a player's action
func DisplayListMenuIDLoop() {
	target := CurListMenu.image
	CurListMenu.PrintEntries()
	previous := CurListMenu.current
	pressed := HandleMenuInput(target)
	PlaceCursor(target)

	switch {
	case pressed.A:
		PlaceUnfilledArrowCursor(target)
	case pressed.Down:
		if CurListMenu.offset+3 < uint(len(CurListMenu.Elm)+1) {
			if previous == 2 {
				CurListMenu.offset++
			}
		}
	case pressed.Up:
		if CurListMenu.offset > 0 {
			if previous == 0 {
				CurListMenu.offset--
			}
		}
	}
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
	util.ClearScreenArea(l.image, 5, 3, 9, 14)
	index := 0
	for i, e := range l.Elm {
		if i < int(l.offset) {
			continue
		}

		nameAtX, nameAtY := ListMenuTopX+1, ListMenuTopY+index*2

		// if a number of entries is more than 4, blink ▼
		if index == 4 {
			text.PlaceChar(l.image, "▼", nameAtX+12, nameAtY-1)
			break
		}

		switch l.ID {
		case PCPokemonListMenu:
			id, _ := ParseListMenuElm(e)
			name := constant.PokemonNameMap[id]
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		case MovesListMenu:
			id, _ := ParseListMenuElm(e)
			name := constant.MoveNameMap[id]
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		case PricedItemListMenu:
			id, _ := ParseListMenuElm(e)
			name := constant.ItemNameMap[id]
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
			price := constant.ItemPriceMap[id]
			text.PlaceChar(l.image, "¥", nameAtX+8, nameAtY+1)
			text.PlaceUintAtOnce(l.image, price, nameAtX+9, nameAtY+1)
		case ItemListMenu:
			id, _ := ParseListMenuElm(e)
			name := constant.ItemNameMap[id]
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		}

		// print cancel
		if int(l.offset)+index == len(l.Elm)-1 && index <= 2 {
			text.PlaceStringAtOnce(l.image, "CANCEL", nameAtX, nameAtY+2)
		}
		index++
	}
}
