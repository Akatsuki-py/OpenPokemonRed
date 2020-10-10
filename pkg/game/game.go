package game

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"

	"github.com/hajimehoshi/ebiten"
)

var player = store.SpriteData[0]

// Game implements ebiten.Game interface.
type Game struct {
	frame uint
}

// Update proceeds the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	if g.frame == 0 {
		setup()
	}
	util.BlackScreen()
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
	screen.DrawImage(store.TileMap, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

func setup() {
	sprite.InitPlayer(sprite.Normal)
	sprite.AddSprite("sailor", 2, 2, [2]byte{0xff, 0xff})
	player = store.SpriteData[0]
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
	joypad.ReadJoypad()
	store.DecFrameCounter()
	audio.FadeOutAudio()
	sprite.VBlank()
	menu.VBlank()
	text.VBlank()
}

func second() {}
