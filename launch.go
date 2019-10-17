package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const companyName = "Raijinsoft"
const apllicationName = "launch"

func main() {
	scriptSystem := NewScriptSystem()
	defer scriptSystem.DestroyScriptSystem()

	if len(os.Args) >= 2 {
		var commands *Commands

		commands = NewCommands()
		commands.AddCommand("wow", "lsd")
		commands.AddOutputCommand("lsd", "lsd -lt")

		scriptSystem.SetGlobal("Commands", commands)
		scriptSystem.DoFiles(setupCommandFilesDir())

		commandName := os.Args[1]

		if commandName == "list" {
			commands.list()
		} else {
			fmt.Println("Running command:", commandName)
			commands.run(commandName)
		}
	} else {
		fmt.Println("No command passed! Use 'launch list' for a list of commands.")
	}
}

func setupCommandFilesDir() string {
	// NOTE: This function is only available in Go 1.13
	configPath, _ := os.UserConfigDir()
	commandFilesDir := filepath.Join(configPath, companyName, apllicationName, "commands")

	err := os.MkdirAll(commandFilesDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	return commandFilesDir
}
