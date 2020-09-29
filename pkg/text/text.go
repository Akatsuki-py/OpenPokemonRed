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

// CurText text which should be displayed
var CurText = ""

// InScroll in scroll
var InScroll bool
var blink = " "
var downArrowBlinkCnt uint = 6 * 10 // FF8B,FF8C

func Blink() {
	placeChar(blink, 18, 16, false)
}

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
			Blink()
			if pressed := placePara(); pressed {
				CurText = string(runes[2:])
			}
		case "c":
			Blink()
			if pressed := placeCont(); pressed {
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
	pressed := manualTextScroll()
	if pressed {
		clearScreenArea()
		store.DelayFrames = 20
		Seek(1, 14)
	}
	return pressed
}

func clearScreenArea() {
	for h := 13; h <= 17; h++ {
		for w := 0; w < 20; w++ {
			placeChar(" ", w, h, false)
		}
	}
}

func placeCont() bool {
	Blink()
	pressed := manualTextScroll()
	if pressed {
		blink = " "
		Blink()
	}
	return pressed
}

func manualTextScroll() bool {
	pressed := WaitForTextScrollButtonPress()
	if pressed {
		audio.PlaySound(audio.SFX_PRESS_AB)
	}
	return pressed
}

// WaitForTextScrollButtonPress wait for AB button press
func WaitForTextScrollButtonPress() bool {
	handleDownArrowBlinkTiming()
	joypad.JoypadLowSensitivity()
	pressed := joypad.Joy5.A || joypad.Joy5.B
	return pressed
}

func handleDownArrowBlinkTiming() {
	downArrowBlinkCnt--
	if downArrowBlinkCnt == 0 {
		switch blink {
		case "▼":
			blink = " "
		case " ":
			blink = "▼"
		}
		downArrowBlinkCnt = 6 * 10
	}
}

// ScrollTextUpOneLine scroll text up one line
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
