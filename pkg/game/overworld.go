package game

import (
	"pokered/pkg/joypad"
	"pokered/pkg/script"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/util"
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
		player.DeltaX, player.DeltaY = 0, 0
		joypad.Joypad()
		directionPressed := false
		switch {
		case joypad.JoyHeld.Start:
			script.SetScriptID(script.WidgetStartMenu)
			return
		case joypad.JoyHeld.Down:
			player.DeltaY = 1
			player.Direction = util.Down
		case joypad.JoyHeld.Up:
			player.DeltaY = -1
			player.Direction = util.Up
		case joypad.JoyHeld.Right:
			player.DeltaX = 1
			player.Direction = util.Right
		case joypad.JoyHeld.Left:
			player.DeltaX = -1
			player.Direction = util.Left
		}

		h := joypad.JoyHeld
		directionPressed = h.Up || h.Down || h.Right || h.Left
		if directionPressed {
			p.WalkCounter = 16
			sprite.UpdateSprites()
			sprite.AdvancePlayerSprite()
		} else {
			p.RightHand = false
		}
	}
}
