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
