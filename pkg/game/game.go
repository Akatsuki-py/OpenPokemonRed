package game

import (
	"pokered/pkg/audio"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/overworld"
	pal "pokered/pkg/palette"
	scr "pokered/pkg/screen"
	"pokered/pkg/script"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/widget"
	"pokered/pkg/world"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	frame uint
}

// Update proceeds the game state.
func (g *Game) Update() error {
	if g.frame == 0 {
		initialize()
		// initializeWorld()
	}
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
	screen.DrawImage(pal.Filter(scr.TileMap(), store.Palette), nil)
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
		overworld.ExecOverworld()
	case Script:
		execScript()
	}
}

func vBlank() {
	joypad.ReadJoypad()
	store.DecFrameCounter()
	audio.FadeOutAudio()

	if script.InTitle() {
		scr.AddLayerOnTop("title", script.TitleScreen, 0, 0)
		script.TitleScreen = util.NewImage()
	}

	if script.InOakSpeech() {
		scr.AddLayerOnTop("oakspeech", script.OakSpeechScreen, 0, 0)
	}

	if isOverworld() || store.SpriteData[0] != nil {
		p := store.SpriteData[0]
		world.VBlank(p.MapXCoord, p.MapYCoord, p.DeltaX, p.DeltaY, p.WalkCounter, p.Direction)
		sprite.VBlank()
	}
	menu.VBlank()
	widget.VBlank()
	text.VBlank()

	scr.VBlank()
}

func second() {}
