package sprite

import (
	"fmt"
	"image/png"
	"pokered/pkg/audio"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

// InitPlayer initialize player sprite
func InitPlayer(state uint) {
	FS, _ := fs.New()

	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		name := "red"
		switch state {
		case Cycling:
			name = "red_cycling"
		case Seel:
			name = "seel"
		}

		f, _ := FS.Open(fmt.Sprintf("/%s_%d.png", name, i))
		defer f.Close()
		img, _ := png.Decode(f)
		imgs[i], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	s := &store.Sprite{
		ID:           1,
		ScreenXPixel: 16 * util.PlayerX,
		ScreenYPixel: 16*util.PlayerY - 4,
		MapXCoord:    util.PlayerX,
		MapYCoord:    util.PlayerY,
		VRAM: store.SpriteImage{
			Index:  1,
			Images: imgs,
		},
	}
	store.SpriteData[0] = s
}

// ChangePlayerSprite change player sprite image
func ChangePlayerSprite(state uint) {
	p := store.SpriteData[0]
	if p == nil || p.ID == 0 {
		return
	}

	FS, _ := fs.New()

	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		name := "red"
		switch state {
		case Cycling:
			name = "red_cycling"
		case Seel:
			name = "seel"
		}

		f, _ := FS.Open(fmt.Sprintf("/%s_%d.png", name, i))
		defer f.Close()
		img, _ := png.Decode(f)
		imgs[i], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}
	p.VRAM.Images = imgs
}

// UpdatePlayerSprite update sprite direction and anim counter
// if in moving, increment anim counter
// if player is starting moving, change direction and increment anim counter
func UpdatePlayerSprite() {
	p := store.SpriteData[0]
	if p == nil || p.ID == 0 {
		return
	}

	if p.WalkCounter > 0 {
		p.AnimationFrame++
		if p.AnimationCounter() == 4 {
			p.AnimationFrame = 0
		}
	}
	p.VRAM.Index = int(p.Direction + (p.AnimationFrame >> 2))
}

// AdvancePlayerSprite advance player's walk by a frame
func AdvancePlayerSprite() {
	p := store.SpriteData[0]
	if p == nil || p.ID == 0 {
		return
	}
	p.WalkCounter--
	if p.WalkCounter == 0 {
		p.RightHand = !p.RightHand
		p.MapXCoord += p.DeltaX
		p.MapYCoord += p.DeltaY
	}

	store.SCX += p.DeltaX
	store.SCY += p.DeltaY

	for i, s := range store.SpriteData {
		if i == 0 {
			continue
		}
		if s == nil || s.ID == 0 {
			return
		}
		s.ScreenXPixel -= p.DeltaX
		s.ScreenYPixel -= p.DeltaY
	}
}

// CollisionCheckForPlayer check if collision occurs in player moving ahead
func CollisionCheckForPlayer() bool {
	collision := false
	p := store.SpriteData[0]
	if p == nil || p.ID == 0 {
		return false
	}
	for offset, s := range store.SpriteData {
		if offset == 0 {
			continue
		}
		if s == nil || s.ID == 0 {
			break
		}

		switch p.Direction {
		case util.Up:
			if p.MapXCoord == s.MapXCoord && p.MapYCoord-1 == s.MapYCoord {
				collision = true
			}
		case util.Down:
			if p.MapXCoord == s.MapXCoord && p.MapYCoord+1 == s.MapYCoord {
				collision = true
			}
		case util.Left:
			if p.MapXCoord-1 == s.MapXCoord && p.MapYCoord == s.MapYCoord {
				collision = true
			}
		case util.Right:
			if p.MapXCoord+1 == s.MapXCoord && p.MapYCoord == s.MapYCoord {
				collision = true
			}
		}

		if collision {
			break
		}
	}
	if collision {
		audio.PlaySound(audio.SFX_COLLISION)
	}
	return collision
}
