package script

import (
	"fmt"
	"pokered/pkg/audio"
	"pokered/pkg/data/txt"
	"pokered/pkg/menu"
	"pokered/pkg/palette"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/widget"

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
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

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
	case counter == 80:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.OakSpeechText1, false)
		palette.GBFadeOutToWhite(true)
		store.PushScriptID(store.OakSpeech1)
	}
}

func oakSpeech1() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter <= 15:
		if counter == 0 {
			util.WhiteScreen(store.TileMap)
		}
		x := int((176 - counter*8) / 8)
		util.DrawImage(store.TileMap, lectureImage.nidorino[0], x, centerY)
	case counter == 16:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.OakSpeechText2A, false)
		store.PushScriptID(store.OakSpeech2)
	}
}

func oakSpeech2() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 33:
		if counter == 0 {
			audio.PlaySound(audio.SFX_CRY_NIDORINO)
		}
	case counter == 33:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.OakSpeechText2B, false)
		store.PushScriptID(store.IntroducePlayer)
	}
}

func introducePlayer() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter <= 15:
		if counter == 0 {
			util.WhiteScreen(store.TileMap)
		}
		x := int((168 - counter*8) / 8)
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 16:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.IntroducePlayerText, false)
		store.PushScriptID(store.ChoosePlayerName)
	}
}

// ref: ChoosePlayerName
func choosePlayerName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 21:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(56+(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 21:
		reset = true
		store.SetScriptID(store.ChoosePlayerName2)

		// ref: DisplayIntroNameTextBox
		width, height := 10, 9
		elm := []string{
			"NEW NAME",
			"RED",
			"ASH",
			"JACK",
		}
		menu.NewSelectMenu(elm, 0, 0, width, height, true, true)
	}
}

func choosePlayerName2() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()

	switch {
	case pressed.A:
		m.Close()
		switch m.Item() {
		case "NEW NAME":
			store.SetScriptID(store.CustomPlayerName)
		default:
			store.Player.Name = m.Item()
			store.SetScriptID(store.AfterChoosePlayerName)
		}
	}
}

func customPlayerName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 15:
		util.WhiteScreen(store.TileMap)
	case counter == 15:
		reset = true
		widget.DrawNameScreen(widget.PlayerName)
		store.SetScriptID(store.WidgetPlayerNamingScreen)
	}
}

func widgetPlayerNamingScreen() {
	name, ok := handleNamingScreen()
	if ok {
		store.Player.Name = name
		store.SetScriptID(store.AfterCustomPlayerName)
	}
}

// after choose NAME
func afterChoosePlayerName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter == 0:
		util.WhiteScreen(store.TileMap)
	case counter < 22:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(104-(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 22:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.YourNameIsText, false)
		store.PushScriptID(store.IntroduceRival)
	}
}

func afterCustomPlayerName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	util.WhiteScreen(store.TileMap)
	switch {
	case counter == 18:
		reset = true
		util.DrawImage(store.TileMap, lectureImage.red[0], 7, centerY)
		text.DoPrintTextScript(text.TextBoxImage, txt.YourNameIsText, false)

		store.PushOtScript(fadeoutScreen)
		store.PushScriptID(store.IntroduceRival)
	}
}

func fadeoutScreen() {
	palette.GBFadeOutToWhite(true)
}

// introduce rival
func introduceRival() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter <= 10:
		util.WhiteScreen(store.TileMap)
	case counter <= 20:
		util.DrawImage(store.TileMap, lectureImage.rival[0], centerX, centerY)
	case counter <= 30:
		util.DrawImage(store.TileMap, lectureImage.rival[1], centerX, centerY)
	case counter <= 40:
		util.DrawImage(store.TileMap, lectureImage.rival[2], centerX, centerY)
	case counter <= 50:
		util.DrawImage(store.TileMap, lectureImage.rival[3], centerX, centerY)
	case counter <= 60:
		util.DrawImage(store.TileMap, lectureImage.rival[4], centerX, centerY)
	case counter <= 70:
		util.DrawImage(store.TileMap, lectureImage.rival[5], centerX, centerY)
	case counter == 80:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.IntroduceRivalText, false)
		store.PushScriptID(store.ChooseRivalName)
	}
}

// ref: ChooseRivalName
func chooseRivalName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 21:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(56+(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.rival[5], x, centerY)
	case counter == 21:
		reset = true
		store.SetScriptID(store.ChooseRivalName2)

		// ref: DisplayIntroNameTextBox
		width, height := 10, 9
		elm := []string{
			"NEW NAME",
			"BLUE",
			"GARY",
			"JOHN",
		}
		menu.NewSelectMenu(elm, 0, 0, width, height, true, true)
	}
}

func chooseRivalName2() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()

	switch {
	case pressed.A:
		m.Close()
		switch m.Item() {
		case "NEW NAME":
			store.SetScriptID(store.CustomRivalName)
		default:
			store.Rival.Name = m.Item()
			store.SetScriptID(store.AfterChooseRivalName)
		}
	}
}

func customRivalName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 15:
		util.WhiteScreen(store.TileMap)
	case counter == 15:
		reset = true
		widget.DrawNameScreen(widget.RivalName)
		store.SetScriptID(store.WidgetRivalNamingScreen)
	}
}

func widgetRivalNamingScreen() {
	name, ok := handleNamingScreen()
	if ok {
		store.Rival.Name = name
		store.SetScriptID(store.AfterCustomRivalName)
	}
}

func afterChooseRivalName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter == 0:
		util.WhiteScreen(store.TileMap)
	case counter < 22:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(104-(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.rival[5], x, centerY)
	case counter == 22:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.HisNameIsText, false)
		store.PushScriptID(store.LetsGoPlayer)
	}
}

func afterCustomRivalName() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	util.WhiteScreen(store.TileMap)
	switch {
	case counter == 18:
		reset = true
		util.DrawImage(store.TileMap, lectureImage.rival[5], 7, centerY)
		text.DoPrintTextScript(text.TextBoxImage, txt.HisNameIsText, false)
		store.PushScriptID(store.LetsGoPlayer)
	}
}

func letsGoPlayer() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter <= 10:
		util.WhiteScreen(store.TileMap)
	case counter <= 20:
		util.DrawImage(store.TileMap, lectureImage.red[2], centerX, centerY)
	case counter <= 30:
		util.DrawImage(store.TileMap, lectureImage.red[2], centerX, centerY)
	case counter <= 40:
		util.DrawImage(store.TileMap, lectureImage.red[1], centerX, centerY)
	case counter <= 50:
		util.DrawImage(store.TileMap, lectureImage.red[1], centerX, centerY)
	case counter <= 60:
		util.DrawImage(store.TileMap, lectureImage.red[0], centerX, centerY)
	case counter <= 70:
		util.DrawImage(store.TileMap, lectureImage.red[0], centerX, centerY)
	case counter == 80:
		reset = true
		text.DoPrintTextScript(text.TextBoxImage, txt.OakSpeechText3, false)
		store.PushScriptID(store.ShrinkPlayer)
	}
}

func shrinkPlayer() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter == 0:
		audio.PlaySound(audio.SFX_SHRINK)
	case counter < 5:
	case counter < 35:
		util.DrawImage(store.TileMap, lectureImage.redShrink[0], centerX, centerY)
	case counter < 65:
		util.DrawImage(store.TileMap, lectureImage.redShrink[1], centerX, centerY)
	case counter < 95:
		util.DrawImage(store.TileMap, lectureImage.redShrink[2], centerX, centerY)
	case counter < 125:
		util.DrawImage(store.TileMap, lectureImage.redSprite[0], centerX, centerY)
	case counter < 140:
		util.DrawImage(store.TileMap, lectureImage.redSprite[1], centerX, centerY)
	case counter < 155:
		util.DrawImage(store.TileMap, lectureImage.redSprite[2], centerX, centerY)
	case counter < 205:
		util.WhiteScreen(store.TileMap)
	case counter == 205:
		reset = true
		InitializeOverworld()
	}
}
