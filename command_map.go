package main

import (
	"errors"
	"fmt"
)

func commandMap(cnf *config) error {
	resp, err := cnf.pokeapiClient.ListLocationArea(cnf.nextLocationAreaURL)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println("-", location.Name)
	}
	cnf.nextLocationAreaURL = resp.Next
	cnf.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cnf *config) error {
	if cnf.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cnf.pokeapiClient.ListLocationArea(cnf.prevLocationAreaURL)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println("-", location.Name)
	}
	cnf.nextLocationAreaURL = resp.Next
	cnf.prevLocationAreaURL = resp.Previous
	return nil
}
