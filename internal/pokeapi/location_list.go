package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    dat, ok := c.pokeCache.Get(url)
//    fmt.Printf(">> %v \n", ok)
    if ok {
        locationsResp := RespShallowLocations{}
        err := json.Unmarshal(dat, &locationsResp)
        if err != nil {
            return RespShallowLocations{}, err
        }
        fmt.Println(">> pulled from cache")
        return locationsResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespShallowLocations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowLocations{}, err
    }
    defer resp.Body.Close()

    dat, err = io.ReadAll(resp.Body)
    if err != nil {
        return RespShallowLocations{}, err
    }

    err = c.pokeCache.Add(url, dat)
    if err != nil {
        return RespShallowLocations{}, err
    }

    locationsResp := RespShallowLocations{}
    err = json.Unmarshal(dat, &locationsResp)
    if err != nil {
        return RespShallowLocations{}, err
    }

    return locationsResp, nil
}
