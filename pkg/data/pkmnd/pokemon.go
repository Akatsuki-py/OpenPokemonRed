package pkmnd

type stat struct {
	HP      uint
	Attack  uint
	Defense uint
	Speed   uint
	SpAtk   uint
	SpDef   uint // unused in gen1
}

// Name get name from PokemonID
func Name(id uint) string {
	if h := Header(id); h != nil {
		return h.Name
	}
	return "UNKNOWN"
}
