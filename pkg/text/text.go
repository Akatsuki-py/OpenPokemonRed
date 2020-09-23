package text

type TextBox interface {
	XYZ() (int, int, int)
}

func DrawTextBox(box TextBox) {
}
