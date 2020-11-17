/*
	For Japanese
		種族値: BaseStat
		個体値: DV
		努力値: EV
*/

package pkmn

import "math"

// CalcHP calc HP stat
func CalcHP(base, dv, ev, level uint) uint {
	return calcStat(base, dv, ev, level) + level + 10
}

// CalcStat calc Atk,Def,Spd,Sp stat
func CalcStat(base, dv, ev, level uint) uint {
	return calcStat(base, dv, ev, level) + 5
}

func calcStat(base, dv, ev, level uint) uint {
	tmp1 := float64((base + dv) * 2)                                         // (base+dv)×2
	tmp2 := math.Min(63, math.Floor(math.Floor(1+math.Sqrt(float64(ev)))/4)) // min(63,floor(floor(1+√ev)÷4))
	result := uint(math.Floor((tmp1 + tmp2) * float64(level) / 100))         // floor{(tmp1+tmp2)×lv÷100}
	return result
}
