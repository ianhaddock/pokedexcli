package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "log"
)

func commandMap() error {

    response, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        log.Fatal(err)
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

//    var results map[string]any

    type Places struct {
            Name string `json:"name"`
            Url string `json:"url"`
        }

    type Results struct {
        Count float64 `json:"count"`
        Next string `json:"next"`
        Previous string `json:"previous"`
        Places []Places `json:"results"`
    }

    results := Results{}

    err = json.Unmarshal(body, &results)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Next: %s\n", results.Next)
    fmt.Printf("Previous: %v\n", results.Previous)

    fmt.Printf("\nLocations: %s\n", results.Places[0].Name)
/*    for _, maps := range locations {
        fmt.Println(maps)
        }*/
    return nil
}

