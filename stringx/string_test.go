package stringx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	res := Format("this is a {{.name}}, {{.age}} years old", map[string]any{"name": "Bob", "age": 10})
	assert.Equal(t, "this is a Bob, 10 years old", res)
}

func TestTrimField(t *testing.T) {
	type B string
	type A struct {
		Name string
		C    int
		like string
		Addr string
		DDD  *string
		EEE  B
	}
	a := A{
		Name: "aaa ",
		like: " ccc ",
		Addr: " bbb ",
		C:    1,
		EEE:  "eee ",
	}
	TrimField(&a)
	assert.Equal(t, "aaa", a.Name)
	assert.Equal(t, "bbb", a.Addr)
	assert.Equal(t, " ccc ", a.like)
	assert.Equal(t, B("eee"), a.EEE)
	assert.Equal(t, 1, a.C)

	var b A
	TrimField(&b)

	m := map[string]int{}
	TrimField(&m)
}

func TestCase(t *testing.T) {
	str1 := "hello_world"
	str2 := "HelloWorld"
	assert.Equal(t, str2, Case2Camel(str1))
	assert.Equal(t, str1, Camel2Case(str2))
}

func TestZHToUnit(t *testing.T) {
	raw := "测试中文abc"
	res := ZHToUnicode(raw)
	t.Log(res)
	r, err := UnicodeToZH(res)
	assert.NoError(t, err)
	t.Log(r)
	assert.Equal(t, r, raw)
}

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
