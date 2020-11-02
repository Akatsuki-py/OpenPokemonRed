package menu

import (
	"pokered/pkg/data/item"
	"pokered/pkg/data/move"
	"pokered/pkg/data/pkmn"
	"pokered/pkg/joypad"
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

func defaultListMenu() ListMenu {
	return ListMenu{
		z: 0,
	}
}

func (l *ListMenu) Close() {
	l.z = 0
}

func (l *ListMenu) Item() string {
	if l.current >= uint(len(l.Elm)) {
		return Cancel
	}
	return l.Elm[l.current]
}

// NewListMenuID initialize list menu
func NewListMenuID(id ListMenuID, elm []string) {
	image := util.NewImage()
	util.SetBit(&store.D730, 6)
	text.DisplayTextBoxID(image, text.LIST_MENU_BOX)
	CurListMenu = ListMenu{
		ID:    id,
		Elm:   elm,
		z:     MaxZIndex() + 1,
		image: image,
	}
}

// DisplayListMenuIDLoop wait for a player's action
func DisplayListMenuIDLoop() joypad.Input {
	target := CurListMenu.image
	CurListMenu.PrintEntries()
	previous := CurListMenu.current
	pressed := HandleListMenuInput(target)
	PlaceCursor(target, &CurListMenu)

	switch {
	case pressed.A:
		PlaceUnfilledArrowCursor(target, &CurListMenu)
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
	return pressed
}

// HandleListMenuInput メニューでのキー入力に対処するハンドラ
func HandleListMenuInput(target *ebiten.Image) joypad.Input {
	l := &CurListMenu
	PlaceCursor(target, l)
	store.DelayFrames = 3
	// TODO: AnimatePartyMon

	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.Any() {
		return joypad.Input{} // TODO: blink
	}

	maxItem := uint(len(l.Elm) - 1)
	if maxItem > 2 {
		maxItem = 2
	} else {
		maxItem++
	}
	l.current = handleMenuInput(l.current, maxItem, l.wrap)
	return joypad.Joy5
}

// PrintEntries print list menu entries in text box
// ref: PrintListMenuEntries
func (l *ListMenu) PrintEntries() {
	util.ClearScreenArea(l.image, 5, 3, 9, 14)
	index := 0
	if len(l.Elm) == 0 {
		text.PlaceStringAtOnce(l.image, "CANCEL", ListMenuTopX+1, ListMenuTopY)
		return
	}

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
			name := pkmn.Name(id)
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		case MovesListMenu:
			id, _ := ParseListMenuElm(e)
			name := move.Name(id)
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		case PricedItemListMenu:
			id, _ := ParseListMenuElm(e)
			name := item.Name(id)
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
			price := item.Price(id)
			text.PlaceChar(l.image, "¥", nameAtX+8, nameAtY+1)
			text.PlaceUintAtOnce(l.image, price, nameAtX+9, nameAtY+1)
		case ItemListMenu:
			id, _ := ParseListMenuElm(e)
			name := item.Name(id)
			text.PlaceStringAtOnce(l.image, name, nameAtX, nameAtY)
		}

		// print cancel
		if int(l.offset)+index == len(l.Elm)-1 && index <= 2 {
			text.PlaceStringAtOnce(l.image, "CANCEL", nameAtX, nameAtY+2)
		}
		index++
	}
}
