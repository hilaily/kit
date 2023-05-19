package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/hilaily/kit/mapx"
)

var (
	clientMap      = mapx.NewSafeMap[time.Duration, *http.Client](2)
	defaultTimeout = 5 * time.Second
)

// BuildURLWithQueryParams ...
func BuildURLWithQueryParams(baseURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	query := parsedURL.Query()
	for key, value := range params {
		if value != "" {
			query.Set(key, value)
		}
	}
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

// Get Send a get request
func Get(schemaHostPath string, params *url.Values, headers map[string]string, dst interface{}, timeout ...time.Duration) ([]byte, error) {
	to := defaultTimeout
	if len(timeout) > 0 {
		to = timeout[0]
	}

	_url, err := AddParamsToURL(schemaHostPath, params)
	if err != nil {
		return nil, err
	}
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
	body, readErr := io.ReadAll(resp.Body)
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

// Download ...
func Download(fileURL string, params *url.Values, headers map[string]string, filepath string, timeout ...time.Duration) (int64, error) {
	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		return 0, fmt.Errorf("download file fail, %s, %w", fileURL, err)
	}
	defer file.Close()

	// Put content on file
	_url, err := AddParamsToURL(fileURL, params)
	if err != nil {
		return 0, err
	}
	_req, err := http.NewRequest(http.MethodGet, _url.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("[httpx], make request, %w", err)
	}
	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	client := http.DefaultClient
	if len(timeout) > 0 {
		client = getClient(timeout[0])
	}

	resp, err := client.Do(_req)
	if err != nil {
		return 0, fmt.Errorf("http get file fail, %s, %w", fileURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		en, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("%s", string(en))
	}
	size, err := io.Copy(file, resp.Body)
	return size, err
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

// AddParamsToURL ...
func AddParamsToURL(originURL string, params *url.Values) (*url.URL, error) {
	_url, err := url.Parse(originURL)
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
	return _url, nil
}
