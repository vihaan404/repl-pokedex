package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcom to the pokdex")
	for {
		fmt.Print("pokdex >")
		scanner.Scan()

		command := scanner.Text()
		cleaned := cleanInput(command)
		if len(cleaned[0]) == 0 {
			continue
		}

		allCommands := allCommand()

		if command, ok := allCommands[cleaned[0]]; ok {
			command.callback()
		} else {
			fmt.Println("invalid command")
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	discription string
	callback    func() error
}

func allCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			discription: "gives help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			discription: "exit the clipokedex",
			callback:    commandExit,
		},
	}
}
