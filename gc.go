package botandb

import "time"

// GC - Garbage Collection.
func (b *BotanDB) gc(frequency time.Duration) {
	for {
		time.Sleep(frequency)

		b.cleanUp()
	}
}

// cleanUp starts the process of cleaning up storage.
func (b *BotanDB) cleanUp() {
	for _, shard := range b.shards {
		shard.Lock()
		for key, entry := range shard.data {
			if entry.Expiration.Before(time.Now()) {
				delete(shard.data, key)
			}
		}
		shard.Unlock()
	}
}
