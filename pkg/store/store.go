package store

import (
	_ "pokered/pkg/data/statik"

	"github.com/rakyll/statik/fs"
)

// Debug set when debug mode.
var Debug bool

// GameFrame frame count
var GameFrame uint = 0

var CurMapScript int = 0

var Palette int = 5
var FadeCounter int = 0

// FS statik filesystem
var FS, _ = fs.New()

// DelayFrames VBlank以外を拒否
var DelayFrames uint

// FrameCounter VBlankごとにデクリメント
// used to control letter print speed
var FrameCounter uint = 0

// DecFrameCounter decrement FrameCounter
// this function is called at every vBlank
func DecFrameCounter() {
	if FrameCounter > 0 {
		FrameCounter--
	}
}

// TMName wcf4b
var TMName = ""

// BagItems items in bag
// [A@1, B@2, ...]
var BagItems = []string{}

// EventMap event ID -> flag
var EventMap = map[uint]bool{}

// TwoOptionResult two option index
var TwoOptionResult = uint(0)
