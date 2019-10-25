package main

import (
	funk "github.com/thoas/go-funk"
)

var reservedCommandNames []string = []string{
	"list",
}

// IsReservedCommandName Checks if a command is a reserved name
func IsReservedCommandName(name string) bool {
	return funk.Contains(reservedCommandNames, name)
}
