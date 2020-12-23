package widget

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/pkmn"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	dexPageFrame string = "/dexpage.png"
)

// ShowPokedexData ref: ShowPokedexData
func ShowPokedexData(target *ebiten.Image, id uint) {
	page := util.OpenImage(store.FS, dexPageFrame)
	util.DrawImage(target, page, 0, 0)

	monName := pkmnd.Name(id)
	text.PlaceStringAtOnce(target, monName, 9, 6)

	header := pkmnd.Header(id)
	text.PlaceStringAtOnce(target, header.DexEntry.Species, 9, 4)
	text.PlaceUintAtOnce(statusScreen, header.ID, 4, 8)

	pic := pkmn.Picture(targetMon.ID, true)
	util.DrawImage(statusScreen, pic, 1, 1)
	audio.Cry(targetMon.ID)

}
