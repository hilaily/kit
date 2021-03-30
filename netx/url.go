package netx

import "net/url"

// URLQueryParams add query params for url
func URLQueryParams(_url string, params map[string]string) (string, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return "", err
	}
	p := u.Query()
	for k, v := range params {
		p.Add(k, v)
	}
	u.RawQuery = p.Encode()
	return u.String(), nil
}
