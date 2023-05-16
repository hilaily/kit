package httpx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server := http.NewServeMux()
	server.HandleFunc("/", handler)

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	data, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(data))

	header := w.Header()
	for k, v := range header {
		t.Log(k, v)
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	data := &bytes.Buffer{}
	data.WriteString("a test file")
	DownloadHander(resp, data, "test.txt")
}
