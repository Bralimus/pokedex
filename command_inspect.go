package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, caught := cfg.pokedex[name]
	if !caught {
		fmt.Printf("You haven't caught %s yet\n", name)
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for i := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon.Types[i].Type.Name)
	}

	return nil

}
