package dbx

import "fmt"

// BuildDelete 删除语句
func BuildDelete(tableName string, where map[string]interface{}) (string, []any) {
	ph, params := BuildWhereParams(where)
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", SpecialField(tableName), ph)
	return query, params
}
