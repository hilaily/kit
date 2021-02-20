package mapx

import (
	"bytes"
	"encoding/json"
)

// OrderedMap 固定顺序的 Map
type OrderedMap []*sortMapNode

// Put ...
func (c *OrderedMap) Put(key string, val interface{}) {
	index, _, ok := c.get(key)
	if ok {
		(*c)[index].Val = val
	} else {
		node := &sortMapNode{Key: key, Val: val}
		*c = append(*c, node)
	}
}

// Get ...
func (c *OrderedMap) Get(key string) (interface{}, bool) {
	_, val, ok := c.get(key)
	return val, ok
}

// MarshalJSON ...
func (c *OrderedMap) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("{")
	l := len(*c)
	for i, node := range *c {
		v := node.Val
		str := ""
		if m, ok := v.(*OrderedMap); ok {
			s, _ := m.MarshalJSON()
			str = string(s)
		} else {
			b, _ := json.Marshal(node.Val)
			str = string(b)
		}

		buf.WriteString(`"`)
		buf.WriteString(node.Key)
		buf.WriteString(`":`)
		buf.WriteString(str)
		if i != l-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

type sortMapNode struct {
	Key string
	Val interface{}
}

func (c *OrderedMap) get(key string) (int, interface{}, bool) {
	for index, node := range *c {
		if node.Key == key {
			return index, node.Val, true
		}
	}
	return -1, nil, false
}
