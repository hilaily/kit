package stringx

import "testing"

func TestGenRankStr(t *testing.T) {
	for i := 0; i <= 10; i++ {
		res := GenRankStr(10)
		t.Log(res)
	}
}

func TestGenRankStrLowercase(t *testing.T) {
	for i := 0; i <= 10; i++ {
		res := GenRankStrLowercase(10)
		t.Log(res)
	}
}
