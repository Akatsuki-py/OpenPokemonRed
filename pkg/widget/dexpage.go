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
	text.PlaceStringAtOnce(dexPageScreen, util.PaddingLeft(header.ID, 3, "0"), 4, 8)

	// height
	text.PlaceStringAtOnce(dexPageScreen, util.PaddingLeft(header.DexEntry.HT[0], 2, " "), 12, 6)
	text.PlaceStringAtOnce(dexPageScreen, util.PaddingLeft(header.DexEntry.HT[0], 2, "0"), 15, 6)

	// weight
	wt0, wt1 := header.DexEntry.WT/10, header.DexEntry.WT%10
	text.PlaceStringAtOnce(dexPageScreen, util.PaddingLeft(wt0, 3, " "), 12, 8)
	text.PlaceStringAtOnce(dexPageScreen, ".", 15, 8)
	text.PlaceUintAtOnce(dexPageScreen, wt1, 16, 8)

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
	switch store.Player.Starter {
	case pkmnd.CHARMANDER:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabCharmanderText, true)
	case pkmnd.SQUIRTLE:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabSquirtleText, true)
	case pkmnd.BULBASAUR:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabBulbasaurText, true)
	}
}

func DexPageScreen() *ebiten.Image {
	return dexPageScreen
}
