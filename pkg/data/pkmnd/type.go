package pkmnd

// Pokemon type
const (
	Bug  uint = iota + 1
	Dark      // unused in gen1
	Dragon
	Electric
	Fairy // unused in gen1
	Fighting
	Fire
	Flying
	Ghost
	Grass
	Ground
	Ice
	Normal
	Poison
	Psychic
	Rock
	Steel // unused in gen1
	Water
)

func TypeString(t uint) string {
	switch t {
	case Bug:
		return "BUG"
	case Dark:
		return "DARK"
	case Dragon:
		return "DRAGON"
	case Electric:
		return "ELECTRIC"
	case Fairy:
		return "FAIRY"
	case Fighting:
		return "FIGHTING"
	case Fire:
		return "FIRE"
	case Flying:
		return "FLYING"
	case Ghost:
		return "GHOST"
	case Grass:
		return "GRASS"
	case Ground:
		return "GROUND"
	case Ice:
		return "ICE"
	case Normal:
		return "NORMAL"
	case Poison:
		return "POISON"
	case Psychic:
		return "PSYCHIC"
	case Rock:
		return "ROCK"
	case Steel:
		return "STEEL"
	case Water:
		return "WATER"
	default:
		return ""
	}
}
