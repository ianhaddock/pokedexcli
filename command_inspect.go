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
            fmt.Printf("Height: %d\n", pokemon.Height)
            fmt.Printf("Weight: %d\n", pokemon.Weight)
            fmt.Println("Stats:")
            pkStats := pokemon.Stats
//            fmt.Printf("  -%vs\n", pkStats)
            for _, name := range pkStats {
                fmt.Printf("  -%s: %d\n", name.Stat.Name, name.BaseStat)
            }
            pkTypes := pokemon.Types
            fmt.Println("Types:")
            for _, types := range pkTypes {
                fmt.Printf("  - %s\n", types.Type.Name)
            }
        } else {
            fmt.Println("you have not caught that pokemon")
        }
    }
    
    return nil
}

