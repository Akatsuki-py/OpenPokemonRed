package text

import (
	"image"
	"pokered/pkg/audio"
	"pokered/pkg/data/txt"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

var CurText = ""

var InScroll bool

// PrintText print string in text window
func PrintText(str string) {
	Seek(1, 14)
	CurText = preprocess(str)
}

// SetString print string
func SetString(str string, x, y util.Tile) {
	Seek(x, y)
	CurText = preprocess(str)
}

// PlaceText print string one by one
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
			key := string(runes[lParen+1 : rParen])
			CurText = string(runes[rParen:])
			if value, ok := txt.RAM[key]; ok {
				CurText = value() + CurText
			} else if value, ok := txt.Asm[key]; ok {
				value()
			}
			return
		}
	case "#":
		CurText = "POKé" + string(runes[1:])
		PlaceText()
		return
	case "\\":
		switch string(runes[1]) {
		case "n":
			placeLine()
			CurText = string(runes[2:])
		case "p":
			if placePara() {
				CurText = string(runes[2:])
			}
		case "c":
			if placeCont() {
				ScrollTextUpOneLine()
				CurText = string(runes[2:])
			}
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
	util.DrawImage(font, x, y)
	if next {
		Next()
	}
}

func placeNext() {}
func placeLine() {
	Seek(1, 16)
}
func placePara() bool {
	placeChar("▼", 18, 16, false)
	ok := manualTextScroll()
	if ok {
		clearScreenArea()
		store.DelayFrames = 20
		Seek(1, 14)
	}
	return ok
}

func clearScreenArea() {
	for h := 13; h <= 17; h++ {
		for w := 0; w < 20; w++ {
			placeChar(" ", w, h, false)
		}
	}
}

func placeCont() bool {
	placeChar("▼", 18, 16, false)
	ok := manualTextScroll()
	if ok {
		placeChar(" ", 18, 16, false)
	}
	return ok
}

func manualTextScroll() bool {
	ok := waitForTextScrollButtonPress()
	if ok {
		audio.PlaySound(audio.SFX_PRESS_AB)
	}
	return ok
}

func waitForTextScrollButtonPress() bool {
	handleDownArrowBlinkTiming()
	joypad.JoypadLowSensitivity()
	return joypad.Joy5.A || joypad.Joy5.B
}

func handleDownArrowBlinkTiming() {}

func ScrollTextUpOneLine() {
	minX, minY := util.TileToPixel(1, 14)
	min := image.Point{minX, minY}
	maxX, maxY := util.TileToPixel(19, 17)
	max := image.Point{maxX, maxY}
	texts, _ := ebiten.NewImageFromImage(store.TileMap.SubImage(image.Rectangle{min, max}), ebiten.FilterDefault)
	util.DrawImage(texts, 1, 13)
	store.TileMap, _ = ebiten.NewImageFromImage(store.TileMap, ebiten.FilterDefault)
	for w := 0; w < 20; w++ {
		placeChar(" ", w, 16, false)
	}
	store.DelayFrames = 5
	InScroll = !InScroll
	Seek(1, 16)
}

func placePrompt() {}
func placePage()   {}
func placeDex()    {}
