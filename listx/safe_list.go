package listx

import "sync"

// NewSafeList ...
func NewSafeList[T any](caps ...int) *SafeList[T] {
	s := &SafeList[T]{
		RWMutex: &sync.RWMutex{},
	}
	_cap := 0
	if len(caps) > 0 {
		_cap = caps[0]
	}
	s.list = make([]T, 0, _cap)
	return s
}

type SafeList[T any] struct {
	*sync.RWMutex
	list []T
}

func (s *SafeList[T]) Append(e T) {
	s.Lock()
	s.list = append(s.list, e)
	s.Unlock()
}

func (s *SafeList[T]) Len() int {
	return len(s.list)
}

func (s *SafeList[T]) List() []T {
	arr := make([]T, 0, len(s.list))
	s.RLock()
	arr = append(arr, s.list...)
	s.RUnlock()
	return arr
}
