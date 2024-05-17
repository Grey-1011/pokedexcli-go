package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// 创建一个 Scanner 以从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
	
		scanner.Scan() 
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
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