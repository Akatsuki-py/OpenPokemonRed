package game

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/script"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

func execOverworld() {
	p := store.SpriteData[0]
	if p == nil {
		return
	}

	if p.WalkCounter > 0 {
		sprite.UpdateSprites()
		sprite.AdvancePlayerSprite()
	} else {
		p.DeltaX, p.DeltaY = 0, 0
		joypad.Joypad()
		directionPressed := false
		switch {
		case joypad.JoyHeld.Start:
			audio.PlaySound(audio.SFX_START_MENU)
			script.SetID(script.WidgetStartMenu)
			return
		case joypad.JoyHeld.Down:
			p.DeltaY = 1
			p.Direction = util.Down
		case joypad.JoyHeld.Up:
			p.DeltaY = -1
			p.Direction = util.Up
		case joypad.JoyHeld.Right:
			p.DeltaX = 1
			p.Direction = util.Right
		case joypad.JoyHeld.Left:
			p.DeltaX = -1
			p.Direction = util.Left
		}

		h := joypad.JoyHeld
		directionPressed = h.Up || h.Down || h.Right || h.Left
		if directionPressed {
			p.WalkCounter = 16
			sprite.UpdateSprites()
			if sprite.CollisionCheckForPlayer() {
				p.DeltaX, p.DeltaY = 0, 0
			}
			sprite.AdvancePlayerSprite()
		} else {
			sprite.UpdateSprites()
			p.RightHand = false
			return
		}
	}
	moveAhead()
}

func moveAhead() {
	checkWarpsNoCollision()
}

// check if the player has stepped onto a warp after having not collided
// ref: CheckWarpsNoCollision
func checkWarpsNoCollision() {
	curWorld := world.CurWorld
	if len(curWorld.Object.Warps) == 0 {
		checkMapConnections()
		return
	}

	p := store.SpriteData[0]
	if p == nil {
		return
	}
	for _, w := range curWorld.Object.Warps {
		if p.MapXCoord == w.XCoord && p.MapYCoord == w.YCoord {

		}
	}
}

// ref: CheckMapConnections
func checkMapConnections() {}
