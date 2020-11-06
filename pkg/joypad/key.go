package joypad

import ebiten "github.com/hajimehoshi/ebiten/v2"

func a() bool {
	return ebiten.IsKeyPressed(ebiten.KeyS)
}

func b() bool {
	return ebiten.IsKeyPressed(ebiten.KeyA)
}

func start() bool {
	return ebiten.IsKeyPressed(ebiten.KeyEnter)
}

func sel() bool {
	return ebiten.IsKeyPressed(ebiten.KeyShift)
}

func keyUp() bool {
	return ebiten.IsKeyPressed(ebiten.KeyUp)
}

func keyDown() bool {
	return ebiten.IsKeyPressed(ebiten.KeyDown)
}

func keyRight() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight)
}

func keyLeft() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft)
}
