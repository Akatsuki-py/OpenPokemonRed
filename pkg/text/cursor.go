package text

import "pokered/pkg/util"

var cursor = 0

// Cursor getter
func Cursor() (x, y util.Tile) {
	y = cursor / 20
	x = cursor % 20
	return x, y
}

// Next set cursor on next position
func Next() {
	cursor++
	if cursor > 20*18 {
		cursor = 0
	}
}

// Seek cursor
func Seek(x, y util.Tile) {
	if y*20+x >= 0 && y*20+x <= 360 {
		cursor = y*20 + x
	}
}
