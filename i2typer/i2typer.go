package i2typer
import (
	"reflect"
	"fmt"
	"time"
	"strconv"
)

// 定义转换时发生的错误类型
type I2TError struct {
	Thrower string  // 抛出异常的函数名
	format string  // 异常信息和参数
	args []interface{}	
}

// 输出错误信息
func (e *I2TError) Error() string {return fmt.Sprintf(e.format, e.args...)}

func newI2Error(caller, format string, args ...interface{}) error {
	return &I2TError{Thrower: caller, format: format, args: args}
}

// 将interface{}转换为int
//
// 支持将实际类型为bool、int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为int
//
// string类型会调用strconv相关接口直接转换int，
// []byte类型则会先转换为字符串，再转换为int
//
// 如果data为非数字字符串或[]byte先转换为非数字字符串时，则报错
//
// time.Time会转换为对应的时间戳数值，另外如果想要获取时间戳数值，建议调用tm2.I2TimeStamp获取
func I2Int(data interface{}) (int, error) {
	caller := "I2Int"
	var str string
	var dataType string
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return ret, nil
	case bool:
		ret, _ := data.(bool)
		if ret {
			return 1, nil 
		} else {
			return 0, nil
		}
	case int64:
		ret, _ := data.(int64)
		return int(ret), nil
	case float32:
		ret, _ := data.(float32)
		return int(ret), nil
	case float64:
		ret, _ := data.(float64)
		return int(ret), nil
	case byte:
		ret, _ := data.(byte)
		return int(ret), nil
	case time.Time:
		ret := data.(time.Time)
		return int(ret.Unix()), nil
	case []byte:
		str = bytes2String(data)
		dataType = "[]byte"
		break
	case string:
		str = data.(string)
		dataType = "string"
		break
	default:
		return 0, newI2Error(caller, "data-type is %s, only support type: bool, int, int64, float32, float64, type, []byte, string, timeTime",  reflect.TypeOf(data))
	}
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0, newI2Error(caller, "data-type is %s, and call strconv.Atoi(%s) error: %s", dataType, str, err)
	} else {
		return result, nil
	}
}

// 将interface{}转换为int64
//
// 支持将实际类型为bool、int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为int64
//
// string类型会调用strconv相关接口直接转换int64，
// []byte类型则会先转换为字符串，再转换为int64
//
// 如果data为非数字字符串或[]byte先转换为非数字字符串时，则报错
//
// time.Time会转换为对应的时间戳数值，另外如果想要获取时间戳数值，建议调用tm2.I2TimeStamp获取
func I2Int64(data interface{}) (int64, error) {
	caller := "I2Int64"
	var str string
	var dataType string
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return int64(ret), nil
	case bool:
		ret, _ := data.(bool)
		if ret {
			return 1, nil 
		} else {
			return 0, nil
		}
	case int64:
		ret, _ := data.(int64)
		return ret, nil
	case float32:
		ret, _ := data.(float32)
		return int64(ret), nil
	case float64:
		ret, _ := data.(float64)
		return int64(ret), nil
	case byte:
		ret, _ := data.(byte)
		return int64(ret), nil
	case time.Time:
		ret := data.(time.Time)
		return ret.Unix(), nil
	case []byte:
		str = bytes2String(data)
		dataType = "[]byte"
		break
	case string:
		str = data.(string)
		dataType = "string"
		break
	default:
		return 0, newI2Error(caller, "data-type is %s, only support type: bool, int, int64, float32, float64, type, []byte, string, timeTime",  reflect.TypeOf(data))
	}
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, newI2Error(caller, "data-type is %s, and call strconv.ParseInt(%s, 10, 64) error: %s", dataType, str, err)
	} else {
		return result, nil
	}
}

// 将interface{}转换为float64
//
// 支持将实际类型为bool、int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为float64
//
// string类型会调用strconv相关接口直接转换float64，
// []byte类型则会先转换为字符串，再转换为float64
//
// 如果data为非数字字符串或[]byte先转换为非数字字符串时，则报错
//
// time.Time会转换为对应的时间戳数值，另外如果想要获取时间戳数值，建议调用tm2.I2TimeStamp获取
func I2Float64(data interface{}) (float64, error) {
	caller := "I2Float64"
	var str string
	var dataType string
	switch data.(type) {
	case int:
		ret, _ := data.(int)
		return float64(ret), nil
	case bool:
		ret, _ := data.(bool)
		if ret {
			return 1, nil 
		} else {
			return 0, nil
		}
	case int64:
		ret, _ := data.(int64)
		return float64(ret), nil
	case float32:
		ret, _ := data.(float32)
		return float64(ret), nil
	case float64:
		ret, _ := data.(float64)
		return ret, nil
	case byte:
		ret, _ := data.(byte)
		return float64(ret), nil
	case time.Time:
		ret := data.(time.Time)
		return float64(ret.Unix()), nil
	case []byte:
		str = bytes2String(data)
		dataType = "[]byte"
		break
	case string:
		str = data.(string)
		dataType = "string"
		break
	default:
		return 0, newI2Error(caller, "data-type is %s, only support type: bool, int, int64, float32, float64, type, []byte, string, timeTime",  reflect.TypeOf(data))
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, newI2Error(caller, "data-type is %s, and call strconv.ParseFloat(%s, 64) error: %s", dataType, str, err)
	} else {
		return result, nil
	}
}

// 将interface{}转换为string
//
// 支持将所有类型转换为string
//
// time.Time会转换为对应的时间字符串，另外，如果想要获取时间字符串、日期字符串或者日期时间字符串，建议调用tm2.I2DateStr、tm2.I2TimeStr、 tm2.I2DateTimeStr获取
//
// 其它类型通过fmt直接格式化为字符串
func I2String(data interface{}) string {
	//caller := "I2String"
	switch data.(type) {
	case string:
		ret, _ := data.(string)
		return ret 
	case time.Time:
		return time2String(data)
	case []byte:
		return bytes2String(data)
	default:
		ret := fmt.Sprintf("%v", data) 
		return ret
	}
}

func bytes2String(data interface{}) string {
	value := data.([]byte)
	return string(value)
}

func time2String(data interface{}) string {
	tm, _ := data.(time.Time)
	ret := tm.Format("2006-01-02 15:04:05")
	return ret
}