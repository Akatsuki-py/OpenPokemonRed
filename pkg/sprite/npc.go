package sprite

import (
	"pokered/pkg/store"
	"pokered/pkg/util"
)

const (
	changeDirection byte = 0xe0
	walk            byte = 0xfe
	stay            byte = 0xff
)

// movement byte 2
const (
	forceDown      byte = 0xd0
	forceUp        byte = 0xd1
	forceLeft      byte = 0xd2
	forceRight     byte = 0xd3
	forceUpDown    byte = 0x01
	forceLeftRight byte = 0x02
)

var NPCMovementDirections []byte

// UpdateNPCSprite update NPC sprite state
// ref: UpdateNonPlayerSprite
func UpdateNPCSprite(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}

	if s.Scripted {
		DoScriptedNPCMovement(offset)
		return
	}
	updateNPCSprite(offset)
}

// DoScriptedNPCMovement update NPC sprite in "NPC movement script"
func DoScriptedNPCMovement(offset uint) {
	// TODO: implement
}

// If movement status is OK, try walking.
// ref: UpdateNPCSprite
func updateNPCSprite(offset uint) {
	s := store.SpriteData[offset]
	if s.MovmentStatus == Uninitialized {
		initializeSpriteStatus(offset)
		return
	}
	if !checkSpriteAvailability(offset) {
		return
	}

	if util.ReadBit(s.MovmentStatus, 7) {
		makeNPCFacePlayer(offset)
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

	// sprite status is OK
	p := store.SpriteData[0]
	if p != nil && p.WalkCounter > 0 {
		return
	}

	var move byte
	switch s.MovementBytes[0] {
	case walk, stay:
		// take random movement
		move = util.Random()
	default:
		// scripted NPC
		s.MovementBytes[0]++

		move = byte(stay)
		if int(s.MovementBytes[0]) < len(NPCMovementDirections) {
			move = NPCMovementDirections[s.MovementBytes[0]]
		}

		switch move {
		case changeDirection:
			// TODO: ChangeFacingDirection
		case stay:
			s.MovementBytes[0] = stay
			util.ResBit(store.D730, 0)
			// TODO: [wSimulatedJoypadStatesIndex] = 0
			return
		case walk:
			move = util.Random()
		}
	}

	// determine NPC direction
	var direction util.Direction
	switch s.MovementBytes[1] {
	case forceDown:
		direction = util.Down
	case forceUp:
		direction = util.Up
	case forceLeft:
		direction = util.Left
	case forceRight:
		direction = util.Right
	default:
		switch {
		case move < 0x40:
			direction = util.Down
			if s.MovementBytes[1] == forceLeftRight {
				direction = util.Left
			}
		case move < 0x80:
			direction = util.Up
			if s.MovementBytes[1] == forceLeftRight {
				direction = util.Right
			}
		case move < 0xc0:
			direction = util.Left
			if s.MovementBytes[1] == forceUpDown {
				direction = util.Up
			}
		default:
			direction = util.Right
			if s.MovementBytes[1] == forceUpDown {
				direction = util.Down
			}
		}
	}
	tryWalking(offset, direction)
}

// tryWalking UpdateNPCSprite から呼び出される
func tryWalking(offset uint, direction util.Direction) bool {
	return true
}

func initializeSpriteStatus(offset uint) {
	s := store.SpriteData[offset]
	s.MovmentStatus = OK
}

func checkSpriteAvailability(offset uint) bool {
	s := store.SpriteData[offset]
	// TODO: IsObjectHidden

	// disable sprite when it is out of screen
	if s.MovementBytes[0] >= walk {
		tooLeft := s.ScreenXPixel < 0
		tooRight := s.ScreenXPixel > 160
		tooUp := s.ScreenYPixel > 144
		tooDown := s.ScreenYPixel < 0
		if tooLeft || tooRight || tooUp || tooDown {
			DisableSprite(offset)
			return false
		}
		s.VRAM.Index = Index(offset)
	}

	// if player is in walk, disable sprite
	p := store.SpriteData[0]
	if p != nil && p.WalkCounter > 0 {
		return false
	}
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

	s.WalkCounter--
	if s.WalkCounter != 0 {
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
	case util.Up:
		s.Direction = util.Down
	case util.Down:
		s.Direction = util.Up
	case util.Left:
		s.Direction = util.Right
	case util.Right:
		s.Direction = util.Left
	}
	notYetMoving(offset)
}

func animScriptedNPCMovement(offset uint) {
	s := store.SpriteData[offset]

	switch s.Direction {
	case util.Up, util.Down, util.Left, util.Right:
		advanceScriptedNPCAnimFrameCounter(offset)
		s.VRAM.Index = int(s.Direction + (s.AnimationFrame >> 2))
	default:
		return
	}
}

func advanceScriptedNPCAnimFrameCounter(offset uint) {
	s := store.SpriteData[offset]
	s.AnimationFrame++
	if s.AnimationFrame>>2 == 4 {
		s.AnimationFrame = 0
	}
}
