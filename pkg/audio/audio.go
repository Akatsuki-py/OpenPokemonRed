package audio

import "github.com/hajimehoshi/ebiten/audio"

const (
	sampleRate = 44100
)

var (
	audioContext, _ = audio.NewContext(sampleRate)
)
