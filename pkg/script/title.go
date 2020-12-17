package script

import (
	"image"
	"net/http"
	"pokered/pkg/audio"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/data/worldmap/song"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/palette"
	"pokered/pkg/screen"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/world"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	blankHeight int = 32 // 0-31 112-143
)

var TitleScreen *ebiten.Image

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

	whiteOutCounter int

	title = Title{
		logo:        util.OpenImage(store.FS, "/title_logo.png"),
		redVersion:  util.OpenImage(store.FS, "/red_version.png"),
		red:         util.OpenImage(store.FS, "/title_red_1.png"),
		redWithBall: util.OpenImage(store.FS, "/title_red_0.png"),
		redBall:     util.OpenImage(store.FS, "/title_red_ball.png"),
		mon:         util.OpenImage(store.FS, "/charmander_1.png"),
	}
)

var titleMons = []uint{
	pkmnd.CHARMANDER,
	pkmnd.BULBASAUR,
	pkmnd.WEEDLE,
	pkmnd.NIDORAN_M,
	pkmnd.SCYTHER,
	pkmnd.PIKACHU,
	pkmnd.CLEFAIRY,
	pkmnd.RHYDON,
	pkmnd.ABRA,
	pkmnd.GASTLY,
	pkmnd.DITTO,
	pkmnd.PIDGEOTTO,
	pkmnd.ONIX,
	pkmnd.PONYTA,
	pkmnd.MAGIKARP,
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
	if TitleScreen == nil {
		TitleScreen = util.NewImage()
	}

	if copyrightImage == nil {
		copyrightImage = util.NewImage()
		text.PlaceStringAtOnce(copyrightImage, "Open#monRed", 1, 10)
		text.PlaceStringAtOnce(copyrightImage, "This is a fan proj.", 1, 13)
		text.PlaceStringAtOnce(copyrightImage, "Plz support the", 1, 15)
		text.PlaceStringAtOnce(copyrightImage, "official one.", 1, 16)
	}
	util.DrawImage(TitleScreen, copyrightImage, 0, 0)

	if copyrightCounter == 180 {
		copyrightCounter = 0
		store.SetScriptID(store.TitleBlank)
	}
	copyrightCounter++
}

func titleBlank() {
	if blankImage == nil {
		blankImage = util.NewImage()
		util.BlackScreen(blankImage)
		util.ClearScreenArea(blankImage, 0, 4, 10, 20)
	}
	util.DrawImage(TitleScreen, blankImage, 0, 0)

	switch {
	case blankCounter == 64:
		audio.PlaySound(audio.SFX_SHOOTING_STAR)
		text.PlaceStringAtOnce(blankImage, "credit", 7, 9)
	case blankCounter >= 65 && blankCounter < 65+180:
		// shooting star
		ctr := blankCounter - 65
		x, y := 152-4*ctr, -16+4*ctr
		if x >= 0 || y <= 144 {
			util.DrawImagePixel(TitleScreen, star, x, y)
		}

		if checkForUserInterruption() {
			blankCounter = 0
			store.SetScriptID(store.TitleIntroScene)
		}
	case blankCounter >= 65+180:
		blankCounter = 0
		store.SetScriptID(store.TitleIntroScene)
	}

	blankCounter++
}

func titleIntroScene() {
	if introCounter == 0 {
		audio.PlayMusic(song.MUSIC_INTRO_TITLE)
	}

	screen.FillWhite()

	switch {
	case introCounter < 80:
		counter := introCounter / 2
		nidorinoX = counter * 2
		gengarX = 13*8 - counter*2

		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 80+25:
		nidorinoHip(80)

	case introCounter < 105+25:
		nidorinoHop(105)

	case introCounter < 130+10:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 140+25:
		nidorinoHip(140)

	case introCounter < 165+25:
		nidorinoHop(165)

	case introCounter < 190+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 220+8:
		gengarRaiseHand(220)

	case introCounter < 228+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 258+16:
		gengarSlash(258)

	case introCounter < 274+25:
		nidorinoBackstep(274)

	case introCounter < 299+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 329+8:
		gengarLowerHand(329)

	case introCounter < 337+60:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 397+25:
		nidorinoHip(397)

	case introCounter < 422+25:
		nidorinoHop(422)

	case introCounter < 449+20:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 469+20:
		nidorinoCrouch(469)

	case introCounter < 489+30:
		if checkForUserInterruption() {
			fadeOutToTitle()
			return
		}

	case introCounter < 519+20:
		nidorinoLunge(519)
	}

	util.DrawImagePixel(TitleScreen, nidorino, nidorinoX, nidorinoY)
	util.DrawImagePixel(TitleScreen, gengar, gengarX, gengarY)

	// upper and lower black belt
	util.BlackScreenArea(TitleScreen, 0, 0, 4, 20)
	util.BlackScreenArea(TitleScreen, 0, 14, 4, 20)

	if introCounter == 539 {
		fadeOutToTitle()
		return
	}

	introCounter++
}

func nidorinoHip(start int) {

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
}

func nidorinoHop(start int) {
	if introCounter == start {
		audio.PlaySound(audio.SFX_INTRO_HOP)
	}

	if (introCounter-start)%5 == 0 {
		counter := (introCounter - start) / 5
		animX, animY := introNidorinoAnimation2[counter][1], introNidorinoAnimation2[counter][0]
		nidorinoX += animX
		nidorinoY += animY
	}
}

func nidorinoBackstep(start int) {
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
}

func nidorinoCrouch(start int) {
	if introCounter == start {
		nidorino = util.OpenImage(store.FS, "/intro_nidorino_1.png")
	}

	if (introCounter-start)%5 == 0 {
		counter := (introCounter - start) / 5
		animX, animY := introNidorinoAnimation6[counter][1], introNidorinoAnimation6[counter][0]
		nidorinoX += animX
		nidorinoY += animY
	}
}

func nidorinoLunge(start int) {
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

func gengarRaiseHand(start int) {
	if introCounter == start {
		audio.PlaySound(audio.SFX_INTRO_RAISE)
		gengar = util.OpenImage(store.FS, "/intro_gengar_1.png")
	}

	counter := (introCounter - start) / 2
	gengarX = 24 - counter*2
}

func gengarSlash(start int) {
	if introCounter == start {
		audio.PlaySound(audio.SFX_INTRO_CRASH)
		gengar = util.OpenImage(store.FS, "/intro_gengar_2.png")
	}

	counter := (introCounter - start) / 2
	gengarX = 16 + counter*2
}

func gengarLowerHand(start int) {
	if introCounter == start+8-1 {
		gengar = util.OpenImage(store.FS, "/intro_gengar_0.png")
	}

	counter := (introCounter - start) / 2
	gengarX = 32 - counter*2
}

func fadeOutToTitle() {
	introCounter = 0
	palette.GBFadeOutToWhite(false)
	store.PushScriptID(store.TitleWhiteOut)
}

func titleWhiteOut() {
	palette.LoadGBPal()
	screen.FillWhite()
	if whiteOutCounter == 20 {
		store.SetScriptID(store.TitlePokemonRed)
	}
	whiteOutCounter++
}

func titlePokemonRed() {
	if title.counter == 88 {
		audio.PlayMusic(song.MUSIC_TITLE_SCREEN)
	}

	palette.LoadGBPal()

	if title.img == nil {
		title.img = util.NewImage()
		util.WhiteScreen(title.img)
		text.PlaceStringAtOnce(title.img, "Github: pokemium", 2, 17)
	}
	util.DrawImage(TitleScreen, title.img, 0, 0)

	bounceLogo()
	slideVersion()

	util.DrawImage(TitleScreen, title.mon, 5, 10)
	util.DrawImagePixel(TitleScreen, title.redWithBall, 82, 80)

	title.counter++

	if title.counter > 88 && checkForUserInterruption() {
		title.counter = 0
		store.SetScriptID(store.TitleMenu)
	}
}

func InitializeOverworld() {
	initTilesets(store.FS)
	world.LoadWorldData(worldmap.PALLET_TOWN)
	world.LastWorld = world.CurWorld
	world.LoadWorldData(worldmap.REDS_HOUSE_2F)
	sprite.InitPlayer(sprite.Normal, 3, 6)
	store.SetScriptID(store.Overworld)
}

func bounceLogo() {
	logoY := 8
	switch {
	case title.counter < 16:
		counter := title.counter
		logoY = 8 - 64 + 4*counter
	case title.counter < 16+4:
		counter := title.counter - 16
		logoY = 8 + 3*counter
	case title.counter < 20+4:
		if title.counter == 20 {
			audio.PlaySound(audio.SFX_INTRO_CRASH)
		}
		counter := title.counter - 20
		logoY = 20 - 3*counter
	case title.counter < 24+2:
		counter := title.counter - 24
		logoY = 8 + 2*counter
	case title.counter < 26+2:
		counter := title.counter - 26
		logoY = 12 - 2*counter
	case title.counter < 28+2:
		counter := title.counter - 28
		logoY = 8 + counter
	case title.counter < 30+2:
		counter := title.counter - 30
		logoY = 10 - counter
	}
	util.DrawImagePixel(TitleScreen, title.logo, 2*8, logoY)
}

func slideVersion() {
	versionX := 56
	switch {
	case title.counter < 32+36:
		return
	case title.counter < 68+20:
		if title.counter == 68 {
			audio.PlaySound(audio.SFX_INTRO_WHOOSH)
		}
		counter := title.counter - 68
		versionX = 136 - counter*4
	}

	util.DrawImagePixel(TitleScreen, title.redVersion, versionX, 8*8)
}

func initTilesets(fs http.FileSystem) {
	result := map[uint]tileset.Tileset{}
	for id, name := range tileset.TilesetNames {
		path := "/" + name + ".png"
		img := util.OpenImage(fs, path)

		width, height := img.Size()
		width /= 8
		height /= 8
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				min, max := image.Point{w * 8, h * 8}, image.Point{(w + 1) * 8, (h + 1) * 8}
				tile := ebiten.NewImageFromImage(img.SubImage(image.Rectangle{min, max}))
				result[uint(id)] = append(result[uint(id)], tile)
			}
		}
	}
	tileset.Tilesets = result
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

func titleMenu() {
	store.SetScriptID(store.TitleMenu2)
	screen.FillWhite()
	height := 3 * 2
	elm := []string{
		"CONTINUE",
		"NEW GAME",
		"OPTION",
	}
	menu.NewSelectMenu(elm, 0, 0, 13, height, true, false, 0)
}

func titleMenu2() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()

	switch {
	case pressed.A:
		switch m.Item() {
		case "CONTINUE":
		case "NEW GAME":
			m.Close()
			store.SetScriptID(store.OakSpeech0)
			// InitializeOverworld()
		case "OPTION":
		}
	case pressed.B:
		m.Close()
		store.SetScriptID(store.TitlePokemonRed)
	}
}
