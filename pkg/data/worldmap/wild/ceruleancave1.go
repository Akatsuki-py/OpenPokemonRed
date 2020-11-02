package wild

import "pokered/pkg/data/pkmn"

var grass = [10]Pokemon{
	{pkmn.GOLBAT, 46},
	{pkmn.HYPNO, 46},
	{pkmn.MAGNETON, 46},
	{pkmn.DODRIO, 49},
	{pkmn.VENOMOTH, 49},
	{pkmn.ARBOK, 52},
	{pkmn.KADABRA, 49},
	{pkmn.PARASECT, 52},
	{pkmn.RAICHU, 53},
	{pkmn.DITTO, 53},
}

var water = [10]Pokemon{}

var Ceruleancave1 = Wild{
	Rate:  0x0a,
	Grass: grass,
	Water: water,
}
