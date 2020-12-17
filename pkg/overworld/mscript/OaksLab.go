package mscript

import (
	"pokered/pkg/event"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

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
	store.CurMapScript = 5
}
