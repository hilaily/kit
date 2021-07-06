package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	fileName := "./tmp/store.json"
	s, err := NewFileCache(1*time.Hour, 1*time.Hour, fileName)
	assert.NoError(t, err)
	s2, err := NewFileCache(1*time.Hour, 1*time.Hour, fileName)
	assert.NoError(t, err)
	testCache(t, s, s2)
}

func TestMemStore(t *testing.T) {
	s := NewMemCache(1*time.Hour, 1*time.Hour)
	testCache(t, s, s)
}

func testCache(t *testing.T, s, s2 IStore) {
	s.Set("test", "test")
	val, exist := s.Get("test")
	assert.True(t, exist)
	assert.Equal(t, "test", val.(string))

	s.Set("test2", "test2", 1*time.Second)
	time.Sleep(time.Second * 2)
	val, exist = s.Get("test2")
	assert.False(t, exist)
	assert.Nil(t, val)

	err := s.Save()
	assert.NoError(t, err)
	assert.NoError(t, err)
	err = s2.Load()
	assert.NoError(t, err)
	val, exist = s2.Get("test")
	assert.True(t, exist)
	assert.Equal(t, "test", val.(string))
}
