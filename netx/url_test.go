package netx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLQueryParams(t *testing.T) {
	_url := "http://a.com"
	url2 := "http://a.com/path"
	url3 := "http://a.com/path?c=3"
	url4 := "http://a.com/path?a=4"
	data := map[string]string{
		"a": "1",
		"b": "2",
	}
	res, _ := URLQueryParams(_url, data)
	assert.Equal(t, "http://a.com?a=1&b=2", res)
	res, _ = URLQueryParams(url2, data)
	assert.Equal(t, "http://a.com/path?a=1&b=2", res)
	res, _ = URLQueryParams(url3, data)
	assert.Equal(t, "http://a.com/path?a=1&b=2&c=3", res)
	res, _ = URLQueryParams(url4, data)
	assert.Equal(t, "http://a.com/path?a=4&a=1&b=2", res)
}
