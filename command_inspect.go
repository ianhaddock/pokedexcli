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

    pokemon, ok := cfg.characterPokedex[name]
    if !ok {
        return errors.New("you have not caught that pokemon")
    }

    fmt.Println("Name:", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")

    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
    }
   
    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Println("  -", typeInfo.Type.Name)
        }
    
    return nil
}

