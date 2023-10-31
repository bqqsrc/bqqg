package sqlfmt

import "strings"

type Rename struct {
	Origin  SqlFmt
	NewName string
}

func (r *Rename) ToSqlAndArgs() (string, []any, error) {
	renameName := "Rename"
	if r.Origin != nil {
		return "", nil, SqlFmtError(renameName, "Origin is nil")
	}
	sql, args, err := r.Origin.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	var build strings.Builder
	build.WriteString("(")
	build.WriteString(sql)
	build.WriteString(" as ")
	build.WriteString(r.NewName)
	build.WriteString(")")
	return build.String(), args, nil
}
