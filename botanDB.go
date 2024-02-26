package botandb

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

const DefaultGCFrequency = 10 * time.Minute
const DefaultTTL = 24 * time.Hour

var (
	// ErrKeyNotFound is returned, when prompted key either does not exist in db, or expired.
	ErrKeyNotFound = errors.Errorf("key not found")
)

// BotanDB is a struct, that implements an inmemory database service.
type BotanDB struct {
	shards []*Shard
}

// NewClient returns new client to interract with BotanDB.
func NewClient(shardsAmount int, gcFrequency time.Duration) *BotanDB {
	botanClient := &BotanDB{
		shards: make([]*Shard, shardsAmount),
	}
	for i := range botanClient.shards {
		botanClient.shards[i] = &Shard{
			data: make(map[string]*Entry),
		}
	}
	if gcFrequency <= 0 {
		gcFrequency = DefaultGCFrequency
	}
	go botanClient.gc(gcFrequency)

	return botanClient
}

// Set saves a key-value pair in botanDB.
// If expiration is set to zero, it would be set to default value.
func (b *BotanDB) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) {
	shard := b.findShard(key)
	if ttl == 0 {
		ttl = DefaultTTL
	}
	shard.Lock()
	defer shard.Unlock()
	entry := &Entry{
		Key:        key,
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
	shard.data[key] = entry
}

// Get retrieves a value by given key from botanDB.
func (b *BotanDB) Get(ctx context.Context, key string) (interface{}, error) {
	shard := b.findShard(key)
	shard.RLock()
	defer shard.RUnlock()

	entry, exists := shard.data[key]
	if !exists || entry.Expiration.Before(time.Now()) {
		return nil, ErrKeyNotFound
	}
	return entry.Value, nil
}

// Delete deletes a key value pair from botanDB.
func (b *BotanDB) Delete(ctx context.Context, key string) error {
	shard := b.findShard(key)
	shard.Lock()
	defer shard.Unlock()

	entry, exists := shard.data[key]
	if !exists || entry.Expiration.Before(time.Now()) {
		return ErrKeyNotFound
	}

	delete(shard.data, key)

	return nil
}

// findShard returns shard that contains given key.
func (b *BotanDB) findShard(key string) *Shard {
	hash := hashFunc(key)
	shardIndex := hash % len(b.shards)
	return b.shards[shardIndex]
}

// hashFunc is a simplest hash function.
func hashFunc(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash
}
