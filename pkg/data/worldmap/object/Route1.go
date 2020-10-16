package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/util"
)

var Route1 = &Object{
	Border: 0x0b,

	Warps: []Warp{},

	Signs: []Sign{
		{9, 27, 3},
	},

	Sprites: []Sprite{
		{sprdata.SPRITE_BUG_CATCHER, 5, 24, [2]byte{util.Walk, 1}, 1},
		{sprdata.SPRITE_BUG_CATCHER, 15, 13, [2]byte{util.Walk, 2}, 2},
	},

	WarpTos: []WarpTo{
		{2, 7, 4},
	},
}
