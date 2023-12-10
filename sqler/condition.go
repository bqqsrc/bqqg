package sqler

import (
	"strings"

	"github.com/bqqsrc/loger"
)

type ConditionFunc func() (string, []interface{})

func (f ConditionFunc) ToSqlAndArgs() (string, []interface{}) {
	return f()
}

type ConditionLike struct {
	key   string
	value string
	escape string
}

func (c *ConditionLike) Set(key, value, escape string) {
	c.key = key
	c.value = value
	c.escape = escape
}

func (c *ConditionLike) Reset() {
	c.key = ""
	c.value = ""
	c.escape = ""
}

func (c *ConditionLike) SetEscape(escape string) {
	c.escape = escape
}

func (c *ConditionLike) ToSqlAndArgs() (string, []interface{}) {
	if c.key == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(c.key)
	build.WriteString(" like ?")
	args := []interface{}{c.value}
	if c.escape != "" {
		build.WriteString(" escape ?")
		args = append(args, c.escape)
		// build.WriteString(" escape ")
		// build.WriteString(c.escape)
	}
	return build.String(), args
}

func GetConditionLike(key, value, escape string) *ConditionLike {
	return &ConditionLike{key: key, value: value, escape: escape}
}

const (
	Equal = iota
	NotEqual
	Great
	GreatEqual
	Less
	LessEqual
)

type ConditionVale struct {
	key   string
	value interface{}
	oper  string
}

func (c *ConditionVale) Set(key string, value interface{}, oper int) {
	c.key = key
	c.value = value
	switch oper {
	case Equal:
		c.oper = "="
		break
	case NotEqual:
		c.oper = "!="
		break
	case Great:
		c.oper = ">"
		break
	case GreatEqual:
		c.oper = ">="
		break
	case Less:
		c.oper = "<"
		break
	case LessEqual:
		c.oper = "<="
		break
	}
}

func (c *ConditionVale) SetEqual() {
	c.oper = "="
}

func (c *ConditionVale) SetNotEqual() {
	c.oper = "!="
}

func (c *ConditionVale) SetLarger() {
	c.oper = ">"
}

func (c *ConditionVale) SetLargerEqual() {
	c.oper = ">="
}
func (c *ConditionVale) SetLess() {
	c.oper = "<"
}
func (c *ConditionVale) SetLessEqual() {
	c.oper = "<="
}

func (c *ConditionVale) Reset() {
	c.key = ""
	c.value = nil
	c.oper = ""
}

func (c *ConditionVale) ToSqlAndArgs() (string, []interface{}) {
	if c.key == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(c.key)
	build.WriteString(c.oper)
	build.WriteString("?")
	args := []interface{}{c.value}
	return build.String(), args
}

func GetConditionVale(key string, value interface{}, oper int) *ConditionVale {
	condition := &ConditionVale{}
	condition.Set(key, value, oper)
	return condition
}

const (
	Int = iota
	Int64 
	TimeStamp
	Float64 
)

type ConditionRange struct {
	key        string
	begin      interface{}
	end        interface{}
	equalBegin bool
	equalEnd   bool
	keyType int
}

func (c *ConditionRange) Reset() {
	c.key = ""
	c.begin = nil
	c.end = nil
	c.equalBegin = false
	c.equalEnd = false
	c.keyType = 1
}

func (c *ConditionRange) Set(key string, begin, end interface{}, equalBegin, equalEnd bool, keyType int) {
	c.key = key
	c.begin = begin
	c.end = end
	c.equalBegin = equalBegin
	c.equalEnd = equalEnd
	c.keyType = keyType
}

func (c *ConditionRange) SetKeyType(keyType int) {
	c.keyType = keyType
}

func (c *ConditionRange) SetBeginNil() {
	c.begin = nil
}

func (c *ConditionRange) SetEndNil() {
	c.end = nil
}

func (c *ConditionRange) SetKeyValue(key string, begin, end interface{}) {
	c.key = key
	c.begin = begin
	c.end = end
}

func (c *ConditionRange) EqualBegin() {
	c.equalBegin = true
}

func (c *ConditionRange) NotEqualBegin() {
	c.equalBegin = false
}

func (c *ConditionRange) EqualEnd() {
	c.equalEnd = true
}

func (c *ConditionRange) NotEqualEnd() {
	c.equalEnd = false
}

func (c *ConditionRange) ToSqlAndArgs() (string, []interface{}) {
	if c.key == "" || (c.begin == nil && c.end == nil) {
		return "", nil
	}
	var build strings.Builder
	args := make([]interface{}, 0)
	index := 0
	if c.begin != nil {
		build.WriteString(c.key)
		if c.equalBegin {
			build.WriteString(" >= ?")
		} else {
			build.WriteString(" > ?")
		}
		index++
		args = append(args, c.begin)
	}
	if c.end != nil {
		if index > 0 {
			build.WriteString(" and ")
		}
		build.WriteString(c.key)
		if c.equalEnd {
			build.WriteString(" <= ?")
		} else {
			build.WriteString(" < ?")
		}
		args = append(args, c.end)
	}
	return build.String(), args
}

func GetConditionRange(key string, begin, end interface{}, equalBegin, equalEnd bool, keyType int) *ConditionRange {
	return &ConditionRange{key: key, begin: begin, end: end, equalBegin: equalBegin, equalEnd: equalEnd, keyType: keyType}
}

type ConditionTwoKey struct {
	key1 string
	key2 string
	oper string
}

func (c *ConditionTwoKey) Reset() {
	c.key1 = ""
	c.key2 = ""
	c.oper = ""
}

func (c *ConditionTwoKey) Set(key1, key2 string, oper int) {
	c.key1 = key1
	c.key2 = key2
	switch oper {
	case Equal:
		c.oper = "="
		break
	case NotEqual:
		c.oper = "!="
		break
	case Great:
		c.oper = ">"
		break
	case GreatEqual:
		c.oper = ">="
		break
	case Less:
		c.oper = "<"
		break
	case LessEqual:
		c.oper = "<="
		break
	}
}

func (c *ConditionTwoKey) SetEqual() {
	c.oper = "="
}

func (c *ConditionTwoKey) SetNotEqual() {
	c.oper = "!="
}

func (c *ConditionTwoKey) SetLarger() {
	c.oper = ">"
}

func (c *ConditionTwoKey) SetLargerEqual() {
	c.oper = ">="
}
func (c *ConditionTwoKey) SetLess() {
	c.oper = "<"
}
func (c *ConditionTwoKey) SetLessEqual() {
	c.oper = "<="
}

func (c *ConditionTwoKey) ToSqlAndArgs() (string, []interface{}) {
	if (c.key1 == "" && c.key2 == "") || c.oper == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(c.key1)
	build.WriteString(c.oper)
	build.WriteString(c.key2)
	return build.String(), nil
}

func GetConditionTwoKey(key1, key2 string, oper int) *ConditionTwoKey {
	condtion := ConditionTwoKey{}
	condtion.Set(key1, key2, oper)
	return &condtion
}

type ConditionBatch struct {
	oper          string
	conditionList []SqlAndArgs
}

func (c *ConditionBatch) Reset() {
	c.oper = "and"
	c.conditionList = nil
}

func (c *ConditionBatch) SetAnd() {
	c.oper = "and"
}

func (c *ConditionBatch) SetOr() {
	c.oper = "or"
}

func (c *ConditionBatch) AddConditions(conditions ...SqlAndArgs) {
	if c.conditionList == nil {
		c.conditionList = conditions
	} else {
		c.conditionList = append(c.conditionList, conditions...)
	}
}

func (c *ConditionBatch) ToSqlAndArgs() (string, []interface{}) {
	funcName := "ConditionBatch.ToSqlAndArgs"
	oper := c.oper
	if oper == "" {
		oper = "and"
	}
	var build strings.Builder
	index := 0
	args := make([]interface{}, 0)
	count := len(c.conditionList)
	loger.Debugf("%s, count: %d\n", funcName, count)
	for i := 0; i < count; i++ {
		if index > 0 {
			build.WriteString(" ")
			build.WriteString(oper)
			build.WriteString(" ")
		} else {
			index++
		}
		_, isBatch := c.conditionList[i].(*ConditionBatch)
		sql, arg := c.conditionList[i].ToSqlAndArgs()
		if isBatch {
			build.WriteString("(")
		}
		build.WriteString(sql)
		if isBatch {
			build.WriteString(")")
		}
		args = append(args, arg...)
		loger.Debugf("%s, sql: %s\nargs: %v\narg: %v\n", funcName, sql, args, arg)
	}
	return build.String(), args
}
func GetConditionBatch() *ConditionBatch {
	return &ConditionBatch{}
}

func (c *ConditionBatch) toWhere() (string, []interface{}) {
	funcName := "ConditionBatch.toWhere"
	sql, args := c.ToSqlAndArgs()
	if sql == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(" where ")
	build.WriteString(sql)
	loger.Debugf("%s, sql: %s\nargs: %v\n", funcName, sql, args)
	return build.String(), args
}

func (c *ConditionBatch) toOn() (string, []interface{}) {
	sql, args := c.ToSqlAndArgs()
	if sql == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(" on ")
	build.WriteString(sql)
	return build.String(), args
}

func (c *ConditionBatch) toHaving() (string, []interface{}) {
	sql, args := c.ToSqlAndArgs()
	if sql == "" {
		return "", nil
	}
	var build strings.Builder
	build.WriteString(" having ")
	build.WriteString(sql)
	return build.String(), args
}
