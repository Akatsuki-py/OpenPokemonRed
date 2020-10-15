package object

import "pokered/pkg/data/sprdata"

// Object Map object data
type Object struct {
	// Border block
	Border  byte
	Warps   []Warp
	Signs   []Sign
	Sprites []Sprite
	WarpTos []WarpTo
}

// Warp this coord can warp to dest
type Warp struct {
	// warp coord
	XCoord, YCoord int
	DestWarpID     uint
	DestMap        int
}

type Sign struct {
	XCoord, YCoord int
	TextID         uint
}

type Sprite struct {
	ID             sprdata.SpriteID
	XCoord, YCoord int
	MovementBytes  [2]byte
	TextID         uint
}

// WarpTo other map can warp to this WarpTo
type WarpTo struct {
	XCoord, YCoord int
	Width          uint
}
