package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	// Note SyncMap has method Lead , Store , Delete by key so will use these
	cacheMap sync.Map // Using sync's Map to access MAP concurrently by key.
}

type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// uses sync.Map Store and Load methods
func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	// We store key and Value , Value is CacheItem struct ( to add value and TTL)
	c.cacheMap.Store(key, CacheItem{value, time.Now().Add(expiration)})
}
func (c *Cache) Get(key string) (interface{}, bool) {
	if val, ok := c.cacheMap.Load(key); ok { // uses sync.Map Load methods
		item := val.(CacheItem)
		if time.Now().Before(item.Expiration) {
			return item.Value, true // if time before TTL return value
		} else {
			c.cacheMap.Delete(key) // else delete key
		}
	}
	return nil, false
}

func (c *Cache) Delete(key string) {
	c.cacheMap.Delete(key)
}
func main() {
	cache := Cache{}

	// Add data to cache
	cache.Set("mykey", "myvalue", 5*time.Second)

	// Retrieve data from cache
	if val, ok := cache.Get("mykey"); ok {
		fmt.Println("Value found in cache:", val)
	} else {
		fmt.Println("Value not found in cache")
		// Next find the Value in Database
		// if Found in Database then Add to Cache with cache.Set ( Key,Value , TTL)
		// if NOT found in Database as well then return Error
	}

	// Wait for cache to expire
	time.Sleep(6 * time.Second)

	// Retrieve expired data from cache
	if val, ok := cache.Get("mykey"); ok {
		fmt.Println("Value found in cache:", val)
	} else {
		fmt.Println("Value not found in cache")
	}
}
