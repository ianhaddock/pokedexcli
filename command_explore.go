package main

import (
    "fmt"
    "errors"
)

func commandExplore(cfg *config, args ...string) error {

    if len(args) != 1 {
        return errors.New("need a location to explore")
    }
        
    name := args[0]
    exploreResp, err := cfg.pokeapiClient.ExploreLocation(name)
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s... \n", name)
    fmt.Println("Found Pokemon:")
    for _, pokemon := range exploreResp.PokemonEncounters {
        fmt.Printf(" - %s \n", pokemon.Pokemon.Name)
    }
    return nil
}

