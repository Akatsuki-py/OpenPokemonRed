package world

import (
	"math"
	"pokered/pkg/data/tileset"
	"pokered/pkg/util"
)

const (
	Passable   int = -1
	NoPassable int = -2
)

// GetTileID get tile ID on which player stands
func GetTileID(x, y util.Tile) (uint, int) {
	blockX, blockY := (x*8)/32, (y*8)/32
	coordX, coordY := (x*8)%32, (y*8)%32
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 {
		return CurBlockset.TilesetID, -1
	}
	blockID := int(CurWorld.Header.Blk(blockOffset))
	index := coordY/2 + coordX/8
	return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][index])
}

// WriteTileID override tile
func WriteTileID(tileID byte, x, y util.Tile) {
	blockX, blockY := (x*8)/32, (y*8)/32
	coordX, coordY := (x*8)%32, (y*8)%32
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 {
		return
	}
	blockID := int(CurWorld.Header.Blk(blockOffset))
	index := coordY/2 + coordX/8
	CurBlockset.Bytes[blockID][index] = tileID
	util.DrawImage(CurWorld.Image, tileset.Tile(CurBlockset.TilesetID, uint(tileID)), x, y)
}

// CurTileID get tile ID on which player stands
func CurTileID(x, y int) (uint, int) {
	blockX, blockY := (x*16)/32, (y*16+8)/32
	coordX, coordY := (x*16)%32, (y*16+8)%32-8
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 {
		return CurBlockset.TilesetID, Passable
	}
	blockID := CurWorld.Header.Blk(blockOffset)

	switch {
	case coordX == 0 && coordY == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][4])
	case coordX == 16 && coordY == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][6])
	case coordX == 0 && coordY == 16:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][12])
	case coordX == 16 && coordY == 16:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][14])
	}

	return CurBlockset.TilesetID, NoPassable
}

// FrontTileID get tile ID in front of player
func FrontTileID(x, y int, direction util.Direction) (uint, int) {
	deltaX, deltaY := 0, 0
	switch direction {
	case util.Up:
		deltaY = -16
	case util.Down:
		deltaY = 16
	case util.Left:
		deltaX = -16
	case util.Right:
		deltaX = 16
	}

	pixelX, pixelY := x*16+deltaX, y*16+8+deltaY
	blockX, blockY := pixelX/32, pixelY/32

	connections := CurWorld.Header.Connections
	if pixelX < 0 && !connections.West.OK {
		return CurBlockset.TilesetID, NoPassable
	}
	if blockX >= int(CurWorld.Header.Width) && !connections.East.OK {
		return CurBlockset.TilesetID, NoPassable
	}
	if pixelY < 0 && !connections.North.OK {
		return CurBlockset.TilesetID, NoPassable
	}
	if blockY >= int(CurWorld.Header.Height) && !connections.South.OK {
		return CurBlockset.TilesetID, NoPassable
	}

	coordX, coordY := int(math.Abs(float64((x*16+deltaX)%32))), int(math.Abs(float64((y*16+8+deltaY)%32-8)))
	blockOffset := blockY*int(CurWorld.Header.Width) + blockX
	if blockOffset < 0 || blockOffset > CurWorld.Header.BlkLen() {
		return CurBlockset.TilesetID, Passable
	}
	blockID := CurWorld.Header.Blk(blockOffset)

	switch {
	case coordX == 0 && coordY == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][4])
	case coordX == 16 && coordY == 0:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][6])
	case coordX == 0 && coordY == 16:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][12])
	case coordX == 16 && coordY == 16:
		return CurBlockset.TilesetID, int(CurBlockset.Bytes[blockID][14])
	}

	return CurBlockset.TilesetID, Passable
}
