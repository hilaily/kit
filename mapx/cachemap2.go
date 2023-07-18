package mapx

import (
	"sync"
	"time"
)

/*
带缓存的 map
*/
var (
	autoDelInterval = 5 * time.Minute
)

type (
	// CacheMap 带缓存的 m
	CacheMap2[T any] struct {
		lock      *sync.RWMutex
		cacheTime time.Duration
		m         map[string]*val2[T]
	}
	val2[T any] struct {
		t int64
		v T
	}
)

// NewCacheMap ...
//
// param d 缓存时长
func NewCacheMap2[T any](d time.Duration, autoDel bool) *CacheMap2[T] {
	c := &CacheMap2[T]{}
	c.lock = &sync.RWMutex{}
	c.cacheTime = d
	c.m = make(map[string]*val2[T])
	if autoDel {
		c.autoDel()
	}
	return c
}

// Set 设置值
//
// Param k: map 的 key
// Param v: map 的 value
func (c *CacheMap2[T]) Set(k string, v T) {
	c.lock.Lock()
	c.m[k] = &val2[T]{time.Now().Add(c.cacheTime).Unix(), v}
	c.lock.Unlock()
}

// SetWithTime set a key val with expire time
func (c *CacheMap2[T]) SetWithTime(k string, v T, t time.Time) {
	c.lock.Lock()
	c.m[k] = &val2[T]{t.Unix(), v}
	c.lock.Unlock()
}

// Get 根据 key 获取 value
//
//	return interface{} value 值
//	return bool value 是否存在
func (c *CacheMap2[T]) Get(k string) (T, bool) {
	now := time.Now().Unix()
	c.lock.RLock()
	v, ok := c.m[k]
	c.lock.RUnlock()
	if !ok || now > v.t {
		var t T
		return t, false
	}
	return v.v, ok
}

// Del 删除 key
func (c *CacheMap2[T]) Del(k string) {
	c.lock.Lock()
	delete(c.m, k)
	c.lock.Unlock()
}

// autoDel 自动删除过期的 key
func (c *CacheMap2[T]) autoDel() {
	f := func() {
		c.lock.RLock()
		now := time.Now().Unix()
		mm := make([]string, 0, len(c.m))
		for k, v := range c.m {
			if v.t < now {
				mm = append(mm, k)
			}
		}
		for _, v := range mm {
			delete(c.m, v)
		}
		c.lock.RUnlock()
	}
	go func() {
		for {
			time.Sleep(autoDelInterval)
			f()
		}
	}()
}
