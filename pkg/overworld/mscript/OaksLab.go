package mscript

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/event"
	"pokered/pkg/joypad"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
)

func init() {
	txt.RegisterAsmText("OaksLabText1", func() string {
		if event.CheckEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB_2) {
			return txt.OaksLabText40
		} else {
			return txt.OaksLabGaryText1
		}
	})
}

func oaksLabScript() {
	switch store.CurMapScript {
	case 0:
		oaksLabScript0()
	case 1:
		oaksLabScript1()
	case 2:
		oaksLabScript2()
	case 3:
		oaksLabScript3()
	case 4:
		oaksLabScript4()
	case 5:
		oaksLabScript5()
	}
}

func oaksLabScript0() {
	if !event.CheckEvent(event.EVENT_OAK_APPEARED_IN_PALLET) {
		return
	}

	showObject(8)
	store.CurMapScript = 1
}

func oaksLabScript1() {
	oak := store.SpriteData[8]
	oak.DoubleSpd = false
	oak.Simulated = []uint{util.Up, util.Up, util.Up}
	store.CurMapScript = 2
}

func oaksLabScript2() {
	oak := store.SpriteData[8]
	if len(oak.Simulated) > 0 || oak.MovmentStatus == sprite.Movement {
		return
	}

	hideObject(8) // walking oak
	showObject(5) // staying oak
	store.CurMapScript = 3
	delayed3F = false
}

func oaksLabScript3() {
	if !delayed3F {
		store.DelayFrames = 3
		delayed3F = true
		return
	}

	p := store.SpriteData[0]
	p.Simulated = make([]uint, 8)
	for i := 0; i < 8; i++ {
		p.Simulated[i] = util.Up
	}

	store.CurMapScript = 4
}

func oaksLabScript4() {
	p := store.SpriteData[0]
	if len(p.Simulated) > 0 || p.MovmentStatus == sprite.Movement {
		return
	}

	event.UpdateEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB, true)
	event.UpdateEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB_2, true)

	blue := store.SpriteData[1]
	blue.Direction = util.Up

	audio.PlayDefaultMusic(worldmap.OAKS_LAB)

	store.CurMapScript = 5
}

var delay30FCtrInoaksLabScript5 = 0

func oaksLabScript5() {
	switch delay30FCtrInoaksLabScript5 {
	case 0:
		joypad.JoyIgnore = joypad.ByteToInput(0xfc)

		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText17, false)
		store.DelayFrames = 30
		delay30FCtrInoaksLabScript5++
		return
	case 1:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText18, false)
		store.DelayFrames = 30
		delay30FCtrInoaksLabScript5++
	case 2:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText19, false)
		store.DelayFrames = 30
		delay30FCtrInoaksLabScript5++
	case 3:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText20, false)
		store.DelayFrames = 30
		delay30FCtrInoaksLabScript5++
		event.UpdateEvent(event.EVENT_OAK_ASKED_TO_CHOOSE_MON, true)
		joypad.JoyIgnore = joypad.ByteToInput(0x00)
		store.CurMapScript = 6
	}
}
