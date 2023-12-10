package sqler

import (
	"github.com/bqqsrc/dber"
	"database/sql"
	"strings"

	"github.com/bqqsrc/loger"
)

type Selecter struct {
	//selectkeys *KeyerBatch
	selectKeys []string
	joins      []*Joiner
	tables     *Fromer
	groups     *Grouper
	conditions *ConditionBatch
	orders     *OrderBatch
	limit      int
	offset     int
	havingConditions *ConditionBatch
}

func (s *Selecter) Reset() {
	s.selectKeys = nil
	s.tables = nil
	s.joins = nil
	s.groups = nil
	s.conditions = nil
	s.orders = nil
	s.limit = 0
	s.offset = 0
}

// func (s *Selecter) SetKeyBatch(keyBatch *KeyerBatch) {
// 	s.selectKeys = keyBatch
// }

// func (s *Selecter) AddKeyers(key ...*Keyer) {
// 	if s.selectKeys == nil {
// 		s.selecterKeys = GetKeyerBatch()
// 	}
// 	s.selectKeys.AddKeys(key...)
// }

func (s *Selecter) SetSelectKeys(keys ...string) {
	// if s.selectKeys == nil {
	// 	s.selecterKeys = GetKeyerBatch()
	// }
	s.selectKeys = keys
}

func (s *Selecter) AddSelectKeys(keys ...string) {
	if s.selectKeys == nil {
		s.selectKeys = keys
	} else {
		s.selectKeys = append(s.selectKeys, keys...)
	}
}

func (s *Selecter) GetSelectKeys() []string {
	return s.selectKeys
}

func (s *Selecter) SetTables(tables ...string) {
	if s.tables == nil {
		s.tables = GetFromer()
	}
	s.tables.AddTables(tables...)
	//s.tables = tables
}

func (s *Selecter) AddTables(tables ...string) {
	if s.tables == nil {
		s.tables = GetFromer()
	}
	s.tables.AddTables(tables...)
}

func (s *Selecter) GetTables() []string {
	if s.tables == nil {
		return nil
	}
	return s.tables.GetTables()
}

func (s *Selecter) GetTableNum() int {
	if s.tables == nil {
		return 0
	}
	return len(s.tables.GetTables())
}

func (s *Selecter) AddTable(table string, args ...interface{}) {
	if s.tables == nil {
		s.tables = GetFromer()
	}
	s.tables.AddTableWithArgs(table, args...)
}

func (s *Selecter) SetJoiners(joins ...*Joiner) {
	s.joins = joins
}

func (s *Selecter) AddJoiners(joins ...*Joiner) {
	if s.joins == nil {
		s.joins = joins
	} else {
		s.joins = append(s.joins, joins...)
	}
}

func (s *Selecter) SetGrouper(groups *Grouper) {
	s.groups = groups
}

func (s *Selecter) AddGroupKeys(keys ...string) {
	if s.groups == nil {
		s.groups = GetGrouper(keys...)
	} else {
		s.groups.AddKeys(keys...)
	}
}

func (s *Selecter) SetGroupKeys(keys ...string) {
	if s.groups == nil {
		s.groups = GetGrouper(keys...)
	} else {
		s.groups.SetKeys(keys...)
	}
}

func (s *Selecter) SetOrderBatch(orders *OrderBatch) {
	s.orders = orders
}

func (s *Selecter) AddOrder(keys string, isAsc bool) {
	if s.orders == nil {
		s.orders = GetOrderBatch()
	}
	s.orders.AddOrder(keys, isAsc)
}

func (s *Selecter) SetOrder(keys string, isAsc bool) {
	if s.orders == nil {
		s.orders = GetOrderBatch()
	}
	s.orders.SetOrder(keys, isAsc)
}
func (s *Selecter) AddOrders(orders ...*Order) {
	if s.orders == nil {
		s.orders = GetOrderBatch()
	}
	s.orders.AddOrders(orders...)
}
func (s *Selecter) SetOrders(orders ...*Order) {
	if s.orders == nil {
		s.orders = GetOrderBatch()
	}
	s.orders.SetOrders(orders...)
}

func (s *Selecter) SetLimit(limit int) {
	s.limit = limit
}

func (s *Selecter) SetOffset(offset int) {
	s.offset = offset
}

func (s *Selecter) SetConditionBatch(conditions *ConditionBatch) {
	s.conditions = conditions
}

func (s *Selecter) AddConditions(conditions ...SqlAndArgs) {
	if s.conditions == nil {
		s.conditions = GetConditionBatch()
	}
	s.conditions.AddConditions(conditions...)
}

func (s *Selecter) GetConditions() *ConditionBatch {
	return s.conditions
}


func (s *Selecter) SetHavingConditions(conditions *ConditionBatch) {
	s.havingConditions = conditions
}

func (s *Selecter) AddHavingConditions(conditions ...SqlAndArgs) {
	if s.havingConditions == nil {
		s.havingConditions = GetConditionBatch()
	}
	s.havingConditions.AddConditions(conditions...)
}

func (s *Selecter) GetHavingConditions() *ConditionBatch {
	return s.havingConditions
}

func (s *Selecter) ToSqlAndArgs() (string, []interface{}) {
	funcName := "Selecter.ToSqlAndArgs"
	var build strings.Builder
	args := make([]interface{}, 0)
	sql := s.toSelect()
	loger.Debugf("%s, 1 sql: %s\nargs: %v\n", funcName, sql, args)
	build.WriteString(sql)
	if s.tables == nil {
		return "", nil
	}
	sql, arg := s.tables.ToSqlAndArgs()
	//sql = s.toFrom()
	if sql == "" {
		return "", nil
	}
	args = append(args, arg...)
	loger.Debugf("%s, 2 sql: %s\nargs: %v\n", funcName, sql, args)
	build.WriteString(sql)
	sql, arg = s.toJoin()
	build.WriteString(sql)
	args = append(args, arg...)
	loger.Debugf("%s, 3 sql: %s\narg: %v\nargs: %v\n", funcName, sql, arg, args)
	if s.conditions != nil {
		sql, arg = s.conditions.toWhere()
		build.WriteString(sql)
		args = append(args, arg...)
		loger.Debugf("%s, 4 sql: %s\narg: %v\nargs: %v\n", funcName, sql, arg, args)
	}
	if s.groups != nil {
		sql, arg = s.groups.ToSqlAndArgs()
		build.WriteString(sql)
		args = append(args, arg...)
		loger.Debugf("%s, 5 sql: %s\narg: %v\nargs: %v\n", funcName, sql, arg, args)
	}
	if s.havingConditions != nil {
		sql, arg = s.havingConditions.toHaving()
		build.WriteString(sql)
		args = append(args, arg...)
		loger.Debugf("%s, 5 sql: %s\narg: %v\nargs: %v\n", funcName, sql, arg, args)
	}
	if s.orders != nil {
		sql, arg = s.orders.ToSqlAndArgs()
		build.WriteString(sql)
		args = append(args, arg...)
		loger.Debugf("%s, 6 sql: %s\narg: %v\nargs: %v\n", funcName, sql, arg, args)
	}
	limit := s.limit
	if limit > 0 {
		build.WriteString(" limit ?")
		args = append(args, limit)
	}
	offset := s.offset
	if offset > 0 {
		build.WriteString(" offset ?")
		args = append(args, offset)
	}
	//build.WriteString(";")
	loger.Debugf("%s, 7 arg: %v\nargs: %v\n", funcName, arg, args)
	return build.String(), args
}

func (s *Selecter) toSelect() string {
	selectKeys := s.selectKeys
	if selectKeys == nil {
		return "select * "
	}
	count := len(selectKeys)
	if count == 0 {
		return "select * "
	}
	var build strings.Builder
	index := 0
	for i := 0; i < count; i++ {
		if index > 0 {
			build.WriteString(",")
		} else {
			build.WriteString("select ")
			index++
		}
		build.WriteString(selectKeys[i])
	}
	return build.String()
}

func (s *Selecter) toJoin() (string, []interface{}) {
	joins := s.joins
	if joins == nil {
		return "", nil
	}
	count := len(joins)
	if count == 0 {
		return "", nil
	}
	var build strings.Builder
	allArgs := make([]interface{}, 0)
	for _, value := range joins {
		sql, args := value.ToSqlAndArgs()
		if sql != "" {
			build.WriteString(" ")
			build.WriteString(sql)
			allArgs = append(allArgs, args...)
		}
	}
	return build.String(), allArgs
}

func (s *Selecter) ToMap(controller, txName string, keyMap map[string]string) (map[string]interface{}, error) {
	sqlStr, args := s.ToSqlAndArgs()
	var row *sql.Rows
	var err error
	if txName != "" {
		row, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		row, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return nil, err
	}
	return rows2Map(row, keyMap)
}

func (s *Selecter) ToMapListWithKeys(controller, txName string, keys []string) ([]map[string]interface{}, error) {
	funcName := "Selecter.ToMapList"
	loger.Debugf("%s, controller: %s\ns: %v\n", funcName, controller, s)
	sqlStr, args := s.ToSqlAndArgs()
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	var row *sql.Rows
	var err error
	if txName != "" {
		row, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		row, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return nil, err
	}
	return rows2MapsWithKeys(-1, row, keys)
}

func (s *Selecter) ToMapList(controller, txName string, keyMap map[string]string) ([]map[string]interface{}, error) {
	funcName := "Selecter.ToMapList"
	loger.Debugf("%s, controller: %s\ns: %v\n", funcName, controller, s)
	sqlStr, args := s.ToSqlAndArgs()
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	var row *sql.Rows
	var err error
	if txName != "" {
		row, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		row, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return nil, err
	}
	return rows2Maps(-1, row, keyMap)
}
func (s *Selecter) ToTable(controller, txName string, keyMap map[string]string) ([]string, [][]interface{}, error) {
	return nil, nil, nil
}

func (s *Selecter) ToInt(controller, txName string) (int, error) {
	funcName := "Selecter.ToInt"
	loger.Debugf("%s, controller: %s\ns: %v\n", funcName, controller, s)
	sqlStr, args := s.ToSqlAndArgs()
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sqlStr, args)
	var rows *sql.Rows
	var err error
	if txName != "" {
		rows, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		rows, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return 0, err
	}
	var num int
	for rows.Next() {
		rows.Scan(&num)
	}
	loger.Debugf("%s, num: %d\n", funcName, num)
	return num, nil
}

func (s *Selecter) ToFloat64(controller, txName string) (float64, error) {
	sqlStr, args := s.ToSqlAndArgs()
	var rows *sql.Rows
	var err error
	if txName != "" {
		rows, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		rows, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return 0, err
	}
	var num float64
	for rows.Next() {
		rows.Scan(&num)
	}
	return num, nil
}

func (s *Selecter) ToString(controller, txName string) (string, error) {
	sqlStr, args := s.ToSqlAndArgs()
	var rows *sql.Rows
	var err error
	if txName != "" {
		rows, err = dber.QueryTxSql(controller, txName, sqlStr, args...)
	} else {
		rows, err = dber.Query(controller, sqlStr, args...)
	}
	if err != nil {
		return "", err
	}
	var ret string
	for rows.Next() {
		rows.Scan(&ret)
	}
	return ret, nil
}

func GetSelecter() *Selecter {
	return &Selecter{}
}
