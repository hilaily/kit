package listx

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

// FetchAllByBatch
// Parameter
//   f: is a function geting data by page
//   batch: is how much result you want to get once
func FetchAllByBatch[T any](
	f func(page, pageSize int) ([]T, error),
	batch int,
) ([]T, error) {
	step := 1
	ret := make([]T, 0)
	for {
		res, err := f(step, batch)
		if err != nil {
			return nil, err
		}
		ret = append(ret, res...)
		if len(res) < batch {
			break
		}
		step++
	}
	return ret, nil
}

// ThrowAllByBatch
// Parameter
//  f: f is a function to deal with the data
//  batch: is how many data you want to deal with once
//  data: is all data
func ThrowAllByBatch[T any](
	f func([]T) error,
	batch int,
	data []T,
) error {
	length := len(data)
	var err error
	for i := 0; i < length; i += batch {
		end := i + batch
		if end > length {
			end = length
		}
		err = f(data[i:end])
		if err != nil {
			return err
		}
	}
	return nil
}
