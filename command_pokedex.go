package main

import (
    "fmt"
    "errors"
)

func commandPokedex(cfg *config, args ...string) error {

    if len(cfg.characterPokedex) < 1 {
        return errors.New("No pokemon caught")
    }
    fmt.Println("Your Pokedex:")
    for _, pokemon := range cfg.characterPokedex {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    
    return nil
}

