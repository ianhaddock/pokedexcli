package main

import (
    "fmt"
    "errors"
)

func commandInspect(cfg *config, args ...string) error {

    if len(args) != 1 {
        return errors.New("Need a pokemon name to inspect.")
    }

    name := args[0]

    for _, pokemon := range cfg.characterPokedex {
        if pokemon.Name == name {
            fmt.Printf("Name: %s\n", pokemon.Name)
        } else {
            fmt.Println("you have not caught that pokemon")
        }
    }
    
    return nil
}

