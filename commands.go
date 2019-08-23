package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cheynewallace/tabby"
)

// Command holds the command name and the command
type Command struct {
	command string
	name    string
}

// Commands holds a list of commands.
type Commands struct {
	commands []Command
}

// AddCommand Add a command to the list of commands.
func (c *Commands) AddCommand(name string, commandString string) {
	command := Command{commandString, name}
	c.commands = append(c.commands, command)
	fmt.Println("Added command", name)
}

func (c *Commands) run(name string) {
	for _, v := range c.commands {
		if v.name == name {
			args := strings.Split(v.command, " ")
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
			fmt.Println("Failed to find command", name)
		}
	}
}

func (c *Commands) list() {
	t := tabby.New()
	t.AddHeader("Name", "Command")

	for _, v := range c.commands {
		t.AddLine(v.name, v.command)
	}

	t.Print()
}
