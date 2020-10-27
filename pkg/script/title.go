package script

import (
	"fmt"
	"pokered/pkg/audio"
	"pokered/pkg/data/pokemon"
	"pokered/pkg/store"
	"pokered/pkg/text"
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
	star         = util.OpenImage(store.FS, "/star.png")

	title Title
)

var titleMons = []uint{
	pokemon.CHARMANDER,
	pokemon.BULBASAUR,
	pokemon.WEEDLE,
	pokemon.NIDORAN_M,
	pokemon.SCYTHER,
	pokemon.PIKACHU,
	pokemon.CLEFAIRY,
	pokemon.RHYDON,
	pokemon.ABRA,
	pokemon.GASTLY,
	pokemon.DITTO,
	pokemon.PIDGEOTTO,
	pokemon.ONIX,
	pokemon.PONYTA,
	pokemon.MAGIKARP,
}

type Title struct {
	counter     int
	img         *ebiten.Image
	logo        *ebiten.Image
	redVersion  *ebiten.Image
	red         *ebiten.Image
	redWithBall *ebiten.Image
	redBall     *ebiten.Image
	monID       uint
	mon         *ebiten.Image
}

func titleCopyright() {
	if copyrightImage == nil {
		copyrightImage = util.NewImage()
		util.WhiteScreen(copyrightImage)
		text.PlaceStringAtOnce(copyrightImage, "Open#monRed", 1, 10)
		text.PlaceStringAtOnce(copyrightImage, "This is a fan proj.", 1, 13)
		text.PlaceStringAtOnce(copyrightImage, "Plz support the", 1, 15)
		text.PlaceStringAtOnce(copyrightImage, "official one.", 1, 16)
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

	switch {
	case blankCounter == 64:
		audio.PlaySound(audio.SFX_SHOOTING_STAR)
		text.PlaceStringAtOnce(blankImage, "credit", 7, 9)
	case blankCounter >= 65:
		ctr := blankCounter - 65
		x, y := 152-4*ctr, -16+4*ctr
		if x >= 0 || y <= 144 {
			util.DrawImagePixel(store.TileMap, star, x, y)
		}
	}

	blankCounter++
}

func titlePokemonRed() {
	audio.PlayMusic(audio.MUSIC_TITLE_SCREEN)

	if title.img == nil {
		title.img = util.NewImage()
		util.WhiteScreen(title.img)

		title.logo = util.OpenImage(store.FS, "/title_logo.png")
		util.DrawImage(title.img, title.logo, 2, 1)

		title.redVersion = util.OpenImage(store.FS, "/red_version.png")
		util.DrawImage(title.img, title.redVersion, 7, 8)

		text.PlaceStringAtOnce(title.img, "Github: pokemium", 2, 17)
	}
	util.DrawImage(store.TileMap, title.img, 0, 0)

	if title.mon == nil {
		title.monID = pokemon.CHARMANDER
		name := pokemon.Name(title.monID)
		path := fmt.Sprintf("/%s_1.png", name)
		title.mon = util.OpenImage(store.FS, path)
	}
	util.DrawImage(store.TileMap, title.mon, 5, 10)

	if title.red == nil {
		title.red = util.OpenImage(store.FS, "/title_red_1.png")
		title.redWithBall = util.OpenImage(store.FS, "/title_red_0.png")
		title.redBall = util.OpenImage(store.FS, "/title_red_ball.png")
	}
	util.DrawImagePixel(store.TileMap, title.redWithBall, 82, 80)

	title.counter++
}
