package databasehelper

import (
	//	"bqqgc/errer"
	"database/sql"
	// "io/ioutil"
	// "os"
	// "strings"
	// "github.com/bqqsrc/bqqg/sqlfmt"
	// "github.com/bqqsrc/bqqg/loger"
	// "log"
	// "github.com/bqqsrc/bqqg/database"
)

func rows2Maps(num int, rows *sql.Rows, keyTable map[string]string) ([]map[string]any, error) {
	funcName := "rows2Maps"
	defer rows.Close()
	if rows == nil {
		return nil, CallerErr(funcName, "rows is nil")
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, CallerErr(funcName, "rows.Columns error, err: %s", err)
	}
	count := len(columns)
	// loger.Debugf("%s count: %d\ncolumns: %v\n", funcName, count, columns)
	var values = make([]any, count)
	for i, _ := range values {
		var valueI any
		values[i] = &valueI
	}
	ret := make([]map[string]any, 0)
	index := 0
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, CallerErr(funcName, "rows.Scan error, err: %s", err)
		}
		rawValue := make(map[string]any)
		for index, colName := range columns {
			//var tmpValue = *(values[index].(*any))
			// tmp1 := values[index].(*any)
			// tmp2 := *tmp1
			// switch tmp2.(type) {
			// case int:
			// 	loger.Debug("is int %d", tmp2.(int))
			// 	break
			// case string:
			// 	loger.Debug("is string %s", tmp2.(string))
			// 	break
			// case int64:
			// 	loger.Debug("is int64 %d", tmp2.(int64))
			// 	break
			// case float64:
			// 	loger.Debug("is float64 %d", tmp2.(float64))
			// 	break
			// default:
			// 	loger.Debug("is %s", reflect.TypeOf(tmp2))
			// 	break

			// }
			value := *(values[index].(*any))
			if newKey, ok := keyTable[colName]; ok {
				colName = newKey
			}
			//tmptmpValue := tmpValue.([]byte)
			//tmptmptmpValue := string(tmptmpValue)
			//rawValue[colName] = tmptmptmpValue
			//rawValue[colName] = tmpValue
			switch value.(type) {
			case []uint8:
				bytes := value.([]byte)
				valueStr := string(bytes)
				rawValue[colName] = valueStr
				break
			default:
				rawValue[colName] = value
				break

			}
		}
		ret = append(ret, rawValue)
		index++
		if num > 0 && index >= num {
			return ret, nil
		}
	}
	// loger.Debugf("%s ret: %v\n", funcName, ret)
	return ret, nil
}
