package maps

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
