package main

import "fmt"

func callbackHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {	
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}