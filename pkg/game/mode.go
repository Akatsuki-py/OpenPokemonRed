package game

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
	return false
}
