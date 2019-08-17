package main

import "fmt"

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

func (c *Commands) dump() {
	fmt.Println("Dumping")
	for _, v := range c.commands {
		fmt.Println(v.command, "=>", v.name)
	}
}
