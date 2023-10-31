// sqlfmt定义了一系列sql语句拼接的工具
package sqlfmt

// 转换至sql和参数的接口类型
type SqlFmt interface {

	// 转换到sql字符串语句和参数
	// 当error为nil时表示转换成功，error不为nil时转换失败
	ToSqlAndArgs() (string, []any, error)
}

// SqlFormat是对函数类型func() (string, []any, error)的再定义
// 它实现了ToSqlAndArgs，所以一个SqlFormat指针同时也是一个SqlFmt
type SqlFormat func() (string, []any, error)

// 返回值：直接返回s本身的返回值
func (s SqlFormat) ToSqlAndArgs() (string, []any, error) {
	return s()
}

// 转换一个func() (string, []any, error)为SqlFormat
// 参数：f可以是任意func() (string, []any, error)的函数
// 返回值：直接返回f本身
func ToSqlFormat(f func() (string, []any, error)) SqlFormat {
	return f
}
