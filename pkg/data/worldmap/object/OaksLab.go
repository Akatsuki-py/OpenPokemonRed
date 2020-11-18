package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

var OaksLab = &Object{
	Border: 0x03,

	Warps: []Warp{
		{4, 11, 2, worldmap.LAST_MAP},
		{5, 11, 2, worldmap.LAST_MAP},
	},

	Signs: []Sign{},

	Sprites: []Sprite{
		{sprdata.SPRITE_BLUE, 4, 3, [2]byte{util.Stay, util.None}, 1},
		{sprdata.SPRITE_BALL, 6, 3, [2]byte{util.Stay, util.None}, 2},
		{sprdata.SPRITE_BALL, 7, 3, [2]byte{util.Stay, util.None}, 3},
		{sprdata.SPRITE_BALL, 8, 3, [2]byte{util.Stay, util.None}, 4},
		{sprdata.SPRITE_OAK, 5, 2, [2]byte{util.Stay, byte(util.Down)}, 5},
		{sprdata.SPRITE_BOOK_MAP_DEX, 2, 1, [2]byte{util.Stay, util.None}, 6},
		{sprdata.SPRITE_BOOK_MAP_DEX, 3, 1, [2]byte{util.Stay, util.None}, 7},
		{sprdata.SPRITE_OAK, 5, 10, [2]byte{util.Stay, byte(util.Up)}, 8},
		{sprdata.SPRITE_GIRL, 1, 9, [2]byte{util.Walk, 1}, 9},
		{sprdata.SPRITE_OAK_AIDE, 2, 10, [2]byte{util.Stay, util.None}, 10},
		{sprdata.SPRITE_OAK_AIDE, 8, 10, [2]byte{util.Stay, util.None}, 11},
	},

	WarpTos: []WarpTo{
		{4, 11, 5},
		{5, 11, 5},
	},

	HS: map[int]bool{
		0x05: true,
		0x08: true,
	},
}
