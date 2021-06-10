package stringx

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
	"unsafe"
)

var (
	compresReg *regexp.Regexp
)

// StringToBytes ...
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString ...
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
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(p, "/")
}

// AddURLSchema if url do not have scheme then add one
// sheme should have ://
func AddURLSchema(url string, scheme string) string {
	if !strings.HasPrefix(url, "http") {
		return strings.TrimRight(scheme, "://") + "://" + url
	}
	return url
}

// Dedup remove duplicate element in slice
func Dedup(s []string) []string {
	l := len(s)
	uniq := make(map[string]struct{}, l)
	newArr := make([]string, 0, l)
	for _, v := range s {
		_, ok := uniq[v]
		if ok {
			continue
		} else {
			newArr = append(newArr, v)
			uniq[v] = struct{}{}
		}
	}
	return newArr
}

// EncodePEM encode a pem byte array to a string
func EncodePEM(certPEM []byte) (string, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return "", errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate, %w", err)
	}
	publicKeyDer, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return "", fmt.Errorf("marshal x509, %w", err)
	}
	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	return string(pem.EncodeToMemory(&publicKeyBlock)), nil
}

// ToMap transfer a string slice to a map
func ToMap(data []string) map[string]struct{} {
	m := make(map[string]struct{}, len(data))
	for _, v := range data {
		m[v] = struct{}{}
	}
	return m
}
