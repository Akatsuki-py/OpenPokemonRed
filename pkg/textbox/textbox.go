package textbox

// Area textbox area
type Area struct {
	X, Y, Z, W, H int
}

// Box textbox interface
type Box interface {
	area() Area
}

// DrawTextbox draw text box
func DrawTextbox(box Box) {
}
