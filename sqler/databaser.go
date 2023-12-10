package sqler

import (
	"github.com/bqqsrc/errer"
	"database/sql"

	"github.com/bqqsrc/loger"
)

func rows2Map(rows *sql.Rows, keyTable map[string]string) (map[string]interface{}, error) {
	funcName := "Rows2Map"
	maps, err := rows2Maps(1, rows, keyTable)
	loger.Debugf("%s maps: %v\nerr: %s\n", funcName, maps, err)
	if err != nil {
		return nil, err
	}
	if maps == nil || len(maps) == 0 {
		return nil, nil
	}
	return maps[0], nil
}

func rows2MapWithKeys(rows *sql.Rows, keys []string) (map[string]interface{}, error) {
	funcName := "Rows2Map"
	maps, err := rows2MapsWithKeys(1, rows, keys)
	loger.Debugf("%s maps: %v\nerr: %s\n", funcName, maps, err)
	if err != nil {
		return nil, err
	}
	if maps == nil || len(maps) == 0 {
		return nil, nil
	}
	return maps[0], nil
}

func rows2Maps(num int, rows *sql.Rows, keyTable map[string]string) ([]map[string]interface{}, error) {
	funcName := "rows2Maps"
	defer rows.Close()
	if rows == nil {
		return nil, errer.CallerErr(funcName, "rows is nil")
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, errer.CallerErr(funcName, "rows.Columns error, err: %s", err)
	}
	count := len(columns)
	loger.Debugf("%s count: %d\ncolumns: %v\n", funcName, count, columns)
	var values = make([]interface{}, count)
	for i, _ := range values {
		var valueI interface{}
		values[i] = &valueI
	}
	ret := make([]map[string]interface{}, 0)
	index := 0
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, errer.CallerErr(funcName, "rows.Scan error, err: %s", err)
		}
		rawValue := make(map[string]interface{})
		for index, colName := range columns {
			//var tmpValue = *(values[index].(*interface{}))
			// tmp1 := values[index].(*interface{})
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
			value := *(values[index].(*interface{}))
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
	loger.Debugf("%s ret: %v\n", funcName, ret)
	return ret, nil
}

func rows2MapsWithKeys(num int, rows *sql.Rows, keys []string) ([]map[string]interface{}, error) {
	funcName := "rows2Maps"
	defer rows.Close()
	if rows == nil {
		return nil, errer.CallerErr(funcName, "rows is nil")
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, errer.CallerErr(funcName, "rows.Columns error, err: %s", err)
	}
	count := len(columns)
	loger.Debugf("%s count: %d\ncolumns: %v\n", funcName, count, columns)
	var values = make([]interface{}, count)
	for i, _ := range values {
		var valueI interface{}
		values[i] = &valueI
	}
	ret := make([]map[string]interface{}, 0)
	index := 0
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, errer.CallerErr(funcName, "rows.Scan error, err: %s", err)
		}
		rawValue := make(map[string]interface{})
		for index, key := range keys {
			value := *(values[index].(*interface{}))
			switch value.(type) {
			case []uint8:
				bytes := value.([]byte)
				valueStr := string(bytes)
				rawValue[key] = valueStr
				break
			default:
				rawValue[key] = value
				break

			}
		}
		// for index, colName := range columns {
		// 	//var tmpValue = *(values[index].(*interface{}))
		// 	// tmp1 := values[index].(*interface{})
		// 	// tmp2 := *tmp1
		// 	// switch tmp2.(type) {
		// 	// case int:
		// 	// 	loger.Debug("is int %d", tmp2.(int))
		// 	// 	break
		// 	// case string:
		// 	// 	loger.Debug("is string %s", tmp2.(string))
		// 	// 	break
		// 	// case int64:
		// 	// 	loger.Debug("is int64 %d", tmp2.(int64))
		// 	// 	break
		// 	// case float64:
		// 	// 	loger.Debug("is float64 %d", tmp2.(float64))
		// 	// 	break
		// 	// default:
		// 	// 	loger.Debug("is %s", reflect.TypeOf(tmp2))
		// 	// 	break

		// 	// }
		// 	value := *(values[index].(*interface{}))
		// 	if newKey, ok := keyTable[colName]; ok {
		// 		colName = newKey
		// 	}
		// 	//tmptmpValue := tmpValue.([]byte)
		// 	//tmptmptmpValue := string(tmptmpValue)
		// 	//rawValue[colName] = tmptmptmpValue
		// 	//rawValue[colName] = tmpValue
		// 	switch value.(type) {
		// 	case []uint8:
		// 		bytes := value.([]byte)
		// 		valueStr := string(bytes)
		// 		rawValue[colName] = valueStr
		// 		break
		// 	default:
		// 		rawValue[colName] = value
		// 		break

		// 	}
		// }
		ret = append(ret, rawValue)
		index++
		if num > 0 && index >= num {
			return ret, nil
		}
	}
	loger.Debugf("%s ret: %v\n", funcName, ret)
	return ret, nil
}

func Rows2Array(rows *sql.Rows, keyTable map[string]string) ([]string, [][]interface{}, error) {
	funcName := "Rows2Array"
	defer rows.Close()
	if rows == nil {
		return nil, nil, errer.CallerErr(funcName, "rows is nil")
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, errer.CallerErr(funcName, "rows.Columns error, err: %s", err)
	}
	count := len(columns)
	header := make([]string, count)
	for index, value := range columns {
		if tmpValue, ok := keyTable[value]; ok {
			header[index] = tmpValue
		} else {
			header[index] = value
		}
	}
	var values = make([]interface{}, count)
	for i, _ := range values {
		var valueI interface{}
		values[i] = &valueI
	}
	ret := make([][]interface{}, 0)
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, nil, errer.CallerErr(funcName, "rows.Scan error, err: %s", err)
		}
		rawValue := make([]interface{}, count)
		for index, value := range values {
			rawValue[index] = *(value.(*interface{}))
		}
		ret = append(ret, rawValue)
	}
	return header, ret, nil
}
