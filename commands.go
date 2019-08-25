package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cheynewallace/tabby"
)

// Commands holds a list of commands.
type Commands struct {
	commandsMap map[string]string
}

// MakeCommandsMap Initializes the command map
func (c *Commands) makeCommandsMap() {
	/*
		Use like so:
		commands := &Commands{}
		commands.MakeCommandsMap()
	*/
	c.commandsMap = make(map[string]string)
}

// NewCommands Initializes the command map
func NewCommands() *Commands {
	var commands Commands
	commands.commandsMap = make(map[string]string)
	return &commands
}

// AddCommand adds a command to the list of commands.
func (c *Commands) AddCommand(name string, commandString string) {
	if c.HasCommand(name) {
		fmt.Println(name, "command already exists!")
	} else {
		c.commandsMap[name] = commandString
	}
}

// HasCommand returns whether a map key exists.
func (c *Commands) HasCommand(name string) bool {
	if _, ok := c.commandsMap[name]; ok {
		return ok
	}

	return false
}

func (c *Commands) run(name string) {
	if c.HasCommand(name) {
		args := strings.Split(c.commandsMap[name], " ")
		argsLength := len(args)

		if argsLength > 0 {
			if argsLength == 1 {
				execCommand := exec.Command(args[0])
				execCommand.Start()
			} else {
				execCommand := exec.Command(args[0], args[1:]...)
				execCommand.Start()
			}
		}
	} else {
		fmt.Println("Command not found: ", name)
	}

}

func (c *Commands) list() {
	t := tabby.New()
	t.AddHeader("Name", "Command")

	for key, value := range c.commandsMap {
		t.AddLine(key, value)
	}

	t.Print()
}
