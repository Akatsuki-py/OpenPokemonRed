package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

var RedsHouse1F = &Object{
	Border: 0x0a,

	Warps: []Warp{
		{2, 7, 0, worldmap.LAST_MAP},
		{3, 7, 0, worldmap.LAST_MAP},
		{7, 1, 0, worldmap.REDS_HOUSE_2F},
	},

	Signs: []Sign{
		{3, 1, 2},
	},

	Sprites: []Sprite{
		{sprdata.SPRITE_MOM, 5, 4, [2]byte{util.Stay, byte(util.Left)}, 0},
	},

	WarpTos: []WarpTo{
		{2, 7, 4},
		{3, 7, 4},
		{7, 1, 4},
	},
}
