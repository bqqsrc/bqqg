package sqler

import (
	"github.com/bqqsrc/imaper"
	"strings"
	"github.com/bqqsrc/loger"
)

type Keyer struct {
	table string 
	key string
	alias string
	group *Grouper
	expression interface{}
}

func (k *Keyer) Reset() {
	k.table = ""
	k.key = ""
	k.alias = ""
	k.group = nil
	k.expression = nil
}

func (k *Keyer) SetTable(table string) {
	k.table = table
}

func (k *Keyer) SetKey(key string) {
	k.key = key
}

func (k *Keyer) SetAlias(alias string) {
	k.alias = alias
}

func (k *Keyer) SetGroup(group *Grouper) {
	k.group = group
}

func (k *Keyer) SetExpression(expression interface{}) {
	k.expression = expression
}

func (k *Keyer) GetGrouper() *Grouper {
	return k.group
}

func (k *Keyer) ToSqlAndArgs() (string, []interface{}) {
	funcName := "Keyer.ToSqlAndArgs"
	if k.key == "" && k.expression == nil {
		return "", nil
	}
	if k.key != "" && k.expression != nil {
		loger.Errorf("%s error, both key and expression are not empty, key: %s, expression: %s", funcName, k.key, k.expression)
		return "", nil
	}
	var build strings.Builder
	if k.table != "" {
		build.WriteString(k.table)
		build.WriteString(".")
	}
	if k.key != "" {
		build.WriteString(k.key)
	}
	keyArgs := make([]interface{}, 0)
	if k.expression != nil {
		if expressSqlAndArgs, ok := k.expression.(SqlAndArgs); ok {
			sqlStr, args := expressSqlAndArgs.ToSqlAndArgs()
			build.WriteString(sqlStr)
			keyArgs = append(keyArgs, args...)
		} else {
			if keyStr, ok := imaper.I2String(k.expression); ok {
				build.WriteString(keyStr)
			}
		}
	}
	if k.alias != "" {
		build.WriteString(" as ")
		build.WriteString(k.alias)
	}
	return build.String(), keyArgs
}

func GetKeyer(table, key, alias string) *Keyer {
	return &Keyer{table: table, key: key, alias: alias}
} 

type KeyerBatch struct {
	keys []*Keyer
}

func (k *KeyerBatch) Reset() {
	k.keys = nil
}

func (k *KeyerBatch) SetKeys(keys ...*Keyer) {
	k.keys = keys
}

func (k *KeyerBatch) AddKeys(keys ...*Keyer) {
	if k.keys == nil {
		k.keys = keys 
	} else {
		k.keys = append(k.keys, keys...)
	}
}

func (k *KeyerBatch) AddKey(table, key, alias string) {
	keyer := Keyer{table: table, key: key, alias: alias}
	k.AddKeys(&keyer)
}

func (k *KeyerBatch) GetGroupers() []*Grouper {
	if k.keys == nil || len(k.keys) == 0 {
		return nil
	}
	groupers := make([]*Grouper, 0)
	for _, value := range k.keys {
		groupers = append(groupers, value.group)
	}
	return groupers
}

func (k *KeyerBatch) ToSqlAndArgs() (string, []interface{}) {
	if k.keys == nil || len(k.keys) == 0 {
		return "", nil
	}
	var build strings.Builder
	index := 0
	args := make([]interface{}, 0)
	for _, value := range k.keys {
		sql, arg := value.ToSqlAndArgs() 
		if sql != "" {
			if index > 0 {
				build.WriteString(", ")
			} else {
				build.WriteString("select ")
				index++
			}
			build.WriteString(sql)
			args = append(args, arg...)
		}
	}
	return build.String(), args
}

func GetKeyerBatch() *KeyerBatch {
	return &KeyerBatch{}
} 