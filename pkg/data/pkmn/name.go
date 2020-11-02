package pkmn

// nameMap PokemonID -> Name
var nameMap = map[uint]string{
	RHYDON:     "rhydon",
	CHARMANDER: "charmander",
}

// Name get name from PokemonID
func Name(id uint) string {
	return nameMap[id]
}
