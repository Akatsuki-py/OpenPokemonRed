package audio

import (
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sampleRate     = 44100
	stopSound  int = -1
)

const reloadFadeOut = 10

var audioContext = audio.NewContext(sampleRate)

// FadeOut control fadeout switch and counter
var FadeOut = struct {
	Control uint
	Counter uint
}{}

// NewMusicID Music ID played on current music fadeout is completed
var NewMusicID int

// LastMusicID Music ID played latest
var LastMusicID int

// FadeOutAudio fadeout process called in every vBlank
func FadeOutAudio() {
	preVolume := Volume
	defer func() {
		if CurMusic != nil && CurMusic.IsPlaying() && preVolume != Volume {
			CurMusic.SetVolume(float64(Volume) / 7)
		}
	}()

	if FadeOut.Control == 0 {
		if util.ReadBit(store.D72C, 1) {
			return
		}
		setVolumeMax()
	}

	// fade out
	if FadeOut.Counter > 0 {
		FadeOut.Counter--
		return
	}

	// counterReachedZero
	{
		FadeOut.Counter = reloadFadeOut

		// fadeOutComplete
		if Volume == 0 {
			// start next music
			FadeOut.Control, FadeOut.Counter = 0, 0
			stopMusic()
			PlayMusic(NewMusicID)
			return
		}

		decrementVolume()
	}
}
