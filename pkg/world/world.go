package world

import (
	"pokered/pkg/data/worldmap/header"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// World data
type World struct {
	MapID  uint
	Image  *ebiten.Image
	Header *header.Header
}

var curWorld *World

// LoadWorldData load world data
func LoadWorldData(id uint) {
	h := header.Get(id)
	img, _ := ebiten.NewImage(int(h.Width*32), int(h.Height*32), ebiten.FilterDefault)
	loadBlockset(h.Tileset)

	for y := 0; y < int(h.Height); y++ {
		for x := 0; x < int(h.Width); x++ {
			blockID := h.Blk[y*int(h.Width)+x]
			block := curBlockset.Data[blockID]
			util.DrawImageBlock(img, block, x, y)
		}
	}
	curWorld = &World{
		MapID:  id,
		Image:  img,
		Header: h,
	}
}

// VBlank script executed in VBlank
func VBlank() {
	if curWorld == nil {
		return
	}

	util.DrawImage(store.TileMap, curWorld.Image, 0, 0)
}
