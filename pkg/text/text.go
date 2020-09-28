package text

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

var CurText = ""

var specialChar = [...]string{
	"${pkmn}", "${PLAYER}", "${RIVAL}", "${TARGET}", "${USER}",
}

func PrintText(str string) {
	Seek(1, 14)
	CurText = preprocess(str)
}
func SetString(str string, x, y util.Tile) {
	Seek(x, y)
	CurText = preprocess(str)
}

func PlaceText() {
	if len([]rune(CurText)) == 0 {
		return
	}

	runes := []rune(CurText)
	c := string(runes[0])
	switch c {
	case "$":
		lParen := strings.Index(CurText, "{")
		rParen := strings.Index(CurText, "}")
		if lParen == 1 || rParen > 1 {
			key := string(runes[lParen:rParen])
			CurText = string(runes[rParen:])
			if value, ok := txt.RAM[key]; ok {
				CurText = value() + CurText
			} else if value, ok := txt.Asm[key]; ok {
				value()
			}
			return
		}
	case "\\":
		switch string(runes[1]) {
		case "n":
			placeLine()
			CurText = string(runes[2:])
		case "p":
			CurText = string(runes[2:])
		case "c":
			placeCont()
			CurText = string(runes[2:])
		case "d":
			CurText = string(runes[2:])
		case "▼":
			CurText = string(runes[2:])
		default:
			CurText = string(runes[1:])
		}
	default:
		if IsCorrectChar(c) {
			x, y := Caret()
			placeChar(c, x, y, true)
		}
		CurText = string(runes[1:])
	}
}

func placeChar(char string, x, y util.Tile, next bool) {
	font, ok := fontmap[char]
	if !ok {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(util.TileToFPixel(x, y))
	store.TileMap.DrawImage(font, op)
	if next {
		Next()
	}
}

func placeNext() {}
func placeLine() {
	Seek(1, 16)
}
func placePara() {}
func placeCont() {
	placeChar("▼", 18, 16, false)
	store.DelayFrames = 3
	manualTextScroll()
}

func manualTextScroll() {
	waitForTextScrollButtonPress()

}

func waitForTextScrollButtonPress() {
	handleDownArrowBlinkTiming()
	joypad.JoypadLowSensitivity()
	if !joypad.Joy5.A && !joypad.Joy5.B {

	}
}

func handleDownArrowBlinkTiming() {

}

func placePrompt() {}
func placePage()   {}
func placeDex()    {}
