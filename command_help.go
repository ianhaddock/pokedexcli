package main

import (
    "fmt"
    )


func commandHelp(conf *config) error {
    fmt.Println("\nWelcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")

    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}

