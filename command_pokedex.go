package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}