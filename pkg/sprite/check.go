package sprite

import (
	"pokered/pkg/data/worldmap/ledge"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

// IsStandingOnDoorOrWarp プレイヤーが、ドアタイルかwarpタイルの上に立っているかを調べる
// ref: IsPlayerStandingOnDoorTileOrWarpTile
func IsStandingOnDoorOrWarp(offset int) bool {
	if IsStandingOnDoor(offset) {
		if offset == 0 {
			p := store.SpriteData[offset]
			if p.Direction == util.Down {
				return false
			}
		}

		util.ResBit(&store.D736, 2)
		return true
	}
	if isStandingOnWarp(offset) {
		util.ResBit(&store.D736, 2)
		return true
	}

	return false
}

// IsStandingOnDoor check player is standing on door tile
func IsStandingOnDoor(offset int) bool {
	p := store.SpriteData[offset]
	if store.IsInvalidSprite(0) {
		return false
	}

	return world.StandOnDoor(p.MapXCoord, p.MapYCoord)
}

// isPlayerStandingOnWarp check player is standing on warp tile
func isStandingOnWarp(offset int) bool {
	p := store.SpriteData[offset]
	if store.IsInvalidSprite(0) {
		return false
	}

	return world.StandOnWarp(p.MapXCoord, p.MapYCoord)
}

func isJumpingLedge(offset int) bool {
	if offset > 0 {
		return false
	}

	s := store.SpriteData[offset]
	if s == nil {
		return false
	}

	_, curTileID := world.CurTileID(s.MapXCoord, s.MapYCoord)
	_, frontTileID := world.FrontTileID(s.MapXCoord, s.MapYCoord, s.Direction)

	for _, l := range ledge.LedgeTiles {
		if s.Direction == l.Direction && curTileID == l.CurTileID && frontTileID == l.LedgeTileID {
			s.Simulated = []uint{s.Direction}
			return true
		}
	}

	return false
}
