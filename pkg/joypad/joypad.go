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

// Any button pressed?
func (i *Input) Any() bool {
	return i.Up || i.Down || i.Left || i.Right || i.A || i.B || i.Start || i.Select
}

func ByteToInput(b byte) Input {
	input := Input{}
	// [↓, ↑, ←, →, Start, Select, B, A]
	if util.ReadBit(b, 0) {
		input.A = true
	}
	if util.ReadBit(b, 1) {
		input.B = true
	}
	if util.ReadBit(b, 2) {
		input.Select = true
	}
	if util.ReadBit(b, 3) {
		input.Start = true
	}
	if util.ReadBit(b, 4) {
		input.Right = true
	}
	if util.ReadBit(b, 5) {
		input.Left = true
	}
	if util.ReadBit(b, 6) {
		input.Up = true
	}
	if util.ReadBit(b, 7) {
		input.Down = true
	}
	return input
}

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

	if JoyPressed.Any() {
		store.FrameCounter = 8
		return
	}

	if store.FrameCounter > 0 {
		Joy5 = Input{}
		return
	}

	if JoyHeld.A || JoyHeld.B {
		Joy5 = Input{}
	}

	if !Joy6 {
		Joy5 = Input{}
	}

	store.FrameCounter = 1
	return
}

// ABButtonPress return if AB button is pressed
func ABButtonPress() bool {
	JoypadLowSensitivity()
	pressed := Joy5.A || Joy5.B
	return pressed
}
