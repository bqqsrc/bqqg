package sqlfmt

import "strings"

const (
	LeftJoin = iota
	RightJoin
	InnerJoin
	OuterJoin
)

type Join struct {
	LeftOrRight  int
	InnerOrOuter int
	Tables       *Fields
	Conditions   *AndCondBatch
}

func (jn *Join) Reset() {
	jn.LeftOrRight = LeftJoin
	jn.InnerOrOuter = InnerJoin
	jn.Tables = nil
	jn.Conditions = nil
}

func (jn *Join) SetType(LeftOrRight, innerOrOuter int) {
	jn.LeftOrRight = LeftOrRight
	jn.InnerOrOuter = innerOrOuter
}

func (jn *Join) AddTables(tables ...string) {
	if jn.Tables == nil {
		jn.Tables = &Fields{}
	}
	jn.Tables.AddFields(tables...)
}

func (jn *Join) AddTablesBySqlFmt(tables ...SqlFmt) {
	if jn.Tables == nil {
		jn.Tables = &Fields{}
	}
	jn.Tables.AddSqlFmt(tables...)
}

func (jn *Join) AddConditions(conds ...SqlFmt) {
	if jn.Conditions == nil {
		jn.Conditions = &AndCondBatch{}
	}
	jn.Conditions.AddConditions(conds...)
}

func (jn *Join) ToSqlAndArgs() (string, []any, error) {
	joinName := "Join"
	if jn.Tables == nil {
		return "", nil, SqlFmtError(joinName, "Tables is nil")
	}
	var build strings.Builder
	build.WriteString(" ")
	if jn.LeftOrRight == LeftJoin {
		build.WriteString("left")
	} else {
		build.WriteString("right")
	}
	build.WriteString(" ")
	if jn.InnerOrOuter == InnerJoin {
		build.WriteString("inner")
	} else {
		build.WriteString("outer")
	}
	build.WriteString(" join ")
	sqlStr, args, err := jn.Tables.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	build.WriteString(sqlStr)
	finalArgs := make([]any, 0)
	finalArgs = append(finalArgs, args...)
	if jn.Conditions != nil {
		if sqlStr, args, err = cond2On(jn.Conditions); err != nil {
			return "", nil, err
		}
		build.WriteString(sqlStr)
		finalArgs = append(finalArgs, args...)
	}
	return build.String(), finalArgs, nil
}
