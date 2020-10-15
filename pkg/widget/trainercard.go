package widget

import (
	"pokered/pkg/event"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	_ "pokered/pkg/data/statik"

	"github.com/hajimehoshi/ebiten"
)

var trainerCard *ebiten.Image

const (
	tcPath     string = "/trainercard.png"
	avatarPath string = "/trainercard_avatar.png"
)

var leader = [8]string{
	"brock",
	"misty",
	"lt_surge",
	"erika",
	"koga",
	"sabrina",
	"blaine",
	"giovanni",
}

const faceSuffix = "_face"
const badgeSuffix = "_badge"
const pngSuffix = ".png"

// DrawTrainerCard initialize trainer card gfx data
func DrawTrainerCard() {
	trainerCard = util.OpenImage(store.FS, tcPath)
	if trainerCard == nil {
		return
	}

	drawName()
	drawMoney()
	drawTime()
	drawAvatar()
	drawBadges()
}

func drawName() {
	name := store.Player.Name
	text.PlaceStringAtOnce(trainerCard, name, 7, 2)
}

func drawMoney() {
	text.PlaceChar(trainerCard, "¥", 8, 4)
	money := store.Player.Money
	text.PlaceUintAtOnce(trainerCard, money, 9, 4)
}

func drawTime() {
	minutes := store.Player.Time / 60
	hours := minutes / 60

	minutes %= 60
	if hours >= 255 {
		hours, minutes = 255, 0
	}

	text.PlacePlayTime(trainerCard, hours, minutes, 10, 6)
}

func drawAvatar() {
	avatar := util.OpenImage(store.FS, avatarPath)
	util.DrawImagePixel(trainerCard, avatar, 120, 8)
}

func drawBadges() {
	badges := event.Badges()
	for i := 0; i < 8; i++ {
		x := 24 + 32*(i%4)
		y := 96
		if i > 3 {
			y = 120
		}

		path := "/" + leader[i]
		if badges[i] {
			path += badgeSuffix + pngSuffix
		} else {
			path += faceSuffix + pngSuffix
		}

		badge := util.OpenImage(store.FS, path)
		util.DrawImagePixel(trainerCard, badge, x, y)
	}
}

// CloseTrainerCard release trainer card gfx data
func CloseTrainerCard() {
	trainerCard = nil
}
