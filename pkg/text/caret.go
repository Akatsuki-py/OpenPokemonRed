package text

import "pokered/pkg/util"

var caret = 0

// Caret getter
func Caret() (x, y util.Tile) {
	y = caret / 20
	x = caret % 20
	return x, y
}

// Next set caret on next position
func Next() {
	caret++
	if caret > 20*18 {
		caret = 0
	}
	printCharDelay()
}

// Seek caret
func Seek(x, y util.Tile) {
	if y*20+x >= 0 && y*20+x <= 360 {
		caret = y*20 + x
	}
}
