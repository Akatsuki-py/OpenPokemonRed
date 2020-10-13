package pokemon

// nameMap PokemonID -> Name
var nameMap = map[uint]string{}

// Name get name from PokemonID
func Name(id uint) string {
	return nameMap[id]
}
