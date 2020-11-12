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
		store.SetScriptID(store.ExecText)
		palette.GBFadeOutToWhite(true)
		store.PushScriptID(store.OakSpeech1)
		text.PrintText(text.TextBoxImage, txt.OakSpeechText1)
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
		x := int((168 - counter*8) / 8)
		util.DrawImage(store.TileMap, lectureImage.nidorino[0], x, centerY)
	case counter == 16:
		reset = true
		store.SetScriptID(store.ExecText)
		store.PushScriptID(store.OakSpeech2)
		text.PrintText(text.TextBoxImage, txt.OakSpeechText2A)
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
		store.SetScriptID(store.ExecText)
		store.PushScriptID(store.OakSpeech3)
		text.PrintText(text.TextBoxImage, txt.OakSpeechText2B)
	}
}

func oakSpeech3() {
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
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 16:
		reset = true
		store.SetScriptID(store.ExecText)
		store.PushScriptID(store.OakSpeech4)
		text.PrintText(text.TextBoxImage, txt.IntroducePlayerText)
	}
}

// ref: ChoosePlayerName
func oakSpeech4() {
	reset := false
	defer func() {
		if reset {
			counter = 0
			return
		}
		counter++
	}()

	switch {
	case counter < 18:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(56+(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 18:
		reset = true
		store.SetScriptID(store.OakSpeech5)

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

func oakSpeech5() {
	m := menu.CurSelectMenu()
	pressed := menu.HandleSelectMenuInput()

	switch {
	case pressed.A:
		m.Close()
		switch m.Item() {
		case "NEW NAME":
			store.SetScriptID(store.OakSpeech6)
		default:
			store.Player.Name = m.Item()
			store.SetScriptID(store.OakSpeech7)
		}
	}
}

func oakSpeech6() {
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

// after choose NAME
func oakSpeech7() {
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
	case counter < 19:
		util.ClearScreenArea(store.TileMap, 0, 4, 7, 20)
		x := int(96-(counter/3)*8) / 8
		util.DrawImage(store.TileMap, lectureImage.red[0], x, centerY)
	case counter == 19:
		reset = true
		store.SetScriptID(store.ExecText)
		text.PrintText(text.TextBoxImage, txt.YourNameIsText)
		store.PushScriptID(store.OakSpeech9)
	}
}

// after NEW NAME
func oakSpeech8() {
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
		store.SetScriptID(store.ExecText)
		text.PrintText(text.TextBoxImage, txt.YourNameIsText)
		store.PushScriptID(store.OakSpeech9)
	}
}

// introduce rival
func oakSpeech9() {
	palette.GBFadeOutToWhite(false)
	store.PushScriptID(store.OakSpeech10)
}

func oakSpeech10() {
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
		store.SetScriptID(store.ExecText)
		palette.GBFadeOutToWhite(true)
		text.PrintText(text.TextBoxImage, txt.IntroduceRivalText)
	}
}
