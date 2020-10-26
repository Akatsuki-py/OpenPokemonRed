package script

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

const (
	Halt uint = iota
	ExecText
	WidgetStartMenu
	WidgetStartMenu2
	WidgetBag
	WidgetTrainerCard
	WidgetNamingScreen
	FadeOutToBlack
	LoadMapData
	TitleCopyright
	TitleBlank
	TitlePokemonRed
)

var scriptQueue = Queue{
	Buffer: [10]uint{Halt},
	Length: 0,
}

// ID current script ID
func ID() uint {
	if scriptQueue.Length == 0 {
		return Halt
	}
	return scriptQueue.Buffer[0]
}

// SetID change script ID
func SetID(id uint) {
	scriptQueue = Queue{
		Buffer: [10]uint{id},
		Length: 1,
	}
}

// PushID change script ID
func PushID(id uint) {
	if scriptQueue.Length == 10 {
		return
	}
	scriptQueue.Buffer[scriptQueue.Length] = id
	scriptQueue.Length++
}

func PopID() {
	if scriptQueue.Length == 0 {
		return
	}
	newBuffer := [10]uint{}
	for i := 0; i < scriptQueue.Length; i++ {
		if i == 9 {
			break
		}
		newBuffer[i] = scriptQueue.Buffer[i+1]
	}
	scriptQueue.Buffer = newBuffer
	scriptQueue.Length--
}

// ScriptMap script ID -> script
var scriptMap = newScriptMap()

func newScriptMap() map[uint]func() {
	result := map[uint]func(){}
	result[Halt] = halt
	result[ExecText] = execText
	result[WidgetStartMenu] = widgetStartMenu
	result[WidgetStartMenu2] = widgetStartMenu2
	result[WidgetBag] = widgetBag
	result[WidgetTrainerCard] = widgetTrainerCard
	result[WidgetNamingScreen] = widgetNamingScreen
	result[FadeOutToBlack] = fadeOutToBlack
	result[LoadMapData] = loadMapData
	result[TitleCopyright] = titleCopyright
	result[TitleBlank] = titleBlank
	result[TitlePokemonRed] = titlePokemonRed
	return result
}

// Current return current script
func Current() func() {
	s, ok := scriptMap[ID()]
	if !ok {
		util.NotRegisteredError("scriptMap", ID())
		return halt
	}
	return s
}

func halt() {}

func execText() {
	if len([]rune(text.CurText)) == 0 {
		SetID(Halt)
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
		SetID(Halt)
	}
}

func fadeOutToBlack() {
	if store.FadeCounter <= 0 {
		SetID(Halt)
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
		PopID()
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

	SetID(Halt)
}
