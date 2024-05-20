package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 指针允许我们直接操作内存中的数据，而无需复制或重复数据
func startRepl(cfg *config) {
	// 创建一个 Scanner 以从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
	
		scanner.Scan() 
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
			// 停止当前循环跌代, 并继续到下一次迭代
		}
		
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		// availableCommands: map[string]cliCommand
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		// command 是 map 的 value
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "Lists the next page of location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous page of location areas",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore <area_name>",
			description: "Lists the pokemon in a location area",
			callback: callbackExplore,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	// strings.Fields()函数会将字符串按照空白字符进行分割，并返回一个包含分割后单词的切片
	words := strings.Fields(lowered)
	return words
}