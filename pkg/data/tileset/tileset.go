package tileset

import (
	"github.com/hajimehoshi/ebiten"
)

var TilesetNames = [...]string{
	"cavern",
	"cemetery",
	"club",
	"facility",
	"forest",
	"gate",
	"gym",
	"house",
	"interior",
	"lab",
	"lobby",
	"mansion",
	"overworld",
	"plateau",
	"pokecenter",
	"reds_house",
	"ship_port",
	"ship",
	"underground",
}

type Tileset []*ebiten.Image

var Tilesets map[uint]Tileset

// Tile get tile data
func Tile(tilesetID, tileID uint) *ebiten.Image {
	ts, ok := Tilesets[tilesetID]
	if !ok {
		return nil
	}
	return ts[tileID]
}
