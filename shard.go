package botandb

import (
	"sync"
	"time"
)

// Entry is a single entry in database.
type Entry struct {
	Key        string
	Value      interface{}
	Expiration time.Time
}

// Shard is a partition of database.
type Shard struct {
	sync.RWMutex
	data map[string]*Entry
}
