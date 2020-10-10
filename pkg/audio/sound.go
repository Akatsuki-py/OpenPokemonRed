package audio

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/rakyll/statik/fs"

	_ "pokered/pkg/data/statik"
	"pokered/pkg/util"
)

const (
	SFX_GET_ITEM_2 uint = iota
	SFX_TINK
	SFX_HEAL_HP
	SFX_HEAL_AILMENT
	SFX_START_MENU
	SFX_PRESS_AB
	SFX_COLLISION
)

// WAV wav file
type WAV struct {
	stream *wav.Stream
	player *audio.Player
}

var soundMap = newSoundMap()

func newSoundMap() map[uint]*WAV {
	soundMap := map[uint]*WAV{}
	FS, _ := fs.New()

	soundMap[SFX_TINK] = newWav(FS, "/tink.wav")
	soundMap[SFX_START_MENU] = newWav(FS, "/start_menu.wav")
	soundMap[SFX_PRESS_AB] = newWav(FS, "/press_ab.wav")
	soundMap[SFX_COLLISION] = newWav(FS, "/collision.wav")

	return soundMap
}

func newWav(fs http.FileSystem, path string) *WAV {
	w := &WAV{}
	f, _ := fs.Open(path)
	defer f.Close()
	w.stream, _ = wav.Decode(audioContext, f)
	w.player, _ = audio.NewPlayer(audioContext, w.stream)
	return w
}

// PlaySound play sfx
func PlaySound(soundID uint) {
	sound, ok := soundMap[soundID]
	if !ok {
		util.NotRegisteredError("soundMap", soundID)
		return
	}
	if sound.player.IsPlaying() {
		sound.player.Seek(0)
	} else {
		sound.player.Seek(0)
		sound.player.Play()
	}
}

func closeSE(se *WAV) {
	se.stream.Close()
}
