package pkmn

import (
	"fmt"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/store"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func Picture(id uint, flip bool) *ebiten.Image {
	name := pkmnd.Name(id)
	path := fmt.Sprintf("/%s_1.png", name)
	pic := util.OpenImage(store.FS, path)
	if flip {
		dst := ebiten.NewImage(pic.Bounds().Dx(), pic.Bounds().Dy())
		op := &ebiten.DrawImageOptions{}
		util.FlipLR(op, pic.Bounds().Dx(), pic.Bounds().Dy())
		dst.DrawImage(pic, op)
		pic = dst
	}
	return pic
}
