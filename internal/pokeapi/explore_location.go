package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
)


// Explore Location -
func (c *Client) ExploreLocation(loc string) (RespExploreLocation, error) {

    url := baseURL + "/location-area/" + loc

    dat, ok := c.pokeCache.Get(url)
//    fmt.Printf(">> %v \n", ok)
    if ok {
        locationsResp := RespExploreLocation{}
        err := json.Unmarshal(dat, &locationsResp)
        if err != nil {
            return RespExploreLocation{}, err
        }
        fmt.Println(">> pulled from cache")
        return locationsResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespExploreLocation{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespExploreLocation{}, err
    }
    defer resp.Body.Close()

    dat, err = io.ReadAll(resp.Body)
    if err != nil {
        return RespExploreLocation{}, err
    }

    c.pokeCache.Add(url, dat)

    locationsResp := RespExploreLocation{}
    err = json.Unmarshal(dat, &locationsResp)
    if err != nil {
        return RespExploreLocation{}, err
    }

    return locationsResp, nil
}
