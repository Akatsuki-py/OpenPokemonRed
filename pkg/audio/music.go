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
	MUSIC_PALLET_TOWN int = iota + 1
	MUSIC_POKECENTER
	MUSIC_GYM
	MUSIC_CITIES1
	MUSIC_CITIES2
	MUSIC_CELADON
	MUSIC_CINNABAR
	MUSIC_VERMILION
	MUSIC_LAVENDER
	MUSIC_SS_ANNE
	MUSIC_MEET_PROF_OAK
	MUSIC_MEET_RIVAL
	MUSIC_MUSEUM_GUY
	MUSIC_SAFARI_ZONE
	MUSIC_PKMN_HEALED
	MUSIC_ROUTES1
	MUSIC_ROUTES2
	MUSIC_ROUTES3
	MUSIC_ROUTES4
	MUSIC_INDIGO_PLATEAU
	MUSIC_GYM_LEADER_BATTLE
	MUSIC_TRAINER_BATTLE
	MUSIC_WILD_BATTLE
	MUSIC_FINAL_BATTLE
	MUSIC_DEFEATED_TRAINER
	MUSIC_DEFEATED_WILD_MON
	MUSIC_DEFEATED_GYM_LEADER
	MUSIC_TITLE_SCREEN
	MUSIC_CREDITS
	MUSIC_HALL_OF_FAME
	MUSIC_OAKS_LAB
	MUSIC_JIGGLYPUFF_SONG
	MUSIC_BIKE_RIDING
	MUSIC_SURFING
	MUSIC_GAME_CORNER
	MUSIC_INTRO_BATTLE
	MUSIC_DUNGEON1
	MUSIC_DUNGEON2
	MUSIC_DUNGEON3
	MUSIC_CINNABAR_MANSION
	MUSIC_POKEMON_TOWER
	MUSIC_SILPH_CO
	MUSIC_MEET_EVIL_TRAINER
	MUSIC_MEET_FEMALE_TRAINER
	MUSIC_MEET_MALE_TRAINER
	MUSIC_INTRO_TITLE
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
	musicMap[MUSIC_INTRO_TITLE] = newMusic(store.FS, "/1-00 Intro.ogg", "0:00.000")         // 11.771
	musicMap[MUSIC_TITLE_SCREEN] = newMusic(store.FS, "/1-01 Title Screen.ogg", "0:04.520") // 0:59.695
	musicMap[MUSIC_PALLET_TOWN] = newMusic(store.FS, "/1-02 Pallet Town Theme.ogg", "0:08.040")
	musicMap[MUSIC_MEET_PROF_OAK] = newMusic(store.FS, "/1-03 Professor Oak.ogg", "0:13.560")
	musicMap[MUSIC_OAKS_LAB] = newMusic(store.FS, "/1-04 Professor Oak's Laboratory.ogg", "0:04.389")
	musicMap[MUSIC_MEET_RIVAL] = newMusic(store.FS, "/1-05 A Rival Appears.ogg", "0:02.277")
	musicMap[MUSIC_ROUTES1] = newMusic(store.FS, "/1-06 Road To Viridian City_ Leaving Pallet Town.ogg", "0:00.000")
	musicMap[MUSIC_WILD_BATTLE] = newMusic(store.FS, "/1-07 Battle! (Wild Pokémon).ogg", "0:15.657")
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
	if id == stopSound {
		if CurMusic != nil && CurMusic.IsPlaying() {
			stopMusic()
		}
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
