package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4}
	t.Log(arr[1:3])
	res := GroupIt(99, 10)
	assert.Equal(t, 10, len(res))
	assert.Equal(t, 0, res[0].Start)
	assert.Equal(t, 10, res[0].End)
	assert.Equal(t, 90, res[9].Start)
	assert.Equal(t, 99, res[9].End)
	for _, v := range res {
		t.Log(v.Start, v.End)
	}

	res = GroupIt(13, 100)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, 0, res[0].Start)
	assert.Equal(t, 13, res[0].End)
	for _, v := range res {
		t.Log(v.Start, v.End)
	}
}
