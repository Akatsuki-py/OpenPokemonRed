package widget

import (
	"image/png"
	"pokered/pkg/store"
	"pokered/pkg/util"

	_ "pokered/pkg/data/statik"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

var trainerCard *ebiten.Image

const (
	tcPath string = "/trainercard.png"
)

// InitTrainerCard initialize trainer card gfx data
func InitTrainerCard() {
	FS, _ := fs.New()
	f, err := FS.Open(tcPath)
	if err != nil {
		util.NotFoundFileError(tcPath)
		return
	}
	defer f.Close()

	img, _ := png.Decode(f)
	trainerCard, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

// DrawTrainerCard draw trainer card on screen
func DrawTrainerCard() {
	if trainerCard == nil {
		return
	}
	util.DrawImage(store.TileMap, trainerCard, 0, 0)
}

// ExitTrainerCard release trainer card gfx data
func ExitTrainerCard() {
	trainerCard = nil
}
