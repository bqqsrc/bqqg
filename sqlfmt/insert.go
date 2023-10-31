package sqlfmt

import (
	"strings"
)

type Insert struct {
	KeyValue map[string]any
	Table    string
}

func (i *Insert) Reset() {
	i.KeyValue = nil
	i.Table = ""
}

func (i *Insert) AddKeyValue(key string, value any) {
	if i.KeyValue == nil {
		i.KeyValue = make(map[string]any)
	}
	i.KeyValue[key] = value
}

func (i *Insert) ToSqlAndArgs() (string, []any, error) {
	insertName := "Insert"
	if i.Table == "" {
		return "", nil, SqlFmtError(insertName, "Table is empty")
	}
	keyValue := i.KeyValue
	if keyValue == nil || len(keyValue) == 0 {
		return "", nil, SqlFmtError(insertName, "KeyValue is nil or empty")
	}
	var build strings.Builder
	build.WriteString("insert into ")
	build.WriteString(i.Table)
	build.WriteString("(")
	var tmpBuild strings.Builder
	index := 0
	args := make([]any, 0)
	for key, value := range keyValue {
		if index > 0 {
			build.WriteString(",")
			tmpBuild.WriteString(",")
		} else {
			index++
		}
		build.WriteString(key)
		if valueSqlAndArgs, ok := value.(SqlFmt); ok {
			sql, tmpArgs, err := valueSqlAndArgs.ToSqlAndArgs()
			if err != nil {
				return "", nil, err
			}
			if sql != "" {
				tmpBuild.WriteString("(")
				tmpBuild.WriteString(sql)
				tmpBuild.WriteString(")")
				args = append(args, tmpArgs...)
			}
		} else {
			tmpBuild.WriteString("?")
			args = append(args, value)
		}
	}
	build.WriteString(") values (")
	build.WriteString(tmpBuild.String())
	build.WriteString(")")
	return build.String(), args, nil
}

type InsertFormatter interface {
	ToInsert() (string, []any, error)
}

type Inserter struct {
	InsertFmt InsertFormatter
	Table     string
}

func (i *Inserter) ToSqlAndArgs() (string, []any, error) {
	inserterName := "Inserter"
	if i.InsertFmt == nil || i.Table == "" {
		return "", nil, SqlFmtError(inserterName, "InsertFmt is nil or Table is empty, InsertFmt: %v, Table: %s", i.InsertFmt, i.Table)
	}
	sql, args, err := i.InsertFmt.ToInsert()
	if err != nil {
		return "", nil, err
	}
	var build strings.Builder
	build.WriteString("insert into ")
	build.WriteString(i.Table)
	build.WriteString(sql)
	return build.String(), args, nil
}
