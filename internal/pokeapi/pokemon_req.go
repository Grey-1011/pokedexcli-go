package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		pokemon := Pokemon{}
		// 用于将 JSON 数据解码到 Go 语言中的结构体或其他数据类型中。
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// .Do() 用于发送 HTTP 请求并返回 HTTP 响应
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	
	pokemon := Pokemon{}
	// 用于将 JSON 数据解码到 Go 语言中的结构体或其他数据类型中。
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// add cache
	c.cache.Add(fullURL, dat)

	return pokemon, nil
}


