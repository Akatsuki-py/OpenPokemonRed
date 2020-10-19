package audio

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"

	_ "pokered/pkg/data/statik"
	"pokered/pkg/store"
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
	SFX_LEDGE
	SFX_GO_INSIDE
	SFX_GO_OUTSIDE
)

// WAV wav file
type WAV struct {
	stream *wav.Stream
	player *audio.Player
}

var soundMap = newSoundMap()

func newSoundMap() map[uint]*WAV {
	soundMap := map[uint]*WAV{}

	soundMap[SFX_TINK] = newWav(store.FS, "/tink.wav")
	soundMap[SFX_START_MENU] = newWav(store.FS, "/start_menu.wav")
	soundMap[SFX_PRESS_AB] = newWav(store.FS, "/press_ab.wav")
	soundMap[SFX_COLLISION] = newWav(store.FS, "/collision.wav")
	soundMap[SFX_LEDGE] = newWav(store.FS, "/ledge.wav")
	soundMap[SFX_GO_INSIDE] = newWav(store.FS, "/go_inside.wav")
	soundMap[SFX_GO_OUTSIDE] = newWav(store.FS, "/go_outside.wav")

	return soundMap
}

func newWav(fs http.FileSystem, path string) *WAV {
	w := &WAV{}
	f, err := fs.Open(path)
	if err != nil {
		util.NotFoundFileError(path)
		return w
	}

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
