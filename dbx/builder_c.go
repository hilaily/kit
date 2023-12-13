package dbx

import (
	"fmt"
	"strings"
)

type IModel interface {
	Fields() ([]string, []interface{})
}

func BuildCreate(tableName string, data map[string]interface{}) (string, []any) {
	cols, ph, params := BuildInsertParams(data)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", SpecialField(tableName), cols, ph)
	return query, params
}

// GetBulkInsertSQL gen bulk insert sql
// NOTE: items is a slice, and the element must implement IModel
// INSERT INTO user (uid, name, money) VALUES (77, "name1", 77), (88, "name2", 88);
func BuildBulkCreate[T IModel](tableName string, data []T) (string, []interface{}) {
	if len(data) == 0 {
		return "", nil
	}
	fields, _ := data[0].Fields()
	cols := strings.Builder{}
	onePH := strings.Builder{}
	for k, v := range fields {
		cols.WriteByte('`')
		cols.WriteString(v)
		cols.WriteByte('`')
		onePH.WriteByte('?')
		if k < len(fields)-1 {
			cols.WriteByte(',')
			onePH.WriteByte(',')
		}
	}
	placeHolder := make([]string, 0, len(data))
	vals := make([]interface{}, 0, len(data))
	for k := range data {
		placeHolder = append(placeHolder, fmt.Sprintf("(%s)", onePH.String()))
		_, dataVal := data[k].Fields()
		vals = append(vals, dataVal...)
	}
	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", tableName, cols.String(), strings.Join(placeHolder, ","))
	return query, vals
}

// GetBulkInsertSQLOnDuplicate gen bulk insert sql on duplicate
// NOTE: items is a slice, and the element must implement IModel
// SQL 示例
// INSERT INTO user (uid, name, money) VALUES (77, "name1", 77), (88, "name2", 88) ON DUPLICATE KEY UPDATE money=money, `name`=VALUES(`name`);
func BuildBulkCreateSQLOnDuplicate[T IModel](tableName string, data []T, notUpdateColumn, updateColumn []string) (string, []interface{}) {
	if len(notUpdateColumn) == 0 && len(updateColumn) == 0 {
		return BuildBulkCreate(tableName, data)
	}

	fields, _ := data[0].Fields()
	cols := strings.Builder{}
	onePH := strings.Builder{}
	for k, v := range fields {
		cols.WriteByte('`')
		cols.WriteString(v)
		cols.WriteByte('`')
		onePH.WriteByte('?')
		if k < len(fields)-1 {
			cols.WriteByte(',')
			onePH.WriteByte(',')
		}
	}
	placeHolder := make([]string, 0, len(data))
	vals := make([]interface{}, 0, len(data))
	for k := range data {
		placeHolder = append(placeHolder, fmt.Sprintf("(%s)", onePH.String()))
		_, dataVal := data[k].Fields()
		vals = append(vals, dataVal...)
	}
	var updateQuery []string
	for _, v := range notUpdateColumn {
		col := SpecialField(v)
		updateQuery = append(updateQuery, col+"="+col)
	}
	for _, v := range updateColumn {
		col := SpecialField(v)
		updateQuery = append(updateQuery, col+"=VALUES("+col+")")
	}
	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s ON DUPLICATE KEY UPDATE %s",
		tableName, cols.String(), strings.Join(placeHolder, ","), strings.Join(updateQuery, ","))
	return query, vals
}

func BuildCreateOnduplicateWithAdd(tablename string, insertData map[string]interface{}, addData map[string]interface{}) (string, []any) {
	cols, ph, params := BuildInsertParams(insertData)
	query := strings.Builder{}
	query.WriteString(fmt.Sprintf("INSERT %s (%s) VALUES(%s)", SpecialField(tablename), cols, ph))
	if len(addData) > 0 {
		query.WriteString(" ON DUPLICATE KEY UPDATE ")
		keys := make([]string, 0, len(addData))
		for k, v := range addData {
			// set k = k + ?
			keys = append(keys, fmt.Sprintf(" set %s = %s + ? ", k, k))
			params = append(params, v)
		}
		query.WriteString(strings.Join(keys, ","))
	}
	qStr := query.String()
	return qStr, params
}

// MakeOnDuplicateSQL ...
// Return string for example: INSERT INTO `tablename` (`f2`,`f3`,`f1`) VALUES (?,?,?) ON DUPLICATE KEY UPDATE `f2` = ?,`f3` = ?
func BuildOnDuplicateSQL(table string, insert map[string]interface{}, update map[string]interface{}) (string, []interface{}) {
	b := &strings.Builder{}
	insertSQL, insertParams := BuildInsertSQL(table, insert)
	b.WriteString(insertSQL)
	b.WriteString(" ON DUPLICATE KEY UPDATE ")
	whereSQL, whereParams := BuildKeyVal(update, ",")
	b.WriteString(whereSQL)
	return b.String(), append(insertParams, whereParams...)
}

// BuildInsertSQL
// Return string for example: INSERT INTO `tablename` (`f1`,`f2`,`f3`) VALUES (?,?,?)
// Return []interface{} for example: [v1 v2 v3]
func BuildInsertSQL(table string, data map[string]interface{}) (string, []interface{}) {
	b := strings.Builder{}
	cols, ph, params := BuildInsertParams(data)
	b.WriteString("INSERT INTO ")
	b.WriteString(SpecialField(table))
	b.WriteString(" (")
	b.WriteString(cols)
	b.WriteString(") ")
	b.WriteString("VALUES (")
	b.WriteString(ph)
	b.WriteString(")")
	return b.String(), params
}

// BuildInsertParams 组装 sql 参数
// Return string for example: `f2`,`f3`,`f1`
// Return string for example: ?,?,?
// Return []interface{} for example: [v2 v3 v1]
func BuildInsertParams(data map[string]interface{}) (string, string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", "", make([]interface{}, 0)
	}
	cols := make([]string, 0, length)
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		k := strings.Replace(k, "`", "", -1)
		cols = append(cols, fmt.Sprintf("`%s`", k))
		placeholder = append(placeholder, "?")
		params = append(params, v)
	}
	return strings.Join(cols, ","), strings.Join(placeholder, ","), params
}
