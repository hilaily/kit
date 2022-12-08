package listx

func Index[E comparable](list []E, element E) int {
	for i, vs := range list {
		if element == vs {
			return i
		}
	}
	return -1
}

// ToMap transfer a string slice to a map
func ToMap[E comparable](data []E) map[E]struct{} {
	m := make(map[E]struct{}, len(data))
	for _, v := range data {
		m[v] = struct{}{}
	}
	return m
}

// Has represent if a element is in the list
func Has[E comparable](list []E, element E) bool {
	return Index(list, element) > 0
}

// Dedup remove duplicate element in slice
func Dedup[E comparable](s []E) []E {
	l := len(s)
	uniq := make(map[E]struct{}, l)
	newArr := make([]E, 0, l)
	for _, v := range s {
		_, ok := uniq[v]
		if ok {
			continue
		} else {
			newArr = append(newArr, v)
			uniq[v] = struct{}{}
		}
	}
	return newArr
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ToInterface[T any](data []T) []interface{} {
	arr := make([]interface{}, len(data))
	for i, v := range data {
		arr[i] = v
	}
	return arr
}

// FetchAllByBatch
// Parameter
//
//	f: is a function getting data by page
//	batch: is how much result you want to get once
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
//
//	f: f is a function to deal with the data
//	batch: is how many data you want to deal with once
//	data: is all data
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

// GroupIt
// Parameter
//
//	total: total count of the list
//	count: number of element in a group
//
// Example:
//
//	GroupIt(23,10) will return [{Start: 0, End: 10}, {Start: 10, End 23}]
func GroupIt(total, count int) []*Range {
	ret := make([]*Range, 0, total/count+1)
	start := 0
	end := 0
	for {
		end = start + count
		if end > total {
			end = total
		}
		ret = append(ret, &Range{Start: start, End: end})
		start = end
		if end >= total {
			break
		}
	}
	return ret
}

type Range struct {
	Start int
	End   int
}
