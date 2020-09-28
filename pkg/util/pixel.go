package util

type Tile = int
type Coord = int
type Pixel = int

// TileToPixel convert pokered tile into ebiten screen pixel
func TileToPixel(x, y Tile) (Pixel, Pixel) {
	return x * 8, y * 8
}

// TileToFPixel convert pokered tile into ebiten screen pixel
func TileToFPixel(x, y Tile) (float64, float64) {
	return float64(x * 8), float64(y * 8)
}

// CoordToPixel convert pokered coord into ebiten screen pixel
func CoordToPixel(x, y Coord) (Pixel, Pixel) {
	return x * 16, y * 16
}
