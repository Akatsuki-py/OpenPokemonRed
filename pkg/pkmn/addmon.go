package pkmn

import (
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/store"
	"strings"
)

func AddPlayerPartyMon() bool {
	return false
}
func AddRivalPartyMon() bool {
	return false
}
func addPartyMon(isRival bool) bool {
	if isRival {
		if store.RivalPartyMonLen() >= 6 {
			return false
		}
	} else {
		if store.PartyMonLen() >= 6 {
			return false
		}
	}
	return false
}

func NewBoxMon(id, level uint) *store.BoxMon {
	header := pkmnd.Header(id)
	return &store.BoxMon{
		ID:        id,
		BoxLevel:  level,
		Status:    pkmnd.OK,
		Type:      header.Type,
		CatchRate: header.CatchRate,
		Moves:     [4]store.Move{},
		Exp:       205,
		DVs: store.DVStat{
			Attack:  9,
			Defense: 8,
			Speed:   8,
			SpAtk:   8,
			SpDef:   8,
		},
		Nick: strings.ToUpper(header.Name),
	}
}

// NewMoves is temporary implement
func NewMoves(id, level uint) [4]store.Move {
	return [4]store.Move{}
}
