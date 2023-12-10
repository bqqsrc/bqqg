package sqler

import (
	"github.com/bqqsrc/bqqg/database"
	"github.com/bqqsrc/bqqg/log"
)

type Txer struct {
	txSqlList []SqlAndArgs
}

func (t *Txer) Reset() {
	t.txSqlList = nil
}

func (t *Txer) AddSqlAndArgs(sqlAndArgs ...SqlAndArgs) {
	if t.txSqlList == nil {
		t.txSqlList = sqlAndArgs
	} else {
		t.txSqlList = append(t.txSqlList, sqlAndArgs...)
	}
}

// func (t *Txer) ToSqlListAndArgsGroup() ([]string, [][]interface{}) {
// 	if t.txSqlList == nil {
// 		return nil, nil
// 	}
// 	sqlList := make([]string, 0)
// 	argsList := make([][]interface{}, 0)
// 	for _, value := range t.txSqlList {
// 		sql, args := value.ToSqlAndArgs()
// 		if sql == "" {
// 			continue
// 		}
// 		sqlList = append(sqlList, sql)
// 		argsList = append(argsList, args)
// 	}
// 	return sqlList, argsList
// }

func (t *Txer) ToSqlAndArgsBatch() *database.SqlAndArgsBatch {
	batch := &database.SqlAndArgsBatch{}
	for _, value := range t.txSqlList {
		sql, args := value.ToSqlAndArgs()
		if sql == "" {
			continue
		}
		batch.AddSqlAndArgs(sql, args...)
	}
	return batch
}

// func (t *Txer) ToSqlAndArgsGroup() database.SqlAndArgsGroup {
// 	group := database.SqlAndArgsGroup{}
// 	for _, value := range t.txSqlList {
// 		sql, args := value.ToSqlAndArgs()
// 		if sql == "" {
// 			continue
// 		}
// 		group.AddSqlAndArgs(sql, args...)
// 	}
// 	return group
// }

func (t *Txer) ExecSql(controller string) (int64, error) {
	funcName := "Txer.ExecSql"
	batch := t.ToSqlAndArgsBatch()
	log.Debugf("%s, TestTest batch is %v", funcName, batch)
	err := database.ExecTxs(controller, batch)
	return 0, err
}

func GetTxer() *Txer {
	return &Txer{}
}
