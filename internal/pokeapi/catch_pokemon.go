package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
)


// Catch Pokemon -
func (c *Client) Catch(name string) (Pokemon, error) {

    url := baseURL + "/pokemon/" + name

    dat, ok := c.pokeCache.Get(url)
//    fmt.Printf(">> %v \n", ok)
    if ok {
        pokemonResp := Pokemon{}
        err := json.Unmarshal(dat, &pokemonResp)
        if err != nil {
            return Pokemon{}, err
        }
        fmt.Println(">> pulled from cache")
        return pokemonResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return Pokemon{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }
    defer resp.Body.Close()

    dat, err = io.ReadAll(resp.Body)
    if err != nil {
        return Pokemon{}, err
    }

    c.pokeCache.Add(url, dat)

    pokemonResp := Pokemon{}
    err = json.Unmarshal(dat, &pokemonResp)
    if err != nil {
        return Pokemon{}, err
    }

    return pokemonResp, nil
}
