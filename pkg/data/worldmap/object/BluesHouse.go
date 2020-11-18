package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

var BluesHouse = &Object{
	Border: 0x0a,

	Warps: []Warp{
		{2, 7, 1, worldmap.LAST_MAP},
		{3, 7, 1, worldmap.LAST_MAP},
	},

	Signs: []Sign{},

	Sprites: []Sprite{
		{sprdata.SPRITE_DAISY, 2, 3, [2]byte{util.Stay, byte(util.Right)}, 1},
		{sprdata.SPRITE_DAISY, 6, 4, [2]byte{util.Walk, byte(1)}, 0},
		{sprdata.SPRITE_BOOK_MAP_DEX, 3, 3, [2]byte{util.Stay, byte(util.None)}, 3},
	},

	WarpTos: []WarpTo{
		{2, 7, 4},
		{3, 7, 4},
	},

	HS: map[int]bool{
		0x02: true,
	},
}
