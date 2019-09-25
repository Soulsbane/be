package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/brettski/go-termtables"
)

type command struct {
	commandName   string
	commandString string
	showOutput    bool
}

// Commands holds a list of commands.
type Commands struct {
	commandsArray []command
}

// NewCommands Initializes the command map
func NewCommands() *Commands {
	var newCommands Commands

	newCommands.commandsArray = []command{
		command{
			commandName:   "",
			commandString: "",
			showOutput:    false,
		},
	}

	return &newCommands
}

// AddCommand adds a command to the list of commands.
func (c *Commands) AddCommand(name string, commandString string) {
	if c.HasCommand(name) {
		fmt.Println(name, "command name already exists!")
	} else {
		var newCommand = command{name, commandString, false}
		c.commandsArray = append(c.commandsArray, newCommand)
	}
}

// AddOutputCommand adds a command to the list of commands.
func (c *Commands) AddOutputCommand(name string, commandString string) {
	if c.HasCommand(name) {
		fmt.Println(name, "command name already exists!")
	} else {
		var newCommand = command{name, commandString, true}
		c.commandsArray = append(c.commandsArray, newCommand)
	}
}

// HasCommand returns whether a map key exists.
func (c *Commands) HasCommand(name string) bool {
	for i := range c.commandsArray {
		if c.commandsArray[i].commandName == name {
			return true
		}
	}

	return false
}

func (c *Commands) getCommandIndex(name string) int {
	for i, val := range c.commandsArray {
		if val.commandName == name {
			return i
		}
	}

	return -1
}

func (c *Commands) handleSingleCommand(args []string, showOutput bool) {
	if showOutput {
		output, err := exec.Command(args[0]).Output()

		if err != nil {
			panic(err)
		}

		fmt.Println(string(output))
	} else {
		execCommand := exec.Command(args[0])
		execCommand.Start()
	}
}

func (c *Commands) handleMultiCommand(args []string, showOutput bool) {
	if showOutput {
		output, err := exec.Command(args[0], args[1:]...).Output()

		if err != nil {
			panic(err)
		}

		fmt.Println(string(output))
	} else {
		execCommand := exec.Command(args[0], args[1:]...)
		execCommand.Start()
	}
}

func (c *Commands) runCommandAtIndex(index int) {
	fmt.Println(c.commandsArray[index].commandName)
}

func (c *Commands) run(name string) {
	index := c.getCommandIndex(name)

	if index >= 0 {
		command := c.commandsArray[index]
		args := strings.Split(command.commandString, " ")
		argsLength := len(args)

		if argsLength > 0 {
			if argsLength == 1 {
				c.handleSingleCommand(args, command.showOutput)
			} else {
				c.handleMultiCommand(args, command.showOutput)
			}
		}
	} else {
		fmt.Println("Command not found: ", name)
	}
}

func (c *Commands) list() {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Command")

	for i := range c.commandsArray {
		table.AddRow(c.commandsArray[i].commandName, c.commandsArray[i].commandString)
	}

	fmt.Println(table.Render())
}
