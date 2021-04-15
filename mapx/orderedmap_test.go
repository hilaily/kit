package mapx

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	m := OrderedMap{}
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	res := make([]int, 0)
	m.Range(func(k string, val interface{}) error {
		fmt.Println(k, val)
		v, _ := val.(int)
		res = append(res, v)
		return nil
	})
	fmt.Println(res)
}

func TestMap(t *testing.T) {
	m := OrderedMap{}
	m.Put("a", 1)
	m.Put("d", "2")
	m.Put("b", "3")
	m.Put("c", []string{"4", "4"})
	res, _ := m.MarshalJSON()
	mm := make(map[string]interface{}, 0)
	err := json.Unmarshal(res, &mm)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(res))
}
