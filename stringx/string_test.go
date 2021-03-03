package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	data := []struct {
		data   string
		expect string
	}{
		{data: "a b c", expect: "abc"},
		{data: `a 
b c`, expect: "abc"},
		{data: "a	b", expect: "ab"},
	}
	for _, v := range data {
		res := Compress(v.data)
		if res != v.expect {
			t.Errorf("res: %s, expect: %s", res, v.expect)
		}
	}
}

func TestURLJoin(t *testing.T) {
	data := [][]string{
		{"http://a.com/", "/b", "/c"},
		{"http://a.com/", "b/", "/c"},
		{"http://a.com", "/b", "c"},
		{"http://a.com", "b", "/c"},
	}
	expect := "http://a.com/b/c"
	for _, v := range data {
		e := URLJoin(v[0], v[1:]...)
		if e != expect {
			t.Errorf("err: res %s, %v, %s\n", e, v[0], v[0:])
		}
	}
}

func TestAddURLSchema(t *testing.T) {
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("a/b", "http://"))
	assert.Equal(t, "http://a/b", AddURLSchema("a/b", "http"))
	assert.Equal(t, "https://a/b", AddURLSchema("a/b", "https"))
}
