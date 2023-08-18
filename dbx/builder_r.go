package dbx

import (
	"reflect"
	"strings"

	"github.com/hilaily/kit/stringx"
)

// BuildWhereParams 构造 where 部分 sql 语句
// Return string for example: `f1` = ? AND `f2` = ? AND `f3` = ?
func BuildWhereParams(data map[string]interface{}) (string, []interface{}) {
	return BuildKeyVal(data, " AND ")
}

// BuildKeyVal
// Return 1 string: `c1` = ? <connector> `c2` = ?
func BuildKeyVal(data map[string]interface{}, connector string) (string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", make([]interface{}, 0)
	}
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		k := strings.Replace(k, "`", "", -1)
		placeholder = append(placeholder, "`"+k+"` = ?")
		params = append(params, v)
	}
	return strings.Join(placeholder, connector), params
}

// MakeColsParams 拼接字段部分 sql 语句
// Exaple: c1, c2, c3
func BuildColsParams(cols []string) string {
	c := make([]string, 0, len(cols))
	for _, v := range cols {
		c = append(c, SpecialField(v))
	}
	return strings.Join(c, ",")
}

// SpecialField 处理表名，列名的反引号
func SpecialField(s string) string {
	if strings.Contains(s, "`") {
		return s
	}
	return "`" + s + "`"
}

func GenColumnsByStructTag(val interface{}, tag ...string) []string {
	_tag := "db"
	if len(tag) > 0 {
		_tag = tag[0]
	}
	t := reflect.TypeOf(val)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("val is not a struct")
	}
	fieldNum := t.NumField()
	arr := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		// get val by tag
		tagVal := t.Field(i).Tag.Get(_tag)
		if tagVal == "-" {
			continue
		}
		if tagVal == "" {
			tagVal = stringx.Camel2Case(t.Field(i).Name)
		}
		arr = append(arr, SpecialField(tagVal))
	}
	return arr
}
