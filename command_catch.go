package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/Bralimus/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	calculateCatch(cfg, pokemon)

	return nil
}

func calculateCatch(cfg *config, pokemon pokeapi.Pokemon) {
	randomNum := rand.Float64() // Generates number between 0.0 and 1.0
	catchChance := 1.0
	if pokemon.BaseExperience > 500 {
		catchChance = 0.10
	} else {
		catchChance = float64(500-pokemon.BaseExperience) / 500
	}

	if randomNum < catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now insepct it with the inspect command.")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}
