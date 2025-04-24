package main

import (
   // "errors"
)

func commandListCache(cfg *config) error {

    cfg.pokeapiClient.ListCache()

    return nil
}

