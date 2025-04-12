package main

import (
    "strings"
    )


func cleanInput(text string) []string {

    var output string
    var words []string

    output = strings.ToLower(text)
    words = strings.Fields(output)

    return words
}
