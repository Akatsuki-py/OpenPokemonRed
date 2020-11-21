package store

import (
	"pokered/pkg/data/move"
	"pokered/pkg/data/pkmnd"
)

// Move data stored in pokemon move slot
type Move struct {
	ID    uint
	CurPP uint // CurPP <= MaxPP
	MaxPP uint // affected by PPUP
}

// BoxMon data of mon in box
type BoxMon struct {
	ID        uint
	HP        uint
	BoxLevel  uint
	Status    pkmnd.NonVolatileStatus
	Type      [2]uint
	CatchRate byte
	Moves     [4]Move
	OTID      uint
	Exp       int
	EVs       EVStat
	DVs       DVStat
	OTName    string
	Nick      string
}

// EVStat Effort Value Japanese:努力値
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
	*BoxMon
	Level   uint
	MaxHP   uint
	Attack  uint
	Defense uint
	Speed   uint
	SpAtk   uint
	SpDef   uint
}

// PartyMons party mon data in game
var PartyMons = [6]PartyMon{*defaultPartyMon(), *defaultPartyMon()}

// PartyMonLen return a number of party pokemons
func PartyMonLen() int {
	for i, mon := range PartyMons {
		if !mon.Initialized {
			return i
		}
	}
	return 6
}

func defaultPartyMon() *PartyMon {
	return &PartyMon{
		Initialized: true,
		BoxMon: &BoxMon{
			ID:        pkmnd.CHARMANDER,
			HP:        22,
			BoxLevel:  6,
			Status:    pkmnd.OK,
			Type:      [2]uint{pkmnd.Fire},
			CatchRate: 255,
			Moves: [4]Move{
				{move.SCRATCH, 35, 35},
				{move.GROWL, 40, 40},
			}, // scratch growl
			OTID:   48024,
			Exp:    205,
			EVs:    EVStat{},
			DVs:    DVStat{},
			OTName: "RED",
			Nick:   "CHARMANDER",
		},
		Level:   6,
		MaxHP:   22,
		Attack:  12,
		Defense: 11,
		Speed:   13,
		SpAtk:   11,
	}
}
