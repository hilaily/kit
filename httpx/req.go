package httpx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get Send a get request
func Get(schemaHostPath string, params *url.Values, headers map[string]string, dst interface{}) ([]byte, error) {
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
	resp, err := http.DefaultClient.Do(_req)
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
