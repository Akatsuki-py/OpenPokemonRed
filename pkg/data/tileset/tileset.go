package tileset

import (
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

var tilesetsString = [...]string{
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

var tilesets = newTilesets()

func newTilesets() map[uint]Tileset {
	result := map[uint]Tileset{}
	for id, tilesetString := range tilesetsString {
		path := "/" + tilesetString + ".png"
		img := util.OpenImage(path)
		result[uint(id)] = append(result[uint(id)], img)
	}
	return result
}

// Tile get tile data
func Tile(tilesetID, tileID uint) *ebiten.Image {
	ts, ok := tilesets[tilesetID]
	if !ok {
		return nil
	}
	return ts[tileID]
}
