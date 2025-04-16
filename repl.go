package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    )

type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

type config struct {
    next        string
    previous    string
}

func startRepl() {

    scanner := bufio.NewScanner(os.Stdin)
    conf := config{}

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
            err := command.callback(&conf)
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


func getCommands() map[string]cliCommand {

    commands := map[string]cliCommand{
        "exit": {
            name:           "exit",
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
        },
        "map": {
            name:           "map",
            description:    "Displays 20 locations areas",
            callback:       commandMap,
        },
        "mapb": {
            name:           "mapb",
            description:    "Displays previous 20 locations areas",
            callback:       commandMapb,
        },
    }
    return commands
}
