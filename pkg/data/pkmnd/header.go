package pkmnd

import "pokered/pkg/data/move"

// PHeader Pokemon Header
type PHeader struct {
	ID            uint
	Name          string
	IconGen1      uint
	BaseStatsGen1 stat
	BaseStats     stat
	Type          [2]uint
	CatchRate     byte
	BaseExp       uint
	Lv0MoveIDs    [4]uint
	GrowthRate    uint
	Learnset      []uint
	Evos          []Evo
	LvMoves       [][2]uint // (Level, MoveID)[]
}

type Evo struct {
	ID uint
	// if this is zero, evo is taken by item or trade
	Level  uint
	ItemID uint
	Trade  bool
}

func Header(id uint) *PHeader {
	switch id {
	case 1:
		return &Bulbasaur
	case 4:
		return &Charmander
	case 63:
		return &AbraHeader
	}
	return nil
}

var AbraHeader = PHeader{
	ID:            63,
	Name:          "abra",
	IconGen1:      MonMon,
	BaseStatsGen1: stat{25, 20, 15, 90, 105, 105},
	BaseStats:     stat{25, 20, 15, 90, 105, 55},
	Type:          [2]uint{Psychic},
	CatchRate:     200,
	BaseExp:       73,
	Lv0MoveIDs:    [4]uint{move.TELEPORT},
	GrowthRate:    3,
	Learnset:      []uint{},
	Evos: []Evo{
		{KADABRA, 16, 0, false},
	},
	LvMoves: [][2]uint{},
}

var Bulbasaur = PHeader{
	ID:            1,
	Name:          "bulbasaur",
	IconGen1:      GrassMon,
	BaseStatsGen1: stat{45, 49, 49, 45, 65, 65},
	BaseStats:     stat{45, 49, 49, 45, 65, 65},
	Type:          [2]uint{Grass, Poison},
	CatchRate:     45,
	BaseExp:       64,
	Lv0MoveIDs:    [4]uint{move.TACKLE, move.GROWL},
	GrowthRate:    3,
	Learnset:      []uint{},
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

var Charmander = PHeader{
	ID:            4,
	Name:          "charmander",
	IconGen1:      MonMon,
	BaseStatsGen1: stat{39, 52, 43, 65, 50, 50},
	BaseStats:     stat{39, 52, 43, 65, 60, 50},
	Type:          [2]uint{Fire},
	CatchRate:     45,
	BaseExp:       65,
	Lv0MoveIDs:    [4]uint{move.SCRATCH, move.GROWL},
	GrowthRate:    3,
	Learnset:      []uint{},
	Evos: []Evo{
		{CHARMELEON, 16, 0, false},
	},
	LvMoves: [][2]uint{
		{9, move.EMBER},
		{15, move.LEER},
		{22, move.RAGE},
		{30, move.SLASH},
		{38, move.FLAMETHROWER},
		{46, move.FIRE_SPIN},
	},
}
