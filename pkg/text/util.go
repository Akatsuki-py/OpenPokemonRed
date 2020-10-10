package text

import (
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// PlacePlayTime place play time
func PlacePlayTime(target *ebiten.Image, hours, minutes uint, colonX, colonY util.Tile) {
	// hours
	delta := 1
	switch {
	case hours >= 100:
		delta = 3
	case hours >= 10:
		delta = 2
	}
	PlaceUintAtOnce(target, hours, colonX-delta, colonY)

	// minutes
	switch {
	case minutes >= 10:
		PlaceUintAtOnce(target, minutes, colonX+1, colonY)
	default:
		PlaceUintAtOnce(target, 0, colonX+1, colonY)
		PlaceUintAtOnce(target, minutes, colonX+2, colonY)
	}
}
