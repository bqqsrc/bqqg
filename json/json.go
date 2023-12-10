package json

import (
	jsonn "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type MyJsonObject map[string]interface{}

func GetJsonFromFile(jsonFile string) MyJsonObject {
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("read %s error: %s", jsonFile, err.Error())
	}
	var result MyJsonObject

	if err := jsonn.Unmarshal(data, &result); err != nil {
		log.Fatalf("Unmarshal data err, error is %s", err.Error())
	}
	return result
}

func GetJsonFromStr(jsonStr string) MyJsonObject {
	var result MyJsonObject
	if err := jsonn.Unmarshal([]byte(jsonStr), &result); err != nil {
		log.Fatalf("Unmarshal data err, error is %s", err.Error())
	}
	return result
}

func GetInt(jsonObject MyJsonObject, key string) int {
	valueFloat, ok := jsonObject[key].(float64)
	if ok {
		return int(valueFloat)
	}
	valueBool, ok := jsonObject[key].(bool)
	if ok {
		if valueBool {
			return 1
		} else {
			return 0
		}
	}
	valueStr, ok := jsonObject[key].(string)
	if ok {
		if valueStr == "" {
			return 0
		}
		result, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("can't convert jsonObject of key: %s, value: %s to int", key, valueStr)
		} else {
			return result
		}
	}
	log.Fatalf("can't convert jsonObject of key: %s, value: %s to int", key, valueStr)
	return 0
}

func GetBool(jsonObject MyJsonObject, key string) bool {
	valueBool, ok := jsonObject[key].(bool)
	if ok {
		return valueBool
	}
	valueFloat, ok := jsonObject[key].(float64)
	if ok {
		if valueFloat == 0 {
			return false
		} else {
			return true
		}
	}
	valueStr, ok := jsonObject[key].(string)
	if ok {
		if valueStr == "" {
			return false
		} else {
			return true
		}
	}
	log.Fatalf("can't convert jsonObject of key: %s, value: %s to bool", key, valueStr)
	return false
}

func GetString(jsonObject MyJsonObject, key string) string {
	valueStr, ok := jsonObject[key].(string)
	if ok {
		return valueStr
	}
	valueFloat64, ok := jsonObject[key].(float64)
	if ok {
		return fmt.Sprintf("%v", valueFloat64)
	}
	valueBool, ok := jsonObject[key].(bool)
	if ok {
		return fmt.Sprintf("%t", valueBool)
	}
	return ""
}

func GetFloat64(jsonObject MyJsonObject, key string) float64 {
	valueFloat, ok := jsonObject[key].(float64)
	if ok {
		return valueFloat
	}
	valueBool, ok := jsonObject[key].(bool)
	if ok {
		if valueBool {
			return 1
		} else {
			return 0
		}
	}
	valueStr, ok := jsonObject[key].(string)
	if ok {
		if valueStr == "" {
			return 0
		}
		result, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			log.Fatalf("can't convert jsonObject of key: %s, value: %s to float64", key, valueStr)
		} else {
			return result
		}
	}
	log.Fatalf("can't convert jsonObject of key: %s, value: %s to float64", key, valueStr)
	return 0.0
}

func GetJson(jsonObject MyJsonObject, key string) MyJsonObject {
	resultObject, ok := jsonObject[key].(map[string]interface{})
	if ok {
		return resultObject
	}
	log.Fatalf("can't convert jsonObject of key: %s to MyJsonObject", key)
	return nil
}

func GetStringArr(jsonObject MyJsonObject, key string) []string {
	resultArr, ok := jsonObject[key].([]interface{})
	if ok {
		var result []string
		for i := 0; i < len(resultArr); i++ {
			temp, tempOK := resultArr[i].(string)
			if tempOK {
				result = append(result, temp)
			} else {
				log.Fatalf("can't convert jsonObject of key: %s to []string", key)
			}
		}
		return result
	}
	log.Fatalf("can't convert jsonObject of key: %s to []string", key)
	return nil
}

func GetIntArr(jsonObject MyJsonObject, key string) []int {
	resultArr, ok := jsonObject[key].([]interface{})
	if ok {
		var result []int
		for i := 0; i < len(resultArr); i++ {
			temp, tempOK := resultArr[i].(float64)
			if tempOK {
				result = append(result, int(temp))
			} else {
				log.Fatalf("can't convert jsonObject of key: %s to []int", key)
			}
		}
		return result
	}
	log.Fatalf("can't convert jsonObject of key: %s to []int", key)
	return nil
}

func GetBoolArr(jsonObject MyJsonObject, key string) []bool {
	resultArr, ok := jsonObject[key].([]interface{})
	if ok {
		var result []bool
		for i := 0; i < len(resultArr); i++ {
			temp, tempOK := resultArr[i].(bool)
			if tempOK {
				result = append(result, temp)
			} else {
				log.Fatalf("can't convert jsonObject of key: %s to []bool", key)
			}
		}
		return result
	}
	log.Fatalf("can't convert jsonObject of key: %s to []bool", key)
	return nil
}

func GetFloat64Arr(jsonObject MyJsonObject, key string) []float64 {
	resultArr, ok := jsonObject[key].([]interface{})
	if ok {
		var result []float64
		for i := 0; i < len(resultArr); i++ {
			temp, tempOK := resultArr[i].(float64)
			if tempOK {
				result = append(result, temp)
			} else {
				log.Fatalf("can't convert jsonObject of key: %s to []float64", key)
			}
		}
		return result
	}
	log.Fatalf("can't convert jsonObject of key: %s to []float64", key)
	return nil
}
