package game

import (
	"pokered/pkg/store"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	Overworld()
	Text()
	VBlank()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(store.TileMap, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160 * 2, 144 * 2
}

func Overworld() {}

func Text() {}

func VBlank() {}
