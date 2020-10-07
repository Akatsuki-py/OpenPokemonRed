package game

import (
	"pokered/pkg/menu"
	"pokered/pkg/text"
)

const (
	Script uint = iota
	Overworld
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
	if isOverworld() {
		return Overworld
	}
	return Script
}

func isText() bool {
	return len([]rune(text.CurText)) > 0
}

func isMenu() bool {
	if menu.Quantity.Quantity > 0 {
		return true
	}
	return menu.MaxZIndex() > 0
}

func isOverworld() bool {
	return false
}
