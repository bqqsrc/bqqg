package databasehelper

import (
	//	"bqqgc/errer"
	// "database/sql"
	// "io/ioutil"
	// "os"
	// "strings"
	"github.com/bqqsrc/bqqg/log"
	"github.com/bqqsrc/bqqg/sqlfmt"

	// "log"
	"github.com/bqqsrc/bqqg/database"
)

func Deleter2Result(d *sqlfmt.Delete) (int64, error) {
	funcName := "Deleter2Result"
	sqlStr, args, err := d.ToSqlAndArgs()
	if err != nil {
		return -1, err
	}
	if sqlStr == "" {
		log.Errorf("Error, %s, sql is empty", funcName)
		return -1, nil
	}
	log.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	if ret, err := database.Exec(defaultRegister, sqlStr, args...); err != nil {
		return -1, err
	} else {
		return ret.LastInsertId()
	}
}
