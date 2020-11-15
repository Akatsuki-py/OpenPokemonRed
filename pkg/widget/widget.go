package widget

import (
	"pokered/pkg/screen"
)

// VBlank script executed in VBlank
func VBlank() {
	if trainerCard != nil {
		screen.AddLayerOnTop("widget/trainercard", trainerCard, 0, 0)
	}
	if name.screen != nil {
		screen.AddLayerOnTop("widget/name", name.screen, 0, 0)
	}
	if partyMenu != nil {
		screen.AddLayerOnTop("widget/partymenu", partyMenu, 0, 0)
	}
}
