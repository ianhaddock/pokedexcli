package pokeapi

import (
    "net/http"
    "time"
    "github.com/ianhaddock/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
    httpClient http.Client
    pokeCache pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
    interval := time.Duration(5 * time.Second)
    pc := pokecache.NewCache(interval)

    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        pokeCache: pc,
    }
}

