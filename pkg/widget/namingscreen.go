package widget

import (
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

var namingScreen *ebiten.Image
var isLowercase bool
var NamingCursor = [2]uint{}

var cursorMask = newCursorMask()

const (
	PLAYER_NAME uint = iota
	RIVAL_NAME
	NICKNAME
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

// DrawNamingScreen initialize naming screen gfx data
func DrawNamingScreen(id uint) {
	namingScreen = util.NewImage()
	util.WhiteScreen(namingScreen)
	isLowercase = false
	drawKeyboard()

	switch id {
	case PLAYER_NAME:
		text.PlaceStringAtOnce(namingScreen, "YOUR NAME?", 0, 1)
	case RIVAL_NAME:
		text.PlaceStringAtOnce(namingScreen, "RIVAL's NAME?", 0, 1)
	case NICKNAME:
		text.PlaceStringAtOnce(namingScreen, "NICKNAME?", 1, 3)
	}
}

// UpdateNamingScreen update naming screen gfx data
func UpdateNamingScreen() {
	placeCursor()
}

// CloseNamingScreen release naming screen gfx data
func CloseNamingScreen() {
	namingScreen = nil
}

func placeCursor() {
	util.DrawImage(namingScreen, cursorMask, 0, 4)
	x, y := 1, 1
	text.PlaceChar(namingScreen, "▶︎", 1+2*x, 5+2*y)
}

func drawKeyboard() {
	keyboard := util.OpenImage("/uppercase.png")
	if isLowercase {
		keyboard = util.OpenImage("/lowercase.png")
	}
	util.DrawImage(namingScreen, keyboard, 0, 4)
}
