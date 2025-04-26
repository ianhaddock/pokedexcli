package main

import (
    "time"
    "github.com/ianhaddock/pokedexcli/internal/pokeapi"
)

func main() {
    // NewClient(timeout inteval, cache interval)
    pokeClient := pokeapi.NewClient(time.Second * 5, time.Minute * 5)

    cfg := &config{
        pokeapiClient: pokeClient,
        characterExp: 100,
        characterPokedex: make(map[string]pokeapi.Pokemon),
    }

    startRepl(cfg)
}
