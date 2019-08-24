package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func main() {
	//var commands Commands
	L := lua.NewState()
	defer L.Close()

	const script = `print("Adding command from Lua") Commands:addCommand("subl", "subl")`

	commands := NewCommands()
	L.SetGlobal("Commands", luar.New(L, commands))

	commands.AddCommand("wow", "lsd")
	commands.list()
	fmt.Println("")

	if err := L.DoString(script); err != nil {
		panic(err)
	}

	commands.run("wow")
	commands.run("subl")
	fmt.Println("")
	commands.list()
}
