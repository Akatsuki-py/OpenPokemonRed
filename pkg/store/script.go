package store

const (
	Halt uint = iota
	Overworld
	ExecText
	WidgetStartMenu
	WidgetStartMenu2
	WidgetBag
	WidgetTrainerCard
	WidgetPlayerNamingScreen
	WidgetPartyMenu
	FadeOutToBlack
	FadeOutToWhite
	LoadMapData
	TitleCopyright
	TitleBlank
	TitleIntroScene
	TitleWhiteOut
	TitlePokemonRed
	TitleMenu
	TitleMenu2
	OakSpeech0
	OakSpeech1
	OakSpeech2
	OakSpeech3
	OakSpeech4
	OakSpeech5
	OakSpeech6
	OakSpeech7
	OakSpeech8
	OakSpeech9
)

type ScriptQueue struct {
	Buffer [10]interface{}
	Length int
}

var scriptQueue = ScriptQueue{
	Buffer: [10]interface{}{Overworld},
	Length: 0,
}

// ScriptID current script ID
func ScriptID() interface{} {
	if scriptQueue.Length == 0 {
		return Overworld
	}
	return scriptQueue.Buffer[0]
}

// ScriptLength return queue length of script ID
func ScriptLength() int {
	return scriptQueue.Length
}

// SetScriptID change script ID
func SetScriptID(id uint) {
	scriptQueue = ScriptQueue{
		Buffer: [10]interface{}{id},
		Length: 1,
	}
}

func SetOtScript(f func()) {
	scriptQueue = ScriptQueue{
		Buffer: [10]interface{}{f},
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

func PushOtScript(f func()) {
	if scriptQueue.Length == 10 {
		return
	}
	scriptQueue.Buffer[scriptQueue.Length] = f
	scriptQueue.Length++
}

func PopScript() {
	if scriptQueue.Length == 0 {
		return
	}
	newBuffer := [10]interface{}{}
	for i := 0; i < scriptQueue.Length; i++ {
		if i == 9 {
			break
		}
		newBuffer[i] = scriptQueue.Buffer[i+1]
	}
	scriptQueue.Buffer = newBuffer
	scriptQueue.Length--
}
