package sqlfmt

import "fmt"

// string类型的再定义
//
// 可以表示任何字段，sql语句等，包括键名字段、表名字段等
//
// ToSqlAndArgs会转换为字符串值本身
type Field string

func (f *Field) ToSqlAndArgs() (string, []any, error) {
	if *f == "" {
		return "", nil, SqlFmtError("Fields", "Field is empty")
	}
	return string(*f), nil, nil
}

// 指定表名的字段
//
// Table可以为空，Key不能为空
//
// 最后返回字符串格式如：table1.key1
type TableField struct {
	Table string
	Key   string
}

func (f *TableField) ToSqlAndArgs() (string, []any, error) {
	tableFieldName := "TableField"
	if f == nil {
		return "", nil, SqlFmtError(tableFieldName, "TableField is empty")
	}
	if f.Key == "" {
		return "", nil, SqlFmtError(tableFieldName, "TableField has no a Key")
	}
	if f.Table == "" {
		return f.Key, nil, nil
	}
	return fmt.Sprintf("%s.%s", f.Table, f.Key), nil, nil
}
