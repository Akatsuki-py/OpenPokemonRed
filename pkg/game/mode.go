package game

import "pokered/pkg/text"

const (
	Overworld uint = iota
	Text
)

func mode() uint {
	if isText() {
		return Text
	}
	return Overworld
}

func isText() bool {
	return len([]rune(text.CurText)) > 0
}
