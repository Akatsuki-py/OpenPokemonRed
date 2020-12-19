package store

type WalkBikeSurf = int

const (
	WalkState WalkBikeSurf = iota
	BikeState
	SurfState
)

type PlayerState struct {
	Name    string
	Money   uint
	Time    uint
	State   WalkBikeSurf
	Starter uint
}

var Player = PlayerState{
	Name:  "NINTEN",
	Money: 0,
	Time:  0,
	State: WalkState,
}
