package main

import "fmt"

func main() {
	const script = `print("Adding command from Lua") Commands:addCommand("subl", "subl")`

	scriptSystem := NewScriptSystem()

	scriptSystem.commands.AddCommand("wow", "lsd")
	scriptSystem.commands.list()
	fmt.Println("")

	scriptSystem.DoString(script)
	scriptSystem.commands.run("wow")
	scriptSystem.commands.run("subl")
	fmt.Println("")
	scriptSystem.commands.list()
}
