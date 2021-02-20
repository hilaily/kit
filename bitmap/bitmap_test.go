package bitmap

import "testing"

func TestMask(t *testing.T) {
	m := GetMask()
	t.Log("---", m)
	t.Log(BinaryString(m[0]))
	m[0] = 12
	t.Log(BinaryString(m[0]))
	m1 := GetMask()
	t.Log(BinaryString(m1[0]))
}

func TestSet1(t *testing.T) {
	val := uint64(0)
	val = Set(val, 7)
	t.Log(BinaryString(val))
}

func TestUnset1(t *testing.T) {
	val := uint64(115)
	t.Log(BinaryString(val))
	val = Unset(val, 4)
	t.Log(BinaryString(val))
}

func TestBin(t *testing.T) {
	arr := []uint64{2186509048025116382, 9223372036584016704, 9223372036854644735, 9223372036854775807, 4244635647}
	for _, v := range arr {
		t.Log(BinaryString(v))
	}
}
