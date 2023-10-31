package sqlfmt4type

type SqlFmt4Type interface {
	ToInsert() (string, []any, error)
	ToDelete() (string, []any, error)
	ToUpdate() (string, []any, error)
	ToSelect() (string, []any, error)
}
