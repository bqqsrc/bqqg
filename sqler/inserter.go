package sqler

import (
	"github.com/bqqsrc/dber"
	"errors"
	"strings"

	"github.com/bqqsrc/loger"
)

type Inserter struct {
	keyValue map[string]interface{}
	table    string
}

func (i *Inserter) Reset() {
	i.keyValue = nil
	i.table = ""
}

func (i *Inserter) SetKeyValues(keyValue map[string]interface{}) {
	i.keyValue = keyValue
}

func (i *Inserter) AddKeyValue(key string, value interface{}) {
	if i.keyValue == nil {
		i.keyValue = make(map[string]interface{})
	}
	if _, ok := i.keyValue[key]; ok {
		loger.Errorf("Inserter redeclare key: %s", key)
		return
	}
	i.keyValue[key] = value
}

func (i *Inserter) SetTable(table string) {
	i.table = table
}

func (i *Inserter) ToSqlAndArgs() (string, []interface{}) {
	if i.table == "" {
		return "", nil
	}
	keyValue := i.keyValue
	if keyValue == nil || len(keyValue) == 0 {
		return "", nil
	}
	var build strings.Builder
	build.WriteString("insert into ")
	build.WriteString(i.table)
	build.WriteString("(")
	var tmpBuild strings.Builder
	index := 0
	args := make([]interface{}, 0)
	for key, value := range keyValue {
		if index > 0 {
			build.WriteString(",")
			tmpBuild.WriteString(",")
		} else {
			index++
		}
		build.WriteString(key)
		if valueSqlAndArgs, ok := value.(SqlAndArgs); ok {
			sql, tmpArgs := valueSqlAndArgs.ToSqlAndArgs()
			if sql != "" {
				tmpBuild.WriteString("(")
				tmpBuild.WriteString(sql)
				tmpBuild.WriteString(")")
				args = append(args, tmpArgs...)
			}
		} else {
			tmpBuild.WriteString("?")
			args = append(args, value)
		}
	}
	build.WriteString(") values (")
	build.WriteString(tmpBuild.String())
	build.WriteString(")")
	return build.String(), args
}

func (i *Inserter) ExecSql(controller string) (int64, error) {
	funcName := "Inserter.ExecSql"
	sql, args := i.ToSqlAndArgs()
	if sql == "" {
		loger.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sql, args)
	if ret, err := dber.Exec(controller, sql, args...); err != nil {
		return -1, err
	} else {
		return ret.LastInsertId()
	}
}

func (i *Inserter) ExecSqlTx(controller, name string, commit bool) error {
	funcName := "Inserter.ExecSqlTx"
	sql, args := i.ToSqlAndArgs()
	if sql == "" {
		loger.Errorf("Error, %s, sql is empty", funcName)
		return errors.New("sql is empty")
	}
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sql, args)
	if _, err := dber.ExecTxSql(controller, name, sql, args...); err != nil {
		return err
	} else {
		if commit {
			return dber.CommitTx(controller, name)
		}
		return nil
	}
}

func GetInserter(table string) *Inserter {
	return &Inserter{table: table}
}
