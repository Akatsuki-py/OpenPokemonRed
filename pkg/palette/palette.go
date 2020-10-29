package palette

import (
	"pokered/pkg/store"
)

func LoadGBPal() {
	store.Palette = 5
}

func GBFadeOutToBlack() {
	store.SetScriptID(store.FadeOutToBlack)
	store.FadeCounter = 4
}

func GBFadeOutToWhite() {
	store.SetScriptID(store.FadeOutToWhite)
	store.FadeCounter = 4
}
