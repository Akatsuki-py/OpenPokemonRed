package pal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func newShapeImage(from *ebiten.Image) *ebiten.Image {
	bounds := from.Bounds()
	img, _ := ebiten.NewImage(bounds.Dx(), bounds.Dy(), ebiten.FilterDefault)
	return img
}

// Filter get palette filtered img
func Filter(img *ebiten.Image, p int) *ebiten.Image {
	switch p {
	case 1:
		return filter1(img)
	case 2:
		return filter2(img)
	case 3:
		return filter3(img)
	case 4:
		return filter4(img)
	case 5:
		return filter5(img)
	}
	return filter5(img)
}

func filter1(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)

	sheet, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0, 40, 0, 0xff})
	result.DrawImage(sheet, nil)

	return result
}

func filter2(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result, _ := ebiten.NewImageFromImage(target, ebiten.FilterDefault)

	sheet, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0, 40, 0, 0xdf})
	result.DrawImage(sheet, nil)

	return result
}

func filter3(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result, _ := ebiten.NewImageFromImage(target, ebiten.FilterDefault)

	sheet, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0, 40, 0, 0xcf})
	result.DrawImage(sheet, nil)

	return result
}

func filter4(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result, _ := ebiten.NewImageFromImage(target, ebiten.FilterDefault)

	sheet, _ := ebiten.NewImage(8*20, 8*18, ebiten.FilterDefault)
	sheet.Fill(color.NRGBA{0, 40, 0, 0xbf})
	result.DrawImage(sheet, nil)

	return result
}

func filter5(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := newShapeImage(target)

	sheet := newShapeImage(target)
	sheet.Fill(color.NRGBA{95, 125, 100, 0xff})
	result.DrawImage(sheet, nil)

	src0, _ := ebiten.NewImageFromImage(target, ebiten.FilterDefault)
	op0 := &ebiten.DrawImageOptions{}
	op0.ColorM.Scale(0.65, 0.85, 0.65, 0.95)
	result.DrawImage(src0, op0)

	return result
}
