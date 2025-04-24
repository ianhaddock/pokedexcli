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


func (c *Cache) reapLoop() {   // should have interval time.Duration

    fmt.Println("reap loop started.")

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
