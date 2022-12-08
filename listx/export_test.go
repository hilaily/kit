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
