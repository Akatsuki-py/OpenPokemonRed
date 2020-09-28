package text

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

var CurText = ""

var specialChar = [...]string{
	"${pkmn}", "${PLAYER}", "${RIVAL}", "${TARGET}", "${USER}",
}

func SetText(str string) {
	CurText = str
}

func PrintText() {
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
			CurText = string(runes[2:])
		case "p":
			CurText = string(runes[2:])
		case "c":
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
			placeChar(c)
		}
		CurText = string(runes[1:])
	}
}

func placeChar(char string) {
	charcode, ok := charmap[char]
	if !ok {
		return
	}
	font, ok := fontmap[charcode]
	if !ok {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(util.TileToFPixel(Caret()))
	store.TileMap.DrawImage(font, op)
	Next()
}

func placeNext()   {}
func placeLine()   {}
func placePara()   {}
func placeCont()   {}
func placePrompt() {}
func placePage()   {}
func placeDex()    {}
