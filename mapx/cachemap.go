package mapx

import (
	"sync"
	"time"
)

/*
	带缓存的 map
*/

type (
	// CacheMap 带缓存的 m
	CacheMap struct {
		lock      *sync.RWMutex
		cacheTime time.Duration
		m         map[string]*val
	}
	val struct {
		t int64
		v interface{}
	}
)

// NewCacheMap ...
//
// param d 缓存时长
func NewCacheMap(d time.Duration, autoDel bool) *CacheMap {
	c := &CacheMap{}
	c.lock = &sync.RWMutex{}
	c.cacheTime = d
	c.m = make(map[string]*val)
	if autoDel {
		c.autoDel()
	}
	return c
}

// Set 设置值
//
// param k map 的 key
// param v map 的 value
func (c *CacheMap) Set(k string, v interface{}) {
	c.lock.Lock()
	c.m[k] = &val{time.Now().Add(c.cacheTime).Unix(), v}
	c.lock.Unlock()
}

// SetWithTime set a key val with expire time
func (c *CacheMap) SetWithTime(k string, v interface{}, t time.Time) {
	c.lock.Lock()
	c.m[k] = &val{t.Unix(), v}
	c.lock.Unlock()
}

// Get 根据 key 获取 value
//
// return interface{} value 值
// return bool value 是否存在
func (c *CacheMap) Get(k string) (interface{}, bool) {
	now := time.Now().Unix()
	c.lock.RLock()
	v, ok := c.m[k]
	c.lock.RUnlock()
	if !ok || now > v.t {
		return nil, false
	}
	return v.v, ok
}

// Del 删除 key
func (c *CacheMap) Del(k string) {
	c.lock.Lock()
	delete(c.m, k)
	c.lock.Unlock()
}

// autoDel 自动删除过期的 key
func (c *CacheMap) autoDel() {
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
