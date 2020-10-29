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
	result[store.WidgetNamingScreen] = widgetNamingScreen
	result[store.FadeOutToBlack] = fadeOutToBlack
	result[store.FadeOutToWhite] = fadeOutToWhite
	result[store.LoadMapData] = loadMapData
	result[store.TitleCopyright] = titleCopyright
	result[store.TitleBlank] = titleBlank
	result[store.TitleIntroScene] = titleIntroScene
	result[store.TitleWhiteOut] = titleWhiteOut
	result[store.TitlePokemonRed] = titlePokemonRed
	return result
}

// Current return current script
func Current() func() {
	s, ok := scriptMap[store.ScriptID()]
	if !ok {
		util.NotRegisteredError("scriptMap", store.ScriptID())
		return halt
	}
	return s
}

func halt() {}

func execText() {
	if len([]rune(text.CurText)) == 0 {
		store.SetScriptID(store.Overworld)
	}

	if text.InScroll {
		text.ScrollTextUpOneLine(text.Image)
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

	text.CurText = text.PlaceStringOneByOne(text.Image, text.CurText)
	if len([]rune(text.CurText)) == 0 {
		store.SetScriptID(store.Overworld)
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
		store.PopScriptID()
	}
}

func fadeOutToWhite() {
	if store.FadeCounter <= 0 {
		store.SetScriptID(store.Overworld)
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
		store.PopScriptID()
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
