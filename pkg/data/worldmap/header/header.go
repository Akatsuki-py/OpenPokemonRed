package header

import (
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

// Header Map header
type Header struct {
	// Tileset ID
	Tileset uint

	// Block(32×32) Height
	Height uint
	// Block(32×32) Width
	Width uint

	// Block data
	blk []byte

	// Map Text data
	Text        []string
	Connections Connections
}

type Connections struct {
	North Connection
	South Connection
	West  Connection
	East  Connection
}

type Connection struct {
	OK        bool
	DestMapID int
	Coords    []uint
}

// Get Map Header
func Get(id int) *Header {
	switch id {
	case worldmap.AGATHAS_ROOM:
		return AgathasRoom
	case worldmap.PALLET_TOWN:
		return PalletTown
	case worldmap.REDS_HOUSE_1F:
		return RedsHouse1F
	case worldmap.REDS_HOUSE_2F:
		return RedsHouse2F
	case worldmap.ROUTE_1:
		return Route1
	case worldmap.ROUTE_21:
		return Route21
	case worldmap.OAKS_LAB:
		return OaksLab
	case worldmap.BLUES_HOUSE:
		return BluesHouse
	default:
		util.NotRegisteredError("header.Get", id)
	}
	return nil
}

// Blk get block data
func (h *Header) Blk(index int) byte {
	if int(index) >= len(h.blk) {
		return 0
	}
	return h.blk[index]
}

// BlkLen get block data length
func (h *Header) BlkLen() int {
	return len(h.blk)
}
