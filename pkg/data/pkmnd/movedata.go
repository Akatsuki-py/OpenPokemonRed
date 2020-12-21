package pkmnd

const Inevitable = 1000

type MoveData struct {
	ID        uint
	Accuracy  uint
	BasePower uint
	PP        uint
	Type      uint
}

var moveMap = map[uint]MoveData{}

func GetMoveData(id uint) MoveData {
	m, ok := moveMap[id]
	if !ok {
		return tackle
	}
	return m
}

// MovePP returns move PP
func MovePP(id uint) uint {
	m := GetMoveData(id)
	return m.PP
}

var tackle = MoveData{
	ID:        TACKLE,
	Accuracy:  95,
	BasePower: 35,
	PP:        35,
	Type:      Normal,
}

func init() {
	moveMap[TACKLE] = tackle

	var growl = MoveData{
		ID:       GROWL,
		Accuracy: 100,
		PP:       40,
		Type:     Normal,
	}
	moveMap[GROWL] = growl

	var leechSeed = MoveData{
		ID:       LEECH_SEED,
		Accuracy: 90,
		PP:       10,
		Type:     Grass,
	}
	moveMap[LEECH_SEED] = leechSeed

	var vineWhip = MoveData{
		ID:        VINE_WHIP,
		Accuracy:  100,
		BasePower: 45,
		PP:        25,
		Type:      Grass,
	}
	moveMap[VINE_WHIP] = vineWhip

	var poisonpowder = MoveData{
		ID:       POISONPOWDER,
		Accuracy: 75,
		PP:       35,
		Type:     Poison,
	}
	moveMap[POISONPOWDER] = poisonpowder

	var razorLeaf = MoveData{
		ID:        RAZOR_LEAF,
		Accuracy:  95,
		BasePower: 55,
		PP:        25,
		Type:      Grass,
	}
	moveMap[RAZOR_LEAF] = razorLeaf

	var growth = MoveData{
		ID:       GROWTH,
		Accuracy: Inevitable,
		PP:       20,
		Type:     Normal,
	}
	moveMap[GROWTH] = growth

	var sleepPowder = MoveData{
		ID:       SLEEP_POWDER,
		Accuracy: 75,
		PP:       15,
		Type:     Grass,
	}
	moveMap[SLEEP_POWDER] = sleepPowder

	var solarbeam = MoveData{
		ID:        SOLARBEAM,
		Accuracy:  100,
		BasePower: 120,
		PP:        10,
		Type:      Grass,
	}
	moveMap[SOLARBEAM] = solarbeam

	var scratch = MoveData{
		ID:        SCRATCH,
		Accuracy:  100,
		BasePower: 40,
		PP:        35,
		Type:      Normal,
	}
	moveMap[SCRATCH] = scratch

	var ember = MoveData{
		ID:        EMBER,
		Accuracy:  100,
		BasePower: 40,
		PP:        25,
		Type:      Fire,
	}
	moveMap[EMBER] = ember

	var leer = MoveData{
		ID:       LEER,
		Accuracy: 100,
		PP:       30,
		Type:     Normal,
	}
	moveMap[LEER] = leer

	var rage = MoveData{
		ID:       RAGE,
		Accuracy: 100,
		PP:       30,
		Type:     Normal,
	}
	moveMap[RAGE] = rage

	var slash = MoveData{
		ID:        SLASH,
		Accuracy:  100,
		BasePower: 70,
		PP:        15,
		Type:      Normal,
	}
	moveMap[SLASH] = slash

	var flamethrower = MoveData{
		ID:        FLAMETHROWER,
		Accuracy:  100,
		BasePower: 90,
		PP:        15,
		Type:      Fire,
	}
	moveMap[FLAMETHROWER] = flamethrower

	var fireSpin = MoveData{
		ID:        FIRE_SPIN,
		Accuracy:  85,
		BasePower: 35,
		PP:        15,
		Type:      Fire,
	}
	moveMap[FIRE_SPIN] = fireSpin

	var tailWhip = MoveData{
		ID:       TAIL_WHIP,
		Accuracy: 100,
		PP:       30,
		Type:     Normal,
	}
	moveMap[TAIL_WHIP] = tailWhip

	var bubble = MoveData{
		ID:        BUBBLE,
		Accuracy:  100,
		BasePower: 20,
		PP:        30,
		Type:      Water,
	}
	moveMap[BUBBLE] = bubble

	var waterGun = MoveData{
		ID:        WATER_GUN,
		Accuracy:  100,
		BasePower: 40,
		PP:        25,
		Type:      Water,
	}
	moveMap[WATER_GUN] = waterGun

	var bite = MoveData{
		ID:        BITE,
		Accuracy:  100,
		BasePower: 60,
		PP:        25,
		Type:      Normal, // Dark (>= gen2)
	}
	moveMap[BITE] = bite

	var withdraw = MoveData{
		ID:       WITHDRAW,
		Accuracy: Inevitable,
		PP:       40,
		Type:     Water,
	}
	moveMap[WITHDRAW] = withdraw

	var skullBash = MoveData{
		ID:        SKULL_BASH,
		Accuracy:  100,
		BasePower: 130,
		PP:        10,
		Type:      Normal,
	}
	moveMap[SKULL_BASH] = skullBash

	var hydroPump = MoveData{
		ID:        HYDRO_PUMP,
		Accuracy:  80,
		BasePower: 110,
		PP:        5,
		Type:      Water,
	}
	moveMap[HYDRO_PUMP] = hydroPump
}
