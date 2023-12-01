package database

import (
	//	"bqqgc/errer"
	"database/sql"
	"github.com/bqqsrc/bqqg/log"
	"io/ioutil"
	"os"
	"strings"
)

type sqlAndArgs struct {
	sqlStr string
	args   []any
}

type SqlAndArgsBatch struct {
	batch []sqlAndArgs
}

func (s *SqlAndArgsBatch) AddSqlAndArgs(sqlStr string, args ...any) {
	if s.batch == nil {
		s.batch = make([]sqlAndArgs, 0)
	}
	s.batch = append(s.batch, sqlAndArgs{sqlStr, args})
}

type SqlController struct {
	name   string
	driver string
	db     *sql.DB
	txs    map[string]*sql.Tx //TODO验证一些可不可以同时开启多个tx
}

func (s *SqlController) GetDB() *sql.DB {
	return s.db
}

var controllerMap map[string]*SqlController

func GetController(name string) (*SqlController, bool) {
	if controllerMap != nil {
		controller, ok := controllerMap[name]
		return controller, ok
	}
	return nil, false
}

func RegistController(name, driverName string, sqlName string) (*SqlController, error) {
	funcName := "RegistSqler"
	if controllerMap == nil {
		controllerMap = make(map[string]*SqlController)
	}
	if _, ok := controllerMap[name]; ok {
		return nil, CallerErr(funcName, "redeclare SqlController, name: %s, driverName: %s, args1: %s", name, driverName, sqlName)
	}
	switch driverName {
	case "sqlite3", "mysql":
		db, err := sql.Open(driverName, sqlName)
		//TODO 这里连接上了之后，添加一个查看所有表格的操作，以确定是否真的连接成功
		if err != nil {
			return nil, CallerErr(funcName, "sql.Open(%s, %s) error, err is: %s", driverName, sqlName, err)
		}
		if err = checkDB(db); err != nil {
			return nil, CallerErr(funcName, "checkDB error, check the db passwd, user, host, etc... driverName: %s, sqlName: %s, err is: %s", driverName, sqlName, err)
		}
		sqler := SqlController{name: name, driver: driverName, db: db}
		controllerMap[name] = &sqler
		return &sqler, nil
	default:
		return nil, CallerErr(funcName, "unsupport driverName: %s, only can be %s", driverName, "sqlite3")
	}
}

func checkDB(db *sql.DB) error {
	funcName := "checkDB"
	stmt, err := db.Prepare("show tables")
	defer closeStmt(stmt)
	if err != nil {
		return CallerErr(funcName, "db.Prepare(show tables) error, err: %s", err)
	}
	_, err = stmt.Query()
	if err != nil {
		return CallerErr(funcName, "stmt.Query error, err: %s", err)
	}
	return nil
}

func UnRegistController(name string) error {
	funcName := "UnRegistSqler"
	sqler, ok := controllerMap[name]
	if !ok {
		return nil
	}
	if len(sqler.txs) > 0 {
		return CallerErr(funcName, "can't UnRegistSqler while txs more than one, name: %s", name)
	}
	if err := sqler.db.Close(); err != nil {
		return CallerErr(funcName, "sqler.db.Close error, name: %s, err: %s", name, err)
	}
	delete(controllerMap, name)
	return nil
}

func (s *SqlController) UnRegistController() error {
	name := s.name
	return UnRegistController(name)
}

func Exec(controllerName, sqlStr string, args ...any) (sql.Result, error) {
	funcName := "Exec"
	if controller, ok := GetController(controllerName); !ok {
		return nil, CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.Exec(sqlStr, args...)
	}
}

func (s *SqlController) Exec(sqlStr string, args ...any) (sql.Result, error) {
	funcName := "SqlController.Exec"
	if err := s.checkController(funcName); err != nil {
		return nil, err
	}
	stmt, err := s.db.Prepare(sqlStr)
	defer closeStmt(stmt)
	if err != nil {
		return nil, CallerErr(funcName, "s.db.Prepare(%s) error, err: %s", sqlStr, err)
	}
	ret, err := stmt.Exec(args...)
	if err != nil {
		return nil, CallerErr(funcName, "stmt.Exec error, err: %s", err)
	}
	return ret, nil
}

func Query(controllerName, sqlStr string, args ...any) (*sql.Rows, error) {
	funcName := "Query"
	if controller, ok := GetController(controllerName); !ok {
		return nil, CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.Query(sqlStr, args...)
	}
}
func (s *SqlController) Query(sqlStr string, args ...any) (*sql.Rows, error) {
	//TODO考虑如何实现多并发的数据查询方式
	//验证在事务过程中，执行查询能否查到还未提交的数据
	funcName := "SqlController.Query"
	if err := s.checkController(funcName); err != nil {
		return nil, err
	}
	stmt, err := s.db.Prepare(sqlStr)
	// log.Debugf("sqlStr is %s, args is %v", sqlStr, args)
	defer closeStmt(stmt)
	if err != nil {
		return nil, CallerErr(funcName, "s.db.Prepare(%s) error, err: %s", sqlStr, err)
	}
	res, err := stmt.Query(args...)
	if err != nil {
		return nil, CallerErr(funcName, "stmt.Query error, err: %s", err)
	}
	return res, nil
}

func ExecTxs(controllerName string, sqlAndArgsArr *SqlAndArgsBatch) error {
	funcName := "ExecTxs"
	log.Debugf("%s, Test", funcName)
	if controller, ok := GetController(controllerName); !ok {
		return CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.ExecTxs(sqlAndArgsArr)
	}
}
func (s *SqlController) ExecTxs(sqlAndArgsArr *SqlAndArgsBatch) error {
	funcName := "SqlController.ExecTxs"
	name := "ExecTxs"
	if sqlAndArgsArr == nil {
		return CallerErr(funcName, "sqlAndArgsArr is nil")
	}
	batch := sqlAndArgsArr.batch
	if batch != nil {
		for _, value := range sqlAndArgsArr.batch {
			log.Debugf("%s, sql: %s, args: %v", funcName, value.sqlStr, value.args)
			if _, err := s.ExecTxSql(name, value.sqlStr, value.args...); err != nil {
				return err
			}
		}
		return s.CommitTx(name)
	}
	return nil
}
func ExecTxSql(controllerName, txName, sqlStr string, args ...any) (sql.Result, error) {
	funcName := "ExecTxs"
	if controller, ok := GetController(controllerName); !ok {
		return nil, CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.ExecTxSql(txName, sqlStr, args...)
	}
}
func (s *SqlController) ExecTxSql(name, sqlStr string, args ...any) (sql.Result, error) {
	funcName := "SqlController.ExecTxSql"
	err := s.checkController(funcName)
	if err != nil {
		return nil, err
	}
	if s.txs == nil {
		s.txs = make(map[string]*sql.Tx)
	}
	tx, ok := s.txs[name]
	if !ok {
		tx, err = s.db.Begin()
		if err != nil {
			return nil, CallerErr(funcName, "s.db.Begin error, name: %s, driver: %s, err: %s", s.name, s.driver, err)
		}
		s.txs[name] = tx
	}
	if ret, err := tx.Exec(sqlStr, args...); err != nil {
		if newErr := tx.Rollback(); newErr != nil {
			log.Criticalf("%s, tx.Rollback error, err: %s", funcName, newErr)
		}
		delete(s.txs, name)
		//TODO 考虑添加一些方法，在出现队列中有待执行的sql语句时，会继续运行
		return nil, CallerErr(funcName, "tx.Exec(%s, %v) error, err: %s", sqlStr, args, err)
	} else {
		return ret, nil
	}
}
func QueryTxSql(controllerName, txName, sqlStr string, args ...any) (*sql.Rows, error) {
	funcName := "QueryTxSql"
	if controller, ok := GetController(controllerName); !ok {
		return nil, CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.QueryTxSql(txName, sqlStr, args...)
	}
}
func (s *SqlController) QueryTxSql(name, sqlStr string, args ...any) (*sql.Rows, error) {
	funcName := "SqlController.QueryTxSql"
	err := s.checkController(funcName)
	if err != nil {
		return nil, err
	}
	if s.txs == nil {
		s.txs = make(map[string]*sql.Tx)
	}
	tx, ok := s.txs[name]
	if !ok {
		tx, err = s.db.Begin()
		if err != nil {
			return nil, CallerErr(funcName, "s.db.Begin error, name: %s, driver: %s, err: %s", s.name, s.driver, err)
		}
		s.txs[name] = tx
	}
	if ret, err := tx.Query(sqlStr, args...); err != nil {
		if newErr := tx.Rollback(); newErr != nil {
			log.Criticalf("%s, tx.Rollback error, err: %s", funcName, newErr)
		}
		delete(s.txs, name)
		//TODO 考虑添加一些方法，在出现队列中有待执行的sql语句时，会继续运行
		return nil, CallerErr(funcName, "tx.Exec(%s, %v) error, err: %s", sqlStr, args, err)
	} else {
		return ret, nil
	}
}
func CommitTx(controllerName, name string) error {
	funcName := "ExecTxs"
	if controller, ok := GetController(controllerName); !ok {
		return CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.CommitTx(name)
	}
}
func (s *SqlController) CommitTx(name string) error {
	funcName := "SqlController.CommitTx"
	if err := s.checkController(funcName); err != nil {
		return err
	}
	if s.txs == nil {
		return CallerErr(funcName, "s.txs is nil")
	}
	tx, ok := s.txs[name]
	if !ok {
		return CallerErr(funcName, "%s is not found in s.txs", name)
	}
	defer delete(s.txs, name)
	if err := tx.Commit(); err != nil {
		return CallerErr(funcName, "tx.Commit error, name: %s, driver: %s, err: %s", s.name, s.driver, err)
	}
	return nil
}
func ExecSqlScripts(controllerName, script string) error {
	funcName := "ExecTxs"
	if controller, ok := GetController(controllerName); !ok {
		return CallerErr(funcName, "%s not found in controllerMap", controllerName)
	} else {
		return controller.ExecSqlScripts(script)
	}
}
func (s *SqlController) ExecSqlScripts(script string) error {
	funcName := "SqlController.ExecSqlScripts"
	file, err := os.Lstat(script)
	if os.IsNotExist(err) {
		return CallerErr(funcName, "script not found, %s", script)
	}
	if file.IsDir() {
		return CallerErr(funcName, "script is dir, %s", script)
	}
	data, err := ioutil.ReadFile(script)
	if err != nil {
		return CallerErr(funcName, "ioutil.ReadFile(%s) error, err: %s", script, err)
	}
	scriptStr := string(data)
	return s.ExecSqls(scriptStr)
	// scriptArr := strings.Split(scriptStr, ";")
	// sqlArgs := SqlAndArgsBatch{}
	// for _, value := range scriptArr {
	// 	value = strings.TrimSpace(value)
	// 	if value != "" {
	// 		sqlArgs.AddSqlAndArgs(value)
	// 	}
	// }
	// return s.ExecTxs(&sqlArgs)
}

func (s *SqlController) ExecSqls(script string) error {
	//funcName := "SqlController.ExecSqls"
	scriptArr := strings.Split(script, ";")
	sqlArgs := SqlAndArgsBatch{}
	for _, value := range scriptArr {
		value = strings.TrimSpace(value)
		if value != "" {
			sqlArgs.AddSqlAndArgs(value)
		}
	}
	return s.ExecTxs(&sqlArgs)
}

func (s *SqlController) checkController(funcName string) error {
	if s == nil || s.db == nil {
		return CallerErr(funcName, "SqlController is nil or SqlController.db is nil")
	}
	return nil
}

func closeStmt(stmt *sql.Stmt) {
	if err := stmt.Close(); err != nil {
		log.Errorf("closeStmt error, %s", err)
	}
}
