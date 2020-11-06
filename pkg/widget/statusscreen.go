package widget

import (
	"fmt"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// DrawHP draw HP bar and HP value
// ref: DrawHP_
func DrawHP(target *ebiten.Image, hp, maxHP uint, x, y util.Tile, isRight bool) {
	px := int(hp * 48 / maxHP)
	DrawHPBar(target, px, x, y)

	if isRight {
		x += 9
	} else {
		x += 2
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
