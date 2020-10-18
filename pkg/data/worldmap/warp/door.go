package warp

import "pokered/pkg/data/tileset"

// DoorTileIDs tilesetID -> door tiles
var DoorTileIDs = map[uint][]byte{
	tileset.Overworld: {0x1b, 0x58},
	tileset.Forest:    {0x3a},
	tileset.Mart:      {0x5e},
	tileset.House:     {0x54},
	tileset.Gate:      {0x3b},
	tileset.Ship:      {0x1e},
	tileset.Lobby:     {0x1c, 0x38, 0x1a},
	tileset.Mansion:   {0x1a, 0x1c, 0x53},
	tileset.Lab:       {0x34},
	tileset.Facility:  {0x43, 0x58, 0x1b},
	tileset.Plateau:   {0x3b, 0x1b},
}
