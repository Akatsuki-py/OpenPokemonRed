package menu

import "pokered/pkg/text"

var ItemQuantity, maxItemQuantity uint

// NewQuantityMenu create quantity menu
func NewQuantityMenu(id ListMenuID) {
	switch id {
	case ItemListMenu:
		text.DrawTextBoxWH(15, 9, 3, 1)
		text.PlaceStringAtOnce("×01", 16, 10)
	case PricedItemListMenu:
		text.DrawTextBoxWH(7, 9, 11, 1)
		text.PlaceStringAtOnce("×01", 8, 10)
	}

	ItemQuantity = 1
}

func DisplayChooseQuantityMenu() {}
