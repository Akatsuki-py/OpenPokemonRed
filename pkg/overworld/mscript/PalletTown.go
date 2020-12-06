package mscript

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/txt"
	"pokered/pkg/event"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/world"
)

var delayed3F bool

func init() {
	txt.RegisterAsmText("OakAppears", func() {
		p := store.SpriteData[0]
		p.Direction = util.Down
	})
}

func palletTownScript() {
	switch store.CurMapScript {
	case 0:
		palletTownScript0()
	case 1:
		palletTownScript1()
	case 2:
		palletTownScript2()
	case 3:
		palletTownScript3()
	case 4:
		palletTownScript4()
	case 5:
		palletTownScript5()
	case 6:
		palletTownScript6()
	}
}

func palletTownScript0() {
	// イベント消化済み
	if event.CheckEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB) {
		return
	}

	// プレイヤーがマサラタウンの南からマップ外に出ようとしている時は return
	p := store.SpriteData[0]
	if p == nil || p.MapYCoord != 1 {
		return
	}

	audio.PlayMusic(audio.MUSIC_MEET_PROF_OAK)

	event.UpdateEvent(event.EVENT_OAK_APPEARED_IN_PALLET, true)
	store.CurMapScript = 1
}

func palletTownScript1() {
	text.DoPrintTextScript(text.TextBoxImage, txt.OakAppearsText, false)
	world.CurWorld.Object.HS[0x01] = false
	store.SpriteData[1].Hidden = false
	store.CurMapScript = 2
}

func palletTownScript2() {
	if !delayed3F {
		oak := store.SpriteData[1]
		oak.Direction = util.Up
		store.DelayFrames = 3
		delayed3F = true
		return
	}
	store.CurMapScript = 3
}

func palletTownScript3() {}
func palletTownScript4() {}
func palletTownScript5() {}
func palletTownScript6() {}
