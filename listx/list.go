package listx

func NewList(list []string) *List {
	m := ToMap(list)
	return &List{
		m:    m,
		list: list,
	}
}

type List struct {
	list []string
	m    map[string]struct{}
}

func (l *List) ToMap() map[string]struct{} {
	m := make(map[string]struct{}, len(l.m))
	for k, v := range l.m {
		m[k] = v
	}
	return m
}

func (l *List) Has(element string) bool {
	_, ok := l.m[element]
	return ok
}
