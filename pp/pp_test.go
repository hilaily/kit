package pp

import "testing"

var (
	a = map[string]interface{}{
		"a": "adfafdasf",
		"b": 1231,
	}
)

func TestDump(t *testing.T) {
	Dump("test1: ", a)
}

func TestIDump(t *testing.T) {
	IDump("test2: ", a)
}
