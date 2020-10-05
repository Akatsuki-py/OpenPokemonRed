package store

import (
	"github.com/hajimehoshi/ebiten"
)

// SpriteData wSpriteStateData1, wSpriteStateData2
var SpriteData [16]*Sprite

// Sprite data
type Sprite struct {
	ID                         uint      // C1x0 0:none 1:player 2~:others
	MovmentStatus              uint      // C1x1
	ScreenXPixel, ScreenYPixel int       // Pixel C1x4, C1x5
	AnimationFrame             uint      // C1x7, C1x8 update on UpdateSprites
	Direction                  Direction // C1x9
	WalkAnimationCounter       int       // C2x0 本来はNPCのみ
	MapXCoord, MapYCoord       int       // Coord C2x4, C2x5
	Delay                      uint      // C2x8
	VRAM                       struct {
		Index  int // C1x2
		Images []*ebiten.Image
	}
	Scripted       bool
	MovementBytes  [2]byte // movement byte 1,2
	DeltaX, DeltaY int
}

func (s *Sprite) AnimationCounter() uint {
	return s.AnimationFrame >> 2
}
func (s *Sprite) intraAnimationCounter() uint {
	return s.AnimationFrame % 4
}
