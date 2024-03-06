package main

import "os"

func commandExit(cnf *config) error {
	os.Exit(0)
	return nil
}
