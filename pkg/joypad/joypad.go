package joypad

import (
	"pokered/pkg/store"
	"pokered/pkg/util"
)

type Input struct {
	Up, Down, Left, Right bool
	A, B, Start, Select   bool
}

var Joy5 Input

var Joy6 bool
var Joy7 bool

var JoyInput Input

var JoyLast Input

// JoyReleased 今回の_Joypad処理でONからOFFに変わったボタン
var JoyReleased Input

// JoyPressed 今回の_Joypad処理でOFFからONに変わったボタン
var JoyPressed Input

// JoyHeld 現在押されているボタン
var JoyHeld Input

// JoyIgnore true のものはキーが無視される
var JoyIgnore Input

// ReadJoypad read joypad input
func ReadJoypad() {
	JoyInput = Input{
		Up:     keyUp(),
		Down:   keyDown(),
		Left:   keyLeft(),
		Right:  keyRight(),
		A:      a(),
		B:      b(),
		Start:  start(),
		Select: sel(),
	}
}

// Joypad process joypad input
func Joypad() {
	if JoyInput.A && JoyInput.B && JoyInput.Start && JoyInput.Select {
		// trySoftReset
	}

	JoyReleased = Input{
		Up:     util.XOR(JoyLast.Up, JoyInput.Up) && JoyLast.Up,
		Down:   util.XOR(JoyLast.Down, JoyInput.Down) && JoyLast.Down,
		Left:   util.XOR(JoyLast.Left, JoyInput.Left) && JoyLast.Left,
		Right:  util.XOR(JoyLast.Right, JoyInput.Right) && JoyLast.Right,
		A:      util.XOR(JoyLast.A, JoyInput.A) && JoyLast.A,
		B:      util.XOR(JoyLast.B, JoyInput.B) && JoyLast.B,
		Start:  util.XOR(JoyLast.Start, JoyInput.Start) && JoyLast.Start,
		Select: util.XOR(JoyLast.Select, JoyInput.Select) && JoyLast.Select,
	}

	JoyPressed = Input{
		Up:     util.XOR(JoyLast.Up, JoyInput.Up) && JoyInput.Up,
		Down:   util.XOR(JoyLast.Down, JoyInput.Down) && JoyInput.Down,
		Left:   util.XOR(JoyLast.Left, JoyInput.Left) && JoyInput.Left,
		Right:  util.XOR(JoyLast.Right, JoyInput.Right) && JoyInput.Right,
		A:      util.XOR(JoyLast.A, JoyInput.A) && JoyInput.A,
		B:      util.XOR(JoyLast.B, JoyInput.B) && JoyInput.B,
		Start:  util.XOR(JoyLast.Start, JoyInput.Start) && JoyInput.Start,
		Select: util.XOR(JoyLast.Select, JoyInput.Select) && JoyInput.Select,
	}

	JoyLast = JoyInput

	if util.ReadBit(store.D730, 5) {
		discardButtonPresses()
	}

	JoyHeld = JoyLast

	if JoyIgnore.Up {
		JoyHeld.Up, JoyPressed.Up = false, false
	}
	if JoyIgnore.Down {
		JoyHeld.Down, JoyPressed.Down = false, false
	}
	if JoyIgnore.Left {
		JoyHeld.Left, JoyPressed.Left = false, false
	}
	if JoyIgnore.Right {
		JoyHeld.Right, JoyPressed.Right = false, false
	}
	if JoyIgnore.A {
		JoyHeld.A, JoyPressed.A = false, false
	}
	if JoyIgnore.B {
		JoyHeld.B, JoyPressed.B = false, false
	}
	if JoyIgnore.Start {
		JoyHeld.Start, JoyPressed.Start = false, false
	}
	if JoyIgnore.Select {
		JoyHeld.Select, JoyPressed.Select = false, false
	}
}

func discardButtonPresses() {
	JoyReleased, JoyPressed, JoyHeld = Input{}, Input{}, Input{}
}

func JoypadLowSensitivity() {
	Joypad()

	Joy5 = JoyPressed
	if Joy7 {
		Joy5 = JoyHeld
	}

	if JoyPressed.Up || JoyPressed.Down || JoyPressed.Left || JoyPressed.Right || JoyPressed.A || JoyPressed.B || JoyPressed.Start || JoyPressed.Select {
		store.FrameCounter = 30
		return
	}

	if store.FrameCounter > 0 {
		Joy5 = Input{}
		return
	}

	if !((!JoyHeld.A && !JoyHeld.B) || Joy6) {
		Joy5 = Input{}
	}

	store.FrameCounter = 5
	return
}
