package audio

import (
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten/audio"
)

const (
	sampleRate     = 44100
	STOP_SOUND int = -1
)

var audioContext, _ = audio.NewContext(sampleRate)

var FadeOut = struct {
	Control uint
	Counter uint
	Reload  uint
}{}
var NewMusicID int

func FadeOutAudio() {
	if FadeOut.Control == 0 {
		if util.ReadBit(store.D72C, 1) {
			return
		}
		offVolume()
	}

	// fade out
	if FadeOut.Counter > 0 {
		FadeOut.Counter--
		return
	}

	// counterReachedZero
	FadeOut.Counter = FadeOut.Reload
}
