package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error { // Have c *config only to match function signature not to use it
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
