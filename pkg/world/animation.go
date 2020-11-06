package world

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/store"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// ref: hMovingBGTilesCounter1, wMovingBGTilesCounter2
var animCounter1, animCounter2 int

const (
	flower int = iota + 1
	water0
	water1
	water2
)

const (
	flower0Path = "/flower_0.png"
	flower1Path = "/flower_1.png"
	flower2Path = "/flower_2.png"
	water00Path = "/water0_0.png"
	water01Path = "/water0_1.png"
	water02Path = "/water0_2.png"
	water03Path = "/water0_3.png"
	water04Path = "/water0_4.png"
	water10Path = "/water1_0.png"
	water11Path = "/water1_1.png"
	water12Path = "/water1_2.png"
	water13Path = "/water1_3.png"
	water14Path = "/water1_4.png"
	water20Path = "/water2_0.png"
	water21Path = "/water2_1.png"
	water22Path = "/water2_2.png"
	water23Path = "/water2_3.png"
	water24Path = "/water2_4.png"
)

var (
	flowers = [3]*ebiten.Image{
		util.OpenImage(store.FS, flower0Path),
		util.OpenImage(store.FS, flower1Path),
		util.OpenImage(store.FS, flower2Path),
	}
	waters = [15]*ebiten.Image{
		util.OpenImage(store.FS, water00Path),
		util.OpenImage(store.FS, water01Path),
		util.OpenImage(store.FS, water02Path),
		util.OpenImage(store.FS, water03Path),
		util.OpenImage(store.FS, water04Path),
		util.OpenImage(store.FS, water10Path),
		util.OpenImage(store.FS, water11Path),
		util.OpenImage(store.FS, water12Path),
		util.OpenImage(store.FS, water13Path),
		util.OpenImage(store.FS, water14Path),
		util.OpenImage(store.FS, water20Path),
		util.OpenImage(store.FS, water21Path),
		util.OpenImage(store.FS, water22Path),
		util.OpenImage(store.FS, water23Path),
		util.OpenImage(store.FS, water24Path),
	}
)

type animTile struct {
	specie int
	x, y   util.Tile
}

var curAnimTiles []animTile

func checkAnimTiles(blockID uint, blockX, blockY int) {
	tilesetID := CurBlockset.TilesetID
	tiles := CurBlockset.Bytes[blockID]
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tileID := tiles[y*4+x]
			switch tileID {
			case 0x3:
				switch tilesetID {
				case tileset.Overworld, tileset.Gym:
					curAnimTiles = append(curAnimTiles, animTile{
						specie: flower,
						x:      blockX*4 + x,
						y:      blockY*4 + y,
					})
				}
			case 0x14:
				specie := 0
				switch tilesetID {
				case tileset.Overworld:
					specie = water2
				case tileset.Cavern, tileset.Gym, tileset.Plateau, tileset.ShipPort, tileset.Ship:
					specie = water1
				}
				if specie > 0 {
					curAnimTiles = append(curAnimTiles, animTile{
						specie: specie,
						x:      blockX*4 + x,
						y:      blockY*4 + y,
					})
				}
			}
		}
	}
}

// updateMovingTiles execute map animation, such as water and flower
// ref: UpdateMovingBgTiles
func updateMovingTiles() {
	animCounter1++
	switch animCounter1 {
	case 20:
		// water
		animCounter2 = (animCounter2 + 1) % 8

		var index int
		switch animCounter2 {
		case 0:
			index = 0
		case 1, 7:
			index = 1
		case 2, 6:
			index = 2
		case 3, 4:
			index = 3
		case 5:
			index = 4
		}

		for _, t := range curAnimTiles {
			switch t.specie {
			case water0:
				util.DrawImage(CurWorld.Image, waters[index], t.x, t.y)
			case water1:
				util.DrawImage(CurWorld.Image, waters[5+index], t.x, t.y)
			case water2:
				util.DrawImage(CurWorld.Image, waters[10+index], t.x, t.y)
			}
		}

		// don't execute flower animation when player is in cavern
		if IsCurTileset(tileset.Cavern) {
			animCounter1 = 0
		}

	case 21:
		// flower
		animCounter1 = 0

		var index int
		switch animCounter2 % 4 {
		case 0, 1:
			index = 0
		case 2:
			index = 1
		case 3:
			index = 2
		}

		for _, t := range curAnimTiles {
			if t.specie == flower {
				util.DrawImage(CurWorld.Image, flowers[index], t.x, t.y)
			}
		}
	}
}
