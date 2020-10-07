package script

import (
	"fmt"
	"pokered/pkg/menu"
	"pokered/pkg/store"
	"pokered/pkg/widget"
)

const (
	Halt uint = iota
	WidgetStartMenu
	WidgetStartMenu2
)

// ScriptID current script ID
var ScriptID = Halt

// ScriptMap script ID -> script
var scriptMap = newScriptMap()

func newScriptMap() map[uint]func() {
	result := map[uint]func(){}
	return result
}

func Current() func() {
	s, ok := scriptMap[ScriptID]
	if !ok {
		return halt
	}
	return s
}

func halt() {}

func widgetStartMenu() {
	ScriptID = WidgetStartMenu2
	widget.DisplayStartMenu()
}

func widgetStartMenu2() {
	item := store.PopMenuItem()
	switch item {
	case "EXIT":
		fmt.Println("exit")
	case menu.Cancelled:
		fmt.Println("cancel")
	}
}
