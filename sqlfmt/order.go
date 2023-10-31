package sqlfmt

import "fmt"

type Order struct {
	Key *TableField
	Asc bool
}

func (od *Order) Reset() {
	od.Key = nil
	od.Asc = false
}

func (od *Order) Set(table, key string, asc bool) {
	if od.Key == nil {
		od.Key = &TableField{Table: table, Key: key}
	} else {
		od.Key.Table = table
		od.Key.Key = key
	}
	od.Asc = asc
}

func (od *Order) ToSqlAndArgs() (string, []any, error) {
	orderName := "Order"
	if od.Key == nil {
		return "", nil, SqlFmtError(orderName, "Key is nil")
	}
	sqlStr, args, err := od.Key.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	ascStr := "desc"
	if od.Asc {
		ascStr = "asc"
	}

	return fmt.Sprintf("%s %s", sqlStr, ascStr), args, nil
}
