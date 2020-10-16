package sprite

import (
	"fmt"
	"image/png"
	"pokered/pkg/audio"
	"pokered/pkg/data/tilecoll"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"

	"github.com/hajimehoshi/ebiten"
)

// InitPlayer initialize player sprite
func InitPlayer(state uint, x, y int) {
	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		name := "red"
		switch state {
		case Cycling:
			name = "red_cycling"
		case Seel:
			name = "seel"
		}

		path := fmt.Sprintf("/%s_%d.png", name, i)
		f, err := store.FS.Open(path)
		if err != nil {
			util.NotFoundFileError(path)
		}
		defer f.Close()

		img, _ := png.Decode(f)
		imgs[i], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	s := &store.Sprite{
		ScreenXPixel: 16 * 4,
		ScreenYPixel: 16*4 - 4,
		MapXCoord:    x,
		MapYCoord:    y,
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
	if store.IsInvalidSprite(0) {
		return
	}

	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		name := "red"
		switch state {
		case Cycling:
			name = "red_cycling"
		case Seel:
			name = "seel"
		}

		path := fmt.Sprintf("/%s_%d.png", name, i)
		f, err := store.FS.Open(path)
		if err != nil {
			util.NotFoundFileError(path)
		}
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
	if store.IsInvalidSprite(0) {
		return
	}

	if p.WalkCounter > 0 {
		p.AnimationFrame++
		if p.AnimationCounter() == 4 {
			p.AnimationFrame = 0
		}
	}
	p.VRAM.Index = int(p.Direction + p.AnimationCounter())
}

// AdvancePlayerSprite advance player's walk by a frame
func AdvancePlayerSprite() {
	p := store.SpriteData[0]
	if store.IsInvalidSprite(0) {
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
		if store.IsInvalidSprite(uint(i)) {
			return
		}
		s.ScreenXPixel -= p.DeltaX
		s.ScreenYPixel -= p.DeltaY
	}
}

// CollisionCheckForPlayer check if collision occurs in player moving ahead
func CollisionCheckForPlayer() bool {
	collision := false
	defer func() {
		if collision {
			audio.PlaySound(audio.SFX_COLLISION)
		}
	}()

	p := store.SpriteData[0]
	if store.IsInvalidSprite(0) {
		return false
	}

	for offset, s := range store.SpriteData {
		if offset == 0 {
			continue
		}
		if store.IsInvalidSprite(uint(offset)) {
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
			return collision
		}
	}

	tilesetID, frontTileID := world.FrontTileID(p.MapXCoord, p.MapYCoord, p.ScreenXPixel, p.ScreenYPixel, p.Direction)
	if frontTileID > 0 && !util.Contains(tilecoll.Get(tilesetID), byte(uint(frontTileID))) {
		collision = true
		return collision
	}

	_, curTileID := world.CurTileID(p.MapXCoord, p.MapYCoord, p.ScreenXPixel, p.ScreenYPixel)
	if tilecoll.IsCollisionPair(tilesetID, byte(curTileID), byte(frontTileID), false) {
		collision = true
		return collision
	}
	return collision
}
