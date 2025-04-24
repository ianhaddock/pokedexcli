package main

import (
    "time"
    "github.com/ianhaddock/pokedexcli/internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(time.Second * 5)

    cfg := &config{
        pokeapiClient: pokeClient,
    }

    startRepl(cfg)
}
