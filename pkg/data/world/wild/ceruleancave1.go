package wild

import "pokered/pkg/data/pokemon"

var grass = [10]Pokemon{
	{pokemon.GOLBAT, 46},
	{pokemon.HYPNO, 46},
	{pokemon.MAGNETON, 46},
	{pokemon.DODRIO, 49},
	{pokemon.VENOMOTH, 49},
	{pokemon.ARBOK, 52},
	{pokemon.KADABRA, 49},
	{pokemon.PARASECT, 52},
	{pokemon.RAICHU, 53},
	{pokemon.DITTO, 53},
}

var water = [10]Pokemon{}

var Ceruleancave1 = Wild{
	Rate:  0x0a,
	Grass: grass,
	Water: water,
}
