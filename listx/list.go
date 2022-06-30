package listx

func NewList[E comparable](list []E) *List[E] {
	m := ToMap(list)
	return &List[E]{
		m:    m,
		list: list,
	}
}

type List[E comparable] struct {
	list []E
	m    map[E]struct{}
}

func (l *List[E]) ToMap() map[E]struct{} {
	m := make(map[E]struct{}, len(l.m))
	for k, v := range l.m {
		m[k] = v
	}
	return m
}

func (l *List[E]) Has(element E) bool {
	_, ok := l.m[element]
	return ok
}

func (l *List[E]) Index(list []E, element E) int {
	for i, vs := range list {
		if element == vs {
			return i
		}
	}
	return -1
}
