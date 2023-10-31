package sqlfmt

import (
	"fmt"
)

type sqlFmtError struct {
	sqlfmtName string
	msgFormat  string // 错误的format字符串
	args       []any  // 错误的字符串
}

func (e *sqlFmtError) Error() string {
	fmtStr := fmt.Sprintf("sqlfmt error: %s err: %s", e.sqlfmtName, e.msgFormat)
	return fmt.Sprintf(fmtStr, e.args...)
}

// sqlfmt转换为sql字符串语句的错误类型
// 参数： name 错误名称
//
//	format 错误信息的format语句
//	args 错误参数
func SqlFmtError(name, format string, args ...any) error {
	return &sqlFmtError{name, format, args}
}
