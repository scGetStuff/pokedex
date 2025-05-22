package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{map[string]CacheEntry{}, interval, &sync.Mutex{}}
	go c.reapLoop()

	return c
}

// spec was not clear, so I am creating new or overwriting existing
func (c *Cache) Add(key string, val []byte) {
	c.MU.Lock()
	defer c.MU.Unlock()

	item, ok := c.Stuff[key]
	if !ok {
		item = CacheEntry{}
	}

	item.CreatedAt = time.Now()
	item.Val = val
	c.Stuff[key] = item
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.MU.Lock()
	defer c.MU.Unlock()

	item, ok := c.Stuff[key]
	if !ok {
		return []byte{}, false
	}

	return item.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		_ = t
		fmt.Println("Current time: ", t.Format(time.DateTime))
		now := time.Now()
		for k, v := range c.Stuff {
			delta := now.Sub(v.CreatedAt)
			fmt.Println("Seconds: ", delta.Seconds())
			if delta > c.Interval {
				c.MU.Lock()
				delete(c.Stuff, k)
				c.MU.Unlock()
			}
		}
	}
}
