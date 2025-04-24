package pokecache

import (
    "time"
    "sync"
    "fmt"
)

// types

type Cache struct {
    cache map[string]cacheEntry
    mu *sync.RWMutex
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

// funcs

func NewCache(interval time.Duration) Cache {
    c := Cache{
        cache: make(map[string]cacheEntry),
        mu: &sync.RWMutex{},
    }
    go c.reapLoop(interval)

    return c
}

func (c *Cache) ListCache() {
    for string, entry := range c.cache {
        fmt.Printf("cache entry: %s \n", string)
        fmt.Printf("created at: %s \n", entry.createdAt)
    }
}


func (c *Cache) reapLoop(interval time.Duration) {

    ticker := time.NewTicker(interval) 
    defer ticker.Stop()

    for range ticker.C {
        currentTime := time.Now()

        fmt.Printf(">> cache size: %d \n", len(c.cache))
        for string, entry := range c.cache {
            fmt.Printf("cache entry: %s \n", string)
            fmt.Printf("created at: %s \n", entry.createdAt)
            elapsed := currentTime.Sub(entry.createdAt)
            fmt.Printf("Elapsed time: %s \n", elapsed)
            if elapsed > interval {
                fmt.Println("removing entry")
                c.mu.Lock()
                delete(c.cache, string)
                c.mu.Unlock()
            }
        }
    }
}

func (c *Cache) Add(key string, val []byte) {

    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
    //fmt.Printf("Added: %s \n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {

    c.mu.Lock()
    defer c.mu.Unlock()
    v, ok := c.cache[key]
    if ok {
        return v.val, true
    }
    return []byte{}, false
}
