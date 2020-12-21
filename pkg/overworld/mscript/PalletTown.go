package mscript

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap/song"
	"pokered/pkg/event"
	"pokered/pkg/joypad"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
)

func init() {
	txt.RegisterAsmText("OakAppears", func() string {
		p := store.SpriteData[0]
		p.Direction = util.Down
		return ""
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
	if p.MapXCoord != 10 && p.MapXCoord != 11 {
		return
	}

	audio.PlayMusic(song.MUSIC_MEET_PROF_OAK)

	event.UpdateEvent(event.EVENT_OAK_APPEARED_IN_PALLET, true)
	store.CurMapScript = 1
	joypad.JoyIgnore = joypad.ByteToInput(0xfc)
}

func palletTownScript1() {
	text.DoPrintTextScript(text.TextBoxImage, txt.OakAppearsText, false)
	showObject(1)
	store.CurMapScript = 2
}

func palletTownScript2() {
	p := store.SpriteData[0]
	oak := store.SpriteData[1]
	if !delayed3F {
		oak.Direction = util.Up
		store.DelayFrames = 3
		delayed3F = true
		return
	}

	switch p.MapXCoord {
	case 10:
		oak.Simulated = []uint{util.Right, util.Up, util.Right, util.Up, util.Up}
	case 11:
		oak.Simulated = []uint{util.Right, util.Up, util.Right, util.Up, util.Right, util.Up}
	}

	store.CurMapScript = 3
}

func palletTownScript3() {
	oak := store.SpriteData[1]
	if len(oak.Simulated) > 0 || oak.MovmentStatus == sprite.Movement {
		joypad.JoyIgnore = joypad.ByteToInput(0xff)
		return
	}
	joypad.JoyIgnore = joypad.ByteToInput(0xfc)
	text.DoPrintTextScript(text.TextBoxImage, txt.OakWalksUpText, false)
	store.CurMapScript = 4

}
func palletTownScript4() {
	p, oak := store.SpriteData[0], store.SpriteData[1]
	oak.DoubleSpd = true
	simulated := []uint{util.Down, util.Down, util.Down, util.Down, util.Down, util.Left, util.Down, util.Down, util.Down, util.Down, util.Down, util.Right, util.Right, util.Right, util.Up}
	switch p.MapXCoord {
	case 10:
		copied := make([]uint, len(simulated))
		copy(copied, simulated)
		p.Simulated = append([]uint{util.Down}, copied...)
		copied = make([]uint, len(simulated))
		copy(copied, simulated)
		oak.Simulated = copied
	case 11:
		copied := make([]uint, len(simulated))
		copy(copied, simulated)
		p.Simulated = append([]uint{util.Stop, util.Left, util.Down}, copied...)
		copied = make([]uint, len(simulated))
		copy(copied, simulated)
		oak.Simulated = append([]uint{util.Left, util.Stop}, copied...)
	}

	store.CurMapScript = 5
}
func palletTownScript5() {
	p, oak := store.SpriteData[0], store.SpriteData[1]
	if len(p.Simulated) > 0 || p.MovmentStatus == sprite.Movement {
		return
	}
	if len(oak.Simulated) > 0 || oak.MovmentStatus == sprite.Movement {
		return
	}
}
func palletTownScript6() {}
