package mapx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAutoDel(t *testing.T) {
	autoDelInterval = time.Second * 3

	a := NewCacheMap2[string](5*time.Second, true)
	a.Set("1", "1")
	a.SetWithTime("2", "2", time.Now().Add(time.Hour))
	time.Sleep(7 * time.Second)

	_, ok := a.Get("1")
	assert.False(t, ok)
	assert.Equal(t, 1, len(a.m))
	t.Logf("%v", a.m)
	_, ok = a.Get("2")
	assert.True(t, ok)
}
