package main

import (
    "fmt"
    "errors"
    "math/rand"
)

func commandCatch(cfg *config, args ...string) error {

    if len(args) != 1 {
        return errors.New("need a pokeman name to catch")
    }
        
    name := args[0]
    catchResp, err := cfg.pokeapiClient.Catch(name)
    if err != nil {
        return err
    }

    charSkill := rand.Intn(cfg.characterExp * 4)
    fmt.Printf("BaseExperience: %d\n", catchResp.BaseExperience)
    fmt.Printf("RandomCatchVal: %d\n", charSkill)

    fmt.Printf("Throwing a Pokeball at %s... \n", name)
    if charSkill > catchResp.BaseExperience {
        fmt.Printf("%s was caught!\n", name)
        cfg.characterPokedex[catchResp.Name] = catchResp
    } else {
        fmt.Printf("%s escaped!\n", name)
    }

    fmt.Println("Caught pokemon:")
    for _, pokemon := range cfg.characterPokedex {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    
    return nil
}

