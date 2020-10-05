package sprite

import (
	"pokered/pkg/store"
	"pokered/pkg/util"
)

func UpdateNonPlayerSprite(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}

	if s.Scripted {
		DoScriptedNPCMovement(offset)
		return
	}
	UpdateNPCSprite(offset)
}

func DoScriptedNPCMovement(offset uint) {}

func UpdateNPCSprite(offset uint) {
	s := store.SpriteData[offset]
	if s.MovmentStatus == Uninitialized {
		initializeSpriteStatus(offset)
		return
	}
	if !checkSpriteAvailability(offset) {
		return
	}

	switch s.MovmentStatus {
	case Delay:
		updateSpriteMovementDelay(offset)
		return
	case Movement:
		updateSpriteInWalkingAnimation(offset)
		return
	}
}

// tryWalking UpdateNPCSprite から呼び出される
func tryWalking() bool {
	return true
}

func initializeSpriteStatus(offset uint) {
	s := store.SpriteData[offset]
	s.MovmentStatus = OK
	s.VRAM.Index = -1
}

func checkSpriteAvailability(offset uint) bool {
	s := store.SpriteData[offset]
	// TODO: IsObjectHidden

	// disable sprite when it is out of screen
	if s.MovementBytes[0] >= 0xfe {
		p := store.SpriteData[0]
		tooLeft := p.MapXCoord > s.MapXCoord
		tooRight := s.MapXCoord > p.MapXCoord+9
		tooUp := p.MapYCoord > s.MapYCoord
		tooDown := s.MapYCoord > p.MapYCoord+8
		if tooLeft || tooRight || tooUp || tooDown {
			DisableSprite(offset)
			return false
		}
	}

	// if player is in walk, disable sprite
	if WalkCounter > 0 {
		return false
	}

	UpdateSpriteImage(offset)
	return true
}

// update delay value (c2x8) for sprites in the delayed state (c1x1)
func updateSpriteMovementDelay(offset uint) {
	s := store.SpriteData[offset]
	movementByte1 := s.MovementBytes[0]
	switch movementByte1 {
	case 0xfe, 0xff:
		s.Delay--
	default:
		s.Delay = 0
		s.MovmentStatus = OK
	}
	notYetMoving(offset)
}

// increment animation counter
func updateSpriteInWalkingAnimation(offset uint) {
	s := store.SpriteData[offset]
	s.ScreenXPixel += s.DeltaX
	s.ScreenYPixel += s.DeltaY

	s.WalkAnimationCounter--
	if s.WalkAnimationCounter != 0 {
		return
	}

	if s.MovementBytes[0] < 0xfe {
		s.MovmentStatus = OK
		return
	}

	s.Delay = uint(util.Random() & 0x7f)
	s.MovmentStatus = Delay
	s.DeltaX, s.DeltaY = 0, 0
}

func notYetMoving(offset uint) {
	s := store.SpriteData[offset]
	s.AnimationFrame %= 4
	UpdateSpriteImage(offset)
}

// make NPC face player when player talk to NPC
func makeNPCFacePlayer(offset uint) {
	// D72D[5] is set on talking to SS.anne's captain
	if util.ReadBit(store.D72D, 5) {
		notYetMoving(offset)
		return
	}

	p, s := store.SpriteData[0], store.SpriteData[offset]
	switch p.Direction {
	case store.Up:
		s.Direction = store.Down
	case store.Down:
		s.Direction = store.Up
	case store.Left:
		s.Direction = store.Right
	case store.Right:
		s.Direction = store.Left
	}
	notYetMoving(offset)
}
