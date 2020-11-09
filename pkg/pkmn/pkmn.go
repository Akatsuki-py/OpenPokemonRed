package pkmn

const (
	Green uint = iota
	Yellow
	Red
)

// GetHealthBarColor return HP bar color
func GetHealthBarColor(hp, maxHP uint) uint {
	px := int(hp * 48 / maxHP)
	switch {
	case px >= 27:
		return Green
	case px >= 10:
		return Yellow
	default:
		return Red
	}
}

// PartyMonSpeeds animation speed in pokemon SD icon
// if this value is 5 frame, each icon take 5 frame (0(0:5) -> 1(5:10) -> 0(10:15) -> 1(15:20) -> ...)
func PartyMonSpeeds(hp, maxHP uint) uint {
	c := GetHealthBarColor(hp, maxHP)
	switch c {
	case Green:
		return 2 // 5/3
	case Yellow:
		return 16 / 3
	case Red:
		return 32 / 3
	default:
		return 2 // 5/3
	}
}
