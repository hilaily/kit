package mapx

import "sync"

func NewSafeMap[K comparable, V any](caps ...int) *SafeMap[K, V] {
	s := &SafeMap[K, V]{
		RWMutex: &sync.RWMutex{},
	}
	_cap := 0
	if len(caps) > 0 {
		_cap = caps[0]
	}
	s.m = make(map[K]V, _cap)
	return s
}

type SafeMap[K comparable, V any] struct {
	*sync.RWMutex
	m map[K]V
}

func (s *SafeMap[K, V]) Set(k K, v V) {
	s.Lock()
	s.m[k] = v
	s.Unlock()
}

func (s *SafeMap[K, V]) Get(k K) (V, bool) {
	s.RLock()
	v, ok := s.m[k]
	s.RUnlock()
	return v, ok
}

func (s *SafeMap[K, V]) Len() int {
	return len(s.m)
}

func (s *SafeMap[K, V]) Range(f func(k K, v V) error) error {
	s.RLock()
	for k, v := range s.m {
		err := f(k, v)
		if err != nil {
			return err
		}
	}
	s.RUnlock()
	return nil
}

func (s *SafeMap[K, V]) ToMap() map[K]V {
	s.RLock()
	mm := make(map[K]V, len(s.m))
	for k := range s.m {
		mm[k] = s.m[k]
	}
	s.RUnlock()
	return mm
}
