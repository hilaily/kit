package httpx

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleResp(t *testing.T) {
	rd := bytes.NewBufferString(`{"a":1, "b":"ok"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(rd),
	}
	r := &rBody{}
	_, err := HandleResp(resp, r)
	assert.NoError(t, err)
	assert.Equal(t, 1, r.A)
	assert.Equal(t, "ok", r.B)

	rd = bytes.NewBufferString(`{"a":1, "b":"ok"}`)
	resp = &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       ioutil.NopCloser(rd),
	}
	r = &rBody{}
	_, err = HandleResp(resp, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `{"a"`)
	t.Log(err.Error())

	// parser resp err
	rd = bytes.NewBufferString(`{"a":1, "b":"ok"`)
	resp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(rd),
	}
	r = &rBody{}
	_, err = HandleResp(resp, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `{"a"`)
	t.Log(err.Error())

}

func TestGet(t *testing.T) {
	val := url.Values{}
	val.Set("aa", "bb")
	b, err := Get("https://baidu.com", &val, map[string]string{"a": "b"}, nil)
	assert.NoError(t, err)
	assert.NotZero(t, len(b))
}

type rBody struct {
	A int    `json:"a"`
	B string `json:"b"`
}
