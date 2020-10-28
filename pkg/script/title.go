package script

import (
	"fmt"
	"pokered/pkg/audio"
	"pokered/pkg/data/pokemon"
	"pokered/pkg/joypad"
	"pokered/pkg/palette"
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

	introCounter         int
	gengar               = util.OpenImage(store.FS, "/intro_gengar_0.png")
	nidorino             = util.OpenImage(store.FS, "/intro_nidorino_0.png")
	nidorinoX, nidorinoY = 0, 72
	gengarX, gengarY     = 13 * 8, 7 * 8
	title                Title
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

var introNidorinoAnimation1 = [...][2]int{
	{0, 0}, {-2, 2}, {-1, 2}, {1, 2}, {2, 2},
}

var introNidorinoAnimation2 = [...][2]int{
	{0, 0}, {-2, -2}, {-1, -2}, {1, -2}, {2, -2},
}

var introNidorinoAnimation3 = [...][2]int{
	{0, 0}, {-12, 6}, {-8, 6}, {8, 6}, {12, 6},
}

var introNidorinoAnimation4 = [...][2]int{
	{0, 0}, {-8, -4}, {-4, -4}, {4, -4}, {8, -4},
}

var introNidorinoAnimation5 = [...][2]int{
	{0, 0}, {-8, 4}, {-4, 4}, {4, 4}, {8, 4},
}

var introNidorinoAnimation6 = [...][2]int{
	{0, 0}, {2, 0}, {2, 0}, {0, 0},
}

var introNidorinoAnimation7 = [...][2]int{
	{-8, -16}, {-7, -14}, {-6, -12}, {-4, -10},
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
	case blankCounter >= 65 && blankCounter < 65+180:
		// shooting star
		ctr := blankCounter - 65
		x, y := 152-4*ctr, -16+4*ctr
		if x >= 0 || y <= 144 {
			util.DrawImagePixel(store.TileMap, star, x, y)
		}

		if checkForUserInterruption() {
			blankCounter = 0
			SetID(TitleIntroScene)
		}
	case blankCounter >= 65+180:
		blankCounter = 0
		SetID(TitleIntroScene)
	}

	blankCounter++
}

func titleIntroScene() {
	if introCounter == 0 {
		audio.PlayMusic(audio.MUSIC_INTRO_TITLE)
	}

	util.WhiteScreen(store.TileMap)

	switch {
	case introCounter < 80:
		counter := introCounter / 2
		nidorinoX = counter * 2
		gengarX = 13*8 - counter*2

		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 80+25:
		// nidorino hip
		start := 80

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HIP)
		}

		if introCounter%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation1[counter][1], introNidorinoAnimation1[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 105+25:
		// nidorino hop
		start := 105

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HOP)
		}

		if introCounter%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation2[counter][1], introNidorinoAnimation2[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 130+10:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 140+25:
		// nidorino hip
		start := 140

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HIP)
		}

		if introCounter%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation1[counter][1], introNidorinoAnimation1[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 165+25:
		// nidorino hop
		start := 165

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HOP)
		}

		if introCounter%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation2[counter][1], introNidorinoAnimation2[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 190+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 220+8:
		// gengar raise hand

		start := 220
		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_RAISE)
			gengar = util.OpenImage(store.FS, "/intro_gengar_1.png")
		}

		counter := (introCounter - start) / 2
		gengarX = 24 - counter*2

	case introCounter < 228+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 258+16:
		// gengar slash
		start := 258
		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_CRASH)
			gengar = util.OpenImage(store.FS, "/intro_gengar_2.png")
		}

		counter := (introCounter - start) / 2
		gengarX = 16 + counter*2

	case introCounter < 274+25:
		// nidorino back step
		start := 274

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HIP)
			nidorino = util.OpenImage(store.FS, "/intro_nidorino_1.png")
		}

		if (introCounter-start)%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation3[counter][1], introNidorinoAnimation3[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 299+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 329+8:
		start := 329

		if introCounter == start+8-1 {
			gengar = util.OpenImage(store.FS, "/intro_gengar_0.png")
		}

		counter := (introCounter - start) / 2
		gengarX = 32 - counter*2

	case introCounter < 337+60:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 397+25:
		// nidorino hip
		start := 397

		if introCounter == start {
			nidorino = util.OpenImage(store.FS, "/intro_nidorino_0.png")
			audio.PlaySound(audio.SFX_INTRO_HIP)
		}

		if (introCounter-start)%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation1[counter][1], introNidorinoAnimation1[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 422+25:
		// nidorino hop
		start := 422

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_HOP)
		}

		if (introCounter-start)%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation2[counter][1], introNidorinoAnimation2[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 449+20:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 469+20:
		start := 469

		if introCounter == start {
			nidorino = util.OpenImage(store.FS, "/intro_nidorino_1.png")
		}

		if (introCounter-start)%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation6[counter][1], introNidorinoAnimation6[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}

	case introCounter < 489+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
		}

	case introCounter < 519+20:
		start := 519

		if introCounter == start {
			audio.PlaySound(audio.SFX_INTRO_LUNGE)
			nidorino = util.OpenImage(store.FS, "/intro_nidorino_2.png")
		}

		if (introCounter-start)%5 == 0 {
			counter := (introCounter - start) / 5
			animX, animY := introNidorinoAnimation7[counter][1], introNidorinoAnimation7[counter][0]
			nidorinoX += animX
			nidorinoY += animY
		}
	}

	util.DrawImagePixel(store.TileMap, nidorino, nidorinoX, nidorinoY)
	util.DrawImagePixel(store.TileMap, gengar, gengarX, gengarY)

	// upper and lower black belt
	util.BlackScreenArea(store.TileMap, 0, 0, 4, 20)
	util.BlackScreenArea(store.TileMap, 0, 14, 4, 20)

	if introCounter == 705 {
		fadeOutToTitle()
	}

	introCounter++
}

func fadeOutToTitle() {
	introCounter = 0
	palette.GBFadeOutToWhite()
	SetID(FadeOutToWhite)
	PushID(TitlePokemonRed)
}

func titlePokemonRed() {
	if title.counter == 0 {
		audio.PlayMusic(audio.MUSIC_TITLE_SCREEN)
	}
	palette.LoadGBPal()

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

func checkForUserInterruption() bool {
	joypad.JoypadLowSensitivity()

	if joypad.JoyHeld.Up && joypad.JoyHeld.B && joypad.JoyHeld.Select {
		return true
	}

	if joypad.Joy5.Start || joypad.Joy5.A {
		return true
	}

	return false
}
