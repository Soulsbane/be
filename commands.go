package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Command holds the command name and the command
type Command struct {
	command string
	name    string
}

// Commands this is it
type Commands struct {
	commands []Command
}

func (c *Commands) addCommand(name string, commandString string) {
	command := Command{commandString, name}
	c.commands = append(c.commands, command)
}

func (c *Commands) run(name string) {
	for _, v := range c.commands {
		if v.name == name {
			fmt.Println("Found the command:", v.command)

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
			fmt.Println("Failed to find command!")
		}
	}
}

func (c *Commands) dump() {
	fmt.Println("Dumping")
	for _, v := range c.commands {
		fmt.Println(v.command, "=>", v.name)
	}
}
