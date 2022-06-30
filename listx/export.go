package listx

// ToMap transfer a string slice to a map
func ToMap(data []string) map[string]struct{} {
	m := make(map[string]struct{}, len(data))
	for _, v := range data {
		m[v] = struct{}{}
	}
	return m
}

// Has represent if a element is in the list
func Has(list []string, element string) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

// Dedup remove duplicate element in slice
func Dedup(s []string) []string {
	l := len(s)
	uniq := make(map[string]struct{}, l)
	newArr := make([]string, 0, l)
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
