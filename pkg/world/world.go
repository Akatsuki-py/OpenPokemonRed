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

// map exterior range(block)
const exterior int = 3

// LoadWorldData load world data
func LoadWorldData(id int) {
	h, o := header.Get(id), object.Get(id)
	o.Initialized = false
	img, _ := ebiten.NewImage(int(h.Width*32)+2*exterior*32, int(h.Height*32)+2*exterior*32, ebiten.FilterDefault)
	loadBlockset(h.Tileset)

	for y := 0; y < int(h.Height)+2*exterior; y++ {
		for x := 0; x < int(h.Width)+2*exterior; x++ {
			switch {
			case y < int(exterior):
				northCon := h.Connections.North
				if northCon.OK {
					northMapH, northMapO := header.Get(northCon.DestMapID), object.Get(northCon.DestMapID)
					if x < int(exterior) || x > int(h.Width)+exterior-1 {
						block := CurBlockset.Data[northMapO.Border]
						util.DrawImageBlock(img, block, x, y)
						continue
					}
					blockID := northMapH.Blk(int((northMapH.Height-uint(exterior-y))*northMapH.Width) + (x - exterior))
					block := CurBlockset.Data[blockID]
					util.DrawImageBlock(img, block, x, y)
				} else {
					block := CurBlockset.Data[o.Border]
					util.DrawImageBlock(img, block, x, y)
				}

			case y > int(h.Height)+exterior-1:
				southCon := h.Connections.South
				if southCon.OK {
					southMapH := header.Get(southCon.DestMapID)
					if x < int(exterior) || x > int(h.Width)+1 {
						block := CurBlockset.Data[o.Border]
						util.DrawImageBlock(img, block, x, y)
						continue
					}
					blockID := southMapH.Blk(int((uint(y)-h.Height-uint(exterior))*southMapH.Width) + (x - exterior))
					block := CurBlockset.Data[blockID]
					util.DrawImageBlock(img, block, x, y)
				} else {
					block := CurBlockset.Data[o.Border]
					util.DrawImageBlock(img, block, x, y)
				}

			case x < int(exterior) || x > int(h.Width)+2:
				block := CurBlockset.Data[o.Border]
				util.DrawImageBlock(img, block, x, y)

			default:
				blockID := h.Blk((y-exterior)*int(h.Width) + (x - exterior))
				block := CurBlockset.Data[blockID]
				util.DrawImageBlock(img, block, x, y)
			}
		}
	}

	CurWorld = &World{
		MapID:  id,
		Image:  img,
		Header: h,
		Object: o,
	}
}

// CurTileID get tile ID on which player stands
func CurTileID(x, y, pixelX, pixelY int) (uint, int) {
	blockX, blockY := ((x-4)*16+pixelX)/32, ((y-4)*16+pixelY+4)/32
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 {
		return CurBlockset.TilesetID, -1
	}
	blockID := CurWorld.Header.Blk(blockOffset)

	switch {
	case x%2 == 0 && y%2 == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+0])
	case x%2 == 1 && y%2 == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+2])
	case x%2 == 0 && y%2 == 1:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+8])
	case x%2 == 1 && y%2 == 1:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+10])
	}

	return CurBlockset.TilesetID, 0
}

// FrontTileID get tile ID in front of player
func FrontTileID(x, y, pixelX, pixelY int, direction util.Direction) (uint, int) {
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

	blockX, blockY := ((x-4)*16+pixelX+deltaX)/32, ((y-4)*16+pixelY+4+deltaY)/32
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 || blockOffset > CurWorld.Header.BlkLen() {
		return CurBlockset.TilesetID, -1
	}
	blockID := CurWorld.Header.Blk(blockOffset)

	switch {
	case px%2 == 0 && py%2 == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+0])
	case px%2 == 1 && py%2 == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+2])
	case px%2 == 0 && py%2 == 1:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+8])
	case px%2 == 1 && py%2 == 1:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[uint(blockID)*16+10])
	}

	return CurBlockset.TilesetID, 0
}

// VBlank script executed in VBlank
func VBlank(XCoord, YCoord, deltaX, deltaY, walkCounter int, direction uint) {
	if CurWorld == nil {
		return
	}

	x := -32*exterior - XCoord*16 + 64
	y := -32*exterior - YCoord*16 + 64
	if walkCounter > 0 {
		x -= deltaX * (16 - walkCounter)
		y -= deltaY * (16 - walkCounter)
	}
	util.DrawImagePixel(store.TileMap, CurWorld.Image, x, y)
}
