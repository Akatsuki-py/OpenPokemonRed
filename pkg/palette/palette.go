package pal

import (
	"pokered/pkg/script"
	"pokered/pkg/store"
)

func LoadGBPal() {
	store.Palette = 5
}

func GBFadeOutToBlack() {
	script.SetID(script.FadeOutToBlack)
	store.FadeCounter = 4
}
