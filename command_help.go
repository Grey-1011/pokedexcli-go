package main

import "fmt"
/* 定义函数时, 变成参数(variadic parameter)的语法是 
* `...type`, 这表示盖参数可以接受任意数量的指定类型的值.
*/
func callbackHelp(cfg *config, args ...string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {	
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}