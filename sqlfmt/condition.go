package sqlfmt

import (
	"fmt"
	"strings"
)

func toCondKeyValueWithOper(key *TableField, oper string, value any, call string) (string, []any, error) {
	if key == nil {
		return "", nil, SqlFmtError(call, "Key is nil")
	}
	sql, args, err := key.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	var build strings.Builder
	build.WriteString(sql)
	build.WriteString(oper)
	build.WriteString("?")
	if args == nil {
		args = []any{value}
	} else {
		args = append(args, value)
	}
	return build.String(), args, nil
}

func toCond2KeyWithOper(key1 *TableField, oper string, key2 *TableField, call string) (string, []any, error) {
	if key1 == nil || key2 == nil {
		return "", nil, SqlFmtError(call, "key1 or key2 is nil, key1: %v, key2: %v", key1, key2)
	}
	key1Sql, args1, err := key1.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	key2Sql, args2, err := key2.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	var build strings.Builder
	var finalArgs []any
	build.WriteString(key1Sql)
	build.WriteString(oper)
	build.WriteString(key2Sql)
	if args1 != nil {
		finalArgs = args1
	}
	if args2 != nil {
		if finalArgs != nil {
			finalArgs = append(finalArgs, args2...)
		} else {
			finalArgs = args2
		}
	}
	return build.String(), finalArgs, nil
}

// 键like值条件，最终会转换为sql格式：
//
// key like ?
//
// 或 table.key like ? （如果Key设置了Table）
//
// 而value则作为参数返回
//
// 如果设置了EscapeValue，则转换为 key like ? escape ?
//
// 同理EscapeValue会作为参数返回
type CondKeyLikeValue struct {
	Key         *TableField
	Value       any
	EscapeValue any
}

func (c *CondKeyLikeValue) ToSqlAndArgs() (string, []any, error) {
	sql, args, err := toCondKeyValueWithOper(c.Key, " like ", c.Value, "CondKeyLikeValue")
	if err != nil {
		return "", nil, err
	}
	if sql != "" && c.EscapeValue != "" {
		var build strings.Builder
		build.WriteString(sql)
		build.WriteString(" escape ?")
		args = append(args, c.EscapeValue)
		return build.String(), args, nil
	}
	return sql, args, nil
}

// 键等于值条件，最终会转换为sql格式：
//
// key=?
//
// 或 table.key=? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyEqValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyEqValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, "=", c.Value, "CondKeyEqValue")
}

// 键不等于值条件，最终会转换为sql格式：
//
// key!=?
//
// 或 table.key!=? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyNeValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyNeValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, "!=", c.Value, "CondKeyNeValue")
}

// 键大于值条件，最终会转换为sql格式：
//
// key>?
//
// 或 table.key>? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyGtValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyGtValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, ">", c.Value, "CondKeyGtValue")
}

// 键大于等于值条件，最终会转换为sql格式：
//
// key>=?
//
// 或 table.key>=? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyGeValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyGeValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, ">=", c.Value, "CondKeyGeValue")
}

// 键小于值条件，最终会转换为sql格式：
//
// key<?
//
// 或 table.key<? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyLtValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyLtValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, "<", c.Value, "CondKeyLtValue")
}

// 键小于等于值条件，最终会转换为sql格式：
//
// key<=?
//
// 或 table.key<=? （如果Key设置了Table）
//
// 而value则作为参数返回
type CondKeyLeValue struct {
	Key   *TableField
	Value any
}

func (c *CondKeyLeValue) ToSqlAndArgs() (string, []any, error) {
	return toCondKeyValueWithOper(c.Key, "<=", c.Value, "CondKeyLeValue")
}

// 键取范围条件，最终会转换为sql格式：
//
// key > ? and key < ?
//
// 上面的大于号可以通过设置EqualBegin来变成>=
//
// 同理，小于号可以通过设置EqualEnd变成<=
//
// 如果Table设置了Table值，则key会变成table.key
//
// 而BeginValue和EndValue则作为参数返回
type CondKeyRangeBeginEnd struct {
	Key        *TableField
	BeginValue any
	EndValue   any
	EqualBegin bool
	EqualEnd   bool
}

func (c *CondKeyRangeBeginEnd) ToSqlAndArgs() (string, []any, error) {
	condKeyRangeBeginEndName := "CondKeyRangeBeginEnd"
	if c.Key == nil {
		return "", nil, SqlFmtError(condKeyRangeBeginEndName, "c.Key is nil")
	}
	tableKey, args, err := c.Key.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	if c.BeginValue == nil && c.EndValue == nil {
		return "", nil, SqlFmtError(condKeyRangeBeginEndName, "both BeginValue and EndValue is nil")
	}
	var build strings.Builder
	if args == nil {
		args = make([]any, 0)
	}
	hasBegin := false
	if c.BeginValue != nil {
		build.WriteString(tableKey)
		if c.EqualBegin {
			build.WriteString(">=?")
		} else {
			build.WriteString(">?")
		}
		hasBegin = true
		args = append(args, c.BeginValue)
	}
	if c.EndValue != nil {
		if hasBegin {
			build.WriteString(" and ")
		}
		build.WriteString(tableKey)
		if c.EqualEnd {
			build.WriteString("<=?")
		} else {
			build.WriteString("<?")
		}
		args = append(args, c.EndValue)
	}
	return build.String(), args, nil
}

// 键取范围条件，最终会转换为sql格式：
//
// key between ? and ?
//
// 或 table.key between ? and ? （如果Key设置了Table）
//
// 而BeginValue和EndValue则作为参数返回
type CondKeyBetwwenBeginEnd struct {
	Key        *TableField
	BeginValue any
	EndValue   any
}

func (c *CondKeyBetwwenBeginEnd) ToSqlAndArgs() (string, []any, error) {
	condKeyBetwwenBeginEndName := "CondKeyBetwwenBeginEnd"
	if c.Key == nil {
		return "", nil, SqlFmtError(condKeyBetwwenBeginEndName, "c.Key is nil")
	}
	tableKey, args, err := c.Key.ToSqlAndArgs()
	if err != nil {
		return "", nil, err
	}
	if c.BeginValue == nil || c.EndValue == nil {
		return "", nil, SqlFmtError(condKeyBetwwenBeginEndName, "BeginValue or EndValue is nil, BeginValue is %v, EndValue is %v", c.BeginValue, c.EndValue)
	}
	var build strings.Builder
	build.WriteString(tableKey)
	build.WriteString(" between ? and ?")
	if args == nil {
		args = []any{c.BeginValue, c.EndValue}
	} else {
		args = append(args, c.BeginValue, c.EndValue)
	}
	return build.String(), args, nil
}

// 键等于键条件，最终会转换为sql格式：
//
// key1=key2
//
// 或 table1.key1=table2.key2 （如果Key设置了Table）
type CondKeyEqKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyEqKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, "=", c.Key2, "CondKeyEqKey")
}

// 键不等于键条件，最终会转换为sql格式：
//
// key1!=key2
//
// 或 table1.key1!=table2.key2 （如果Key设置了Table）
type CondKeyNeKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyNeKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, "!=", c.Key2, "CondKeyNeKey")
}

// 键大于键条件，最终会转换为sql格式：
//
// key1>key2
//
// 或 table1.key1>table2.key2 （如果Key设置了Table）
type CondKeyGtKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyGtKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, ">", c.Key2, "CondKeyGtKey")
}

// 键大于等于键条件，最终会转换为sql格式：
//
// key1>=key2
//
// 或 table1.key1>=table2.key2 （如果Key设置了Table）
type CondKeyGeKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyGeKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, ">=", c.Key2, "CondKeyGeKey")
}

// 键小于键条件，最终会转换为sql格式：
//
// key1<key2
//
// 或 table1.key1<table2.key2 （如果Key设置了Table）
type CondKeyLtKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyLtKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, "<", c.Key2, "CondKeyLtKey")
}

// 键小于等于键条件，最终会转换为sql格式：
//
// key1<=key2
//
// 或 table1.key1<=table2.key2 （如果Key设置了Table）
type CondKeyLeKey struct {
	Key1 *TableField
	Key2 *TableField
}

func (c *CondKeyLeKey) ToSqlAndArgs() (string, []any, error) {
	return toCond2KeyWithOper(c.Key1, "<=", c.Key2, "CondKeyLeKey")
}

func toBatchCondition(call, oper string, conds []SqlFmt) (string, []any, error) {
	if conds == nil {
		return "", nil, SqlFmtError(call, "Conds is nil")
	}
	count := len(conds)
	if count == 0 {
		return "", nil, SqlFmtError(call, "Conds has no element")
	}
	var build strings.Builder
	hasFirst := false
	finalArgs := make([]any, 0)
	var sqlStr string
	var err error
	var args []any
	for i := 0; i < count; i++ {
		if hasFirst {
			build.WriteString(oper)
		}
		sqlStr, args, err = conds[i].ToSqlAndArgs()
		if err != nil {
			return "", nil, err
		}
		if sqlStr != "" {
			if !hasFirst {
				build.WriteString("(")
				hasFirst = true
			}
			build.WriteString(sqlStr)
		}
		if args != nil {
			finalArgs = append(finalArgs, args...)
		}
	}
	if hasFirst {
		build.WriteString(")")
	}
	return build.String(), finalArgs, nil
}

// or批量条件
//
// 可以将各种条件组成一个or条件组
//
// 最后会转换为sql语句：
//
// (条件1 or 条件2 or 条件3 ...)
type OrCondBatch struct {
	Conds []SqlFmt
}

// 添加条件SqlFmt
func (c *OrCondBatch) AddConditions(conds ...SqlFmt) {
	if c.Conds == nil {
		c.Conds = conds
	} else {
		c.Conds = append(c.Conds, conds...)
	}
}

// 重置条件
func (c *OrCondBatch) Reset() {
	c.Conds = nil
}

func (c *OrCondBatch) ToSqlAndArgs() (string, []any, error) {
	return toBatchCondition("OrCondBatch", " or ", c.Conds)
}

// and批量条件
//
// 可以将各种条件组成一个and条件组
//
// 最后会转换为sql语句：
//
// (条件1 and 条件2 and 条件3 ...)
type AndCondBatch struct {
	Conds []SqlFmt
}

// 添加条件SqlFmt
func (c *AndCondBatch) AddConditions(conds ...SqlFmt) {
	if c.Conds == nil {
		c.Conds = conds
	} else {
		c.Conds = append(c.Conds, conds...)
	}
}

// 重置条件
func (c *AndCondBatch) Reset() {
	c.Conds = nil
}

func (c *AndCondBatch) ToSqlAndArgs() (string, []any, error) {
	return toBatchCondition("AndCondBatch", " and ", c.Conds)
}

func cond2Sql(key string, cond SqlFmt) (string, []any, error) {
	if cond == nil {
		return "", nil, SqlFmtError("cond2Sql", "cond is nil, key: %s", key)
	}
	if sql, args, err := cond.ToSqlAndArgs(); err != nil {
		return "", nil, err
	} else {
		return fmt.Sprintf(" %s %s", key, sql), args, nil
	}
}

func cond2Where(cond SqlFmt) (string, []any, error) {
	return cond2Sql("where", cond)
}

func cond2On(cond SqlFmt) (string, []any, error) {
	return cond2Sql("on", cond)
}

func cond2Having(cond SqlFmt) (string, []any, error) {
	return cond2Sql("having", cond)
}
