package game

import (
	"pokered/pkg/menu"
	"pokered/pkg/text"
)

const (
	Overworld uint = iota
	Text
	Menu
)

func mode() uint {
	if isText() {
		return Text
	}
	if isMenu() {
		return Menu
	}
	return Overworld
}

func isText() bool {
	return len([]rune(text.CurText)) > 0
}

func isMenu() bool {
	if menu.ItemQuantity > 0 {
		return true
	}
	return menu.MaxZIndex() > 0
}
