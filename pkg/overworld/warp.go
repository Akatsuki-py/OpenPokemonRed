package overworld

import (
	"pokered/pkg/data/tileset"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/data/worldmap/header"
	"pokered/pkg/joypad"
	"pokered/pkg/palette"
	"pokered/pkg/script"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

// check if the player has stepped onto a warp after having not collided
// ref: CheckWarpsNoCollision
func checkWarpsNoCollision() {
	curWorld := world.CurWorld
	if len(curWorld.Object.Warps) == 0 {
		checkMapConnections()
		return
	}

	p := store.SpriteData[0]
	if p == nil {
		return
	}
	for _, w := range curWorld.Object.Warps {
		if p.MapXCoord == w.XCoord && p.MapYCoord == w.YCoord {
			util.SetBit(&store.D736, 2)
			if store.Enable.NormalWarp && sprite.IsStandingOnDoorOrWarp(0) {
				store.Enable.NormalWarp = false
				warpFound(w.DestMap, w.DestWarpID)
				return
			}

			if !extraWarpCheck() {
				return
			}

			// if the extra check passed
			joypad.Joypad()
			if joypad.JoyHeld.Down || joypad.JoyHeld.Up || joypad.JoyHeld.Left || joypad.JoyHeld.Right {
				p.WalkCounter, p.AnimationFrame = 0, 0
				warpFound(w.DestMap, w.DestWarpID)
			}

		}
	}

	checkMapConnections()
}

// ref: ExtraWarpCheck
func extraWarpCheck() bool {
	result := false
	curMap, curTileset := world.CurWorld.MapID, world.CurWorld.Header.Tileset

	switch curMap {
	case worldmap.ROCKET_HIDEOUT_B1F, worldmap.ROCKET_HIDEOUT_B2F, worldmap.ROCKET_HIDEOUT_B4F, worldmap.ROCK_TUNNEL_1F:
		result = sprite.IsWarpTileInFrontOfPlayer()

	default:
		switch curTileset {
		case tileset.Overworld, tileset.Ship, tileset.ShipPort, tileset.Plateau:
			result = sprite.IsWarpTileInFrontOfPlayer()
		case tileset.RedsHouse:
			result = sprite.IsPlayerFacingEdgeOfMap()
			_, curTileID := sprite.PlayerCurTileID()
			result = result && (curTileID == 0x14)
		default:
			result = sprite.IsPlayerFacingEdgeOfMap()
		}
	}
	return result
}

// ref: CheckMapConnections
func checkMapConnections() {
	curWorld := world.CurWorld
	p := store.SpriteData[0]
	if p == nil {
		return
	}

	if p.Direction == util.Up && p.MapYCoord == -1 {
		for i, XCoord := range curWorld.Header.Connections.North.Coords {
			if p.MapXCoord == int(XCoord) {
				destMapID := curWorld.Header.Connections.North.DestMapID
				DestMapHeader := header.Get(destMapID)
				loadWorldData(destMapID, -1)
				p.MapXCoord = int(DestMapHeader.Connections.South.Coords[i])
				p.MapYCoord = int(DestMapHeader.Height*2 - 1)
				return
			}
		}
	}

	if p.Direction == util.Down && p.MapYCoord == int(curWorld.Header.Height*2) {
		for i, XCoord := range curWorld.Header.Connections.South.Coords {
			if p.MapXCoord == int(XCoord) {
				destMapID := curWorld.Header.Connections.South.DestMapID
				DestMapHeader := header.Get(destMapID)
				loadWorldData(destMapID, -1)
				p.MapXCoord = int(DestMapHeader.Connections.North.Coords[i])
				p.MapYCoord = 0
				return
			}
		}
	}
}

func warpFound(mapID, warpID int) {
	if world.CheckIfInOutsideMap() {
		world.LastWorld = world.CurWorld
		if mapID != worldmap.ROCK_TUNNEL_1F {
		}
	} else {
		// indoorMaps
		if mapID == worldmap.LAST_MAP {
			mapID = world.LastWorld.MapID
		}
	}
	playMapChangeSound()
	palette.GBFadeOutToBlack()

	world.WarpTo = [2]int{mapID, warpID}
	script.PushID(script.LoadMapData)
}
