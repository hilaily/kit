package stringx

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unsafe"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	compresReg *regexp.Regexp
)

// Camel2Case ...
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Case2Camel ...
func Case2Camel(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	caser := cases.Title(language.English)
	name = caser.String(name)
	return strings.ReplaceAll(name, " ", "")
}

// Ucfirst ...
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Lcfirst
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("out of memory")
		}
	}()
	b.WriteString(s)
	return b
}

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
		compresReg = regexp.MustCompile(`\s+`)
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
		return strings.TrimSuffix(scheme, "://") + "://" + url
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

// ZHToUnicode turn Chinese into unicode
func ZHToUnicode(raw string) string {
	textQuoted := strconv.QuoteToASCII(raw)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}

// UnicodeToZH turn unitcode into Chinese
func UnicodeToZH(raw string) (string, error) {
	return strconv.Unquote(strings.ReplaceAll(strconv.Quote(raw), `\\u`, `\u`))
}

// TrimField trim struct string field
func TrimField(ptr any) {
	rv := reflect.ValueOf(ptr).Elem()
	num := rv.NumField()
	for i := 0; i < num; i++ {
		val := rv.Field(i)
		if val.Kind() == reflect.String && val.CanSet() {
			oldValue := val.String()
			val.SetString(strings.TrimSpace(oldValue))
		}
	}
}
