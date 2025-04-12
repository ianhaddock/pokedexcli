package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    )

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        output := cleanInput(scanner.Text())
        if len(output) > 0 {
            fmt.Printf("Your command was: %s\n", output[0])
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
