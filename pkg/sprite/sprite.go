package sprite

import (
	"fmt"
	"image/png"
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

// Movment status
const (
	Uninitialized byte = iota
	OK
	Delay
	Movement
)

const (
	Normal uint = iota
	Cycling
	Seel
)

func AddSprite(name string, x, y util.Tile) {
	FS, _ := fs.New()
	imgs := make([]*ebiten.Image, 10)
	for i := 0; i < 10; i++ {
		f, err := FS.Open(fmt.Sprintf("/%s_%d.png", name, i))
		if err != nil {
			break
		}
		defer f.Close()
		img, _ := png.Decode(f)
		imgs[i], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	n := NumSprites()
	s := &store.Sprite{
		ID:           n,
		ScreenXPixel: 8 * x,
		ScreenYPixel: 8*y - 4,
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
	for store.SpriteData[i] != nil && store.SpriteData[i].ID > 0 {
		i++
	}
	return i
}

// UpdateSprites update sprite data
func UpdateSprites() {
	for offset, s := range store.SpriteData {
		if s == nil || s.ID == 0 {
			break
		}
		if offset == 0 {
			UpdatePlayerSprite()
			continue
		}
		UpdateNPCSprite(uint(offset))
	}
}

// Index update sprite image index
func Index(offset uint) int {
	s := store.SpriteData[offset]
	if s == nil {
		return -1
	}
	length := len(s.VRAM.Images)
	if length == 1 {
		return -1
	}

	animCounter := s.AnimationFrame >> 2

	// ref:
	index := 0
	switch animCounter + uint(s.Direction) {

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
	return index
}

// DisableSprite hide sprite
func DisableSprite(offset uint) {
	s := store.SpriteData[offset]
	s.VRAM.Index = -1
}

// MoveSprite forcely move sprite by movement data
// set wNPCMovementDirections
func MoveSprite(offset uint, movement []byte) {
	copy(NPCMovementDirections, movement)
	util.SetBit(store.D730, 0)
	joypad.JoyIgnore = joypad.ByteToInput(0xff)
}

func DrawSprite(offset uint) {
	s := store.SpriteData[offset]
	index := Index(offset)
	util.DrawImagePixel(store.TileMap, s.VRAM.Images[index], s.ScreenXPixel, s.ScreenYPixel)
}

func VBlank() {
	for i, s := range store.SpriteData {
		if s == nil || s.ID == 0 {
			break
		}
		DrawSprite(uint(i))
	}
}
