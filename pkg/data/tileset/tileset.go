package tileset

import (
	"image"
	"pokered/pkg/store"
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

var tilesets map[uint]Tileset

func InitTilesets() {
	result := map[uint]Tileset{}
	for id, tilesetString := range tilesetsString {
		path := "/" + tilesetString + ".png"
		img := util.OpenImage(store.FS, path)

		width, height := img.Size()
		width /= 8
		height /= 8
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				min, max := image.Point{w * 8, h * 8}, image.Point{(w + 1) * 8, (h + 1) * 8}
				tile, err := ebiten.NewImageFromImage(img.SubImage(image.Rectangle{min, max}), ebiten.FilterDefault)
				if err != nil {
					panic(err)
				}
				result[uint(id)] = append(result[uint(id)], tile)
			}
		}
	}
	tilesets = result
}

// Tile get tile data
func Tile(tilesetID, tileID uint) *ebiten.Image {
	ts, ok := tilesets[tilesetID]
	if !ok {
		return nil
	}
	return ts[tileID]
}
