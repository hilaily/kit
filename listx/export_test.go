package listx

import (
	"testing"
)

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4}
	Reverse(a)
	t.Log(a)
	Reverse(a)
	t.Log(a)
}

func TestToMap(t *testing.T) {
	arr := []string{"1", "2"}
	t.Log(ToMap(arr))

	arr1 := []int{1, 2}
	t.Log(ToMap(arr1))
}
