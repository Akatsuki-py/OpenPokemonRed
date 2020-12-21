package overworld

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/overworld/mscript"
	"pokered/pkg/palette"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

// ExecOverworld exec overworld loop
// ref: OverworldLoop
func ExecOverworld() {
	p := store.SpriteData[0]
	if p == nil {
		return
	}

	palette.LoadGBPal()

	if store.Flag.D736.InLedgeOrFishingAnim {
		sprite.HandleMidJump()
	}

	if p.WalkCounter > 0 {
		sprite.UpdateSprites()
		sprite.AdvancePlayerSprite()

		if p.WalkCounter == 0 {
			if (p.DeltaX + p.DeltaY) != 0 {
				store.Flag.Enable.NormalWarp = true
			}
		}
	} else {
		joypadOverworld()

		directionPressed := false
		switch {
		case joypad.JoyPressed.Start:
			audio.PlaySound(audio.SFX_START_MENU)
			store.SetScriptID(store.WidgetStartMenu)
			return
		case joypad.JoyPressed.A:
			if offset := sprite.GetFrontSpriteOrSign(0); offset > 0 {
				sprite.MakeNPCFacePlayer(uint(offset))
				DisplayDialogue(offset)
				return
			}
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
			if sprite.CollisionCheckForPlayer() {
				p.DeltaX, p.DeltaY = 0, 0
			}
		} else {
			sprite.UpdateSprites()
			p.RightHand = false
			return
		}
	}
	moveAhead()
}

// simulatedJoypad
func joypadOverworld() {
	p := store.SpriteData[0]
	p.DeltaX, p.DeltaY = 0, 0

	runMapScript()

	joypad.Joypad()

	if p.Direction == util.Down && sprite.IsStandingOnDoor(0) {
		joypad.JoyHeld = joypad.Input{Down: true}
		return
	}

	if len(p.Simulated) == 0 {
		return
	}
	if len(p.Simulated) == 1 && p.Simulated[0] == util.Stop {
		p.Simulated = []uint{}
		return
	}

	switch p.Simulated[0] {
	case util.Down:
		joypad.JoyHeld = joypad.Input{Down: true}
	case util.Up:
		joypad.JoyHeld = joypad.Input{Up: true}
	case util.Right:
		joypad.JoyHeld = joypad.Input{Right: true}
	case util.Left:
		joypad.JoyHeld = joypad.Input{Left: true}
	case util.Stop:
		joypad.JoyHeld = joypad.Input{}
	}
	if len(p.Simulated) > 1 {
		p.Simulated = p.Simulated[1:]
		return
	}
	p.Simulated = []uint{util.Stop}
}

// ref: RunMapScript
func runMapScript() {
	doBoulderAnimation()
	runNPCMovementScript()
	mscript.Run(world.CurWorld.MapID)
}

func doBoulderAnimation() {}

// ref: RunNPCMovementScript
func runNPCMovementScript() {
}

func moveAhead() {
	checkWarpsNoCollision()
}

// function to play a sound when changing maps
func playMapChangeSound() {
	_, tileID := world.GetTileID(8, 8)
	soundID := audio.SFX_GO_OUTSIDE
	if tileID == 0x0b {
		soundID = audio.SFX_GO_INSIDE
	}
	audio.PlaySound(soundID)
}

func loadWorldData(mapID, warpID int) {
	world.LoadWorldData(mapID)

	// ref: LoadDestinationWarpPosition
	if warpID >= 0 {
		warpTo := world.CurWorld.Object.WarpTos[warpID]
		p := store.SpriteData[0]
		p.MapXCoord, p.MapYCoord = warpTo.XCoord, warpTo.YCoord
	}
}

func DisplayDialogue(offset int) {
	texts, textID := world.CurWorld.Header.Text, offset
	text.DisplayTextID(text.TextBoxImage, texts, textID)
	store.SetScriptID(store.ExecText)
}
