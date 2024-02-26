package botandb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientType = NewClient(3, 0)

var testBoolKey = "bkey"
var testBoolValue = true
var testIntKey = "ikey"
var testIntValue = 13
var testStringKey = "skey"
var testStringValue = "test"

var ctxType = context.Background()

// test bool
func TestBool(t *testing.T) {
	err := clientType.Set(ctxType, testBoolKey, testBoolValue, 0)
	assert.Nil(t, err)
	err = clientType.Set(ctxType, testStringKey, testStringValue, 0)
	assert.Nil(t, err)

	t.Run("Get bool like bool", func(t *testing.T) {
		val, err := clientType.GetBool(ctxType, testBoolKey)
		assert.Equal(t, testBoolValue, val)
		assert.Nil(t, err)
	})

	t.Run("Get string like bool", func(t *testing.T) {
		_, err := clientType.GetBool(ctxType, testStringKey)
		assert.Equal(t, ErrTypeConverted, err)
	})
}

// test int
func TestInt(t *testing.T) {
	err := clientType.Set(ctxType, testIntKey, testIntValue, 0)
	assert.Nil(t, err)
	err = clientType.Set(ctxType, testStringKey, testStringValue, 0)
	assert.Nil(t, err)

	t.Run("Get int like int", func(t *testing.T) {
		val, err := clientType.GetInt(ctxType, testIntKey)
		assert.Equal(t, testIntValue, val)
		assert.Nil(t, err)
	})

	t.Run("Get string like int", func(t *testing.T) {
		_, err := clientType.GetInt(ctxType, testStringKey)
		assert.Equal(t, ErrTypeConverted, err)
	})
}

// test string
func TestString(t *testing.T) {
	err := clientType.Set(ctxType, testStringKey, testStringValue, 0)
	assert.Nil(t, err)
	err = clientType.Set(ctxType, testBoolKey, testBoolValue, 0)
	assert.Nil(t, err)

	t.Run("Get string like string", func(t *testing.T) {
		val, err := clientType.GetString(ctxType, testStringKey)
		assert.Equal(t, testStringValue, val)
		assert.Nil(t, err)
	})

	t.Run("Get string like bool", func(t *testing.T) {
		_, err := clientType.GetString(ctxType, testBoolKey)
		assert.Equal(t, ErrTypeConverted, err)
	})
}
