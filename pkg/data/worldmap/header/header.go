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
	Blk []byte
}

// Get Map Header
func Get(id uint) *Header {
	switch id {
	case worldmap.AGATHAS_ROOM:
		return AgathasRoom
	}
	return nil
}
