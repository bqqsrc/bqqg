package databasehelper

import (
	//	"bqqgc/errer"
	// "database/sql"
	// "io/ioutil"
	// "os"
	// "strings"
	"github.com/bqqsrc/bqqg/loger"
	"github.com/bqqsrc/bqqg/sqlfmt"

	// "log"
	"github.com/bqqsrc/bqqg/database"
)

func Updater2Result(i *sqlfmt.Updater) (int64, error) {
	funcName := "Updater2Result"
	sqlStr, args, err := i.ToSqlAndArgs()
	if err != nil {
		return -1, err
	}
	if sqlStr == "" {
		loger.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	if ret, err := database.Exec(defaultRegister, sqlStr, args...); err != nil {
		return -1, err
	} else {
		return ret.LastInsertId()
	}
}
