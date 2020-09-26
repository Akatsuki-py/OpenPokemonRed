package util

type Tile = int
type Coord = int

// TileToPixel convert pokered tile into ebiten screen pixel
func TileToPixel(x, y Tile) (int, int) {
	return x * 8 * 2, y * 8 * 2
}

// TileToFPixel convert pokered tile into ebiten screen pixel
func TileToFPixel(x, y Tile) (float64, float64) {
	return float64(x * 8 * 2), float64(y * 8 * 2)
}

// CoordToPixel convert pokered coord into ebiten screen pixel
func CoordToPixel(x, y Coord) (int, int) {
	return x * 16 * 2, y * 16 * 2
}
