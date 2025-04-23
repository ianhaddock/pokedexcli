package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    "github.com/ianhaddock/pokedexcli/internal/pokeapi"
    "github.com/ianhaddock/pokedexcli/internal/pokecache"
    )

type config struct {
    pokeapiClient pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
}

func startRepl(cfg *config) {

    urlCache := pokecache.NewCache()
    urlCache.Add("test", []byte{})
    val, ok := urlCache.Get("test")
    if ok {
        fmt.Printf("Found in cache: %v \n", val)
    } else {
        fmt.Printf("Not found in cache: %v \n", val)
    }

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


        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg)
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
    callback    func(*config) error
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
    }
    return commands
}
