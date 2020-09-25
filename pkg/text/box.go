package text

type boxArea struct {
	X, Y, Z, W, H int
}

type Box interface {
	area() boxArea
}

func DrawTextBox(box Box) {
}
