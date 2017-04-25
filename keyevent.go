package main

import termbox "github.com/nsf/termbox-go"

type Action uint8
type Keybind struct {
	Mod    termbox.Modifier
	Key    termbox.Key
	Ch     rune
	Action Action
}

// Actions
const (
	ACTION_RELOAD = iota + 1
	ACTION_QUIT
	ACTION_UP
	ACTION_DOWN
	ENTER_NORMAL_MODE
	ENTER_COMMAND_MODE
)

const NO_MOD = 0
const NO_KEY = 0
const NO_CH = 0
const NO_ACTION = 0

var KeybindList = []Keybind{
	{NO_MOD, termbox.KeyCtrlR, NO_CH, ACTION_RELOAD},
	{NO_MOD, termbox.KeyCtrlQ, NO_CH, ACTION_QUIT},
	{NO_MOD, termbox.KeyArrowUp, NO_CH, ACTION_UP},
	{NO_MOD, termbox.KeyArrowDown, NO_CH, ACTION_DOWN},
	{NO_MOD, termbox.KeyCtrlC, NO_CH, ENTER_NORMAL_MODE},
	{NO_MOD, NO_KEY, ':', ENTER_COMMAND_MODE},
}