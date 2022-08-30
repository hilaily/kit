package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/hilaily/kit/mapx"
)

var (
	clientMap      = mapx.NewSafeMap[time.Duration, *http.Client](2)
	defaultTimeout = 5 * time.Second
)

// Get Send a get request
func Get(schemaHostPath string, params *url.Values, headers map[string]string, dst interface{}, timeout ...time.Duration) ([]byte, error) {
	to := defaultTimeout
	if len(timeout) > 0 {
		to = timeout[0]
	}
	_url, err := url.Parse(schemaHostPath)
	if err != nil {
		return nil, fmt.Errorf("[httpx], parse url, %w", err)
	}
	q := _url.Query()
	if params != nil {
		for k, v := range *params {
			if len(v) > 0 {
				q.Set(k, v[0])
			}
		}
	}

	_url.RawQuery = q.Encode()
	_req, err := http.NewRequest(http.MethodGet, _url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("[httpx], make request, %w", err)
	}
	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	resp, err := getClient(to).Do(_req)
	if err != nil {
		return nil, fmt.Errorf("[httpx], send request, %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return HandleResp(resp, dst)
}

// Post Send a post request
func Post(schemaHostPath string, headers map[string]string, body io.Reader, dst interface{}, timeout ...time.Duration) ([]byte, error) {
	to := defaultTimeout
	if len(timeout) > 0 {
		to = timeout[0]
	}

	_url, err := url.Parse(schemaHostPath)
	if err != nil {
		return nil, fmt.Errorf("[httpx], parse url, %w", err)
	}
	_req, err := http.NewRequest(http.MethodPost, _url.String(), body)
	if err != nil {
		return nil, fmt.Errorf("[httpx], make request, %w", err)
	}
	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	resp, err := getClient(to).Do(_req)
	if err != nil {
		return nil, fmt.Errorf("[httpx], send request, %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return HandleResp(resp, dst)
}

// HandleResp check http status code and unmarshal a response body
func HandleResp(resp *http.Response, dst interface{}) ([]byte, error) {
	body, readErr := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("[httpx], http status: code= %d, body= %s", resp.StatusCode, string(body))
	}
	if readErr != nil {
		return body, fmt.Errorf("[httpx], read resp, %w ", readErr)
	}
	if dst == nil {
		return body, nil
	}
	err := json.Unmarshal(body, dst)
	if err != nil {
		return body, fmt.Errorf("[httpx], unmarshal resp, body: %s, err: %w", string(body), err)
	}
	return body, nil
}

func getClient(timeout time.Duration) *http.Client {
	v, ok := clientMap.Get(timeout)
	if ok {
		return v
	}
	c := &http.Client{Timeout: timeout}
	clientMap.Set(timeout, c)
	return c
}
