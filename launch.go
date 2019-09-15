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
	commands.AddOutputCommand("lsd", "lsd -lt")
	commands.list()

	scriptSystem.SetGlobal("Commands", commands)
	fmt.Println("")
	scriptSystem.DoString(script)
	commands.run("subl")
	commands.list()
	/*fmt.Println(commands.HasCommand("lsd"))
	fmt.Println(commands.getCommandIndex("subl"))
	fmt.Println(commands.getCommandIndex("lsddddd"))
	commands.runCommandAtIndex(commands.getCommandIndex("lsd"))*/
}
