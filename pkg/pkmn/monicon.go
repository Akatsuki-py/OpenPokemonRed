package pkmn

import (
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/store"
	"pokered/pkg/util"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	suffix    = "_mon_icon"
	ball      = "ball"
	bird      = "bird"
	bug       = "bug"
	fairy     = "fairy"
	grass     = "grass"
	helix     = "helix"
	mon       = "mon"
	quadruped = "quadruped"
	snake     = "snake"
	water     = "water"
)

var IconGen1 = map[uint][2]*ebiten.Image{
	pkmnd.BallMon: {
		util.OpenImage(store.FS, "/"+ball+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+ball+suffix+"_1.png"),
	},
	pkmnd.BirdMon: {
		util.OpenImage(store.FS, "/"+bird+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+bird+suffix+"_1.png"),
	},
	pkmnd.BugMon: {
		util.OpenImage(store.FS, "/"+bug+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+bug+suffix+"_1.png"),
	},
	pkmnd.FairyMon: {
		util.OpenImage(store.FS, "/"+fairy+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+fairy+suffix+"_1.png"),
	},
	pkmnd.GrassMon: {
		util.OpenImage(store.FS, "/"+grass+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+grass+suffix+"_1.png"),
	},
	pkmnd.HelixMon: {
		util.OpenImage(store.FS, "/"+helix+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+helix+suffix+"_1.png"),
	},
	pkmnd.MonMon: {
		util.OpenImage(store.FS, "/"+mon+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+mon+suffix+"_1.png"),
	},
	pkmnd.QuadrupedMon: {
		util.OpenImage(store.FS, "/"+quadruped+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+quadruped+suffix+"_1.png"),
	},
	pkmnd.SnakeMon: {
		util.OpenImage(store.FS, "/"+snake+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+snake+suffix+"_1.png"),
	},
	pkmnd.WaterMon: {
		util.OpenImage(store.FS, "/"+water+suffix+"_0.png"),
		util.OpenImage(store.FS, "/"+water+suffix+"_1.png"),
	},
}
