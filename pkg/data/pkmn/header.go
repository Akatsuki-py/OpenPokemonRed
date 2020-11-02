package pkmn

import "pokered/pkg/data/move"

// PHeader Pokemon Header
type PHeader struct {
	ID         uint
	DexID      uint
	BaseStats  stat
	Type       [2]uint
	CatchRate  byte
	BaseExp    uint
	Lv0MoveIDs [4]uint
	GrowthRate uint
	Learnset   []uint
	Evos       []Evo
	LvMoves    [][2]uint // (Level, MoveID)[]
}

type Evo struct {
	ID uint
	// if this is zero, evo is taken by item or trade
	Level  uint
	ItemID uint
	Trade  bool
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
	Evos: []Evo{
		{KADABRA, 16, 0, false},
	},
	LvMoves: [][2]uint{},
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
	Evos: []Evo{
		{IVYSAUR, 16, 0, false},
	},
	LvMoves: [][2]uint{
		{7, move.LEECH_SEED},
		{13, move.VINE_WHIP},
		{20, move.POISONPOWDER},
		{27, move.RAZOR_LEAF},
		{34, move.GROWTH},
		{41, move.SLEEP_POWDER},
		{48, move.SOLARBEAM},
	},
}
