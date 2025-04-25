package main

import (
   // "errors"
)

func commandListCache(cfg *config, loc string) error {

    cfg.pokeapiClient.ListCache()

    return nil
}

