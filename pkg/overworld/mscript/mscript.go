package mscript

import "pokered/pkg/data/worldmap"

func Run(mapID int) {
	switch mapID {
	case worldmap.PALLET_TOWN:
		palletTownScript()
	}
}
