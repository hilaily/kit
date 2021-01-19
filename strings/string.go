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

// Pointer return a address of a string
func Pointer(input string) *string {
	return &input
}

// Has represent if a element is in the list
func Has(list []string, element string) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}
