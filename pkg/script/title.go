package script

import (
	"pokered/pkg/audio"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

const (
	blankHeight int = 32 // 0-31 112-143
)

var (
	copyrightCounter int
	copyrightImage   *ebiten.Image

	blankCounter int
	blankImage   *ebiten.Image

	title Title
)

type Title struct {
	counter     int
	img         *ebiten.Image
	logo        *ebiten.Image
	redVersion  *ebiten.Image
	red         *ebiten.Image
	redWithBall *ebiten.Image
	redBall     *ebiten.Image
}

func titleCopyright() {
	if copyrightImage == nil {
		copyrightImage = util.NewImage()
		util.WhiteScreen(copyrightImage)
	}
	util.DrawImage(store.TileMap, copyrightImage, 0, 0)

	if copyrightCounter == 180 {
		copyrightCounter = 0
		SetID(TitleBlank)
	}
	copyrightCounter++
}

func titleBlank() {
	if blankImage == nil {
		blankImage = util.NewImage()
		util.BlackScreen(blankImage)
		util.ClearScreenArea(blankImage, 0, 4, 10, 20)
	}
	util.DrawImage(store.TileMap, blankImage, 0, 0)

	if blankCounter == 64 {
		blankCounter = 0
		SetID(TitlePokemonRed)
	}
	blankCounter++
}

func titlePokemonRed() {
	audio.PlayMusic(audio.MUSIC_TITLE_SCREEN)

	if title.img == nil {
		title.img = util.NewImage()
		util.FillScreen(title.img, 0xff, 0xff, 0xff)

		title.logo = util.OpenImage(store.FS, "/title_logo.png")
		util.DrawImage(title.img, title.logo, 2, 1)

		title.redVersion = util.OpenImage(store.FS, "/red_version.png")
		util.DrawImage(title.img, title.redVersion, 7, 8)
	}
	util.DrawImage(store.TileMap, title.img, 0, 0)

	if title.red == nil {
		title.red = util.OpenImage(store.FS, "/title_red_1.png")
		title.redWithBall = util.OpenImage(store.FS, "/title_red_0.png")
		title.redBall = util.OpenImage(store.FS, "/title_red_ball.png")
	}
	util.DrawImagePixel(store.TileMap, title.redWithBall, 82, 80)
	title.counter++
}
