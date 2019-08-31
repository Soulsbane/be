package main

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// ScriptSystem use Lua for scripting.
type ScriptSystem struct {
	state    *lua.LState
	commands *Commands
	blah     string
}

// NewScriptSystem Initializes the Lua Script System
func NewScriptSystem() *ScriptSystem {
	var scriptSystem ScriptSystem
	scriptSystem.blah = "blah"
	scriptSystem.state = lua.NewState()
	scriptSystem.commands = NewCommands()
	//defer state.Close()
	scriptSystem.state.SetGlobal("Commands", luar.New(scriptSystem.state, scriptSystem.commands))

	return &scriptSystem
}
