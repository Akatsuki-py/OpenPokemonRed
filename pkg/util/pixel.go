package util

// TileToPixel convert pokered tile into ebiten screen pixel
func TileToPixel(x, y int) (int, int) {
	return x * 8 * 2, y * 8 * 2
}

// CoordToPixel convert pokered coord into ebiten screen pixel
func CoordToPixel(x, y int) (int, int) {
	return x * 16 * 2, y * 16 * 2
}
