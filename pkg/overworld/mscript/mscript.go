package mscript

import (
	"pokered/pkg/data/worldmap"
	"pokered/pkg/store"
	"pokered/pkg/world"
)

var delayed3F bool

func Run(mapID int) {
	switch mapID {
	case worldmap.PALLET_TOWN:
		palletTownScript()
	case worldmap.OAKS_LAB:
		oaksLabScript()
	}
}

func hideObject(id int) {
	world.CurWorld.Object.HS[id] = true
	store.SpriteData[id].Hidden = true
}

func showObject(id int) {
	world.CurWorld.Object.HS[id] = false
	store.SpriteData[id].Hidden = false
}
