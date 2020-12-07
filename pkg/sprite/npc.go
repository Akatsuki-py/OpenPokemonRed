package sprite

import (
	"pokered/pkg/data/tilecoll"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

// movement byte 2
const (
	forceUpDown    byte = 0x01
	forceLeftRight byte = 0x02
)

var isNPCWalkFrame = true

// NPCMovementDirections used for scripted NPC
var NPCMovementDirections []byte

// UpdateNPCSprite update NPC sprite state
// ref: UpdateNonPlayerSprite
func UpdateNPCSprite(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}

	if len(s.Simulated) > 0 {
		DoScriptedNPCMovement(offset)
		return
	}
	updateNPCSprite(offset)
}

// DoScriptedNPCMovement update NPC sprite in "NPC movement script"
func DoScriptedNPCMovement(offset uint) {
	s := store.SpriteData[offset]

	switch s.MovmentStatus {
	case Movement:
		updateSpriteInWalkingAnimation(offset)
		return
	}

	direction := s.Simulated[0]
	if len(s.Simulated) > 1 {
		s.Simulated = s.Simulated[1:]
	} else {
		s.Simulated = []uint{}
	}

	var deltaX, deltaY int
	switch direction {
	case util.Up:
		deltaX, deltaY = 0, -1
	case util.Down:
		deltaX, deltaY = 0, 1
	case util.Left:
		deltaX, deltaY = -1, 0
	case util.Right:
		deltaX, deltaY = 1, 0
	}

	s.Direction = direction

	s.WalkCounter = 16
	s.DeltaX, s.DeltaY = deltaX, deltaY

	s.MapXCoord += deltaX
	s.MapYCoord += deltaY

	s.MovmentStatus = Movement
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
	case util.Walk, util.Stay:
		// take random movement
		move = util.Random()
	default:
		// scripted NPC
		s.MovementBytes[0]++

		move = byte(util.Stay)
		if int(s.MovementBytes[0]) < len(NPCMovementDirections) {
			move = NPCMovementDirections[s.MovementBytes[0]]
		}

		switch move {
		case util.ChangeDirection:
			// TODO: ChangeFacingDirection
		case util.Stay:
			s.MovementBytes[0] = util.Stay
			store.Flag.D730.IsNPCScripted = false
			// TODO: [wSimulatedJoypadStatesIndex] = 0
			return
		case util.Walk:
			move = util.Random()
		}
	}

	// determine NPC direction
	var direction util.Direction
	switch s.MovementBytes[1] {
	case byte(util.Down):
		direction = util.Down
	case byte(util.Up):
		direction = util.Up
	case byte(util.Left):
		direction = util.Left
	case byte(util.Right):
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

	var deltaX, deltaY int
	switch direction {
	case util.Up:
		deltaX, deltaY = 0, -1
	case util.Down:
		deltaX, deltaY = 0, 1
	case util.Left:
		deltaX, deltaY = -1, 0
	case util.Right:
		deltaX, deltaY = 1, 0
	}

	tryWalking(offset, direction, deltaX, deltaY)
}

// tryWalking UpdateNPCSprite から呼び出される
func tryWalking(offset uint, direction util.Direction, deltaX, deltaY int) bool {
	s := store.SpriteData[offset]
	s.Direction = direction

	if collisionCheckForNPC(offset) {
		s.Delay = uint(util.Random())
		s.MovmentStatus = Delay
		s.DeltaX, s.DeltaY = 0, 0
		return false
	}

	s.WalkCounter = 16
	s.DeltaX, s.DeltaY = deltaX, deltaY

	s.MapXCoord += deltaX
	s.MapYCoord += deltaY

	s.MovmentStatus = Movement
	return true
}

func initializeSpriteStatus(offset uint) {
	s := store.SpriteData[offset]
	s.MovmentStatus = OK
	if s.Hidden {
		DisableSprite(offset)
	}
}

func checkSpriteAvailability(offset uint) bool {
	s := store.SpriteData[offset]
	if s.Hidden {
		DisableSprite(offset)
		return false
	}

	// disable sprite when it is out of screen
	if s.MovementBytes[0] >= util.Walk {
		tooLeft := s.ScreenXPixel < 0
		tooRight := s.ScreenXPixel > 160
		tooUp := s.ScreenYPixel > 144
		tooDown := s.ScreenYPixel < 0
		if tooLeft || tooRight || tooUp || tooDown {
			DisableSprite(offset)
			return false
		}
		UpdateSpriteImage(offset)
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
		if s.Delay == 0 {
			s.MovmentStatus = OK
		}
	default:
		s.Delay = 0
		s.MovmentStatus = OK
	}
	notYetMoving(offset)
}

// increment animation counter
func updateSpriteInWalkingAnimation(offset uint) {
	defer func() {
		isNPCWalkFrame = !isNPCWalkFrame
	}()
	if !isNPCWalkFrame {
		return
	}

	s := store.SpriteData[offset]
	s.ScreenXPixel += s.DeltaX
	s.ScreenYPixel += s.DeltaY

	s.WalkCounter--
	s.AnimationFrame++
	if s.AnimationCounter() == 4 {
		s.AnimationFrame = 0
	}
	if s.WalkCounter != 0 {
		return
	}

	if s.MovementBytes[0] < 0xfe {
		s.MovmentStatus = OK
		return
	}

	s.Delay = uint(util.Random())
	s.MovmentStatus = Delay
	s.DeltaX, s.DeltaY = 0, 0
}

func notYetMoving(offset uint) {
	s := store.SpriteData[offset]
	s.AnimationFrame %= 4
}

// MakeNPCFacePlayer make NPC face player when player talk to NPC
func MakeNPCFacePlayer(offset uint) {
	// SS.anne's captain
	if store.Flag.D72D.DontFacePlayer {
		notYetMoving(offset)
		return
	}

	p, s := store.SpriteData[0], store.SpriteData[offset]

	// sign
	if s == nil {
		return
	}

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
		s.VRAM.Index = int(s.Direction + s.AnimationCounter())
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

// collisionCheckForNPC check if collision occurs in npc moving ahead
// ref: CanWalkOntoTile
func collisionCheckForNPC(offset uint) bool {
	if store.IsInvalidSprite(offset) {
		return false
	}

	collision := false
	npc := store.SpriteData[offset]

	// if movement byte 1 is STAY(0xff)
	if npc.MovementBytes[0] == util.Stay {
		return true
	}

	for o, s := range store.SpriteData {
		if o == int(offset) {
			continue
		}
		if store.IsInvalidSprite(uint(o)) {
			break
		}

		switch npc.Direction {
		case util.Up:
			if npc.MapXCoord == s.MapXCoord && npc.MapYCoord-1 == s.MapYCoord {
				collision = true
			}
		case util.Down:
			if npc.MapXCoord == s.MapXCoord && npc.MapYCoord+1 == s.MapYCoord {
				collision = true
			}
		case util.Left:
			if npc.MapXCoord-1 == s.MapXCoord && npc.MapYCoord == s.MapYCoord {
				collision = true
			}
		case util.Right:
			if npc.MapXCoord+1 == s.MapXCoord && npc.MapYCoord == s.MapYCoord {
				collision = true
			}
		}

		if collision {
			break
		}
	}

	tilesetID, tileID := world.FrontTileID(npc.MapXCoord, npc.MapYCoord, npc.Direction)
	if tileID >= 0 && !util.Contains(tilecoll.Get(tilesetID), byte(uint(tileID))) {
		collision = true
	}

	return collision
}
