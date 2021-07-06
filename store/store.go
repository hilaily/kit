package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// IStore 通用存储结构
type IStore interface {
	// Load 初始化或者存存储里读取历史数据
	Load() error
	// Save 保存数据
	Save() error
	Get(key string) (val interface{}, exist bool)
	Set(key string, val interface{}, expire ...time.Duration)
}

func NewFileCache(defaultExpiration, cleanupInterval time.Duration, filepath string) (*fileCache, error) {
	f := &fileCache{
		defaultExpire:   defaultExpiration,
		cleanupInterval: cleanupInterval,
		filepath:        filepath,
		once:            &sync.Once{},
	}
	return f, f.Load()
}

func NewMemCache(defaultExpiration, cleanupInterval time.Duration) *memCache {
	return &memCache{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

type memCache struct {
	c *cache.Cache
}

func (m *memCache) Load() error {
	return nil
}

func (m *memCache) Save() error {
	return nil
}

func (m *memCache) Set(key string, val interface{}, expire ...time.Duration) {
	d := time.Duration(0)
	if len(expire) > 0 {
		d = expire[0]
	}
	m.c.Set(key, val, d)
}

func (m memCache) Get(key string) (val interface{}, exist bool) {
	return m.c.Get(key)
}

type fileCache struct {
	defaultExpire   time.Duration
	cleanupInterval time.Duration
	store           *cache.Cache
	filepath        string
	once            *sync.Once
}

func (f fileCache) Save() error {
	data := f.store.Items()
	en, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(f.filepath), 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(f.filepath, en, 0777)
	return err
}

func (f fileCache) Get(key string) (interface{}, bool) {
	return f.store.Get(key)
}

func (f fileCache) Set(key string, val interface{}, expire ...time.Duration) {
	ex := time.Duration(0)
	if len(expire) > 0 {
		ex = expire[0]
	}
	f.store.Set(key, val, ex)
}

func (f *fileCache) Load() error {
	en, e := ioutil.ReadFile(f.filepath)
	if e != nil {
		if !os.IsNotExist(e) {
			return e
		}
	}
	data := make(map[string]cache.Item, 0)
	if len(en) > 0 {
		e = json.Unmarshal(en, &data)
		if e != nil {
			return e
		}
	}
	f.store = cache.NewFrom(f.defaultExpire, f.cleanupInterval, data)
	return nil
}
