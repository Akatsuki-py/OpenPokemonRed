package pkmnd

import "pokered/pkg/data/txt"

// DexEntry contains dex page information
type DexEntry struct {
	Species string
	HT      [2]uint
	WT      uint
	Text    string
}

var abraDexEntry = DexEntry{
	Species: "PSI",
	HT:      [2]uint{2, 11},
	WT:      430,
	Text:    txt.AbraDexEntry,
}

var bulbasaurDexEntry = DexEntry{
	Species: "SEED",
	HT:      [2]uint{2, 4},
	WT:      150,
	Text:    txt.BulbasaurDexEntry,
}

var charmanderDexEntry = DexEntry{
	Species: "LIZARD",
	HT:      [2]uint{2, 0},
	WT:      190,
	Text:    txt.CharmanderDexEntry,
}

var squirtleDexEntry = DexEntry{
	Species: "TINYTURTLE",
	HT:      [2]uint{1, 8},
	WT:      200,
	Text:    txt.SquirtleDexEntry,
}
