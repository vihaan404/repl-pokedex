package main

import "github.com/vihaan404/repl-pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cnf := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cnf)
}
