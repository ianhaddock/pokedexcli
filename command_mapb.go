package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "log"
)

func commandMapb(conf *config) error {

    if conf.previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }

    response, err := http.Get(conf.previous)
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

    conf.next = results.Next
    conf.previous = results.Previous

    for _, location_names := range results.Places {
        fmt.Println(location_names.Name)
        }

    return nil
}

