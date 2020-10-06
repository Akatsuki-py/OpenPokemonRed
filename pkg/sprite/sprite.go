package sprite

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// Movment status
const (
	Uninitialized byte = iota
	OK
	Delay
	Movement
)

// NumSprites a number of sprites at current map
func NumSprites() uint {
	return uint(len(store.SpriteData))
}

// UpdateSprites update sprite data
func UpdateSprites() {
	for offset, s := range store.SpriteData {
		if s == nil || s.ID == 0 {
			continue
		}
		if s.ID == 1 {
			UpdatePlayerSprite()
			continue
		}
		UpdateNPCSprite(uint(offset))
	}
}

// UpdateSpriteImage update sprite image offset
func UpdateSpriteImage(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}
	length := len(s.VRAM.Images)
	if length == 1 {
		s.VRAM.Index = 0
		return
	}

	animCounter := s.AnimationFrame >> 2

	// ref:
	switch animCounter + uint(s.Direction) {

	// down
	case 0:
		s.VRAM.Index = 1
		if length == 4 {
			s.VRAM.Index = 0
		}
	case 1, 2, 3:
		s.VRAM.Index = 0
		if length == 4 {
			s.VRAM.Index = 0
		}

	// up
	case 4:
		s.VRAM.Index = 4
		if length == 4 {
			s.VRAM.Index = 1
		}
	case 5, 6, 7:
		s.VRAM.Index = 3
		if length == 4 {
			s.VRAM.Index = 1
		}

	case 8:
		s.VRAM.Index = 6
		if length == 4 {
			s.VRAM.Index = 2
		}
	case 9, 10, 11:
		s.VRAM.Index = 7
		if length == 4 {
			s.VRAM.Index = 2
		}

	case 12:
		s.VRAM.Index = 8
		if length == 4 {
			s.VRAM.Index = 3
		}
	case 13, 14, 15:
		s.VRAM.Index = 9
		if length == 4 {
			s.VRAM.Index = 3
		}
	}
}

// DisableSprite hide sprite
func DisableSprite(offset uint) {
	s := store.SpriteData[offset]
	s.VRAM.Index = -1
}

// MoveSprite forcely move sprite by movement data
// set wNPCMovementDirections
func MoveSprite(offset uint, movement []byte) {
	copy(NPCMovementDirections, movement)
	util.SetBit(store.D730, 0)
	joypad.JoyIgnore = joypad.ByteToInput(0xff)
}
