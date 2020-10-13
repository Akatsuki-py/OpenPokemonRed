package wild

type pokemon struct {
	ID    uint
	Level uint
}

// Wild map wild pokemon data
type Wild struct {
	Rate  byte
	Grass [10]pokemon
	Water [10]pokemon
}
