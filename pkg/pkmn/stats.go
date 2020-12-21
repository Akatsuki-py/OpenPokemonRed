/*
	For Japanese
		種族値: BaseStat
		個体値: DV
		努力値: EV
*/

package pkmn

import (
	"math"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// CalcHP calc HP stat
func CalcHP(base, dv, ev, level uint) uint {
	return calcStat(base, dv, ev, level) + level + 10
}

// CalcStat calc Atk,Def,Spd,Sp stat
func CalcStat(base, dv, ev, level uint) uint {
	return calcStat(base, dv, ev, level) + 5
}

func CalcHPDV(dvs store.DVStat) uint {
	dv := uint(0)
	if dvs.Attack%2 == 1 {
		dv += 8
	}
	if dvs.Defense%2 == 1 {
		dv += 4
	}
	if dvs.Speed%2 == 1 {
		dv += 2
	}
	if dvs.SpAtk%2 == 1 {
		dv++
	}
	return dv
}

func calcStat(base, dv, ev, level uint) uint {
	tmp1 := float64((base + dv) * 2)                                         // (base+dv)×2
	tmp2 := math.Min(63, math.Floor(math.Floor(1+math.Sqrt(float64(ev)))/4)) // min(63,floor(floor(1+√ev)÷4))
	result := uint(math.Floor((tmp1 + tmp2) * float64(level) / 100))         // floor{(tmp1+tmp2)×lv÷100}
	return result
}

// CalcExpToLevelUp calc exp needed to level up
func CalcExpToLevelUp(lv, exp, growthRate uint) uint {
	if lv == util.MaxLevel {
		return 0
	}

	nextExp := CalcLvExp(lv+1, growthRate)
	return nextExp - exp
}

func pow(lv, p uint) float64 {
	return math.Pow(float64(lv), float64(p))
}

// CalcLvExp calc exp needed to level up from Lv1 to Lv`lv`
func CalcLvExp(lv, growthRate uint) uint {
	if lv == 1 {
		return 0
	}

	switch growthRate {
	case pkmnd.Exp800k:
		return uint(math.Floor(0.8 * pow(lv, 3)))
	case pkmnd.Exp1000k:
		return uint(pow(lv, 3))
	case pkmnd.Exp1050k:
		return uint(math.Floor(1.2*pow(lv, 3) - 15.*pow(lv, 2) + 100.*float64(lv) - 140.))
	default:
		return uint(pow(lv, 3))
	}
}
