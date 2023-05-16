package httpx

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	s, err := Download("", nil, nil, "/tmp/ttt.tar.gz")
	assert.NoError(t, err)
	assert.NotZero(t, s)
	t.Log(s)
}

func TestTimeout(t *testing.T) {
	Convey("TestTimeout", t, func() {
		Convey("test default timeout", func() {
			start := time.Now()
			_, err := Get("https://httpstat.us/200?sleep=60000", nil, nil, nil)
			So(err, ShouldBeError)
			tt := time.Since(start)
			So(int(tt.Seconds()), ShouldEqual, 5)
			t.Log(tt)
		})

		Convey("test custome timeout", func() {
			start := time.Now()
			_, err := Get("https://httpstat.us/200?sleep=60000", nil, nil, nil, 10*time.Second)
			So(err, ShouldBeError)
			tt := time.Since(start)
			So(int(tt.Seconds()), ShouldEqual, 10)
			t.Log(tt)
		})
	})
}

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
