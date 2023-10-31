package databasehelper

import (
	//	"bqqgc/errer"
	"database/sql"
	// "io/ioutil"
	// "os"
	// "strings"
	"github.com/bqqsrc/bqqg/database"
	"github.com/bqqsrc/bqqg/sqlfmt"
	// "log"
)

func Selector2Int(s *sqlfmt.Selector) (int, error) {
	// funcName := "Selecter.ToInt"
	// loger.Debugf("%s, controller: %s\ns: %v\n", funcName, defaultRegister, s)
	sqlStr, args, err := s.ToSqlAndArgs()
	if err != nil {
		return -1, err
	}
	// loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	var rows *sql.Rows
	// if txName != "" {
	// 	rows, err = database.QueryTxSql(controller, txName, sqlStr, args...)
	// } else {
	rows, err = database.Query(defaultRegister, sqlStr, args...)
	// }
	if err != nil {
		return 0, err
	}
	var num int
	for rows.Next() {
		rows.Scan(&num)
	}
	// loger.Debugf("%s, num: %d\n", funcName, num)
	return num, nil
}

func Selector2MapList(s *sqlfmt.Selector) ([]map[string]any, error) {
	// funcName := "Selecter.ToMapList"
	// loger.Debugf("%s, controller: %s\ns: %v\n", funcName, defaultRegister, s)
	sqlStr, args, err := s.ToSqlAndArgs()
	if err != nil {
		return nil, err
	}
	// loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	var row *sql.Rows
	// if txName != "" {
	// 	row, err = database.QueryTxSql(controller, txName, sqlStr, args...)
	// } else {
	row, err = database.Query(defaultRegister, sqlStr, args...)
	// }
	if err != nil {
		return nil, err
	}
	return rows2Maps(-1, row, nil)
}
