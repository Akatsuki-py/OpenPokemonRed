package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

const palletTownWidth uint = 10

var PalletTown = &Object{
	Border: 0x0b,

	Warps: []Warp{
		{5, 5, 0, worldmap.REDS_HOUSE_1F},
		{13, 5, 0, worldmap.BLUES_HOUSE},
		{12, 11, 1, worldmap.OAKS_LAB},
	},

	Signs: []Sign{
		{13, 13, 4},
		{7, 9, 5},
		{3, 5, 6},
		{11, 5, 7},
	},

	Sprites: []Sprite{
		{sprdata.SPRITE_OAK, 8, 5, [2]byte{util.Stay, util.None}, 1},
		{sprdata.SPRITE_GIRL, 3, 8, [2]byte{util.Walk, 0}, 2},
		{sprdata.SPRITE_FISHER2, 11, 14, [2]byte{util.Walk, 0}, 3},
	},

	WarpTos: []WarpTo{
		{5, 5, palletTownWidth},
		{13, 5, palletTownWidth},
		{12, 11, palletTownWidth},
	},
}
