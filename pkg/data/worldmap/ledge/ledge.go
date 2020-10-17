package ledge

import "pokered/pkg/util"

type Ledge struct {
	Direction   uint
	CurTileID   int
	LedgeTileID int
}

var LedgeTiles = [...]Ledge{
	{util.Down, 0x2c, 0x37},
	{util.Down, 0x39, 0x36},
	{util.Down, 0x39, 0x37},
	{util.Left, 0x2c, 0x27},
	{util.Left, 0x39, 0x27},
	{util.Right, 0x2c, 0x0d},
	{util.Right, 0x2c, 0x1d},
	{util.Right, 0x39, 0x0d},
}
