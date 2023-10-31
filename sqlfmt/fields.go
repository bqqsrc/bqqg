package sqlfmt

import (
	"strings"
)

// 字段组表示一组字段，它包含了一个SqlFmt数组
//
// # ToSqlAndArgs会转换为一个字段列表，中间以英文逗号隔开，如：field1, field2, fields
//
// 可以用于sql的select字段列表、表格列表的格式化
type Fields struct {
	FieldList []SqlFmt
}

// 重置整个Fields
func (fs *Fields) Reset() {
	fs.FieldList = nil
}

// 添加fields
func (fs *Fields) AddSqlFmt(fields ...SqlFmt) {
	if fs.FieldList == nil {
		fs.FieldList = fields
	} else {
		fs.FieldList = append(fs.FieldList, fields...)
	}
}

// 添加string类型的Fields
func (fs *Fields) AddFields(fields ...string) {
	if fs.FieldList == nil {
		fs.FieldList = make([]SqlFmt, 0)
	}
	for i, _ := range fields {
		var f Field = Field(fields[i])
		fs.FieldList = append(fs.FieldList, &f)
	}
}

// 转换为sql语句和参数
func (fs *Fields) ToSqlAndArgs() (string, []any, error) {
	fieldsName := "Fields"
	if fs.FieldList == nil {
		return "", nil, SqlFmtError(fieldsName, "FieldList is nil")
	}
	count := len(fs.FieldList)
	if count == 0 {
		return "", nil, SqlFmtError(fieldsName, "FieldList not has element")
	}
	hasFirst := false
	var sqlStr string
	var sqlArgs []any
	var finalSqlArgs []any
	var err error
	var filedsBuild strings.Builder
	for i := 0; i < count; i++ {
		sqlStr, sqlArgs, err = fs.FieldList[i].ToSqlAndArgs()
		if err != nil {
			return "", nil, err
		}
		if hasFirst {
			filedsBuild.WriteString(",")
		} else {
			hasFirst = true
		}
		filedsBuild.WriteString(sqlStr)
		if sqlArgs != nil {
			if finalSqlArgs == nil {
				finalSqlArgs = make([]any, 0)
			}
			finalSqlArgs = append(finalSqlArgs, sqlArgs...)
		}
	}
	return filedsBuild.String(), finalSqlArgs, nil
}
