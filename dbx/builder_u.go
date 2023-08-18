package dbx

import (
	"fmt"
	"strings"
)

// BuildUpdate 更新数据
func BuildUpdate(tableName string, setData map[string]interface{}) (string, []any) {
	ph, params := BuildWhereParams(setData)
	query := fmt.Sprintf("UPDATE %s SET %s", SpecialField(tableName), ph)
	return query, params
}

// UWhere 带 where 的更新
func BuildUpdateWhere(tableName string, setData map[string]interface{}, where map[string]interface{}) (string, []any) {
	ph, params := BuildKeyVal(setData, ",")
	query := &strings.Builder{}
	query.WriteString(fmt.Sprintf("UPDATE %s SET %s ", SpecialField(tableName), ph))
	if len(where) > 0 {
		wph, wparams := BuildWhereParams(where)
		query.WriteString("WHERE ")
		query.WriteString(wph)
		params = append(params, wparams...)
	}
	return query.String(), params

}
