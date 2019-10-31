package main

import (
	"fmt"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

const companyName = "Raijinsoft"
const applicationName = "launch"

func main() {
	scriptSystem := NewScriptSystem()
	defer scriptSystem.DestroyScriptSystem()

	if len(os.Args) >= 2 {
		var commands *Commands

		commands = NewCommands()

		scriptSystem.SetGlobal("Args", createArgsTable(scriptSystem))
		scriptSystem.SetGlobal("Commands", commands)
		scriptSystem.DoFiles(setupCommandFilesDir(), true)

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

func createArgsTable(script *ScriptSystem) *lua.LTable {
	tbl := script.NewTable()

	for _, arg := range os.Args[2:] {
		tbl.Append(lua.LString(arg))
	}

	return tbl
}

func setupCommandFilesDir() string {
	// NOTE: This function is only available in Go 1.13
	configPath, _ := os.UserConfigDir()
	commandFilesDir := filepath.Join(configPath, companyName, applicationName, "commands")

	err := os.MkdirAll(commandFilesDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	return commandFilesDir
}
