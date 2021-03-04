package httpx

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

// WrapHTTPProxy add a proxy for a http client
func WrapHTTPProxy(client *http.Client, proxyURL string, insecure bool) (*http.Client, error) {
	proxyURLIns, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("parse proxyURL: %s, %w", err)
	}

	transport := &http.Transport{}
	transport.Proxy = http.ProxyURL(proxyURLIns)
	if insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client.Transport = transport
	return client, nil
}
