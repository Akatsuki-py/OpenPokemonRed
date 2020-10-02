package audio

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/rakyll/statik/fs"
)

const (
	MUSIC_PALLET_TOWN uint = iota
)

var MusicMap = newMusicMap()

func newMusicMap() map[uint]*mp3.Stream {
	musicMap := map[uint]*mp3.Stream{}
	FS, _ := fs.New()
	musicMap[MUSIC_PALLET_TOWN] = newMP3(FS, "/PalletTown.mp3")
	return musicMap
}

func newMP3(fs http.FileSystem, path string) *mp3.Stream {
	f, err := fs.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	stream, err := mp3.Decode(audioContext, f)
	if err != nil {
		return nil
	}
	return stream
}
