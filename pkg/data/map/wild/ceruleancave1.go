package wild

import "pokered/pkg/data/constant"

var grass = [10]pokemon{
	{constant.GOLBAT, 46},
	{constant.HYPNO, 46},
	{constant.MAGNETON, 46},
	{constant.DODRIO, 49},
	{constant.VENOMOTH, 49},
	{constant.ARBOK, 52},
	{constant.KADABRA, 49},
	{constant.PARASECT, 52},
	{constant.RAICHU, 53},
	{constant.DITTO, 53},
}

var water = [10]pokemon{}

var Ceruleancave1 = Wild{
	Rate:  0x0a,
	Grass: grass,
	Water: water,
}
