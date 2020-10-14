package header

import "pokered/pkg/data/worldmap"

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
}

// Get Map Header
func Get(id uint) *Header {
	switch id {
	case worldmap.AGATHAS_ROOM:
		return AgathasRoom
	case worldmap.PALLET_TOWN:
		return PalletTown
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
