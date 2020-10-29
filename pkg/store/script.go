package store

const (
	Halt uint = iota
	Overworld
	ExecText
	WidgetStartMenu
	WidgetStartMenu2
	WidgetBag
	WidgetTrainerCard
	WidgetNamingScreen
	FadeOutToBlack
	FadeOutToWhite
	LoadMapData
	TitleCopyright
	TitleBlank
	TitleIntroScene
	TitleWhiteOut
	TitlePokemonRed
)

type ScriptQueue struct {
	Buffer [10]uint
	Length int
}

var scriptQueue = ScriptQueue{
	Buffer: [10]uint{Overworld},
	Length: 0,
}

// ScriptID current script ID
func ScriptID() uint {
	if scriptQueue.Length == 0 {
		return Overworld
	}
	return scriptQueue.Buffer[0]
}

// SetScriptID change script ID
func SetScriptID(id uint) {
	scriptQueue = ScriptQueue{
		Buffer: [10]uint{id},
		Length: 1,
	}
}

// PushScriptID change script ID
func PushScriptID(id uint) {
	if scriptQueue.Length == 10 {
		return
	}
	scriptQueue.Buffer[scriptQueue.Length] = id
	scriptQueue.Length++
}

func PopScriptID() {
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
