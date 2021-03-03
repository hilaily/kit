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
		return nil, fmt.Errorf("[httptool], parse url, %w", err)
	}
	if params != nil {
		_url.RawQuery = params.Encode()
	}
	_req, err := http.NewRequest(http.MethodGet, _url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("[httptool], make request, %w", err)
	}
	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(_req)
	if err != nil {
		return nil, fmt.Errorf("[httptool], send request, %w", err)
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
		if readErr == nil && len(body) > 0 {
			return nil, fmt.Errorf("[httptool], http status: code= %d, body= %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("[httptool], http status: code= %d", resp.StatusCode)
	}
	if readErr != nil {
		return nil, fmt.Errorf("[httptool], read resp, %w", readErr)
	}
	if dst == nil {
		return body, nil
	}
	err := json.Unmarshal(body, dst)
	if err != nil {
		return nil, fmt.Errorf("[httptool], unmarshal resp, %w", err)
	}
	return body, nil
}
