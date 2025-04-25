package main

import (
    "fmt"
    "errors"
)

func commandExplore(cfg *config, loc string) error {

    if loc == "" {
        return errors.New("need a location to explore")
    }
        
    exploreResp, err := cfg.pokeapiClient.ExploreLocation(loc)
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s... \n", loc)
    fmt.Println("Found Pokemon:")
    for _, pokemon := range exploreResp.PokemonEncounters {
        fmt.Printf(" - %s \n", pokemon.Pokemon.Name)
    }
    return nil
}

