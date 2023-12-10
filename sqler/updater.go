package sqler

import (
	"github.com/bqqsrc/dber"
	"errors"
	"strings"

	"github.com/bqqsrc/loger"
)

type Updater struct {
	keyValue      map[string]interface{}
	table         string
	conditionList *ConditionBatch
}

func (u *Updater) Reset() {
	u.table = ""
	u.conditionList = nil
	u.keyValue = nil
}

func (u *Updater) SetTable(table string) {
	u.table = table
}

func (u *Updater) AddConditions(conditions ...SqlAndArgs) {
	if u.conditionList == nil {
		u.conditionList = GetConditionBatch()
	}
	u.conditionList.AddConditions(conditions...)
}

func (u *Updater) SetConditions(conditions *ConditionBatch) {
	u.conditionList = conditions
}

func (u *Updater) SetKeyValues(keyValue map[string]interface{}) {
	u.keyValue = keyValue
}

func (u *Updater) AddKeyValue(key string, value interface{}) {
	if u.keyValue == nil {
		u.keyValue = make(map[string]interface{})
	}
	if _, ok := u.keyValue[key]; ok {
		loger.Errorf("Inserter redeclare key: %s", key)
		return
	}
	u.keyValue[key] = value
}

func (u *Updater) ToSqlAndArgs() (string, []interface{}) {
	table := u.table
	if table == "" {
		return "", nil
	}
	keyValue := u.keyValue
	if keyValue == nil || len(keyValue) == 0 {
		return "", nil
	}
	var build strings.Builder
	build.WriteString("update ")
	build.WriteString(table)
	build.WriteString(" set ")
	index := 0
	args := make([]interface{}, 0)
	for key, value := range keyValue {
		if index > 0 {
			build.WriteString(",")
		} else {
			index++
		}
		build.WriteString(key)
		build.WriteString("=?")
		args = append(args, value)
	}
	if u.conditionList != nil {
		where, arg := u.conditionList.toWhere()
		build.WriteString(where)
		args = append(args, arg...)
	}
	return build.String(), args
}

func GetUpdater(table string) *Updater {
	return &Updater{table: table}
}

func (u *Updater) ExecSql(controller string) (int64, error) {
	funcName := "Updater.ExecSql"
	sql, args := u.ToSqlAndArgs()
	if sql == "" {
		loger.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sql, args)
	if ret, err := dber.Exec(controller, sql, args...); err != nil {
		return -1, err
	} else {
		return ret.RowsAffected()
	}
}

func (u *Updater) ExecSqlTx(controller, name string, commit bool) error {
	funcName := "Inserter.ExecSqlTx"
	sql, args := u.ToSqlAndArgs()
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
