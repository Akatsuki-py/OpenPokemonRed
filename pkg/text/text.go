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
var blink = "▼"
var downArrowBlinkCnt uint = 6 * 10 // FF8B,FF8C

// Blink ▼ on display
func Blink(b string) {
	if b == " " || b == "▼" {
		blink = b
	}
	placeChar(blink, 18, 16, false)
}

func resetBlink() {
	blink = "▼"
	downArrowBlinkCnt = 6 * 10
}

// PrintText print string in text window
func PrintText(str string) {
	DisplayTextBoxID(MESSAGE_BOX)
	Seek(1, 14)
	CurText = preprocess(str)
}

// PlaceString print string
func PlaceString(str string, x, y util.Tile) {
	Seek(x, y)
	CurText = preprocess(str)
}

// PlaceStringAtOnce print string at once
func PlaceStringAtOnce(str string, x, y util.Tile) {
	Seek(x, y)
	str = preprocess(str)
	for str != "" {
		str = PlaceChar(str)
	}
}

// PlaceChar place CurText into screen one by one
func PlaceChar(str string) string {
	if len([]rune(str)) == 0 {
		return str
	}

	runes := []rune(str)
	c := string(runes[0])
	switch c {
	case "$":
		lParen := strings.Index(str, "{")
		rParen := strings.Index(str, "}")
		if lParen == 1 || rParen > 1 {
			key := string(runes[lParen+1 : rParen])
			str = string(runes[rParen:])
			if value, ok := txt.RAM[key]; ok {
				str = value() + str
			} else if value, ok := txt.Asm[key]; ok {
				value()
			}
			return str
		}
	case "#":
		str = "POKé" + string(runes[1:])
		return PlaceChar(str)
	case "\\":
		switch string(runes[1]) {
		case "n":
			placeLine()
			str = string(runes[2:])
		case "p":
			Blink("")
			if pressed := placePara(); pressed {
				str = string(runes[2:])
				resetBlink()
			}
		case "c":
			Blink("")
			if pressed := placeCont(); pressed {
				Blink(" ")
				ScrollTextUpOneLine()
				str = string(runes[2:])
				resetBlink()
			}
		case "d":
			str = string(runes[2:])
		case "▼":
			str = string(runes[2:])
		default:
			str = string(runes[1:])
		}
	default:
		if IsCorrectChar(c) {
			x, y := Caret()
			placeChar(c, x, y, true)
		}
		str = string(runes[1:])
	}
	return str
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
	pressed := manualTextScroll()
	if pressed {
		clearScreenArea()
		store.DelayFrames = 20
		Seek(1, 14)
	}
	return pressed
}

func clearScreenArea() {
	for h := 13; h <= 16; h++ {
		for w := 1; w < 19; w++ {
			placeChar(" ", w, h, false)
		}
	}
}

func placeCont() bool {
	pressed := manualTextScroll()
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
	for w := 1; w < 19; w++ {
		placeChar(" ", w, 16, false)
	}
	store.DelayFrames = 5
	InScroll = !InScroll
	Seek(1, 16)
}

func placePrompt() {}
func placePage()   {}
func placeDex()    {}
