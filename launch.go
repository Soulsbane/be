package main

import (
	"fmt"
	"path/filepath"

	"github.com/kirsle/configdir"
)

func main() {
	configPath := configdir.LocalConfig("Raijinsoft/launch")
	commandFilesPath := filepath.Join(configPath, "commands")
	err := configdir.MakePath(commandFilesPath)

	if err != nil {
		panic(err)
	}

	scriptSystem := NewScriptSystem()

	var commands *Commands

	commands = NewCommands()
	commands.AddCommand("wow", "lsd")
	commands.AddOutputCommand("lsd", "lsd -lt")
	commands.list()

	scriptSystem.SetGlobal("Commands", commands)
	fmt.Println("")
	scriptSystem.DoFiles(commandFilesPath)

	commands.run("subl")
	commands.list()
	/*fmt.Println(commands.HasCommand("lsd"))
	fmt.Println(commands.getCommandIndex("subl"))
	fmt.Println(commands.getCommandIndex("lsddddd"))
	commands.runCommandAtIndex(commands.getCommandIndex("lsd"))*/
}
