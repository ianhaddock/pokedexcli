package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
)


// Explore Location -
func (c *Client) ExploreLocation(loc string) (Location, error) {

    url := baseURL + "/location-area/" + loc

    dat, ok := c.pokeCache.Get(url)
//    fmt.Printf(">> %v \n", ok)
    if ok {
        locationsResp := Location{}
        err := json.Unmarshal(dat, &locationsResp)
        if err != nil {
            return Location{}, err
        }
        fmt.Println(">> pulled from cache")
        return locationsResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return Location{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Location{}, err
    }
    defer resp.Body.Close()

    dat, err = io.ReadAll(resp.Body)
    if err != nil {
        return Location{}, err
    }

    c.pokeCache.Add(url, dat)

    locationsResp := Location{}
    err = json.Unmarshal(dat, &locationsResp)
    if err != nil {
        return Location{}, err
    }

    return locationsResp, nil
}
