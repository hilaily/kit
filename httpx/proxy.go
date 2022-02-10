package httpx

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hilaily/kit/stringx"
)

// WrapHTTPProxy add a proxy for a http client
func WrapHTTPProxy(client *http.Client, proxyURL string, insecure bool) (*http.Client, error) {
	if proxyURL == "" {
		return client, nil
	}
	proxyURLIns, err := url.Parse(stringx.AddURLSchema(proxyURL, "http"))
	if err != nil {
		return nil, fmt.Errorf("parse proxyURL: %s, %w", proxyURL, err)
	}

	transport := &http.Transport{}
	transport.Proxy = http.ProxyURL(proxyURLIns)
	if insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client.Transport = transport
	return client, nil
}
