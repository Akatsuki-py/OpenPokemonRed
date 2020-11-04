package util

import (
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func Contains(list interface{}, target interface{}) bool {
	if reflect.TypeOf(list).Kind() == reflect.Slice || reflect.TypeOf(list).Kind() == reflect.Array {
		listvalue := reflect.ValueOf(list)
		for i := 0; i < listvalue.Len(); i++ {
			if target == listvalue.Index(i).Interface() {
				return true
			}
		}
	}
	if reflect.TypeOf(target).Kind() == reflect.String && reflect.TypeOf(list).Kind() == reflect.String {
		return strings.Contains(list.(string), target.(string))
	}
	return false
}

// LF return line feed
func LF() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

// XOR exclusive-OR
func XOR(a, b bool) bool {
	return a != b
}

func ReadBit(data byte, bit uint) bool {
	return data>>bit%2 == 1
}
func SetBit(data *byte, bit uint) {
	*data = *data | (1 << bit)
}
func ResBit(data *byte, bit uint) {
	*data = *data & ^(1 << bit)
}

func BlackScreen(target *ebiten.Image) {
	target.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
}
func WhiteScreen(target *ebiten.Image) {
	target.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
}
func FillScreen(target *ebiten.Image, r, g, b byte) {
	target.Fill(color.NRGBA{r, g, b, 0xff})
}

func BlackScreenArea(target *ebiten.Image, x, y Tile, h, w int) {
	width, height := TileToPixel(Tile(w), Tile(h))
	sheet, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	DrawImage(target, sheet, x, y)
}

// ClearScreenArea clear h×w tiles from (x, y)
func ClearScreenArea(target *ebiten.Image, x, y Tile, h, w uint) {
	width, height := TileToPixel(Tile(w), Tile(h))
	sheet, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
	DrawImage(target, sheet, x, y)
}

func DrawPixel(target *ebiten.Image, x, y int, r, g, b byte) {
	clr := color.RGBA{r, g, b, 0xff}
	target.Set(x, y, clr)
}

func DrawImage(target, src *ebiten.Image, x, y Tile) {
	if src == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(TileToFPixel(x, y))
	target.DrawImage(src, op)
}

func DrawImagePixel(target, src *ebiten.Image, x, y int) {
	if src == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	target.DrawImage(src, op)
}

func DrawImageBlock(target, src *ebiten.Image, x, y int) {
	if src == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*32), float64(y*32))
	target.DrawImage(src, op)
}

func NewImage() *ebiten.Image {
	img, _ := ebiten.NewImage(160, 144, ebiten.FilterDefault)
	return img
}

// Random return random value that is in [0, 255]
func Random() byte {
	rand.Seed(time.Now().UnixNano())
	return byte(rand.Intn(256))
}

func OpenImage(fs http.FileSystem, path string) *ebiten.Image {
	f, err := fs.Open(path)
	if err != nil {
		NotFoundFileError(path)
		return nil
	}
	defer f.Close()

	img, _ := png.Decode(f)
	result, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return result
}
