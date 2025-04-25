package main

import (
    "fmt"
    "os"
    )

func commandExit(conf *config, loc string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    fmt.Println(loc)
    os.Exit(0)
    return nil
}
