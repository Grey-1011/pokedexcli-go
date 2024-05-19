package pokeapi

import (
	"net/http"
	"time"

	"github.com/Grey-1011/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"


// 不使用 default Client, 因为没有设置超时

// no default Client
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}
// export Client struct
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}