package object

import (
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap"
)

var RedsHouse2F = &Object{
	Border: 0x0a,

	Warps: []Warp{
		{7, 1, 2, worldmap.REDS_HOUSE_1F},
	},

	Signs: []Sign{},

	Sprites: []Sprite{},

	WarpTos: []WarpTo{
		{7, 1, 4},
	},

	HO: []HiddenObject{
		{0, 1, txt.RedBedroomPCText, true},
		{3, 5, txt.RedBedroomSNESText, true},
	},
}
