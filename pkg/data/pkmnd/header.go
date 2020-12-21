package pkmnd

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
	case 7:
		return &Squirtle
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
	Lv0MoveIDs:    [4]uint{TELEPORT},
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
	Lv0MoveIDs:    [4]uint{TACKLE, GROWL},
	GrowthRate:    Exp1050k,
	Learnset:      []uint{},
	Evos: []Evo{
		{IVYSAUR, 16, 0, false},
	},
	LvMoves: [][2]uint{
		{7, LEECH_SEED},
		{13, VINE_WHIP},
		{20, POISONPOWDER},
		{27, RAZOR_LEAF},
		{34, GROWTH},
		{41, SLEEP_POWDER},
		{48, SOLARBEAM},
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
	Lv0MoveIDs:    [4]uint{SCRATCH, GROWL},
	GrowthRate:    Exp1050k,
	Learnset:      []uint{},
	Evos: []Evo{
		{CHARMELEON, 16, 0, false},
	},
	LvMoves: [][2]uint{
		{9, EMBER},
		{15, LEER},
		{22, RAGE},
		{30, SLASH},
		{38, FLAMETHROWER},
		{46, FIRE_SPIN},
	},
}

var Squirtle = PHeader{
	ID:            7,
	Name:          "squirtle",
	IconGen1:      WaterMon,
	BaseStatsGen1: StatsGen1{44, 48, 65, 43, 50},
	BaseStats:     Stats{44, 48, 65, 43, 50, 64},
	Type:          [2]uint{Water},
	CatchRate:     45,
	BaseExp:       66,
	Lv0MoveIDs:    [4]uint{TACKLE, TAIL_WHIP},
	GrowthRate:    Exp1050k,
	Learnset:      []uint{},
	Evos: []Evo{
		{WARTORTLE, 16, 0, false},
	},
	LvMoves: [][2]uint{
		{8, BUBBLE},
		{15, WATER_GUN},
		{22, BITE},
		{28, WITHDRAW},
		{35, SKULL_BASH},
		{42, HYDRO_PUMP},
	},
}
