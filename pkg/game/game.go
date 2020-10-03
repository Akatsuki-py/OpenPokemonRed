package game

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/store"

	"github.com/hajimehoshi/ebiten"
)

// Game implements ebiten.Game interface.
type Game struct {
	frame uint
}

// Update proceeds the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	if g.frame == 0 {
		setup()
	}
	exec()
	vBlank()
	g.frame++
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(store.TileMap, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

func setup() {
	audio.PlayMusic(audio.MUSIC_FINAL_BATTLE)
}

func exec() {
	joypad.ReadJoypad()
	if store.DelayFrames > 0 {
		store.DelayFrames--
		return
	}
	switch m := mode(); m {
	case Overworld:
		execOverworld()
	case Text:
		execText()
	case Menu:
		execMenu()
	}
}

func vBlank() {
	joypad.ReadJoypad()
	store.DecFrameCounter()
}
