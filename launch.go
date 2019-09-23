package main

import (
	"fmt"
	"os"
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

	scriptSystem.SetGlobal("Commands", commands)
	scriptSystem.DoFiles(commandFilesPath)

	if len(os.Args) >= 2 {
		commandName := os.Args[1]

		if commandName == "list" {
			commands.list()
		} else {
			fmt.Println("Running command:", commandName)
			commands.run(commandName)
		}
	} else {
		fmt.Println("Invalid option passed!")
	}
}
