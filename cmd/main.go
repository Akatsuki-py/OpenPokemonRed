package main

import (
	"flag"
	"fmt"
	"os"
	"pokered/pkg/game"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	exitCodeOK int = iota
	exitCodeError
)

const (
	version = "Develop"
	title   = "PokemonRed"
)

func main() {
	os.Exit(Run())
}

// Run game
func Run() int {
	var (
		showVersion = flag.Bool("v", false, "show version")
	)
	flag.Parse()
	if *showVersion {
		fmt.Println(title+":", version)
		return exitCodeOK
	}

	g := &game.Game{}
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(160*2, 144*2)
	if err := ebiten.RunGame(g); err != nil {
		return exitCodeError
	}
	return exitCodeOK
}
