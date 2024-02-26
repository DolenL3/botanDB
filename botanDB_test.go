package botandb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var clientBotan = NewClient(0, 0)
var ctxBotan = context.Background()

// TestGet
func TestGet(t *testing.T) {
	testKey := "foo"
	testValue := "bar"

	t.Run("Existing key", func(t *testing.T) {
		err := clientBotan.Set(ctx, testKey, testValue, 0)
		assert.Nil(t, err)

		val, err := clientBotan.Get(ctxBotan, testKey)
		assert.Equal(t, testValue, val)
		assert.Nil(t, err)
	})

	t.Run("Non-existent key", func(t *testing.T) {
		_, err := clientBotan.Get(ctxBotan, "non_existent_key")
		assert.Equal(t, ErrKeyNotFound, err)
	})

	t.Run("Expired key", func(t *testing.T) {
		expiredKey := "expired_key"
		shard := clientBotan.findShard(expiredKey)
		shard.Lock()
		shard.data[expiredKey] = &Entry{
			Key:        expiredKey,
			Value:      "",
			Expiration: time.Now().Add(time.Hour),
		}
		shard.Unlock()

		_, err := clientBotan.Get(ctxBotan, expiredKey)
		assert.Equal(t, ErrKeyNotFound, err)
	})
}

// test_delete
func TestDelete(t *testing.T) {
	testKey := "for_delete"

	err := clientBotan.Set(ctxBotan, testKey, "value", 0)
	assert.Nil(t, err)

	t.Run("Delete existing key", func(t *testing.T) {
		err := clientBotan.Delete(ctxBotan, testKey)
		assert.Nil(t, err)
	})

	t.Run("Delete non-existent key", func(t *testing.T) {
		err := clientBotan.Delete(ctxBotan, testKey)
		assert.Equal(t, ErrKeyNotFound, err)
	})

	t.Run("Delete from map", func(t *testing.T) {
		shard := clientBotan.findShard(testKey)
		shard.RLock()
		_, ok := shard.data[testKey]
		shard.RUnlock()

		assert.False(t, ok)
	})
}
