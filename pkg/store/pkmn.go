package store

// PartyMon data of mon in party
type PartyMon struct{}

// PartyMons party mon data in game
var PartyMons = [6]PartyMon{}

// BoxMon data of mon in box
type BoxMon struct {
	Species    byte
	HP         int
	BoxLevel   int
	Status     byte
	Type       [2]byte
	CatchRate  byte
	Moves      [4]byte // TODO: replace Move struct
	OTID       byte
	Exp        int
	HPExp      int
	AttackExp  int
	DefenseExp int
	Speed      int
	SpecialExp int
	DVs        int
}

// BoxMons box mon data in game
var BoxMons = []BoxMon{}

// DayCareMon data of mon in daycare
type DayCareMon struct{}

// DayCareMons daycare mon data in game
// NOTE: Considering PokemonGSC, multiple mons can be taken.
var DayCareMons = []DayCareMon{}
