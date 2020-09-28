package game

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"

	"github.com/hajimehoshi/ebiten"
)

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	exec()
	vBlank()
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(store.TileMap, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160 * 2, 144 * 2
}

func exec() {
	switch m := mode(); m {
	case Overworld:
		execOverworld()
	case Text:
		execText()
	}
}

func vBlank() {
	joypad.ReadJoypad()
}
