package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    "github.com/ianhaddock/pokedexcli/internal/pokeapi"
    )

type config struct {
    pokeapiClient pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
    characterExp int
    characterPokedex map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        output := cleanInput(scanner.Text())
        //fmt.Println(output)

        if len(output) == 0 {
            continue
        }

        commandName := output[0]
        args := []string{}
        if len(output) > 1 {
            args = output[1:]
        }

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg, args...)
            if err != nil {
                fmt.Println(err)
            }
        //    fmt.Printf("next: %s, prev: %s\n", conf.next, conf.previous)
            continue
        } else {
            fmt.Println("Unknownn command")
            continue
        }
    }
}


func cleanInput(text string) []string {

    var output string
    var words []string

    output = strings.ToLower(text)
    words = strings.Fields(output)

    return words
}


type cliCommand struct {
    name        string
    description string
    callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

    commands := map[string]cliCommand{
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
        },
        "map": {
            name:           "map",
            description:    "Get the next page of locations",
            callback:       commandMap,
        },
        "mapb": {
            name:           "mapb",
            description:    "Get the previous page of locations",
            callback:       commandMapb,
        },
        "exit": {
            name:           "exit",
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
        "cache": {
            name:           "cache",
            description:    "List all cache entries",
            callback:       commandListCache,
        },
        "explore": {
            name:           "explore <location_name>",
            description:    "Explore a location by name",
            callback:       commandExplore,
        },
        "catch": {
            name:           "catch <pokemon_name>",
            description:    "Catch a pokemon by name",
            callback:       commandCatch,
        },
        "inspect": {
            name:           "inspect <pokemon_name>",
            description:    "Inspect a pokemon by name",
            callback:       commandInspect,
        },
        "pokedex": {
            name:           "pokedex",
            description:    "List names of caught pokemon",
            callback:       commandPokedex,
        },
    }
    return commands
}
