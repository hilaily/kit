package strings

import "unsafe"

// StringToBytes
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringPointer
func StringPointer(input string) *string {
	return &input
}
