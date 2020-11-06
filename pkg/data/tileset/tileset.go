package tileset

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
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

	if tileID >= uint(len(ts)) {
		empty := ebiten.NewImage(8, 8)
		return empty
	}
	return ts[tileID]
}
