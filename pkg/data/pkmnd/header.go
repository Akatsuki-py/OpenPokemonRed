package pkmnd

import "pokered/pkg/data/move"

// growth rate
const (
	Exp600k  = iota + 1 // pred none
	Exp800k             // pred 4
	Exp1000k            // pred 0
	Exp1050k            // pred 3
	Exp1250k            // pred 5
	Exp1640k            // pred none
)

// PHeader Pokemon Header
type PHeader struct {
	ID            uint
	Name          string
	IconGen1      uint
	BaseStatsGen1 StatsGen1
	BaseStats     Stats
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

func BaseStatsGen1(id uint) StatsGen1 {
	h := *Header(id)
	return h.BaseStatsGen1
}

var AbraHeader = PHeader{
	ID:            63,
	Name:          "abra",
	IconGen1:      MonMon,
	BaseStatsGen1: StatsGen1{25, 20, 15, 90, 105},
	BaseStats:     Stats{25, 20, 15, 90, 105, 55},
	Type:          [2]uint{Psychic},
	CatchRate:     200,
	BaseExp:       73,
	Lv0MoveIDs:    [4]uint{move.TELEPORT},
	GrowthRate:    Exp1050k,
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
	BaseStatsGen1: StatsGen1{45, 49, 49, 45, 65},
	BaseStats:     Stats{45, 49, 49, 45, 65, 65},
	Type:          [2]uint{Grass, Poison},
	CatchRate:     45,
	BaseExp:       64,
	Lv0MoveIDs:    [4]uint{move.TACKLE, move.GROWL},
	GrowthRate:    Exp1050k,
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
	BaseStatsGen1: StatsGen1{39, 52, 43, 65, 50},
	BaseStats:     Stats{39, 52, 43, 65, 60, 50},
	Type:          [2]uint{Fire},
	CatchRate:     45,
	BaseExp:       65,
	Lv0MoveIDs:    [4]uint{move.SCRATCH, move.GROWL},
	GrowthRate:    Exp1050k,
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
