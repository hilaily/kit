package stringx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	arr := []string{"1", "2", "3"}
	m := ToMap(arr)
	assert.Equal(t, len(m), len(arr))
}

func TestCompress(t *testing.T) {
	data := []struct {
		data   string
		expect string
	}{
		{data: "a b c", expect: "abc"},
		{data: `a 
b c`, expect: "abc"},
		{data: "a	b", expect: "ab"},
	}
	for _, v := range data {
		res := Compress(v.data)
		if res != v.expect {
			t.Errorf("res: %s, expect: %s", res, v.expect)
		}
	}
}

func TestURLJoin(t *testing.T) {
	data := [][]string{
		{"http://a.com/", "/b", "/c"},
		{"http://a.com/", "b/", "/c"},
		{"http://a.com", "/b", "c"},
		{"http://a.com", "b", "/c"},
	}
	expect := "http://a.com/b/c"
	for _, v := range data {
		e := URLJoin(v[0], v[1:]...)
		if e != expect {
			t.Errorf("err: res %s, %v, %s\n", e, v[0], v[0:])
		}
	}
}

func TestAddURLSchema(t *testing.T) {
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("http://a/b", "http"))
	assert.Equal(t, "http://a/b", AddURLSchema("a/b", "http://"))
	assert.Equal(t, "http://a/b", AddURLSchema("a/b", "http"))
	assert.Equal(t, "https://a/b", AddURLSchema("a/b", "https"))
}

func TestDeDup(t *testing.T) {
	data := []struct {
		expect []string
		origin []string
	}{
		{[]string{"1", "2", "3"}, []string{"1", "2", "3", "2", "1"}},
		{[]string{"1", "2", "3"}, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3"}, []string{"1", "1", "2", "2", "3"}},
		{[]string{}, []string{}},
	}
	for _, v := range data {
		res := Dedup(v.origin)
		assert.Equal(t, v.expect, res)
	}
}

func TestEncodePEM(t *testing.T) {
	pem := []byte(`
-----BEGIN CERTIFICATE-----
MIIB1jCCAYCgAwIBAgIJAOuA44qosxpmMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTkwNTEwMTAwNTAwWhcNMjAwNTA5MTAwNTAwWjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKRZ
LqWvnwROzCw6Td/YG/EHaZHnlMkisS+ybk1YCXZyq2Tv8z/ikPMRmWp/WtU1knxC
1VcPCiVIifQos9UNu1kCAwEAAaNTMFEwHQYDVR0OBBYEFDc9jbjXcS6oHuVuXsB9
ZGWr39/FMB8GA1UdIwQYMBaAFDc9jbjXcS6oHuVuXsB9ZGWr39/FMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADQQBFHviM7xFJ3KUg4SPlBm1X2yyAXgQI
oNjuH9WSGQhVm9PxmEZJZnoEpmqYc+tamytRMxLHbermRRaIuMHzQj/J
-----END CERTIFICATE-----`)
	res, err := EncodePEM([]byte(pem))
	assert.NoError(t, err)
	fmt.Printf("%q", res)
}
