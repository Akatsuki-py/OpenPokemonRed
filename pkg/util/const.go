package util

type Direction = uint

// color
var (
	Black     = [3]byte{0x00, 0x00, 0x00}
	GrayBlack = [3]byte{0x60, 0x60, 0x60}
	GrayWhite = [3]byte{0xa8, 0xa8, 0xa8}
	White     = [3]byte{0xf8, 0xf8, 0xf8}
)

// direction
const (
	Stop  Direction = 0xff
	Down  Direction = 0
	Up    Direction = 4
	Left  Direction = 8
	Right Direction = 12
)

// utility string
const (
	Pokedex = "POKéDEX"
	Pokemon = "POKéMON"
)

// movement byte
const (
	None            byte = 0xff
	ChangeDirection byte = 0xe0
	Walk            byte = 0xfe
	Stay            byte = 0xff
)

const MaxLevel = 100
