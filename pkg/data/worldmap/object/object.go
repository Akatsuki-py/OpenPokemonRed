package object

import (
	"pokered/pkg/data/sprdata"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/util"
)

// Object Map object data
type Object struct {
	// Border block
	Initialized bool
	Border      byte
	Warps       []Warp
	Signs       []Sign
	Sprites     []Sprite
	WarpTos     []WarpTo
}

// Warp this coord can warp to dest
type Warp struct {
	// warp coord
	XCoord, YCoord int
	DestWarpID     int
	DestMap        int
}

type Sign struct {
	XCoord, YCoord int
	TextID         int
}

type Sprite struct {
	ID             sprdata.SpriteID
	XCoord, YCoord int
	MovementBytes  [2]byte
	TextID         int
}

// WarpTo other map can warp to this WarpTo
type WarpTo struct {
	XCoord, YCoord int
	Width          uint
}

// Get Map Object
func Get(id int) *Object {
	switch id {
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
	}
	util.NotRegisteredError("object.Get", id)
	return nil
}
