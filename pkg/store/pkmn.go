package store

// NonVolatileStatus type for non-volatile statuses
type NonVolatileStatus uint

// non-volatile statuses
const (
	OK  NonVolatileStatus = 0
	Psn NonVolatileStatus = 3
	Brn NonVolatileStatus = 4
	Frz NonVolatileStatus = 5
	Par NonVolatileStatus = 6
	Slp NonVolatileStatus = 7
)

func (n *NonVolatileStatus) String() string {
	switch *n {
	case Psn:
		return "PSN"
	case Brn:
		return "BRN"
	case Frz:
		return "FRZ"
	case Par:
		return "PAR"
	case Slp:
		return "SLP"
	}
	return ""
}

// Move data stored in pokemon move slot
type Move struct {
	ID uint
	PP uint
}

// BoxMon data of mon in box
type BoxMon struct {
	PokemonID uint
	HP        uint
	BoxLevel  uint
	Status    NonVolatileStatus
	Type      [2]uint
	CatchRate byte
	Moves     [4]Move
	OTID      uint
	Exp       int
	EVs       EVStat
	DVs       DVStat
}

// EVStat Effort Value Japanees:努力値
type EVStat struct {
	HP      uint
	Attack  uint
	Defense uint
	Speed   uint
	SpAtk   uint
	SpDef   uint // unused in gen1
}

// DVStat Determinant values Japanese:個体値
type DVStat struct {
	Attack  uint
	Defense uint
	Speed   uint
	SpAtk   uint
	SpDef   uint // unused in gen1
}

// BoxMons box mon data in game
var BoxMons = []BoxMon{}

// DayCareMon data of mon in daycare
type DayCareMon struct{}

// DayCareMons daycare mon data in game
// NOTE: Considering PokemonGSC, multiple mons can be taken.
var DayCareMons = []DayCareMon{}

// PartyMon data of mon in party
type PartyMon struct {
	Initialized bool
	BoxMon
	Level   uint
	MaxHP   uint
	Attack  uint
	Defense uint
	Speed   uint
	SpAtk   uint
	SpDef   uint
	OTName  string
	Nick    string
}

// PartyMons party mon data in game
var PartyMons = [6]PartyMon{}

// PartyMonLen return a number of party pokemons
func PartyMonLen() int {
	for i, mon := range PartyMons {
		if !mon.Initialized {
			return i
		}
	}
	return 6
}
