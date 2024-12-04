package pokeapi

import (
	"net/http"
	pokeapicache "pokedex/pokeapi/pokeapiCache"
	"time"
)

type Client struct {
	cache      pokeapicache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokeapicache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}
