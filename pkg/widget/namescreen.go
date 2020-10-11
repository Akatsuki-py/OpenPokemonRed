package widget

import (
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

type NameScreen struct {
	id          uint
	screen      *ebiten.Image
	input       string
	isLowercase bool
	cursor      [2]int
}

var name = NameScreen{}

const (
	maxCursorX int = 8
	maxCursorY int = 5
)

const (
	maxName     int = 10
	maxNickname int = 6
)

const (
	uppercaseKeyboard = "/uppercase.png"
	lowercaseKeyboard = "/lowercase.png"
	underscoreUp      = "/name_underscore0.png"
	underscore        = "/name_underscore1.png"
)

var cursorMask = newCursorMask()

// name screen id
const (
	PlayerName uint = iota
	RivalName
	Nickname
)

func newCursorMask() *ebiten.Image {
	img, _ := ebiten.NewImage(20*8, 12*8, ebiten.FilterDefault)
	for h := 0; h < 5; h++ {
		for w := 0; w < 9; w++ {
			text.PlaceChar(img, " ", 1+2*w, 1+2*h)
		}
	}
	text.PlaceChar(img, " ", 1, 11)
	return img
}

// DrawNameScreen initialize naming screen gfx data
func DrawNameScreen(id uint) {
	name.id = id
	name.input = "NINTEN"
	name.screen = util.NewImage()
	util.WhiteScreen(name.screen)
	name.isLowercase = false
	drawKeyboard()

	switch id {
	case PlayerName:
		text.PlaceStringAtOnce(name.screen, "YOUR NAME?", 0, 1)
	case RivalName:
		text.PlaceStringAtOnce(name.screen, "RIVAL's NAME?", 0, 1)
	case Nickname:
		text.PlaceStringAtOnce(name.screen, "NICKNAME?", 1, 3)
	}
}

// UpdateNameScreen update naming screen gfx data
func UpdateNameScreen() {
	drawKeyboard()
	printName()
	placeCursor()
}

// CloseNameScreen release naming screen gfx data
func CloseNameScreen() string {
	input := name.input
	name = NameScreen{}
	return input
}

func placeCursor() {
	util.DrawImage(name.screen, cursorMask, 0, 4)
	x, y := name.cursor[0], name.cursor[1]
	text.PlaceChar(name.screen, "▶︎", 1+2*x, 5+2*y)
}

func drawKeyboard() {
	keyboard := util.OpenImage(uppercaseKeyboard)
	if name.isLowercase {
		keyboard = util.OpenImage(lowercaseKeyboard)
	}
	util.DrawImage(name.screen, keyboard, 0, 4)
}

// SetNameCursor update name cursor position
func SetNameCursor(deltaX, deltaY int) {
	switch deltaY {
	case 1:
		name.cursor[1]++
		if name.cursor[1] == maxCursorY {
			name.cursor[0] = 0
		}
		if name.cursor[1] > maxCursorY {
			name.cursor[0], name.cursor[1] = 0, 0
		}
		return

	case -1:
		name.cursor[1]--
		if name.cursor[1] < 0 {
			name.cursor[0], name.cursor[1] = 0, maxCursorY
		}
		return
	}

	if name.cursor[1] == maxCursorY {
		// if cursor is on UPPER/lower
		name.cursor[0] = 0
		return
	}

	switch deltaX {
	case 1:
		name.cursor[0]++
		if name.cursor[0] > maxCursorX {
			name.cursor[0] = 0
		}
	case -1:
		name.cursor[0]--
		if name.cursor[0] < 0 {
			name.cursor[0] = maxCursorX
		}
	}
}

// ToggleCase toggle UPPER/lower case
func ToggleCase() {
	name.isLowercase = !name.isLowercase
}

func printName() {
	text.PlaceStringAtOnce(name.screen, name.input, 10, 2)
}

func printUnderscores() {

}
