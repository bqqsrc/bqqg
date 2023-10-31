package sqlfmt

import (
	"strings"
)

type Update struct {
	KeyValue   map[string]any
	Table      string
	Conditions *AndCondBatch
}

func (up *Update) Reset() {
	up.KeyValue = nil
	up.Table = ""
	up.Conditions = nil
}

func (up *Update) AddKeyValue(key string, value any) {
	if up.KeyValue == nil {
		up.KeyValue = make(map[string]any)
	}
	up.KeyValue[key] = value
}

func (up *Update) AddConditions(conditions ...SqlFmt) {
	if up.Conditions == nil {
		up.Conditions = &AndCondBatch{}
	}
	up.Conditions.AddConditions(conditions...)
}

func (up *Update) ToSqlAndArgs() (string, []any, error) {
	updateName := "Update"
	if up.Table == "" {
		return "", nil, SqlFmtError(updateName, "Table is empty")
	}
	keyValue := up.KeyValue
	if keyValue == nil || len(keyValue) == 0 {
		return "", nil, SqlFmtError(updateName, "KeyValue is nil or empty")
	}
	var build strings.Builder
	build.WriteString("update ")
	build.WriteString(up.Table)
	build.WriteString(" set ")
	index := 0
	args := make([]any, 0)
	for key, value := range keyValue {
		if index > 0 {
			build.WriteString(",")
		} else {
			index++
		}
		build.WriteString(key)
		build.WriteString("=?")
		args = append(args, value)
	}
	if up.Conditions != nil {
		sqlStr, arg, err := cond2Where(up.Conditions)
		if err != nil {
			return "", nil, err
		}
		build.WriteString(sqlStr)
		args = append(args, arg...)
	}
	return build.String(), args, nil
}

type UpdateFormatter interface {
	ToUpdate() (string, []any, error)
}

type Updater struct {
	UpdateFmt  UpdateFormatter
	Table      string
	Conditions *AndCondBatch
}

func (up *Updater) ToSqlAndArgs() (string, []any, error) {
	updaterName := "Updater"
	if up.UpdateFmt == nil || up.Table == "" {
		return "", nil, SqlFmtError(updaterName, "UpdateFmt is nil or Table is empty, UpdateFmt: %v, Table: %s", up.UpdateFmt, up.Table)
	}
	finalArgs := make([]any, 0)
	sqlStr, args, err := up.UpdateFmt.ToUpdate()
	if err != nil {
		return "", nil, err
	}
	finalArgs = append(finalArgs, args...)
	var build strings.Builder
	build.WriteString("update ")
	build.WriteString(up.Table)
	build.WriteString(" set ")
	build.WriteString(sqlStr)
	//TODO 这里的where是放在括号里还是括号外？
	build.WriteString(" where ")
	if sqlStr, args, err = up.Conditions.ToSqlAndArgs(); err != nil {
		return "", nil, err
	}
	build.WriteString(sqlStr)
	finalArgs = append(finalArgs, args...)
	return build.String(), finalArgs, nil
}
