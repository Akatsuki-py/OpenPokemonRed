package store

import (
	"github.com/hajimehoshi/ebiten"
)

type MovmentStatus uint

const (
	Uninitialized MovmentStatus = iota
	OK
	Delay
	Movement
)

// SpriteData wSpriteStateData1, wSpriteStateData2
var SpriteData [16]*Sprite

// Sprite data
type Sprite struct {
	ID                         uint          // C1x0 0:none 1:player 2~:others
	MovmentStatus              MovmentStatus // C1x1
	ScreenXPixel, ScreenYPixel int           // Pixel C1x4, C1x5
	AnimationFrame             uint          // C1x7, C1x8
	Direction                  Direction     // C1x9
	WalkAnimationCounter       uint          // C2x0
	MapXCoord, MapYCoord       int           // Coord C2x4, C2x5
	Delay                      uint          // C2x8
	Images                     struct {
		Index int // C1x2
		VRAM  []*ebiten.Image
	}
}
