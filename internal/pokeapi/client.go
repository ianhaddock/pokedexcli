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
    pc := pokecache.NewCache()

    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        pokeCache: pc,
    }
}

