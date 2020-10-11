package widget

import (
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// VBlank script executed in VBlank
func VBlank() {
	if trainerCard != nil {
		util.DrawImage(store.TileMap, trainerCard, 0, 0)
	}
	if name.screen != nil {
		util.DrawImage(store.TileMap, name.screen, 0, 0)
	}
}
