package text

import "pokered/pkg/util"

var cursor = 0

func Cursor() (x, y util.Tile) {
	y = cursor / 20
	x = cursor % 20
	return x, y
}

func Next() {
	cursor++
	if cursor > 20*18 {
		cursor = 0
	}
}

func Seek(x, y util.Tile) {
	if x >= 0 && x <= 20 && y >= 0 && y <= 18 {
		cursor = y*20 + x
	}
}
