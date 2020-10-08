package sprite

import (
	"fmt"
	"image/png"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

func AddPlayer(state uint) {
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
