package listx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch2(t *testing.T) {
	data := []int{}
	res, err := FetchAllByBatch(func(page int, pageSize int) ([]int, error) {
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > len(data) {
			end = len(data)
		}
		t.Logf("step: %d, res: %v", page, data[start:end])
		return data[start:end], nil
	}, 3)
	assert.NoError(t, err)
	t.Log(res)
	assert.Equal(t, data, res)

}

func TestFetch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res, err := FetchAllByBatch(func(page int, pageSize int) ([]int, error) {
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > len(data) {
			end = len(data)
		}
		t.Logf("step: %d, res: %v", page, data[start:end])
		return data[start:end], nil
	}, 3)
	assert.NoError(t, err)
	t.Log(res)
	assert.Equal(t, data, res)
}

func TestThrow2(t *testing.T) {
	data := []int{}
	err := ThrowAllByBatch(func(data []int) error {
		t.Logf("res: %v", data[:])
		return nil
	}, 4, data)
	assert.NoError(t, err)
}

func TestThrow(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	err := ThrowAllByBatch(func(data []int) error {
		t.Logf("res: %v", data[:])
		return nil
	}, 4, data)
	assert.NoError(t, err)
}
