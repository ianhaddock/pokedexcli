package pokecache

import (
    "time"
    "sync"
    "fmt"
)

// types

type Cache struct {
    cache map[string]cacheEntry
    mu sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

// funcs

func NewCache() Cache {
    c := Cache{}
    c.cache = make(map[string]cacheEntry)
    c.reapLoop()

    return c
}


func removeStaleCacheEntries(ch <-chan bool) error {
    check := true

    for check == true {
        check = <-ch
        if check {
            fmt.Println("remove stale entries")
        } else {
            fmt.Println("not removing entires")
        }
    }
    return nil
}

func (c *Cache) reapLoop() {   // should have interval time.Duration

    fmt.Println("reap loop started.")

    ch := make(chan bool)

    go removeStaleCacheEntries(ch)

    ch <- true
    ch <- true
    ch <- false

 /*   for str, _ := range c.cache {
        fmt.Printf("cache entry is newer than interval: %s", str)
    } */
}

func (c *Cache) Add(key string, val []byte) error {

    newEntry := cacheEntry{}
    newEntry.createdAt = time.Now()
    newEntry.val = val

    fmt.Printf("Added: %s \n", key)
    c.cache[key] = newEntry

    return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {

    v, ok := c.cache[key]
    if ok {
        return v.val, true
    }
    return []byte{}, false
}
