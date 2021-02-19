package strings

import (
	"path"
	"regexp"
	stdstrings "strings"
	"unsafe"
)

var (
	compresReg *regexp.Regexp
)

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

// Compress for trim space in string
func Compress(str string) string {
	if str == "" {
		return str
	}
	//匹配一个或多个空白符的正则表达式
	if compresReg == nil {
		compresReg = regexp.MustCompile("\\s+")
	}
	return compresReg.ReplaceAllString(str, "")
}

// URLJoin like strings.Join but for url path
func URLJoin(base string, paths ...string) string {
	p := path.Join(paths...)
	return stdstrings.TrimRight(base, "/") + "/" + stdstrings.TrimLeft(p, "/")
}
