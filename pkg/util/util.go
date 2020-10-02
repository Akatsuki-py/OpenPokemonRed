package util

import (
	"image/color"
	"pokered/pkg/store"
	"reflect"
	"runtime"
	"strings"

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
func SetBit(data byte, bit uint) byte {
	return data | (1 << bit)
}
func ResBit(data byte, bit uint) byte {
	return data & ^(1 << bit)
}

func BlackScreen() {
	store.TileMap.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
}
func WhiteScreen() {
	store.TileMap.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
}

// ClearScreenArea clear h×w tiles from (x, y)
func ClearScreenArea(x, y Tile, h, w uint) {
	width, height := TileToPixel(Tile(w), Tile(h))
	sheet, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0xf8, 0xf8, 0xf8, 0xff})
	DrawImage(sheet, x, y)
}

func DrawImage(i *ebiten.Image, x, y Tile) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(TileToFPixel(x, y))
	store.TileMap.DrawImage(i, op)
}
