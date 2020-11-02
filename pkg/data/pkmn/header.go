package pkmn

import "pokered/pkg/data/move"

// PHeader Pokemon Header
type PHeader struct {
	DexID      uint
	BaseStats  stat
	Type       [2]uint
	CatchRate  byte
	BaseExp    uint
	Lv0MoveIDs [4]uint
	GrowthRate uint
	Learnset   []uint
}

var AbraHeader = PHeader{
	DexID:      63,
	BaseStats:  stat{25, 20, 15, 90, 105, 55},
	Type:       [2]uint{Psychic},
	CatchRate:  200,
	BaseExp:    73,
	Lv0MoveIDs: [4]uint{move.TELEPORT},
	GrowthRate: 3,
	Learnset:   []uint{},
}

var Bulbasaur = PHeader{
	DexID:      1,
	BaseStats:  stat{45, 49, 49, 45, 65, 65},
	Type:       [2]uint{Grass, Poison},
	CatchRate:  45,
	BaseExp:    64,
	Lv0MoveIDs: [4]uint{move.TACKLE, move.GROWL},
	GrowthRate: 3,
	Learnset:   []uint{},
}
