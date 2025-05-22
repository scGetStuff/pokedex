package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Stuff    map[string]CacheEntry
	Interval time.Duration
	MU       *sync.Mutex
}
