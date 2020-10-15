package world

import (
	"pokered/pkg/data/worldmap/header"
	"pokered/pkg/data/worldmap/object"
	"pokered/pkg/store"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

// World data
type World struct {
	MapID  int
	Image  *ebiten.Image
	Header *header.Header
	Object *object.Object
}

var CurWorld *World

// LoadWorldData load world data
func LoadWorldData(id int) {
	h := header.Get(id)
	img, _ := ebiten.NewImage(int(h.Width*32), int(h.Height*32), ebiten.FilterDefault)
	loadBlockset(h.Tileset)
	for y := 0; y < int(h.Height); y++ {
		for x := 0; x < int(h.Width); x++ {
			blockID := h.Blk(y*int(h.Width) + x)
			block := curBlockset.Data[blockID]
			util.DrawImageBlock(img, block, x, y)
		}
	}

	o := object.Get(id)
	CurWorld = &World{
		MapID:  id,
		Image:  img,
		Header: h,
		Object: o,
	}
}

// CurTileID get tile ID on which player stands
func CurTileID(x, y, pixelX, pixelY int) (uint, uint) {
	blockX, blockY := (store.SCX+pixelX)/32, (store.SCY+pixelY+4)/32
	blockID := CurWorld.Header.Blk(blockY*int(CurWorld.Header.Width) + blockX)

	switch {
	case x%2 == 0 && y%2 == 0:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+0])
	case x%2 == 1 && y%2 == 0:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+2])
	case x%2 == 0 && y%2 == 1:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+8])
	case x%2 == 1 && y%2 == 1:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+10])
	}

	return curBlockset.TilesetID, 0
}

// FrontTileID get tile ID in front of player
func FrontTileID(x, y, pixelX, pixelY int, direction util.Direction) (uint, uint) {
	deltaX, deltaY := 0, 0
	px, py := x, y
	switch direction {
	case util.Up:
		py--
		deltaY = -16
	case util.Down:
		py++
		deltaY = 16
	case util.Left:
		px--
		deltaX = -16
	case util.Right:
		px++
		deltaX = 16
	}

	blockX, blockY := (store.SCX+pixelX+deltaX)/32, (store.SCY+pixelY+4+deltaY)/32
	blockID := CurWorld.Header.Blk(blockY*int(CurWorld.Header.Width) + blockX)

	switch {
	case px%2 == 0 && py%2 == 0:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+0])
	case px%2 == 1 && py%2 == 0:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+2])
	case px%2 == 0 && py%2 == 1:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+8])
	case px%2 == 1 && py%2 == 1:
		return curBlockset.TilesetID, uint(curBlockset.Bytes[uint(blockID)*16+10])
	}

	return curBlockset.TilesetID, 0
}

// VBlank script executed in VBlank
func VBlank() {
	if CurWorld == nil {
		return
	}

	util.DrawImagePixel(store.TileMap, CurWorld.Image, -store.SCX, -store.SCY)
}
