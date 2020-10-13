package blockset

import "pokered/pkg/data/tileset"

// Get blockset
func Get(tilesetID uint) []byte {
	switch tilesetID {
	case tileset.Cavern:
		return Cavern[:]
	case tileset.Cemetery:
		return Cemetery[:]
	case tileset.Club:
		return Club[:]
	case tileset.Facility:
		return Facility[:]
	case tileset.Forest:
		return Forest[:]
	case tileset.Gate:
		return Gate[:]
	case tileset.Gym:
		return Gym[:]
	case tileset.House:
		return House[:]
	case tileset.Interior:
		return Interior[:]
	case tileset.Lab:
		return Lab[:]
	case tileset.Lobby:
		return Lobby[:]
	case tileset.Mansion:
		return Mansion[:]
	case tileset.Overworld:
		return Overworld[:]
	case tileset.Plateau:
		return Plateau[:]
	case tileset.Pokecenter:
		return Pokecenter[:]
	case tileset.RedsHouse:
		return Reds_house[:]
	case tileset.ShipPort:
		return Ship_port[:]
	case tileset.Ship:
		return Ship[:]
	case tileset.Underground:
		return Underground[:]
	}
	return []byte{}
}
