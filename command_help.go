package main

import (
	"fmt"
)

func commandHelp() error {
	allcommands := allCommand()
	for _, command := range allcommands {
		fmt.Println(command.name, "-", command.discription)
	}
	return nil
}
