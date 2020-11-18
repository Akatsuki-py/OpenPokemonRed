package pkmnd

// NonVolatileStatus type for non-volatile statuses
type NonVolatileStatus uint

// non-volatile statuses
const (
	OK  NonVolatileStatus = 0
	Psn NonVolatileStatus = 3
	Brn NonVolatileStatus = 4
	Frz NonVolatileStatus = 5
	Par NonVolatileStatus = 6
	Slp NonVolatileStatus = 7
)

func (n *NonVolatileStatus) String() string {
	switch *n {
	case Psn:
		return "PSN"
	case Brn:
		return "BRN"
	case Frz:
		return "FRZ"
	case Par:
		return "PAR"
	case Slp:
		return "SLP"
	}
	return ""
}
