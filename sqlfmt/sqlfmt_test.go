package sqlfmt

import (
	"testing"
)

var fileSqlFmtTest = "sqlfmt_test.go"

type Select struct {
	Key   string
	Table string
}

func (s *Select) ToSqlAndArgs() (string, []any, error) {
	return "select " + s.Key + " from " + s.Table, nil, nil
}

func TestSqlFmt(t *testing.T) {
	testFuncName := "TestSqlFmt"
	slt := Select{}
	slt.Key = "name"
	slt.Table = "user"
	sqlStr, _, _ := slt.ToSqlAndArgs()
	if sqlStr != "select name from user" {
		t.Fatalf("%s-%s: slt.ToSqlAndArgs error", fileSqlFmtTest, testFuncName)
	}
}

func FuncToSqlAndArgs() (string, []any, error) {
	return "select pass from user", nil, nil
}

func TestSqlFormat(t *testing.T) {
	testFuncName := "TestSqlFormat"
	fsf := ToSqlFormat(FuncToSqlAndArgs)
	sqlStr, _, _ := fsf.ToSqlAndArgs()
	if sqlStr != "select pass from user" {
		t.Fatalf("%s-%s: fsf.ToSqlAndArgs", fileSqlFmtTest, testFuncName)
	}
}
