package palette

import (
	"pokered/pkg/store"
)

func LoadGBPal() {
	store.Palette = 5
}

func GBFadeOutToBlack() {
	store.FadeCounter = 4
}

func GBFadeOutToWhite() {
	store.FadeCounter = 4
}
