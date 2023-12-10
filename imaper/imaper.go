package imaper

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/bqqsrc/bqqg/log"
)

func GetStringFromMaps(key string, mapArgs ...map[string]interface{}) (string, error) {
	for _, currentMap := range mapArgs {
		if currentMap != nil && currentMap[key] != nil {
			if result, ok := currentMap[key].(string); ok {
				return result, nil
			} else {
				return "", errors.New("value convert to string error")
			}
		}
	}
	return "", errors.New("key not found")
}

func GetIntFromMaps(key string, mapArgs ...map[string]interface{}) (int, error) {
	for _, currentMap := range mapArgs {
		if currentMap != nil && currentMap[key] != nil {
			if result, ok := currentMap[key].(int); ok {
				return result, nil
			} else {
				return 0, errors.New("value convert to int error")
			}
		}
	}
	return 0, errors.New("key not found")
}

func GetMapFromMaps(key string, mapArgs ...map[string]interface{}) (map[string]interface{}, error) {
	for _, currentMap := range mapArgs {
		if currentMap != nil && currentMap[key] != nil {
			if result, ok := currentMap[key].(map[string]interface{}); ok {
				return result, nil
			} else {
				return nil, errors.New("value convert to map[string]interface{} error")
			}
		}
	}
	return nil, errors.New("key not found")
}

func GetMap(mapData map[string]interface{}, keyArgs ...string) (map[string]interface{}, error) {
	if mapData != nil {
		for _, value := range keyArgs {
			if mapData[value] != nil {
				if result, ok := mapData[value].(map[string]interface{}); ok {
					return result, nil // true, true, result
				} else {
					return nil, errors.New("all value convert to map[string]interface{} error") // true, false, nil
				}
			}
		}
	}
	return nil, errors.New("all key not found") // false, true, nil
}

func GetMapList(key string, data map[string]interface{}) []map[string]interface{} {
	if keyIf, ok := data[key]; ok {
		if value, result := keyIf.([]interface{}); result {
			tmpResult := make([]map[string]interface{}, 0)
			for _, tmpValue := range value {
				if tmpMap, result := tmpValue.(map[string]interface{}); result {
					tmpResult = append(tmpResult, tmpMap)
				}
			}
			return tmpResult
		}
	}
	return nil
}

func MustInt(key string, data map[string]interface{}, defaultValue int) int {
	ret, ok := data[key]
	if !ok {
		return defaultValue
	}
	retInt, ok := I2Int(ret)
	if !ok {
		return defaultValue
	}
	return retInt
}

func MustInt64(key string, data map[string]interface{}, defaultValue int64) int64 {
	ret, ok := data[key]
	if !ok {
		return defaultValue
	}
	retInt, ok := I2Int64(ret)
	if !ok {
		return defaultValue
	}
	return retInt
}

func MustFloat64(key string, data map[string]interface{}, defaultValue float64) float64 {
	ret, ok := data[key]
	if !ok {
		return defaultValue
	}
	retFloat, ok := I2Float(ret)
	if !ok {
		return defaultValue
	}
	return retFloat
}

func MustTimeStamp(key string, data map[string]interface{}, defaultValue int64) int64 {
	dateData, ok := data[key]
	if !ok {
		return defaultValue
	}
	log.Debugf("key: %s, data: %v, defaultValue: %d", key, data, defaultValue)
	var dateStr string
	switch dateData.(type) {
	case int64:
		ret, _ := dateData.(int64)
		log.Debugf("ret is %d", ret)
		return ret
	case int:
		value, _ := dateData.(int)
		ret := int64(value)
		log.Debugf("ret is %d", ret)
		return ret
	case string:
		dateStr = dateData.(string)
		break
	case []byte:
		value := dateData.([]byte)
		dateStr = string(value)
		break
	case time.Time:
		value := dateData.(time.Time)
		return value.Unix()
	default:
		log.Errorf("Error, MustTimeStamp error, a %s value can't convert to TimeStamp", reflect.TypeOf(dateData))
		return defaultValue
	}
	if dateStr == "" {
		return defaultValue
	}
	log.Debugf("dateStr is %s", dateStr)
	return String2TimeStamp(dateStr)
}

func I2DateTimeStr(data interface{}) (string, bool) {
	var timeStamp int64
	var dateStr string
	strType := true
	switch data.(type) {
	case int64:
		timeStamp, _ = data.(int64)
		strType = false
		break
	case int:
		value, _ := data.(int)
		timeStamp = int64(value)
		strType = false
		break
	case string:
		dateStr = data.(string)
		break
	case []byte:
		value := data.([]byte)
		dateStr = string(value)
		break
	default:
		log.Errorf("Error, I2DateTimeStr error, a %s value can't convert to TimeStamp", reflect.TypeOf(data))
		return "", false
	}
	var tm time.Time
	if strType {
		var err error
		tm, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			log.Errorf("MustDateTimeStr error, time.Parse(\"2006-01-02 15:04:05\", %s) error, err: %s", dateStr, err)
			ok := false
			timeStamp, ok = I2Int64(dateStr)
			if ok {
				strType = false
			} else {
				return "", false
			}
		}
	}
	if !strType {
		tm = time.Unix(timeStamp, 0)
	}
	return tm.Format("2006-01-02 15:04:05"), true
}

func MustDateTimeStr(key string, data map[string]interface{}, defaultValue string) string {
	dateData, ok := data[key]
	if !ok {
		return defaultValue
	}
	if value, ok := I2DateTimeStr(dateData); !ok {
		return defaultValue
	} else {
		return value
	}
	// var timeStamp int64
	// var dateStr string
	// strType := true
	// switch dateData.(type) {
	// case int64:
	// 	timeStamp, _ = dateData.(int64)
	// 	strType = false
	// 	break
	// case int:
	// 	value, _ := dateData.(int)
	// 	timeStamp = int64(value)
	// 	strType = false
	// 	break
	// case string:
	// 	dateStr = dateData.(string)
	// 	break
	// case []byte:
	// 	value := dateData.([]byte)
	// 	dateStr = string(value)
	// 	break
	// default:
	// 	log.Errorf("Error, MustTimeStamp error, a %s value can't convert to TimeStamp", reflect.TypeOf(dateData))
	// 	return defaultValue
	// }
	// var tm time.Time
	// if strType {
	// 	var err error
	// 	tm, err = time.Parse("2006-01-02 15:04:05", dateStr)
	// 	if err != nil {
	// 		log.Errorf("MustDateTimeStr error, time.Parse(\"2006-01-02 15:04:05\", %s) error, err: %s", dateStr, err)
	// 		return defaultValue
	// 	}
	// } else {
	// 	tm = time.Unix(timeStamp, 0)
	// }
	// return tm.Format("2006-01-02 15:04:05")
}

// 转为东八区时间差
var cstoffset int64 = 28800

func DataTimeAddDate(dateTime string, years, month, days int) string {
	tm, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		log.Errorf("MustDateTimeStr error, time.Parse(\"2006-01-02 15:04:05\", %s) error, err: %s", dateTime, err)
		return ""
	}
	tm = tm.AddDate(years, month, days)
	return tm.Format("2006-01-02 15:04:05")
}

func TimeStampAddDate(timeStamp int64, years, month, days int) string {
	// t := time.Now()
	// name, offset := t.Zone()
	// log.Debugf("TimeStampAddDate name %s, offset %d", name, offset)
	// //	var nsec int64 = int64(offset) * 3600
	tm := time.Unix(timeStamp-cstoffset, 0)
	tm = tm.AddDate(years, month, days)
	return tm.Format("2006-01-02 15:04:05")
}

func MustString(key string, data map[string]interface{}, defaultValue string) string {
	ret, ok := data[key]
	if !ok {
		return defaultValue
	}
	retString, ok := I2String(ret)
	if !ok {
		return defaultValue
	}
	return retString
}

func String2TimeStamp(dateTime string) int64 {
	log.Debugf("dateTime %s", dateTime)
	times, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		log.Errorf("MustTimeStamp error, time.Parse(\"2006-01-02 15:04:05\", %s) error, err: %s", dateTime, err)
		return 0
	}
	return times.Unix()
}

func I2Int(data interface{}) (int, bool) {
	funcName := "I2Int"
	log.Debugf("%s data: %s\n", funcName, data)
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return ret, true
	case int64:
		ret, _ := data.(int64)
		return int(ret), true
	case string:
		ret, _ := data.(string)
		result, err := strconv.Atoi(ret) //ParseInt(ret, 10, 0)
		if err != nil {
			log.Errorf("Error: %s, strconv.Atoi(%s) error: %s", funcName, ret, err)
			return 0, false
		} else {
			return result, true
		}
	case float32:
		ret, _ := data.(float32)
		return int(ret), true
	case float64:
		ret, _ := data.(float64)
		return int(ret), true
	default:
		log.Errorf("Error: %s(%s) error, not support data.(type) %s",
			funcName, data, reflect.TypeOf(data))
		return 0, false
	}
}

func I2Int64(data interface{}) (int64, bool) {
	funcName := "I2Int"
	log.Debugf("%s data: %s\n", funcName, data)
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return int64(ret), true
	case int64:
		ret, _ := data.(int64)
		return ret, true
	case string:
		ret, _ := data.(string)
		result, err := strconv.ParseInt(ret, 10, 64) //strconv.Atoi(ret) //ParseInt(ret, 10, 0)
		if err != nil {
			log.Errorf("Error: %s, strconv.Atoi(%s) error: %s", funcName, ret, err)
			return 0, false
		} else {
			return result, true
		}
	case float32:
		ret, _ := data.(float32)
		return int64(ret), true
	case float64:
		ret, _ := data.(float64)
		return int64(ret), true
	default:
		log.Errorf("Error: %s(%s) error, not support data.(type) %s",
			funcName, data, reflect.TypeOf(data))
		return 0, false
	}
}

func I2Float(data interface{}) (float64, bool) {
	funcName := "I2Float"
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return float64(ret), true
	case int64:
		ret, _ := data.(int64)
		return float64(ret), true
	case float64:
		ret, _ := data.(float64)
		return ret, true
	case float32:
		ret, _ := data.(float32)
		return float64(ret), true
	case string:
		ret, _ := data.(string)
		result, err := strconv.ParseFloat(ret, 64)
		if err != nil {
			log.Errorf("Error:%s, strconv.ParseFloat(%s, 64) error: %s", funcName, ret, err)
			return 0, false
		} else {
			return result, true
		}
	default:
		log.Errorf("Error: %s(%s) error, not support data.(type) %s",
			funcName, data, reflect.TypeOf(data))
		return 0, false
	}
}

func I2String(data interface{}) (string, bool) {
	funcName := "I2String"
	log.Debugf("%s data is %s", funcName, data)
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return strconv.Itoa(ret), true
	case int64:
		ret, _ := data.(int64)
		return strconv.FormatInt(ret, 10), true
	case float32:
		ret, _ := data.(float32)
		return strconv.FormatFloat(float64(ret), 'f', 0, 32), true
	case float64:
		ret, _ := data.(float64)
		return strconv.FormatFloat(ret, 'f', 0, 64), true
	case string:
		ret, _ := data.(string)
		return ret, true
	case time.Time:
		tm, _ := data.(time.Time)
		ret := tm.Format("2006-01-02 15:04:05")
		return ret, true
	default:
		log.Errorf("Error: %s(%s) error, not support data.(type) %s",
			funcName, data, reflect.TypeOf(data))
		return "", false
	}
}

func Float2String(f float64, perc int) (string, bool) {
	return strconv.FormatFloat(f, 'f', perc, 64), true
}
