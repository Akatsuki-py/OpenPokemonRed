package menu

import (
	"fmt"
	"pokered/pkg/joypad"
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

type QuantityMenu struct {
	MenuID                 uint
	Quantity, Choosen, Max uint
	price                  uint
	image                  *ebiten.Image
}

var Quantity = QuantityMenu{}

// NewQuantityMenu create quantity menu
func NewQuantityMenu(id ListMenuID, max, price uint) {
	if max < 1 {
		return
	}

	Quantity = QuantityMenu{
		MenuID:   id,
		Quantity: 1,
		Max:      max,
		price:    price,
		image:    util.NewImage(),
	}
	switch id {
	case ItemListMenu:
		text.DrawTextBoxWH(Quantity.image, 15, 9, 3, 1)
		text.PlaceStringAtOnce(Quantity.image, "×01", 16, 10)
	case PricedItemListMenu:
		text.DrawTextBoxWH(Quantity.image, 7, 9, 11, 1)
		text.PlaceStringAtOnce(Quantity.image, "×01", 8, 10)
	}

	if Quantity.MenuID == PricedItemListMenu {
		Quantity.Max = 99
		printPrice()
	}
}

func DisplayChooseQuantityMenu() {
	joypad.JoypadLowSensitivity()

	switch {
	case joypad.JoyPressed.A:
		Quantity.Choosen = Quantity.Quantity
		Quantity.Quantity = 0
	case joypad.JoyPressed.B:
		Quantity.Choosen = 0
		Quantity.Quantity = 0
	case joypad.JoyPressed.Up:
		incrementQuantity()
		printQuantity()
		if Quantity.MenuID == PricedItemListMenu {
			printPrice()
		}
	case joypad.JoyPressed.Down:
		decrementQuantity()
		printQuantity()
		if Quantity.MenuID == PricedItemListMenu {
			printPrice()
		}
	}

}

func incrementQuantity() {
	Quantity.Quantity++
	if Quantity.Quantity > Quantity.Max {
		Quantity.Quantity = 1
	}
}

func decrementQuantity() {
	Quantity.Quantity--
	if Quantity.Quantity <= 0 {
		Quantity.Quantity = Quantity.Max
	}
}

func printQuantity() {
	x, y := 17, 10
	if Quantity.MenuID == PricedItemListMenu {
		x = 9
	}
	if Quantity.Quantity >= 10 {
		text.PlaceUintAtOnce(Quantity.image, Quantity.Quantity, x, y)
		return
	}
	text.PlaceStringAtOnce(Quantity.image, fmt.Sprintf("0%d", Quantity.Quantity), x, y)
}

func printPrice() {
	x := 12
	util.ClearScreenArea(Quantity.image, x, 10, 1, 6)
	price := Quantity.price * Quantity.Quantity
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
	text.PlaceChar(Quantity.image, "¥", x, 10)
	text.PlaceUintAtOnce(Quantity.image, price, x+1, 10)
}
