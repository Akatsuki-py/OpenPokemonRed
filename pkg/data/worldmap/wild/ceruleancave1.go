package wild

import "pokered/pkg/data/pkmnd"

var grass = [10]Pokemon{
	{pkmnd.GOLBAT, 46},
	{pkmnd.HYPNO, 46},
	{pkmnd.MAGNETON, 46},
	{pkmnd.DODRIO, 49},
	{pkmnd.VENOMOTH, 49},
	{pkmnd.ARBOK, 52},
	{pkmnd.KADABRA, 49},
	{pkmnd.PARASECT, 52},
	{pkmnd.RAICHU, 53},
	{pkmnd.DITTO, 53},
}

var water = [10]Pokemon{}

var Ceruleancave1 = Wild{
	Rate:  0x0a,
	Grass: grass,
	Water: water,
}
