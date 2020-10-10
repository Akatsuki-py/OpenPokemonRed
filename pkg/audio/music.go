package audio

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"

	_ "pokered/pkg/data/statik"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

const (
	MUSIC_PALLET_TOWN int = iota
	MUSIC_MEET_PROF_OAK
	MUSIC_FINAL_BATTLE
)

type Music struct {
	Ogg   *vorbis.Stream
	intro float64
}

// MusicMap MusicID -> Music
var MusicMap = newMusicMap()

// CurMusic music data played now
var CurMusic *audio.Player

func newMusicMap() map[int]Music {
	musicMap := map[int]Music{}
	musicMap[MUSIC_PALLET_TOWN] = newMusic(store.FS, "/1-02 Pallet Town Theme.ogg", "0:32.167")
	musicMap[MUSIC_MEET_PROF_OAK] = newMusic(store.FS, "/1-03 Professor Oak.ogg", "0:13.560")
	musicMap[MUSIC_FINAL_BATTLE] = newMusic(store.FS, "/1-43 Final Battle! (Rival).ogg", "1:15.120")
	return musicMap
}

func parseTime(t string) float64 {
	s := strings.Split(t, ":")
	if len(s) < 2 {
		return 0
	}

	minute, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		minute = 0
	}
	second, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		second = 0
	}
	return 60*minute + second
}

func newMusic(fs http.FileSystem, path string, intro string) Music {
	f, err := fs.Open(path)
	if err != nil {
		util.NotFoundFileError(path)
		return Music{}
	}
	defer f.Close()

	stream, err := vorbis.Decode(audioContext, f)
	if err != nil {
		return Music{}
	}
	return Music{Ogg: stream, intro: parseTime(intro)}
}

// PlayMusic play BGM
func PlayMusic(id int) {
	if id == stopSound {
		if CurMusic != nil && CurMusic.IsPlaying() {
			CurMusic.Close()
		}
		return
	}
	m, ok := MusicMap[id]
	if !ok {
		util.NotRegisteredError("MusicMap", id)
		return
	}
	intro := int64(m.intro * 4 * sampleRate)
	l := audio.NewInfiniteLoopWithIntro(m.Ogg, intro, m.Ogg.Length())
	p, _ := audio.NewPlayer(audioContext, l)
	CurMusic = p
	go p.Play()
}

// StopMusic stop BGM with fadeout
func StopMusic(fadeout uint) {
	FadeOut.Control = fadeout
	NewMusicID = stopSound
}

// StopMusicImmediately stop BGM
func StopMusicImmediately() {
	FadeOut.Control = 0
	NewMusicID = stopSound
	if CurMusic != nil {
		CurMusic.Close()
	}
}
