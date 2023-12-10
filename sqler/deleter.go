package sqler

import (
	"errors"
	"github.com/bqqsrc/bqqg/database"
	"strings"

	"github.com/bqqsrc/bqqg/log"
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
		log.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	if ret, err := database.Exec(controller, sql, args...); err != nil {
		return -1, err
	} else {
		return ret.RowsAffected()
	}
}

func (d *Deleter) ExecSqlTx(controller, name string, commit bool) error {
	funcName := "Inserter.ExecSqlTx"
	sql, args := d.ToSqlAndArgs()
	if sql == "" {
		log.Errorf("Error, %s, sql is empty", funcName)
		return errors.New("sql is empty")
	}
	log.Debugf("%s, sql: %s\nargs: %v\n", funcName, sql, args)
	if _, err := database.ExecTxSql(controller, name, sql, args...); err != nil {
		return err
	} else {
		if commit {
			return database.CommitTx(controller, name)
		}
		return nil
	}
}

func GetDeleter(table string) *Deleter {
	return &Deleter{table: table}
}
