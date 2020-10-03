package text

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/util"
)

type TextBoxID uint

const (
	MESSAGE_BOX TextBoxID = iota
	FIELD_MOVE_MON_MENU
	JP_MOCHIMONO_MENU_TEMPLATE
	USE_TOSS_MENU_TEMPLATE
	JP_SAVE_MESSAGE_MENU_TEMPLATE
	JP_SPEED_OPTIONS_MENU_TEMPLATE
	BATTLE_MENU_TEMPLATE
	SWITCH_STATS_CANCEL_MENU_TEMPLATE
	LIST_MENU_BOX
	BUY_SELL_QUIT_MENU_TEMPLATE
	MONEY_BOX_TEMPLATE
	MON_SPRITE_POPUP
	JP_AH_MENU_TEMPLATE
	MONEY_BOX
	TWO_OPTION_MENU
	BUY_SELL_QUIT_MENU
	JP_POKEDEX_MENU_TEMPLATE
	SAFARI_BATTLE_MENU_TEMPLATE
)

/*
"┌": 0x79,
"─": 0x7A,
"┐": 0x7B,
"│": 0x7C,
"└": 0x7D,
"┘": 0x7E,
*/

// DrawTextBox draw text box
func DrawTextBox(X0, Y0, X1, Y1 util.Tile) {
	// draw upper boarder
	for x := X0; x <= X1; x++ {
		switch x {
		case X0:
			PlaceChar("┌", x, Y0)
		case X1:
			PlaceChar("┐", x, Y0)
		default:
			PlaceChar("─", x, Y0)
		}
	}

	for y := Y0 + 1; y < Y1; y++ {
		for x := X0; x <= X1; x++ {
			switch x {
			case X0:
				PlaceChar("│", x, y)
			case X1:
				PlaceChar("│", x, y)
			default:
				PlaceChar(" ", x, y)
			}
		}
	}

	// draw lower boarder
	for x := X0; x <= X1; x++ {
		switch x {
		case X0:
			PlaceChar("└", x, Y1)
		case X1:
			PlaceChar("┘", x, Y1)
		default:
			PlaceChar("─", x, Y1)
		}
	}
}

// DrawTextBoxWH draw text box using width and height
func DrawTextBoxWH(X0, Y0, w, h util.Tile) {
	// draw upper boarder
	X1, Y1 := X0+w+1, Y0+h+1
	for x := X0; x <= X1; x++ {
		switch x {
		case X0:
			PlaceChar("┌", x, Y0)
		case X1:
			PlaceChar("┐", x, Y0)
		default:
			PlaceChar("─", x, Y0)
		}
	}

	for y := Y0 + 1; y < Y1; y++ {
		for x := X0; x <= X1; x++ {
			switch x {
			case X0:
				PlaceChar("│", x, y)
			case X1:
				PlaceChar("│", x, y)
			default:
				PlaceChar(" ", x, y)
			}
		}
	}

	// draw lower boarder
	for x := X0; x <= X1; x++ {
		switch x {
		case X0:
			PlaceChar("└", x, Y1)
		case X1:
			PlaceChar("┘", x, Y1)
		default:
			PlaceChar("─", x, Y1)
		}
	}
}

func DisplayTextBoxID(id TextBoxID) {
	switch id {
	case MESSAGE_BOX:
		DrawTextBox(0, 12, 19, 17)
	case LIST_MENU_BOX:
		DrawTextBox(4, 2, 19, 12)
	case MON_SPRITE_POPUP:
		DrawTextBox(6, 4, 14, 13) // https://imgur.com/0TKpIiz.png
	case JP_MOCHIMONO_MENU_TEMPLATE:
		DrawTextBox(0, 0, 14, 17)
		PlaceStringAtOnce(txt.JapaneseMochimonoText, 0, 3)
	case USE_TOSS_MENU_TEMPLATE:
		DrawTextBox(13, 10, 19, 14)
		PlaceStringAtOnce(txt.UseTossText, 15, 11)
	}
}
