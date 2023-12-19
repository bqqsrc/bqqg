package sqler

type SqlAndArgs interface {
	ToSqlAndArgs() (string, []interface{})
}

type Execer interface {
	ExecSql(controller string) (int64, error)
}

type ExecTxer interface {
	ExecSqlTx(controller, name string, commit bool) error
}
