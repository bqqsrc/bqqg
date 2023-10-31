package sqlfmt

import (
	"testing"
)

var fileFieldsTest = "fields_test.go"

func TestFields(t *testing.T) {
	testFuncName := "TestFields"
	fs := Fields{}
	fs.AddFields("name", "password", "age")
	sqlStr, _, _ := fs.ToSqlAndArgs()
	expect := "name,password,age"
	if sqlStr != expect {
		t.Fatalf("%s-%s: fs.ToSqlAndArgss error, expected: %s, got: %s", fileFieldsTest, testFuncName, expect, sqlStr)
	}
	fs.Reset()
	var passField Field = "password"
	var nameField Field = "name"
	var ageField Field = "age"
	fs.AddSqlFmt(&passField, &nameField, &ageField)
	sqlStr, _, _ = fs.ToSqlAndArgs()
	expect = "password,name,age"
	if sqlStr != expect {
		t.Fatalf("%s-%s: fs.ToSqlAndArgss error, expected: %s, got: %s", fileFieldsTest, testFuncName, expect, sqlStr)
	}
}

// TODO 下面这两个测试样例，说明了Fields的AddFields是很耗时的
func Benchmark_AddFields(b *testing.B) {
	fs := Fields{}
	for i := 0; i < b.N; i++ {
		fs.AddFields("name", "password", "age") //替换为待测试性能的代码块
	}
}

type MyFields struct {
	FieldList []string
}

// 添加string类型的Fields
func (fs *MyFields) AddFields(fields ...string) {
	if fs.FieldList == nil {
		fs.FieldList = fields
	}
	fs.FieldList = append(fs.FieldList, fields...)
}

func Benchmark_MyFields(b *testing.B) {
	mfs := MyFields{}
	for i := 0; i < b.N; i++ {
		mfs.AddFields("name", "password", "age") //替换为待测试性能的代码块
	}
}
