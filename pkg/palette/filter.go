package palette

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

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
	case 6:
		return filter6(img)
	case 7:
		return filter7(img)
	case 8:
		return filter8(img)
	}
	return filter5(img)
}

func filter1(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImage(8*20, 8*18)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0, 40, 0, 0xff})
	result.DrawImage(sheet, nil)

	return result
}

func filter2(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImageFromImage(target)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0, 80, 0, 0xf7})
	result.DrawImage(sheet, nil)

	return result
}

func filter3(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImageFromImage(target)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0, 80, 0, 0xdf})
	result.DrawImage(sheet, nil)

	return result
}

func filter4(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImageFromImage(target)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0, 80, 0, 0xcf})
	result.DrawImage(sheet, nil)

	return result
}

func filter5(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImage(8*20, 8*18)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{95, 125, 100, 0xff})
	result.DrawImage(sheet, nil)

	src0 := ebiten.NewImageFromImage(target)
	op0 := &ebiten.DrawImageOptions{}
	op0.ColorM.Scale(0.65, 0.85, 0.65, 0.95)
	result.DrawImage(src0, op0)

	return result
}

func filter6(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImageFromImage(target)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0xb0, 0xf8, 0xb0, 0xcf})
	result.DrawImage(sheet, nil)

	return result
}

func filter7(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImageFromImage(target)

	sheet := ebiten.NewImage(8*20, 8*18)
	sheet.Fill(color.NRGBA{0xb0, 0xf8, 0xb0, 0xef})
	result.DrawImage(sheet, nil)

	return result
}

func filter8(target *ebiten.Image) *ebiten.Image {
	if target == nil {
		return target
	}

	result := ebiten.NewImage(8*20, 8*18)
	result.Fill(color.NRGBA{95, 125, 100, 0xff})

	src0 := ebiten.NewImageFromImage(result)
	op0 := &ebiten.DrawImageOptions{}
	op0.ColorM.Scale(0.65, 0.85, 0.65, 0.95)
	result.DrawImage(src0, op0)

	return result
}
