package main

import (
	"errors"
	"fmt"
	//"github.com/Grey-1011/pokedexcli/internal/pokeapi"
)

func callbackMap(cfg *config, args ...string) error {
	// pokeapiClient := pokeapi.NewClient()  <=>  等价于下一行
	pokeapiClient := cfg.pokeapiClient

	resp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		//log.Fatal(err)
		// log.Fatal(err) 会打印错误消息并调用 os.Exit(1)，立即终止程序。
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n",area.Name)
	}
	return nil
}



func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	fmt.Println("Location areas:")
	
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}