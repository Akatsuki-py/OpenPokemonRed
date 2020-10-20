package game

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/widget"
	"pokered/pkg/world"

	"github.com/hajimehoshi/ebiten"
)

// Game implements ebiten.Game interface.
type Game struct {
	frame uint
}

// Update proceeds the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	if g.frame == 0 {
		initialize()
	}
	util.BlackScreen(store.TileMap)
	// debug(g, 10)
	exec()
	vBlank()
	g.frame++

	if g.frame%60 == 0 {
		second()
	}

	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(store.FilteredTileMap(), nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}
func debug(g *Game, frame int) {
	if frame >= 0 && int(g.frame) != frame {
		return
	}
	{
	}
}

func exec() {
	if store.DelayFrames > 0 {
		store.DelayFrames--
		return
	}
	switch m := mode(); m {
	case Overworld:
		execOverworld()
	case Script:
		execScript()
	}
}

func vBlank() {
	p := store.SpriteData[0]

	joypad.ReadJoypad()
	store.DecFrameCounter()
	audio.FadeOutAudio()
	world.VBlank(p.MapXCoord, p.MapYCoord, p.DeltaX, p.DeltaY, p.WalkCounter, p.Direction)
	sprite.VBlank()
	menu.VBlank()
	widget.VBlank()
	text.VBlank()
}

func second() {}
