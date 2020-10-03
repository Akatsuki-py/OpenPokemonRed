package audio

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/rakyll/statik/fs"
)

const (
	MUSIC_PALLET_TOWN uint = iota
	MUSIC_FINAL_BATTLE
)

type Music struct {
	MP3   *mp3.Stream
	intro float64
}

// MusicMap MusicID -> Music
var MusicMap = newMusicMap()

func newMusicMap() map[uint]Music {
	musicMap := map[uint]Music{}
	FS, _ := fs.New()
	musicMap[MUSIC_PALLET_TOWN] = newMusic(FS, "/1-02 Pallet Town Theme.mp3", getMS(0, 32, 167))
	musicMap[MUSIC_FINAL_BATTLE] = newMusic(FS, "/1-43 Final Battle! (Rival).mp3", getMS(1, 15, 120))
	return musicMap
}

func getMS(min, sec, ms uint) float64 {
	return float64(min)*60 + float64(sec) + float64(ms)/1000
}

func newMusic(fs http.FileSystem, path string, intro float64) Music {
	f, err := fs.Open(path)
	if err != nil {
		return Music{}
	}
	defer f.Close()
	stream, err := mp3.Decode(audioContext, f)
	if err != nil {
		return Music{}
	}
	return Music{MP3: stream, intro: intro}
}

// PlayMusic play BGM
func PlayMusic(id uint) {
	m := MusicMap[id]
	intro := int64(m.intro * 4 * sampleRate)
	l := audio.NewInfiniteLoopWithIntro(m.MP3, intro, m.MP3.Length())
	p, _ := audio.NewPlayer(audioContext, l)
	go p.Play()
}
