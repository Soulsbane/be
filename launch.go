package main

import "fmt"

func main() {
	const script = `print("Adding command from Lua")
	Commands:addCommand("subl", "subl")
	--Commands:run("subl")
	`

	scriptSystem := NewScriptSystem()

	var commands *Commands

	commands = NewCommands()
	commands.AddCommand("wow", "lsd")
	commands.list()

	scriptSystem.SetGlobal("Commands", commands)
	fmt.Println("")
	scriptSystem.DoString(script)
	commands.run("subl")
	commands.list()
}
