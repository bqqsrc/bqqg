package sqlfmt

import (
	"strings"
)

type Delete struct {
	Table      string
	Conditions *AndCondBatch
}

func (d *Delete) Reset() {
	d.Table = ""
	d.Conditions = nil
}

func (d *Delete) SetTable(table string) {
	d.Table = table
}

func (d *Delete) AddConditions(conditions ...SqlFmt) {
	if d.Conditions == nil {
		d.Conditions = &AndCondBatch{}
	}
	d.Conditions.AddConditions(conditions...)
}

func (d *Delete) ToSqlAndArgs() (string, []any, error) {
	deleteName := "Delete"
	if d.Table == "" {
		return "", nil, SqlFmtError(deleteName, "Table is nil")
	}
	var build strings.Builder
	build.WriteString("delete from ")
	build.WriteString(d.Table)
	sqlStr, arg, err := cond2Where(d.Conditions)
	if err != nil {
		return "", nil, err
	}
	build.WriteString(sqlStr)
	return build.String(), arg, nil
}
