package audio

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"

	_ "pokered/pkg/data/statik"
	"pokered/pkg/data/worldmap/song"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

type Music struct {
	Ogg    *vorbis.Stream
	intro  float64
	player *audio.Player
}

// MusicMap MusicID -> Music
var MusicMap = newMusicMap()

// CurMusic music data played now
var CurMusic *audio.Player

func newMusicMap() map[int]*Music {
	musicMap := map[int]*Music{}
	musicMap[song.MUSIC_INTRO_TITLE] = newMusic(store.FS, "/1-00 Intro.ogg", "0:00.000")         // 11.771
	musicMap[song.MUSIC_TITLE_SCREEN] = newMusic(store.FS, "/1-01 Title Screen.ogg", "0:04.520") // 0:59.695
	musicMap[song.MUSIC_PALLET_TOWN] = newMusic(store.FS, "/1-02 Pallet Town Theme.ogg", "0:08.040")
	musicMap[song.MUSIC_MEET_PROF_OAK] = newMusic(store.FS, "/1-03 Professor Oak.ogg", "0:13.560")
	musicMap[song.MUSIC_OAKS_LAB] = newMusic(store.FS, "/1-04 Professor Oak's Laboratory.ogg", "0:04.389")
	musicMap[song.MUSIC_MEET_RIVAL] = newMusic(store.FS, "/1-05 A Rival Appears.ogg", "0:02.277")
	musicMap[song.MUSIC_ROUTES1] = newMusic(store.FS, "/1-06 Road To Viridian City_ Leaving Pallet Town.ogg", "0:00.000")
	musicMap[song.MUSIC_WILD_BATTLE] = newMusic(store.FS, "/1-07 Battle! (Wild Pokémon).ogg", "0:15.657")
	musicMap[song.MUSIC_ROUTES2] = newMusic(store.FS, "/1-21 To Bill_ Leaving Cerulean City.ogg", "0:00.000")
	musicMap[song.MUSIC_FINAL_BATTLE] = newMusic(store.FS, "/1-43 Final Battle! (Rival).ogg", "1:15.120")
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

func newMusic(fs http.FileSystem, path string, intro string) *Music {
	f, err := fs.Open(path)
	if err != nil {
		util.NotFoundFileError(path)
		return nil
	}
	defer f.Close()

	stream, err := vorbis.Decode(audioContext, f)
	if err != nil {
		return nil
	}
	return &Music{Ogg: stream, intro: parseTime(intro)}
}

// PlayMusic play BGM
func PlayMusic(id int) {
	if CurMusic != nil && CurMusic.IsPlaying() {
		stopMusic()
	}

	if id == stopSound {
		return
	}

	m, ok := MusicMap[id]
	if !ok {
		util.NotRegisteredError("MusicMap", id)
		return
	}
	LastMusicID = id

	if m.player == nil {
		intro := int64(m.intro * 4 * sampleRate)
		l := audio.NewInfiniteLoopWithIntro(m.Ogg, intro, m.Ogg.Length())
		m.player, _ = audio.NewPlayer(audioContext, l)
	}
	CurMusic = m.player
	go CurMusic.Play()
}

// StopMusic stop BGM with fadeout
func StopMusic(fadeout uint) {
	FadeOut.Control = fadeout
	NewMusicID = stopSound
}

// StopMusicImmediately stop BGM
func StopMusicImmediately() {
	if CurMusic != nil {
		CurMusic.Pause()
		CurMusic.Seek(0)
	}
}

func stopMusic() {
	if CurMusic != nil {
		CurMusic.Pause()
		CurMusic.Seek(0)
	}
}

// PlayDefaultMusic play BGM accroding to Player's state
// ref: PlayDefaultMusic
func PlayDefaultMusic(mapID int) {
	musicID := song.MapMusics[mapID]
	switch store.Player.State {
	case store.BikeState:
	case store.SurfState:
	}

	if musicID == LastMusicID {
		return
	}

	PlayMusic(musicID)
	LastMusicID = 0
}
