package audio

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"

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
	SFX_INTRO_LUNGE
	SFX_INTRO_HIP
	SFX_INTRO_HOP
	SFX_INTRO_RAISE
	SFX_INTRO_CRASH
	SFX_INTRO_WHOOSH
	SFX_SLOTS_STOP_WHEEL
	SFX_SLOTS_REWARD
	SFX_SLOTS_NEW_SPIN
	SFX_SHOOTING_STAR

	SFX_CRY_NIDORINO
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
	soundMap[SFX_INTRO_LUNGE] = newWav(store.FS, "/intro_lunge.wav")
	soundMap[SFX_INTRO_HIP] = newWav(store.FS, "/intro_hip.wav")
	soundMap[SFX_INTRO_HOP] = newWav(store.FS, "/intro_hop.wav")
	soundMap[SFX_INTRO_RAISE] = newWav(store.FS, "/intro_raise.wav")
	soundMap[SFX_INTRO_CRASH] = newWav(store.FS, "/intro_crash.wav")
	soundMap[SFX_INTRO_WHOOSH] = newWav(store.FS, "/intro_whoosh.wav")
	soundMap[SFX_SLOTS_STOP_WHEEL] = newWav(store.FS, "/slots_stop_wheel.wav")
	soundMap[SFX_SLOTS_REWARD] = newWav(store.FS, "/slots_reward.wav")
	soundMap[SFX_SLOTS_NEW_SPIN] = newWav(store.FS, "/slots_new_spin.wav")
	soundMap[SFX_SHOOTING_STAR] = newWav(store.FS, "/shooting_star.wav")

	// cry
	soundMap[SFX_CRY_NIDORINO] = newWav(store.FS, "/033Cry.wav")
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
