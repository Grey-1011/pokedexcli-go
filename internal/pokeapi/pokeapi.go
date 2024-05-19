package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"


// 不使用 default Client, 因为没有设置超时

// no default Client
type Client struct {
	httpClient http.Client
}
// export Client struct
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}