package sprite

import "pokered/pkg/store"

func UpdateNonPlayerSprite(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}

	if s.Scripted {
		DoScriptedNPCMovement(s)
		return
	}
	UpdateNPCSprite(s)
}

func DoScriptedNPCMovement(s *store.Sprite) {}

func UpdateNPCSprite(s *store.Sprite) {
	if s.MovmentStatus == Uninitialized {
		initializeSpriteStatus(s)
		return
	}
	if !checkSpriteAvailability() {
		return
	}
}

// tryWalking UpdateNPCSprite から呼び出される
func tryWalking() bool {
	return true
}

func initializeSpriteStatus(s *store.Sprite) {
	s.MovmentStatus = OK
	s.VRAM.Index = -1
}

func checkSpriteAvailability() bool {
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

func updateSpriteInWalkingAnimation() {}

func notYetMoving(offset uint) {
	s := store.SpriteData[offset]
	s.AnimationFrame %= 4
	UpdateSpriteImage(offset)
}
