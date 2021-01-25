package strings

import "testing"

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
