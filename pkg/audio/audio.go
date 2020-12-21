package audio

import (
	"pokered/pkg/store"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sampleRate     = 44100
	stopSound  int = -1
)

const reloadFadeOut = 10

var baseVolume = 0.025

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
	preVolume := volume
	defer func() {
		if CurMusic != nil && CurMusic.IsPlaying() && preVolume != volume {
			CurMusic.SetVolume(float64(volume) * baseVolume / 7)
		}
	}()

	if FadeOut.Control == 0 {
		if store.Flag.D72C.DisturbAudioFadeout {
			return
		}
		SetVolumeMax()
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
		if volume == 0 {
			// start next music
			FadeOut.Control, FadeOut.Counter = 0, 0
			stopMusic()
			PlayMusic(NewMusicID)
			return
		}

		decrementVolume()
	}
}
