package listx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInterface(t *testing.T) {
	arr := []int{1, 2, 3}
	data := ToInterface(arr)
	typ := reflect.TypeOf(data)
	assert.Equal(t, "[]interface {}", typ.String())
	t.Log(typ)
}

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

func TestSub(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, Sub(arr, 1, 3), []int{1, 2})
	assert.Equal(t, Sub(arr, -1, 3), []int{0, 1, 2})
	assert.Equal(t, Sub(arr, -1, 10), arr)
	assert.Panics(t, func() { Sub(arr, 5, 2) })
}
