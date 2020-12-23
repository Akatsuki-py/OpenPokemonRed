package widget

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/data/txt"
	"pokered/pkg/pkmn"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"strings"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	dexPageFrame string = "/dexpage.png"
)

var dexPageScreen *ebiten.Image

// TargetMonID dex page mon
var TargetMonID uint

// ShowPokedexData ref: ShowPokedexData
func ShowPokedexData() {
	dexPageScreen = util.NewImage()
	page := util.OpenImage(store.FS, dexPageFrame)
	util.DrawImage(dexPageScreen, page, 0, 0)

	monName := strings.ToUpper(pkmnd.Name(TargetMonID))
	text.PlaceStringAtOnce(dexPageScreen, monName, 9, 2)

	header := pkmnd.Header(TargetMonID)
	text.PlaceStringAtOnce(dexPageScreen, header.DexEntry.Species, 9, 4)
	text.PlaceUintAtOnce(dexPageScreen, header.ID, 4, 8)

	pic := pkmn.Picture(TargetMonID, false)
	util.DrawImage(dexPageScreen, pic, 1, 1)
	audio.Cry(TargetMonID)
	text.DoPrintDexTextScript(dexPageScreen, header.DexEntry.Text, false)
}

// CloseDexPage close dex page
func CloseDexPage() {
	dexPageScreen = nil
}

// CloseStarterDexPage close dex page
func CloseStarterDexPage() {
	dexPageScreen = nil
	text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabReceivedMonText, false)
}

func DexPageScreen() *ebiten.Image {
	return dexPageScreen
}
