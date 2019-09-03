package main

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// ScriptSystem use Lua for scripting.
type ScriptSystem struct {
	state    *lua.LState
	commands *Commands
}

// NewScriptSystem Initializes the Lua Script System
func NewScriptSystem() *ScriptSystem {
	var scriptSystem ScriptSystem

	scriptSystem.state = lua.NewState()
	scriptSystem.commands = NewCommands()
	scriptSystem.state.SetGlobal("Commands", luar.New(scriptSystem.state, scriptSystem.commands))

	return &scriptSystem
}

// DestroyScriptSystem Calls lua.LState.Close
func (s *ScriptSystem) DestroyScriptSystem() {
	s.state.Close()
}

// DoString Run the passed code string
func (s *ScriptSystem) DoString(code string) {
	s.state.DoString(code)
}

// DoFile Load the file and run its code
func (s *ScriptSystem) DoFile(fileName string) {
	s.state.DoFile(fileName)
}

// LoadString load the passed code string
func (s *ScriptSystem) LoadString(code string) (*lua.LFunction, error) {
	return s.state.LoadString(code)
}

// LoadFile Load the file
func (s *ScriptSystem) LoadFile(fileName string) (*lua.LFunction, error) {
	return s.state.LoadFile(fileName)
}
