package util

type Direction = uint

const (
	Down  Direction = 0
	Up    Direction = 4
	Left  Direction = 8
	Right Direction = 12
)

const (
	Pokedex = "POKéDEX"
	Pokemon = "POKéMON"
)

const (
	None            byte = 0xff
	ChangeDirection byte = 0xe0
	Walk            byte = 0xfe
	Stay            byte = 0xff
)
