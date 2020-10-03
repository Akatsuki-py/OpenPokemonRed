package main

import (
	"flag"
	"fmt"
	"os"
	"pokered/pkg/game"

	"github.com/hajimehoshi/ebiten"
)

var version string

const (
	exitCodeOK int = iota
	exitCodeError
)

const (
	title = "PokemonRed"
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
		fmt.Println(title+":", getVersion())
		return exitCodeOK
	}

	g := &game.Game{}
	ebiten.SetWindowTitle(title)
	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowSize(160*2, 144*2)
	if err := ebiten.RunGame(g); err != nil {
		return exitCodeError
	}
	return exitCodeOK
}

func getVersion() string {
	if version == "" {
		return "Develop"
	}
	return version
}
