package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

// ScriptMap script ID -> script
var scriptMap = newScriptMap()

func newScriptMap() map[uint]func() {
	result := map[uint]func(){}
	result[store.Overworld] = halt
	result[store.ExecText] = execText
	result[store.WidgetStartMenu] = widgetStartMenu
	result[store.WidgetStartMenu2] = widgetStartMenu2
	result[store.WidgetBag] = widgetBag
	result[store.WidgetTrainerCard] = widgetTrainerCard
	result[store.WidgetPlayerNamingScreen] = widgetPlayerNamingScreen
	result[store.WidgetRivalNamingScreen] = widgetRivalNamingScreen
	result[store.WidgetPartyMenu] = widgetPartyMenu
	result[store.FadeOutToBlack] = fadeOutToBlack
	result[store.FadeOutToWhite] = fadeOutToWhite
	result[store.LoadMapData] = loadMapData
	result[store.TitleCopyright] = titleCopyright
	result[store.TitleBlank] = titleBlank
	result[store.TitleIntroScene] = titleIntroScene
	result[store.TitleWhiteOut] = titleWhiteOut
	result[store.TitlePokemonRed] = titlePokemonRed
	result[store.TitleMenu] = titleMenu
	result[store.TitleMenu2] = titleMenu2
	result[store.OakSpeech0] = oakSpeech0
	result[store.OakSpeech1] = oakSpeech1
	result[store.OakSpeech2] = oakSpeech2
	result[store.IntroducePlayer] = introducePlayer
	result[store.ChoosePlayerName] = choosePlayerName
	result[store.ChoosePlayerName2] = choosePlayerName2
	result[store.CustomPlayerName] = customPlayerName
	result[store.AfterChoosePlayerName] = afterChoosePlayerName
	result[store.AfterCustomPlayerName] = afterCustomPlayerName
	result[store.IntroduceRival] = introduceRival
	result[store.ChooseRivalName] = chooseRivalName
	result[store.ChooseRivalName2] = chooseRivalName2
	result[store.CustomRivalName] = customRivalName
	result[store.AfterChooseRivalName] = afterChooseRivalName
	result[store.AfterCustomRivalName] = afterCustomRivalName
	result[store.LetsGoPlayer] = letsGoPlayer
	result[store.ShrinkPlayer] = shrinkPlayer
	return result
}

// Current return current script
func Current() func() {
	sid := store.ScriptID()

	switch s := sid.(type) {
	case int:
		sc, ok := scriptMap[uint(s)]
		if !ok {
			util.NotRegisteredError("scriptMap", store.ScriptID())
			return halt
		}
		return sc
	case uint:
		sc, ok := scriptMap[s]
		if !ok {
			util.NotRegisteredError("scriptMap", store.ScriptID())
			return halt
		}
		return sc
	case func():
		return func() {
			s()
			nextScript()
		}
	default:
		return scriptMap[0]
	}
}

func nextScript() {
	if store.ScriptLength() > 1 {
		store.PopScript()
		return
	}
	store.SetScriptID(store.Overworld)
}

func halt() {}

func execText() {
	if len([]rune(text.CurText)) == 0 {
		nextScript()
	}

	if text.InScroll {
		text.ScrollTextUpOneLine(text.TextBoxImage)
		return
	}

	if store.FrameCounter > 0 {
		joypad.Joypad()
		if joypad.JoyHeld.A || joypad.JoyHeld.B {
			store.FrameCounter = 0
			return
		}
		store.FrameCounter--
		if store.FrameCounter > 0 {
			store.DelayFrames = 1
			return
		}
		return
	}

	text.CurText = text.PlaceStringOneByOne(text.TextBoxImage, text.CurText)
	if len([]rune(text.CurText)) == 0 {
		nextScript()
	}
}

func fadeOutToBlack() {
	if store.FadeCounter <= 0 {
		store.SetScriptID(store.Overworld)
		return
	}

	store.FadeCounter--

	if store.Palette < 1 {
		store.Palette = 1
		return
	}

	store.Palette--
	store.DelayFrames = 8

	if store.FadeCounter <= 0 {
		store.PopScript()
	}
}

func fadeOutToWhite() {
	if store.FadeCounter <= 0 {
		nextScript()
		return
	}

	store.FadeCounter--

	if store.Palette > 8 {
		store.Palette = 8
		return
	}

	store.Palette++
	store.DelayFrames = 8

	if store.FadeCounter <= 0 {
		nextScript()
	}
}

func loadMapData() {
	mapID, warpID := world.WarpTo[0], world.WarpTo[1]
	if mapID < 0 {
		return
	}
	world.LoadWorldData(mapID)

	// ref: LoadDestinationWarpPosition
	if warpID < 0 {
		return
	}
	warpTo := world.CurWorld.Object.WarpTos[warpID]
	p := store.SpriteData[0]
	p.MapXCoord, p.MapYCoord = warpTo.XCoord, warpTo.YCoord

	store.SetScriptID(store.Overworld)
}
