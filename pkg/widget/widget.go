package widget

import (
	"pokered/pkg/screen"
)

// VBlank script executed in VBlank
func VBlank() {
	if trainerCard != nil {
		screen.AddLayer("widget/trainercard", screen.Widget, trainerCard, 0, 0)
	}
	if name.screen != nil {
		screen.AddLayer("widget/name", screen.Widget, name.screen, 0, 0)
	}
	if partyMenu != nil {
		screen.AddLayer("widget/partymenu", screen.Widget, partyMenu, 0, 0)
	}
	if statusScreen != nil {
		screen.AddLayerOnTop("widget/status", statusScreen, 0, 0)
	}
}
