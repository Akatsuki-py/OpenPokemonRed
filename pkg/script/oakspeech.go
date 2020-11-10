package script

import (
	"fmt"
	"pokered/pkg/audio"
	"pokered/pkg/store"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var counter uint

const (
	centerX, centerY = 6, 4
)

var lectureImage = struct {
	nidorino  [3]*ebiten.Image
	oak       [8]*ebiten.Image
	red       [3]*ebiten.Image
	redShrink [3]*ebiten.Image
	redSprite [3]*ebiten.Image
	rival     [8]*ebiten.Image
}{
	nidorino: [3]*ebiten.Image{
		openImage("nidorino", 0),
		openImage("nidorino", 1),
		openImage("nidorino", 2),
	},
	oak: [8]*ebiten.Image{
		openImage("oak", 0),
		openImage("oak", 1),
		openImage("oak", 2),
		openImage("oak", 3),
		openImage("oak", 4),
		openImage("oak", 5),
		openImage("oak", 6),
		openImage("oak", 7),
	},
	red: [3]*ebiten.Image{
		openImage("red", 0),
		openImage("red", 1),
		openImage("red", 2),
	},
	redShrink: [3]*ebiten.Image{
		openImage("red_shrink", 0),
		openImage("red_shrink", 1),
		openImage("red_shrink", 2),
	},
	redSprite: [3]*ebiten.Image{
		openImage("red_sprite", 0),
		openImage("red_sprite", 1),
		openImage("red_sprite", 2),
	},
	rival: [8]*ebiten.Image{
		openImage("rival", 0),
		openImage("rival", 1),
		openImage("rival", 2),
		openImage("rival", 3),
		openImage("rival", 4),
		openImage("rival", 5),
		openImage("rival", 6),
		openImage("rival", 7),
	},
}

func openImage(name string, index int) *ebiten.Image {
	path := fmt.Sprintf("/%s_lecture_%d.png", name, index)
	return util.OpenImage(store.FS, path)
}

// ref: OakSpeech
func oakSpeech0() {
	defer func() { counter++ }()
	if counter == 0 {
		audio.PlayMusic(audio.MUSIC_ROUTES2)
		util.WhiteScreen(store.TileMap)
	}
	switch {
	case counter <= 10:
		return
	case counter <= 20:
		util.DrawImage(store.TileMap, lectureImage.oak[0], centerX, centerY)
	case counter <= 30:
		util.DrawImage(store.TileMap, lectureImage.oak[1], centerX, centerY)
	case counter <= 40:
		util.DrawImage(store.TileMap, lectureImage.oak[2], centerX, centerY)
	case counter <= 50:
		util.DrawImage(store.TileMap, lectureImage.oak[3], centerX, centerY)
	case counter <= 60:
		util.DrawImage(store.TileMap, lectureImage.oak[4], centerX, centerY)
	case counter <= 70:
		util.DrawImage(store.TileMap, lectureImage.oak[5], centerX, centerY)
	}
}

func fadeInIntroPic() {

}
