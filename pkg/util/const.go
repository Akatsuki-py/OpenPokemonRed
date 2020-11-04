package util

type Direction = uint

var (
	Black     = [3]byte{0x00, 0x00, 0x00}
	GrayBlack = [3]byte{0x60, 0x60, 0x60}
	GrayWhite = [3]byte{0xa8, 0xa8, 0xa8}
	White     = [3]byte{0xf8, 0xf8, 0xf8}
)

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
