package listx

import "testing"

func TestToMap(t *testing.T) {
	arr := []string{"1", "2"}
	t.Log(ToMap(arr))

	arr1 := []int{1, 2}
	t.Log(ToMap(arr1))
}
