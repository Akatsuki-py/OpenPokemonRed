package pkmn

import (
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"strings"
)

func AddPlayerPartyMon(id, level uint) bool {
	offset := store.PartyMonLen()
	if offset >= 6 {
		return false
	}

	sp := uint(util.Random()) % 16
	dvs := store.DVStat{
		Attack:  uint(util.Random()) % 16,
		Defense: uint(util.Random()) % 16,
		Speed:   uint(util.Random()) % 16,
		SpAtk:   sp,
		SpDef:   sp,
	}
	store.PartyMons[offset] = *NewPartyMon(id, level, dvs, 0, store.Player.Name)
	return true
}

func AddRivalPartyMon(id, level uint) bool {
	offset := store.RivalPartyMonLen()
	if offset >= 6 {
		return false
	}

	dvs := store.DVStat{
		Attack:  9,
		Defense: 8,
		Speed:   8,
		SpAtk:   8,
		SpDef:   8,
	}
	store.Rival.PartyMons[offset] = *NewPartyMon(id, level, dvs, 0, store.Rival.Name)
	return true
}

// NewPartyMon creates new PartyMon data
func NewPartyMon(id, level uint, dvs store.DVStat, OTID uint, OTName string) *store.PartyMon {
	header := pkmnd.Header(id)
	boxMon := NewBoxMon(id, level, dvs, OTID, OTName)
	return &store.PartyMon{
		Initialized: true,
		BoxMon:      boxMon,
		Level:       level,
		MaxHP:       boxMon.HP,
		Attack:      CalcStat(header.BaseStatsGen1.Attack, boxMon.DVs.Attack, 0, level),
		Defense:     CalcStat(header.BaseStatsGen1.Defense, boxMon.DVs.Defense, 0, level),
		Speed:       CalcStat(header.BaseStatsGen1.Speed, boxMon.DVs.Speed, 0, level),
		SpAtk:       CalcStat(header.BaseStatsGen1.Special, boxMon.DVs.SpAtk, 0, level),
		SpDef:       CalcStat(header.BaseStatsGen1.Special, boxMon.DVs.SpDef, 0, level),
	}
}

// NewBoxMon creates new BoxMon data
func NewBoxMon(id, level uint, dvs store.DVStat, OTID uint, OTName string) *store.BoxMon {
	header := pkmnd.Header(id)
	return &store.BoxMon{
		ID:        id,
		HP:        CalcHP(header.BaseStatsGen1.HP, CalcHPDV(dvs), 0, level),
		BoxLevel:  level,
		Status:    store.OK,
		Type:      header.Type,
		CatchRate: header.CatchRate,
		Moves:     NewMoves(id, level),
		OTID:      OTID,
		Exp:       int(CalcLvExp(level, header.GrowthRate)),
		DVs:       dvs,
		OTName:    OTName,
		Nick:      strings.ToUpper(header.Name),
	}
}

// NewMoveIDs is temporary implement
func NewMoveIDs(id, level uint) [4]uint {
	header := pkmnd.Header(id)
	result := header.Lv0MoveIDs

	i := 0
	for _, m := range header.LvMoves {
		moveLV, moveID := m[0], m[1]
		if moveLV > level {
			break
		}
		result[i] = moveID
		i = (i + 1) % 4
	}
	return result
}

func NewMoves(id, level uint) [4]store.Move {
	moveIDs := NewMoveIDs(id, level)
	result := [4]store.Move{}
	for i, id := range moveIDs {
		if id == 0 {
			break
		}
		pp := pkmnd.MovePP(id)
		result[i] = store.Move{
			ID:    id,
			CurPP: pp,
			MaxPP: pp,
		}
	}
	return result
}
