package widget

import (
	"fmt"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/pkmn"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var statusScreen *ebiten.Image
var targetMon store.BoxMon

const (
	statusFrame  string = "/status_screen1.png"
	statusFrame2 string = "/status_screen2.png"
)

// DrawHP draw HP bar and HP value
// ref: DrawHP_
func DrawHP(target *ebiten.Image, hp, maxHP uint, x, y util.Tile, isRight bool) {
	px := 0
	if maxHP > 0 {
		px = int(hp * 48 / maxHP)
	}
	DrawHPBar(target, px, x, y)

	if isRight {
		x += 9
	} else {
		x++
		y++
	}

	// hp
	hpStr := fmt.Sprintf("%d", hp)
	if hp < 100 {
		hpStr = " " + hpStr
	}
	text.PlaceStringAtOnce(target, hpStr, x, y)

	// /
	text.PlaceStringAtOnce(target, "/", x+3, y)

	// mapHP
	maxHPStr := fmt.Sprintf("%d", maxHP)
	if maxHP < 100 {
		maxHPStr = " " + maxHPStr
	}
	text.PlaceStringAtOnce(target, maxHPStr, x+4, y)
}

// DrawHPBar draw HP bar and HP gauge
func DrawHPBar(target *ebiten.Image, px int, x, y util.Tile) {
	hpBar := util.OpenImage(store.FS, "/hp_bar.png")
	util.DrawImage(target, hpBar, x, y)

	r, g, b := util.GrayBlack[0], util.GrayBlack[1], util.GrayBlack[2]
	hpDotY := y*8 + 3
	for i := 0; i < px; i++ {
		hpDotX := x*8 + 16 + i
		util.DrawPixel(target, hpDotX, hpDotY, r, g, b)
		util.DrawPixel(target, hpDotX, hpDotY+1, r, g, b)
	}
}

// PrintLevel print mon level with :L
func PrintLevel(target *ebiten.Image, level uint, x, y util.Tile) {
	if level >= 100 {
		text.PlaceUintAtOnce(target, level, x, y)
		return
	}

	text.PlaceChar(target, ":L", x, y)
	text.PlaceUintAtOnce(target, level, x+1, y)
}

// InitStatusScreen init status screen
func InitStatusScreen(offset int) {
	targetMon = *store.PartyMons[offset].BoxMon
	statusScreen = util.NewImage()
	// audio.ReduceVolume()
	util.WhiteScreen(statusScreen)
}

func RenderStatusScreen1() {
	mon := targetMon
	frame := util.OpenImage(store.FS, statusFrame)
	util.DrawImage(statusScreen, frame, 0, 0)

	DrawHP(statusScreen, mon.HP, mon.HP, 11, 3, false)
	status := mon.Status.String()
	if len(status) == 0 {
		status = "OK"
	}
	text.PlaceStringAtOnce(statusScreen, status, 16, 6)
	PrintLevel(statusScreen, mon.BoxLevel, 14, 2)
	text.PlaceUintAtOnce(statusScreen, mon.ID, 3, 7)
	type1 := mon.Type[0]
	if type1 > 0 {
		text.PlaceStringAtOnce(statusScreen, pkmnd.TypeString(type1), 11, 10)
	}
	type2 := mon.Type[1]
	if type2 > 0 {
		text.PlaceStringAtOnce(statusScreen, pkmnd.TypeString(type2), 11, 12)
	}

	text.PlaceStringAtOnce(statusScreen, mon.Nick, 9, 1)
	text.PlaceStringAtOnce(statusScreen, mon.OTName, 12, 16)
	text.PlaceStringAtOnce(statusScreen, util.Padding(mon.OTID, 5, "0"), 12, 14)

	base := pkmnd.BaseStatsGen1(mon.ID)
	atk := pkmn.CalcStat(base.Attack, mon.DVs.Attack, mon.EVs.Attack, mon.BoxLevel)
	text.PlaceStringAtOnce(statusScreen, util.Padding(atk, 3, " "), 6, 10)
	def := pkmn.CalcStat(base.Defense, mon.DVs.Defense, mon.EVs.Defense, mon.BoxLevel)
	text.PlaceStringAtOnce(statusScreen, util.Padding(def, 3, " "), 6, 12)
	spd := pkmn.CalcStat(base.Speed, mon.DVs.Speed, mon.EVs.Speed, mon.BoxLevel)
	text.PlaceStringAtOnce(statusScreen, util.Padding(spd, 3, " "), 6, 14)
	sp := pkmn.CalcStat(base.Special, mon.DVs.SpAtk, mon.EVs.SpAtk, mon.BoxLevel)
	text.PlaceStringAtOnce(statusScreen, util.Padding(sp, 3, " "), 6, 16)
}

func RenderPokemonAndCryOnStatusScreen1() {

}

func CloseStatusScreen() {
	statusScreen = nil
}
