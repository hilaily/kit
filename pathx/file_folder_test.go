package pathx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	path1 := "./testdata/exist.json"
	path2 := "./testdata/empty.json"
	f, err := os.Create(path1)
	assert.NoError(t, err)
	_ = f.Close()
	assert.True(t, IsExist(path1))
	assert.False(t, IsExist(path2))
}

func TestAppendFile(t *testing.T) {
	f := "./testdata/a.json"
	err := AppendFile(f, []byte("1111"), 0777)
	assert.NoError(t, err)
	err = AppendFile(f, []byte("2222"), 0777)
	assert.NoError(t, err)
}

func TestNewEntity(t *testing.T) {
	path := "./testdata/entity.json"
	e := NewEntity(path)
	assert.False(t, e.Exist())
	f, err := os.Create(path)
	assert.NoError(t, err)
	_ = f.Close()
	e.Reload()
	assert.True(t, e.Exist())
	assert.False(t, e.IsDir())
	_ = os.Remove(path)
}
