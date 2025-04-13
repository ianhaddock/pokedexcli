package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    )

func startRepl() {

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

commands := map[string]cliCommand{
    "exit": {
        name:           "exit",
        description:    "Exit the Pokedex",
        callback:       commandExit,
    },
    "help": {
        name:           "help",
        description:    "Displays a help message",
        callback:       helpMessage,
    },
}

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        output := cleanInput(scanner.Text())
        fmt.Println(output)

        commandFound := false
        for _, command := range commands {
            if output[0] == command.name {
                commandFound = true
                err := command.callback()
                if err != nil {
                    fmt.Printf("Error: %s", err)
                }
                break

            }
        }
        if commandFound == false {
            fmt.Println("Command not found")
        }
    }
}


func helpMessage() error {
    fmt.Println("Welcome to the Pokedex!")
    return nil
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func cleanInput(text string) []string {

    var output string
    var words []string

    output = strings.ToLower(text)
    words = strings.Fields(output)

    return words
}
