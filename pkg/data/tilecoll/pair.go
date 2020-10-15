package tilecoll

import "pokered/pkg/data/tileset"

func matchPairs(tile1, tile2 byte, pairs [][2]byte) bool {
	for _, pair := range pairs {
		if matchPair(tile1, tile2, pair) {
			return true
		}
	}
	return false
}

func matchPair(tile1, tile2 byte, pair [2]byte) bool {
	return (tile1 == pair[0] && tile2 == pair[1]) || (tile1 == pair[1] && tile2 == pair[0])
}

// IsCollisionPair check between tile1 and tile2 is passable
func IsCollisionPair(tilesetID uint, tile1, tile2 byte, isWater bool) bool {
	if isWater {
		switch tilesetID {
		case tileset.Forest:
			return matchPairs(tile1, tile2, [][2]byte{{0x14, 0x2e}, {0x48, 0x2e}})
		case tileset.Cavern:
			return matchPair(tile1, tile2, [2]byte{0x14, 0x05})
		}
		return false
	}

	switch tilesetID {
	case tileset.Cavern:
		pairs := [][2]byte{
			{0x20, 0x05},
			{0x41, 0x05},
			{0x2a, 0x05},
			{0x05, 0x21},
		}
		return matchPairs(tile1, tile2, pairs)
	case tileset.Forest:
		pairs := [][2]byte{
			{0x30, 0x2e},
			{0x52, 0x2e},
			{0x55, 0x2e},
			{0x56, 0x2e},
			{0x20, 0x2e},
			{0x5e, 0x2e},
			{0x5f, 0x2e},
		}
		return matchPairs(tile1, tile2, pairs)
	}

	return false
}
