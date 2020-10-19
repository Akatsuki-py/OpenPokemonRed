package world

import (
	"pokered/pkg/data/worldmap/warp"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

func StandOnDoor(xCoord, yCoord int) bool {
	tilesetID, tileID := CurTileID(xCoord, yCoord)

	doors, ok := warp.DoorTileIDs[tilesetID]
	if !ok {
		return false
	}

	for _, d := range doors {
		if d == byte(tileID) {
			if store.DoorFlag.Check {
				store.DoorFlag.Step = true
			}
			return true
		}
	}

	return false
}

func StandOnWarp(xCoord, yCoord int) bool {
	tilesetID, tileID := CurTileID(xCoord, yCoord)

	doors, ok := warp.WarpTileIDs[tilesetID]
	if !ok {
		return false
	}

	for _, d := range doors {
		if d == byte(tileID) {
			return true
		}
	}

	return false
}

// FaceEdgeOfMap check sprite faces edge of the current map
func FaceEdgeOfMap(xCoord, yCoord int, direction uint) bool {
	switch direction {
	case util.Up:
		return yCoord == 0
	case util.Down:
		return yCoord == int(CurWorld.Header.Height*2-1)
	case util.Left:
		return xCoord == 0
	case util.Right:
		return xCoord == int(CurWorld.Header.Width*2-1)
	}

	return false
}

func IsCurTileset(tilesetID uint) bool {
	return tilesetID == CurWorld.Header.Tileset
}
