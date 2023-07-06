package listx

// ShiftHead use a element as new head
// [1,2,3,4,5] shift with index 2 -> [3,4,5,1,2]
func ShiftHead[T any](arr []T, index int) []T {
	l := len(arr)
	if index >= l {
		return arr
	}
	arr = append(arr, arr...)
	arr = arr[index+1 : index+l+1]
	return arr
}

func NewList[E comparable](list []E) *List[E] {
	m := ToMap(list)
	return &List[E]{
		m:     m,
		list:  list,
		count: len(list),
	}
}

type List[E comparable] struct {
	list  []E
	m     map[E]struct{}
	count int
}

func (l *List[E]) Len() int {
	return l.count
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
