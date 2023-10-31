package sqlfmt

import (
	"testing"
)

var fileFieldTest = "field_test.go"

func TestField(t *testing.T) {
	testFuncName := "TestField"
	var f Field = "select age from user"
	sqlStr, _, _ := f.ToSqlAndArgs()
	if sqlStr != "select age from user" {
		t.Fatalf("%s-%s: f.ToSqlAndArgs error", fileFieldTest, testFuncName)
	}
}

// 下面的两个测试样例说明再定义的转换不会很耗时
func Benchmark_Field(b *testing.B) {
	var f Field
	for i := 0; i < b.N; i++ {
		f = "select age from user"
	}
	f.ToSqlAndArgs()
}

func Benchmark_Field2(b *testing.B) {
	var f Field
	str := "select age from user"
	for i := 0; i < b.N; i++ {
		f = Field(str)
	}
	f.ToSqlAndArgs()
}
