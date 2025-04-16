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
    callback    func() error
}

type config struct {
    next        string
    previous    string
}

func startRepl() {

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
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }
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
    }
    return commands
}
