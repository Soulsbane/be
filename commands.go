package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/brettski/go-termtables"
	cfmt "github.com/mingrammer/cfmt"
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
	passedArgs := os.Args[2:]

	if len(passedArgs) > 1 {
		passedArgs = passedArgs[1:]
	}

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
	c.AddCommandIfAvailable(name, commandString, false)
}

// AddOutputCommand adds a command to the list of commands and outputs its result.
func (c *Commands) AddOutputCommand(name string, commandString string) {
	c.AddCommandIfAvailable(name, commandString, true)
}

// AddCommandIfAvailable Returns true if a command name is available otherwise false
func (c *Commands) AddCommandIfAvailable(name string, commandString string, showOutput bool) {
	if IsReservedCommandName(name) {
		cfmt.Errorf("'%s' is a reserved command name!\n", name)
	} else if c.HasCommand(name) {
		cfmt.Errorf("'%s' command already exists!\n", name)
	} else {
		var newCommand = command{name, commandString, showOutput}
		c.commandsArray = append(c.commandsArray, newCommand)
	}
}

// RemoveCommand Removes a command.
func (c *Commands) RemoveCommand(name string) {
	if IsReservedCommandName(name) {
		cfmt.Errorf("'%s' is a reserved command name!\n", name)
	} else {
		if c.HasCommand(name) {
			for i, v := range c.commandsArray {
				if v.commandName == name {
					c.commandsArray = append(c.commandsArray[:i], c.commandsArray[i+1:]...)
				}
			}
		}
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
		cfmt.Errorf("'%s' command was not found!\n", name)
	}
}

func (c *Commands) list() {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Command")

	for i := range c.commandsArray {
		table.AddRow(c.commandsArray[i].commandName, c.commandsArray[i].commandString)
	}

	cfmt.Successln(table.Render())
}
