package json

import (
	"fmt"
	"path"
	"testing"
)

var JsonEtcPath = "./test_data"

func TestMyJsonObjcet(t *testing.T) {
	var myJson MyJsonObject // map[string]interface{}
	myJson = make(MyJsonObject)
	myJson["key1"] = 8
	myJson["key2"] = "test"
	myJson["key3"] = false
	tempJson := make(MyJsonObject)
	tempJson["key1"] = 88
	tempJson["key2"] = []string{"test1", "test2", "test3"}
	myJson["key4"] = tempJson
	if myJson["key1"] != 8 {
		t.Fatalf("TestMyJsonObjcet error, value of key1 is not 8")
	}
	if fmt.Sprintf("%T", myJson["key1"]) != "int" {
		t.Fatalf("TestMyJsonObjcet error, type of key1 is not int")
	}
	if myJson["key2"] != "test" {
		t.Fatalf("TestMyJsonObjcet error, value of key2 is not test")
	}
	if fmt.Sprintf("%T", myJson["key2"]) != "string" {
		t.Fatalf("TestMyJsonObjcet error, type of key2 is not string")
	}
	if myJson["key3"] != false {
		t.Fatalf("TestMyJsonObjcet error, value of key3 is not false")
	}
	if fmt.Sprintf("%T", myJson["key3"]) != "bool" {
		t.Fatalf("TestMyJsonObjcet error, type of key3 is not bool")
	}
	if fmt.Sprintf("%T", myJson["key4"]) != "json_util.MyJsonObject" {
		t.Fatalf("TestMyJsonObjcet error, type of key4 is not json_util.MyJsonObject")
	}
}

func Test_GetJsonFromFile_GetType(t *testing.T) {
	jsonFile := path.Join(JsonEtcPath, "config.json")
	result := GetJsonFromFile(jsonFile)
	checkJsonForGetType(result, t)
}

func Test_GetJsonFromStr_GetType(t *testing.T) {
	jsonStr := `{
        "IntKey1": 8003,
        "IntKey2": 3306, 
        "IntKey3": -7,
        "IntKey4": 0,
        "FloatKey1": 9.4,
        "FloatKey2": 0,
        "FloatKey3": -10.0,
        "StrKey1": "testValue1",
        "StrKey2": "",
        "StrKey3": "123",
        "StrKey4": "-1010",
        "StrKey5": "7.999",
        "ChineseStrKey1": "中文测试",
        "中文键": "中文值",
        "boolKey1": false,
        "boolKey2": true,
        "array1": [],
        "array2": ["value1", "value2", "value3"],
        "array3": [1, 2, 7.9, 8.8],
        "array4": [true, false, true, true],
        "array5": [1, 2, 3, 4, 5],
        "json1": {},
        "json2": {
            "key1": 1,
            "key2": false,
            "key3": "value1"
        }
    }`
	result := GetJsonFromStr(jsonStr)
	checkJsonForGetType(result, t)
}

func checkJsonForGetType(result MyJsonObject, t *testing.T) {
	//Test GetFloat64
	if fmt.Sprintf("%T", result["IntKey1"]) != "float64" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of IntKey1 is not float64")
	}
	if result["IntKey1"] != 8003.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of IntKey1 is not 8003.0")
	}
	if fmt.Sprintf("%T", result["IntKey3"]) != "float64" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of IntKey3 is not float64")
	}
	if result["IntKey3"] != -7.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of IntKey3 is not -7.0")
	}
	var float64Value float64
	float64Value = GetFloat64(result, "IntKey2")
	if fmt.Sprintf("%T", float64Value) != "float64" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetFloat64 IntKey2 is not float64")
	}
	if float64Value != 3306.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 IntKey2 is not 3306.0")
	}
	float64Value = GetFloat64(result, "IntKey3")
	if float64Value != -7.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 IntKey3 is not -7.0")
	}
	float64Value = GetFloat64(result, "IntKey4")
	if float64Value != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 IntKey4 is not 0")
	}
	float64Value = GetFloat64(result, "FloatKey1")
	if float64Value != 9.4 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 FloatKey1 is not 9.4")
	}
	float64Value = GetFloat64(result, "FloatKey2")
	if float64Value != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 FloatKey2 is not 0")
	}
	float64Value = GetFloat64(result, "FloatKey3")
	if float64Value != -10.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 FloatKey3 is not -10.0")
	}
	float64Value = GetFloat64(result, "StrKey2")
	if float64Value != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 StrKey2 is not 0")
	}
	float64Value = GetFloat64(result, "StrKey3")
	if float64Value != 123.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 StrKey3 is not 123.0")
	}
	float64Value = GetFloat64(result, "StrKey4")
	if float64Value != -1010.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 StrKey4 is not -1010.0")
	}
	float64Value = GetFloat64(result, "StrKey5")
	if float64Value != 7.999 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 StrKey5 is not 7.999")
	}
	float64Value = GetFloat64(result, "boolKey1")
	if float64Value != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 boolKey1 is not 0")
	}
	float64Value = GetFloat64(result, "boolKey2")
	if float64Value != 1.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetFloat64 boolKey2 is not 1.0")
	}

	//Test GetString
	if fmt.Sprintf("%T", result["StrKey1"]) != "string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of StrKey1 is not string")
	}
	if result["StrKey1"] != "testValue1" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of StrKey1 is not testValue1")
	}
	if fmt.Sprintf("%T", result["StrKey2"]) != "string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of StrKey2 is not string")
	}
	if result["StrKey2"] != "" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of StrKey2 is not emptyString")
	}
	if fmt.Sprintf("%T", result["ChineseStrKey1"]) != "string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of ChineseStrKey1 is not string")
	}
	if result["ChineseStrKey1"] != "中文测试" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of ChineseStrKey1 is not 中文测试")
	}
	if fmt.Sprintf("%T", result["中文键"]) != "string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of 中文键 is not string")
	}
	if result["中文键"] != "中文值" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of 中文键 is not 中文值")
	}
	var stringValue string
	stringValue = GetString(result, "IntKey1")
	if fmt.Sprintf("%T", stringValue) != "string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetString IntKey1 is not string")
	}
	if stringValue != "8003" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString IntKey1 is not string 8003")
	}
	stringValue = GetString(result, "IntKey3")
	if stringValue != "-7" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString IntKey3 is not string -7.0")
	}
	stringValue = GetString(result, "IntKey4")
	if stringValue != "0" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString IntKey4 is not string 0")
	}
	stringValue = GetString(result, "FloatKey1")
	if stringValue != "9.4" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString FloatKey1 is not string 9.4")
	}
	stringValue = GetString(result, "FloatKey3")
	if stringValue != "-10" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString FloatKey3 is not string -10.0")
	}
	stringValue = GetString(result, "StrKey1")
	if stringValue != "testValue1" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString StrKey1 is not testValue1")
	}
	stringValue = GetString(result, "StrKey2")
	if stringValue != "" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString StrKey2 is not empty string")
	}
	stringValue = GetString(result, "StrKey3")
	if stringValue != "123" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString StrKey3 is not string 123")
	}
	stringValue = GetString(result, "StrKey4")
	if stringValue != "-1010" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString StrKey4 is not string -1010")
	}
	stringValue = GetString(result, "ChineseStrKey1")
	if stringValue != "中文测试" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString ChineseStrKey1 is not string 中文测试")
	}
	stringValue = GetString(result, "中文键")
	if stringValue != "中文值" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString 中文键 is not string 中文值")
	}
	stringValue = GetString(result, "boolKey1")
	if stringValue != "false" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString boolKey1 is not string false")
	}
	stringValue = GetString(result, "boolKey2")
	if stringValue != "true" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetString boolKey2 is not string true")
	}

	//Test GetInt
	var intValue int
	intValue = GetInt(result, "IntKey1")
	if fmt.Sprintf("%T", intValue) != "int" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetInt IntKey1 is not int")
	}
	if intValue != 8003 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt IntKey1 is not 8003")
	}
	intValue = GetInt(result, "IntKey3")
	if intValue != -7 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt IntKey3 is not -7")
	}
	intValue = GetInt(result, "IntKey4")
	if intValue != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt IntKey4 is not 0")
	}
	intValue = GetInt(result, "FloatKey1")
	if intValue != 9 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt FloatKey1 is not 9")
	}
	intValue = GetInt(result, "FloatKey3")
	if intValue != -10 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt FloatKey3 is not -10")
	}
	intValue = GetInt(result, "StrKey2")
	if intValue != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt StrKey2 is not 0")
	}
	intValue = GetInt(result, "StrKey3")
	if intValue != 123 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt StrKey3 is not 123")
	}
	intValue = GetInt(result, "StrKey4")
	if intValue != -1010 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt StrKey4 is not -1010")
	}
	intValue = GetInt(result, "boolKey1")
	if intValue != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt boolKey1 is not 0")
	}
	intValue = GetInt(result, "boolKey2")
	if intValue != 1 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetInt boolKey2 is not 1")
	}

	//Test GetBool
	if fmt.Sprintf("%T", result["boolKey1"]) != "bool" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of boolKey1 is not bool")
	}
	if result["boolKey1"] != false {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of boolKey1 is not false")
	}
	var boolValue bool
	boolValue = GetBool(result, "IntKey1")
	if fmt.Sprintf("%T", boolValue) != "bool" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetBool IntKey1 is not bool")
	}
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool IntKey1 is not true")
	}
	boolValue = GetBool(result, "IntKey3")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool IntKey3 is not true")
	}
	boolValue = GetBool(result, "IntKey4")
	if boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool IntKey4 is not false")
	}
	boolValue = GetBool(result, "FloatKey1")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool FloatKey1 is not true")
	}
	boolValue = GetBool(result, "FloatKey3")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool FloatKey3 is not true")
	}
	boolValue = GetBool(result, "StrKey1")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool StrKey1 is not true")
	}
	boolValue = GetBool(result, "StrKey2")
	if boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool StrKey2 is not false")
	}
	boolValue = GetBool(result, "中文键")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool 中文键 is not true")
	}
	boolValue = GetBool(result, "boolKey1")
	if boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool boolKey1 is not false")
	}
	boolValue = GetBool(result, "boolKey2")
	if !boolValue {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool boolKey2 is not true")
	}

	//Test GetJson
	if fmt.Sprintf("%T", result["json2"]) != "map[string]interface {}" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of json2 is not []interface {}")
	}
	if fmt.Sprintf("%T", GetJson(result, "json2")) != "json_util.MyJsonObject" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetJson json2 is not json_util.MyJsonObject")
	}
	jsonValue := GetJson(result, "json2")
	if jsonValue["key1"] != 1.0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of jsonValue key1 is not 1.0")
	}
	if jsonValue["key3"] != "value1" {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of jsonValue key3 is not value1")
	}
	if GetBool(jsonValue, "key2") {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool jsonValue key2 is not false")
	}
	jsonValue = GetJson(result, "json1")
	if len(jsonValue) != 0 {
		t.Fatalf("TestGetJsonFromFile_GetType error, value of GetBool json1 is not empty")
	}

	//Test GetStrinngArr
	if fmt.Sprintf("%T", result["array2"]) != "[]interface {}" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of array2 is not []interface {}")
	}
	stringArr := GetStringArr(result, "array2")
	if fmt.Sprintf("%T", stringArr) != "[]string" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetStrinngArr array2 is not []string")
	}
	for i := 0; i < len(stringArr); i++ {
		compareStr := fmt.Sprintf("value%d", i+1)
		if stringArr[i] != compareStr {
			t.Fatalf("TestGetJsonFromFile_GetType error, type of GetStrinngArr array2 of %d is not %s", i, compareStr)
		}
	}

	//Test GetIntArr
	if fmt.Sprintf("%T", result["array5"]) != "[]interface {}" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of array5 is not []interface {}")
	}
	intArr := GetIntArr(result, "array5")
	if fmt.Sprintf("%T", intArr) != "[]int" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetIntArr array5 is not []int")
	}
	for i := 0; i < len(intArr); i++ {
		compareInt := i + 1
		if intArr[i] != compareInt {
			t.Fatalf("TestGetJsonFromFile_GetType error, type of GetIntArr array5 of %d is not %d", i, compareInt)
		}
	}
	intArr = GetIntArr(result, "array3")
	compareInt := []int{1, 2, 7, 8}
	for i := 0; i < len(intArr); i++ {
		if intArr[i] != compareInt[i] {
			t.Fatalf("TestGetJsonFromFile_GetType error, type of GetIntArr array3 of %d is not %d", i, compareInt[i])
		}
	}

	//Test GetBoolArr
	if fmt.Sprintf("%T", result["array4"]) != "[]interface {}" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of array4 is not []interface {}")
	}
	boolArr := GetBoolArr(result, "array4")
	if fmt.Sprintf("%T", boolArr) != "[]bool" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetBoolArr array4 is not []bool")
	}
	compareBool := []bool{true, false, true, true}
	for i := 0; i < len(boolArr); i++ {
		if boolArr[i] != compareBool[i] {
			t.Fatalf("TestGetJsonFromFile_GetType error, type of GetBoolArr array4 of %d is not %t", i, compareBool[i])
		}
	}

	//Test GetFloat64Arr
	if fmt.Sprintf("%T", result["array3"]) != "[]interface {}" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of array3 is not []interface {}")
	}
	float64Arr := GetFloat64Arr(result, "array3")
	if fmt.Sprintf("%T", float64Arr) != "[]float64" {
		t.Fatalf("TestGetJsonFromFile_GetType error, type of GetFloat64Arr array3 is not []float64")
	}
	compareArr := []float64{1, 2, 7.9, 8.8}
	for i := 0; i < len(float64Arr); i++ {
		if float64Arr[i] != compareArr[i] {
			t.Fatalf("TestGetJsonFromFile_GetType error, type of GetFloat64Arr array3 of %d is not %f", i, compareArr[i])
		}
	}
}
