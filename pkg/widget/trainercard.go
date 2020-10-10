package widget

import (
	"pokered/pkg/event"
	"pokered/pkg/util"

	_ "pokered/pkg/data/statik"

	"github.com/hajimehoshi/ebiten"
)

var trainerCard *ebiten.Image

const (
	tcPath string = "/trainercard.png"
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
	trainerCard = util.OpenImage(tcPath)
	if trainerCard == nil {
		return
	}

	drawBadges()
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

		badge := util.OpenImage(path)
		util.DrawImagePixel(trainerCard, badge, x, y)
	}
}

// CloseTrainerCard release trainer card gfx data
func CloseTrainerCard() {
	trainerCard = nil
}
