package main

import (
	"time"

	"github.com/Grey-1011/pokedexcli/internal/pokeapi"
)

//"fmt"
//"log"

/*
Go programs are organized into packages.
A package is a directory of Go code
that's all compiled together.
 Functions, types, variables,
 and constants defined in one source file
are visible to all other source files
within the same package (directory).
Go 程序被组织成包。
包是一个包含了一起编译的 Go 代码的目录。
在同一个包（目录）内定义的函数、类型、变量和常量
对于该包内的所有其他源文件都是可见的。
*/

type config struct {
	pokeapiClient 			pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string // 这里用 指针 *string, 只是为了可以为空
	caughtPokemon 			map[string]pokeapi.Pokemon
	// 在 config 中添加存储 捕获到的 pokemon
	// 因为这是整个应用程序的共享状态
	/*
	潜在风险, map 不是线程安全的
	但这里并不会发生并发(同时对map读写), 所以暂时可以这样写
	如果要解决就要使用 mutex
	*/
}


func main() {
	// pokeapiClient := pokeapi.NewClient()

	// resp, err := pokeapiClient.ListLocationAreas()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)

	cfg := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}