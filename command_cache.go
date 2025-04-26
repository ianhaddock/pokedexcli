package main

import (
   // "errors"
)

func commandListCache(cfg *config, args ...string) error {

    cfg.pokeapiClient.ListCache()

    return nil
}

