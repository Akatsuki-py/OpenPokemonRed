package menu

import (
	"fmt"
	"pokered/pkg/joypad"
	"pokered/pkg/text"
	"pokered/pkg/util"
)

var quantityMenuID uint
var ItemQuantity, ChoosenQuantity, maxItemQuantity uint
var itemPrice uint

// NewQuantityMenu create quantity menu
func NewQuantityMenu(id ListMenuID, max, price uint) {
	if max < 1 {
		return
	}
	switch id {
	case ItemListMenu:
		text.DrawTextBoxWH(15, 9, 3, 1)
		text.PlaceStringAtOnce("×01", 16, 10)
	case PricedItemListMenu:
		text.DrawTextBoxWH(7, 9, 11, 1)
		text.PlaceStringAtOnce("×01", 8, 10)
	}

	quantityMenuID = id
	ItemQuantity = 1
	maxItemQuantity = max
	itemPrice = price

	if quantityMenuID == PricedItemListMenu {
		maxItemQuantity = 99
		printPrice()
	}
}

func DisplayChooseQuantityMenu() {
	joypad.JoypadLowSensitivity()

	switch {
	case joypad.JoyPressed.A:
		ChoosenQuantity = ItemQuantity
		ItemQuantity = 0
	case joypad.JoyPressed.B:
		ChoosenQuantity = 0
		ItemQuantity = 0
	case joypad.JoyPressed.Up:
		incrementQuantity()
		printQuantity()
		if quantityMenuID == PricedItemListMenu {
			printPrice()
		}
	case joypad.JoyPressed.Down:
		decrementQuantity()
		printQuantity()
		if quantityMenuID == PricedItemListMenu {
			printPrice()
		}
	}

}

func incrementQuantity() {
	ItemQuantity++
	if ItemQuantity > maxItemQuantity {
		ItemQuantity = 1
	}
}

func decrementQuantity() {
	ItemQuantity--
	if ItemQuantity <= 0 {
		ItemQuantity = maxItemQuantity
	}
}

func printQuantity() {
	x, y := 17, 10
	if quantityMenuID == PricedItemListMenu {
		x = 9
	}
	if ItemQuantity >= 10 {
		text.PlaceUintAtOnce(ItemQuantity, x, y)
		return
	}
	text.PlaceStringAtOnce(fmt.Sprintf("0%d", ItemQuantity), x, y)
}

func printPrice() {
	x := 12
	util.ClearScreenArea(x, 10, 1, 6)
	price := itemPrice * ItemQuantity
	switch {
	case price >= 10000:
		x = 13
	case price >= 1000:
		x = 14
	case price >= 100:
		x = 15
	case price >= 10:
		x = 16
	default:
		x = 17
	}
	text.PlaceChar("¥", x, 10)
	text.PlaceUintAtOnce(price, x+1, 10)
}
