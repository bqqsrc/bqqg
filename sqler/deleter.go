package sqler

import (
	"github.com/bqqsrc/dber"
	"errors"
	"strings"

	"github.com/bqqsrc/loger"
)

type Deleter struct {
	table         string
	conditionList *ConditionBatch
}

func (d *Deleter) Reset() {
	d.table = ""
	d.conditionList = nil
}

func (d *Deleter) SetTable(table string) {
	d.table = table
}

func (d *Deleter) SetConditions(conditions *ConditionBatch) {
	d.conditionList = conditions
}

func (d *Deleter) AddConditions(conditions ...SqlAndArgs) {
	if d.conditionList == nil {
		d.conditionList = GetConditionBatch()
	}
	d.conditionList.AddConditions(conditions...)
}

func (d *Deleter) ToSqlAndArgs() (string, []interface{}) {
	table := d.table
	if table == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString("delete from ")
	build.WriteString(table)
	if d.conditionList != nil {
		sql, args := d.conditionList.toWhere()
		build.WriteString(sql)
		return build.String(), args
	}
	return build.String(), nil
}

func (d *Deleter) ExecSql(controller string) (int64, error) {
	funcName := "Updater.ExecSql"
	sql, args := d.ToSqlAndArgs()
	if sql == "" {
		loger.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	if ret, err := dber.Exec(controller, sql, args...); err != nil {
		return -1, err
	} else {
		return ret.RowsAffected()
	}
}

func (d *Deleter) ExecSqlTx(controller, name string, commit bool) error {
	funcName := "Inserter.ExecSqlTx"
	sql, args := d.ToSqlAndArgs()
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

func GetDeleter(table string) *Deleter {
	return &Deleter{table: table}
}
