package httpx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapHTTPProxy(t *testing.T) {
	client := http.DefaultClient
	client, err := WrapHTTPProxy(client, "", true)
	assert.NoError(t, err)
	resp, err := client.Get("https://amazon.com/")
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	t.Logf("%v\n", resp)
}
