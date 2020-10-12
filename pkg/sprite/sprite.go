package sprite

import (
	"fmt"
	"image/png"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// Movement status
const (
	Uninitialized byte = iota
	OK
	Delay
	Movement
)

// Player sprite state
const (
	Normal uint = iota
	Cycling
	Seel
)

// AddSprite add sprite into SpriteData
func AddSprite(name string, x, y util.Coord, movementBytes [2]byte) {
	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("/%s_%d.png", name, i)
		f, err := store.FS.Open(path)
		if err != nil {
			// NOTE: NotFoundFileError isn't needed
			break
		}
		defer f.Close()

		img, _ := png.Decode(f)
		imgs[i], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	n := NumSprites()
	s := &store.Sprite{
		ScreenXPixel:  16 * x,
		ScreenYPixel:  16*y - 4,
		MapXCoord:     x,
		MapYCoord:     y,
		MovementBytes: movementBytes,
		VRAM: store.SpriteImage{
			Index:  1,
			Images: imgs,
		},
	}
	store.SpriteData[n] = s
}

// NumSprites a number of sprites at current map
func NumSprites() uint {
	i := uint(0)
	for store.SpriteData[i] != nil {
		i++
	}
	return i
}

// UpdateSprites update sprite data
func UpdateSprites() {
	for offset := range store.SpriteData {
		if store.IsInvalidSprite(uint(offset)) {
			break
		}
		if offset == 0 {
			UpdatePlayerSprite()
			continue
		}
		UpdateNPCSprite(uint(offset))
	}
}

// UpdateSpriteImage update sprite image index
func UpdateSpriteImage(offset uint) {
	s := store.SpriteData[offset]
	index := -1
	if s == nil {
		return
	}
	length := len(s.VRAM.Images)
	if length == 1 {
		s.VRAM.Index = 0
		return
	}

	// ref:
	index = 0
	switch s.AnimationCounter() + uint(s.Direction) {

	// down
	case 0, 3:
		index = 1
		if length == 4 {
			index = 0
		}
	case 1, 2:
		index = 0
		if length == 4 {
			index = 0
		}
		if s.RightHand {
			index = 2
		}

	// up
	case 4, 7:
		index = 4
		if length == 4 {
			index = 1
		}
	case 5, 6:
		index = 3
		if length == 4 {
			index = 1
		}
		if s.RightHand {
			index = 5
		}

	case 8, 11:
		index = 6
		if length == 4 {
			index = 2
		}
	case 9, 10:
		index = 7
		if length == 4 {
			index = 2
		}

	case 12, 15:
		index = 8
		if length == 4 {
			index = 3
		}
	case 13, 14:
		index = 9
		if length == 4 {
			index = 3
		}
	}
	s.VRAM.Index = index
}

// DisableSprite hide sprite
func DisableSprite(offset uint) {
	s := store.SpriteData[offset]
	s.VRAM.Index = -1
}

// MoveSpriteForcely move sprite by movement data forcely
// set wNPCMovementDirections
func MoveSpriteForcely(offset uint, movement []byte) {
	copy(NPCMovementDirections, movement)
	util.SetBit(store.D730, 0)
	joypad.JoyIgnore = joypad.ByteToInput(0xff)
}

// drawSprite
func drawSprite(offset uint) {
	s := store.SpriteData[offset]
	UpdateSpriteImage(offset)
	util.DrawImagePixel(store.TileMap, s.VRAM.Images[s.VRAM.Index], s.ScreenXPixel, s.ScreenYPixel)
}

// VBlank script executed in VBlank
func VBlank() {
	for i, s := range store.SpriteData {
		if store.IsInvalidSprite(uint(i)) {
			break
		}
		if s.VRAM.Index < 0 {
			return
		}
		drawSprite(uint(i))
	}
}
