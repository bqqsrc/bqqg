package sqler

import "strings"

type Fromer struct {
	tables []string
	args   []interface{}
}

func (f *Fromer) Reset() {
	f.tables = nil
	f.args = nil
}

func (f *Fromer) AddTables(tables ...string) {
	if f.tables == nil {
		f.tables = tables
	} else {
		f.tables = append(f.tables, tables...)
	}
}

func (f *Fromer) AddTableWithArgs(table string, args ...interface{}) {
	if f.tables == nil {
		f.tables = make([]string, 0)
	}
	f.tables = append(f.tables, table)
	if f.args == nil {
		f.args = args
	} else {
		f.args = append(f.args, args...)
	}
}

func (f *Fromer) ToSqlAndArgs() (string, []interface{}) {
	tables := f.tables
	if tables == nil {
		return "", nil
	}
	count := len(tables)
	if count == 0 {
		return "", nil
	}
	var build strings.Builder
	index := 0
	for i := 0; i < count; i++ {
		if index > 0 {
			build.WriteString(",")
		} else {
			build.WriteString(" from ")
			if count > 1 {
				build.WriteString("(")
			}
			index++
		}
		build.WriteString(tables[i])
	}
	if count > 1 {
		build.WriteString(")")
	}
	return build.String(), f.args
}

func (f *Fromer) GetTables() []string {
	return f.tables
}

func GetFromer() *Fromer {
	return &Fromer{}
}
