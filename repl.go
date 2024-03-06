package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cnf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcom to the pokdex")
	for {
		fmt.Print("pokdex >")
		scanner.Scan()

		scanned := scanner.Text()
		cleaned := cleanInput(scanned)
		if len(cleaned[0]) == 0 {
			continue
		}

		allCommands := allCommand()

		if command, ok := allCommands[cleaned[0]]; ok {
			err := command.callback(cnf)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("invalid command")
			continue
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
	callback    func(cnf *config) error
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
		"map": {
			name:        "map",
			discription: "displays the names of next 20 location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			discription: "displays the names of prev 20 location",
			callback:    commandMapb,
		},
	}
}
