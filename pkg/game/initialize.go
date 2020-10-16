package game

import (
	"image"
	"net/http"
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"

	"github.com/hajimehoshi/ebiten"
)

func initialize() {
	initTilesets(store.FS)
	world.LoadWorldData(worldmap.PALLET_TOWN)
	sprite.InitPlayer(sprite.Normal, 3, 4)
}

func initTilesets(fs http.FileSystem) {
	result := map[uint]tileset.Tileset{}
	for id, name := range tileset.TilesetNames {
		path := "/" + name + ".png"
		img := util.OpenImage(fs, path)

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
	tileset.Tilesets = result
}
