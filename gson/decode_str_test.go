//  Copyright (C) 晓白齐齐,版权所有.
package gson

import (
	"reflect"
	"testing"
)

// 本文件添加的测试样例：直接测试基本数据类型字符串的解析

var decodeStrData = []testBytesInfo{
	{"decode_str.boolv", "true", true, "", []testElemInfo{
		{&bool_1_1, "decode_str.boolv", []testVarInfo{
			{"decode_str.boolv", true, &bool_1_1, reflect.Bool, true},
		}},
	}},
	{"decode_str.boolgv", "false", true, "", []testElemInfo{
		{&bool_1_2, "decode_str.boolgv", []testVarInfo{
			{"decode_str.boolgv", false, &bool_1_2, reflect.Bool, true},
		}},
	}},
	{"decode_str.boolv2", "false", true, "", []testElemInfo{
		{&bool_2_1, "decode_str.boolv2", []testVarInfo{
			{"decode_str.boolv2", false, &bool_2_1, reflect.Bool, true},
		}},
	}},
	{"decode_str.boolgv2", "true", true, "", []testElemInfo{
		{&bool_2_2, "decode_str.boolgv2", []testVarInfo{
			{"decode_str.boolgv2", true, &bool_2_2, reflect.Bool, true},
		}},
	}},
	{"decode_str.intv", "35766645435", true, "", []testElemInfo{
		{&int_1, "decode_str.intv", []testVarInfo{
			{"decode_str.intv", 35766645435, &int_1, reflect.Int, true},
		}},
	}},
	{"decode_str.intgv", "-564567865678", true, "", []testElemInfo{
		{&int_2, "decode_str.intgv", []testVarInfo{
			{"decode_str.intgv", -564567865678, &int_2, reflect.Int, true},
		}},
	}},
	{"decode_str.int8v", "33", true, "", []testElemInfo{
		{&int8_1, "decode_str.int8v", []testVarInfo{
			{"decode_str.int8v", int8(33), &int8_1, reflect.Int8, true},
		}},
	}},
	{"decode_str.int8gv", "56", true, "", []testElemInfo{
		{&int8_2, "decode_str.int8gv", []testVarInfo{
			{"decode_str.int8gv", int8(56), &int8_2, reflect.Int8, true},
		}},
	}},
	{"decode_str.int16v", "-2333", true, "", []testElemInfo{
		{&int16_1, "decode_str.int16v", []testVarInfo{
			{"decode_str.int16v", int16(-2333), &int16_1, reflect.Int16, true},
		}},
	}},
	{"decode_str.int16gv", "32534", true, "", []testElemInfo{
		{&int16_2, "decode_str.int16gv", []testVarInfo{
			{"decode_str.int16gv", int16(32534), &int16_2, reflect.Int16, true},
		}},
	}},
	{"decode_str.int32v", "-3223233", true, "", []testElemInfo{
		{&int32_1, "decode_str.int32v", []testVarInfo{
			{"decode_str.int32v", int32(-3223233), &int32_1, reflect.Int32, true},
		}},
	}},
	{"decode_str.int32gv", "38273849", true, "", []testElemInfo{
		{&int32_2, "decode_str.int32gv", []testVarInfo{
			{"decode_str.int32gv", int32(38273849), &int32_2, reflect.Int32, true},
		}},
	}},
	{"decode_str.int64v", "9864357877878", true, "", []testElemInfo{
		{&int64_1, "decode_str.int64v", []testVarInfo{
			{"decode_str.int64v", int64(9864357877878), &int64_1, reflect.Int64, true},
		}},
	}},
	{"decode_str.int64gv", "-338746798757", true, "", []testElemInfo{
		{&int64_2, "decode_str.int64gv", []testVarInfo{
			{"decode_str.int64gv", int64(-338746798757), &int64_2, reflect.Int64, true},
		}},
	}},
	{"decode_str.uintv", "39845098374", true, "", []testElemInfo{
		{&uint_1, "decode_str.uintv", []testVarInfo{
			{"decode_str.uintv", uint(39845098374), &uint_1, reflect.Uint, true},
		}},
	}},
	{"decode_str.uintgv", "349876493829", true, "", []testElemInfo{
		{&uint_2, "decode_str.uintgv", []testVarInfo{
			{"decode_str.uintgv", uint(349876493829), &uint_2, reflect.Uint, true},
		}},
	}},
	{"decode_str.uint8v", "124", true, "", []testElemInfo{
		{&uint8_1, "decode_str.uint8v", []testVarInfo{
			{"decode_str.uint8v", uint8(124), &uint8_1, reflect.Uint8, true},
		}},
	}},
	{"decode_str.uint8gv", "239", true, "", []testElemInfo{
		{&uint8_2, "decode_str.uint8gv", []testVarInfo{
			{"decode_str.uint8gv", uint8(239), &uint8_2, reflect.Uint8, true},
		}},
	}},
	{"decode_str.uint16v", "1258", true, "", []testElemInfo{
		{&uint16_1, "decode_str.uint16v", []testVarInfo{
			{"decode_str.uint16v", uint16(1258), &uint16_1, reflect.Uint16, true},
		}},
	}},
	{"decode_str.uint16gv", "34534", true, "", []testElemInfo{
		{&uint16_2, "decode_str.uint16gv", []testVarInfo{
			{"decode_str.uint16gv", uint16(34534), &uint16_2, reflect.Uint16, true},
		}},
	}},
	{"decode_str.uint32v", "234543832", true, "", []testElemInfo{
		{&uint32_1, "decode_str.uint32v", []testVarInfo{
			{"decode_str.uint32v", uint32(234543832), &uint32_1, reflect.Uint32, true},
		}},
	}},
	{"decode_str.uint32gv", "373837482", true, "", []testElemInfo{
		{&uint32_2, "decode_str.uint32gv", []testVarInfo{
			{"decode_str.uint32gv", uint32(373837482), &uint32_2, reflect.Uint32, true},
		}},
	}},
	{"decode_str.uint64v", "23847349384939", true, "", []testElemInfo{
		{&uint64_1, "decode_str.uint64v", []testVarInfo{
			{"decode_str.uint64v", uint64(23847349384939), &uint64_1, reflect.Uint64, true},
		}},
	}},
	{"decode_str.uint64gv", "989348493849", true, "", []testElemInfo{
		{&uint64_2, "decode_str.uint64gv", []testVarInfo{
			{"decode_str.uint64gv", uint64(989348493849), &uint64_2, reflect.Uint64, true},
		}},
	}},
	{"decode_str.float32v", "98.678", true, "", []testElemInfo{
		{&float32_1, "decode_str.float32v", []testVarInfo{
			{"decode_str.float32v", float32(98.678), &float32_1, reflect.Float32, true},
		}},
	}},
	{"decode_str.float32gv", "2.567e2", true, "", []testElemInfo{
		{&float32_2, "decode_str.float32gv", []testVarInfo{
			{"decode_str.float32gv", float32(256.7), &float32_2, reflect.Float32, true},
		}},
	}},
	{"decode_str.float64v", "23e123", true, "", []testElemInfo{
		{&float64_1, "decode_str.float64v", []testVarInfo{
			{"decode_str.float64v", 23e123, &float64_1, reflect.Float64, true},
		}},
	}},
	{"decode_str.float64gv", "592.3e-32", true, "", []testElemInfo{
		{&float64_2, "decode_str.float64gv", []testVarInfo{
			{"decode_str.float64gv", 592.3e-32, &float64_2, reflect.Float64, true},
		}},
	}},
	{"decode_str.stringv", "23eE123", true, "", []testElemInfo{
		{&string_1, "decode_str.stringv", []testVarInfo{
			{"decode_str.stringv", "23eE123", &string_1, reflect.String, true},
		}},
	}},
	{"decode_str.stringgv", "592.3Ee-32", true, "", []testElemInfo{
		{&string_2, "decode_str.stringgv", []testVarInfo{
			{"decode_str.stringgv", "592.3Ee-32", &string_2, reflect.String, true},
		}},
	}},
}

// 测试UnmarshalAny，测试一些普通数据类型字符串的测试样例
func Test_UnmarshalAny(t *testing.T) {
	resetBasicData()
	testUnmarshaAny(t, decodeStrKey, decodeStrData)
}

func testUnmarshaAny(t *testing.T, decodeAct string, decodeData []testBytesInfo) {
	byteWhiteList, ok := whiteList[decodeAct]
	if !allTest && (!ok || byteWhiteList == nil || len(byteWhiteList) == 0) {
		return
	}
	for _, value := range decodeData {
		byteKey := value.key
		var elemWhiteList map[string][]string
		if !allTest {
			elemWhiteList, ok = byteWhiteList[byteKey]
			if !ok || elemWhiteList == nil || len(elemWhiteList) == 0 {
				continue
			}
		}
		elems := value.elems
		for _, elem := range elems {
			elemKey := elem.key
			var varWhiteList []string
			if !allTest {
				varWhiteList, ok = elemWhiteList[elemKey]
				if !ok || varWhiteList == nil || len(varWhiteList) == 0 {
					continue
				}
			}
			byteStr := value.byteStr
			tarVar := elem.elemVar
			err := UnmarshalAny([]byte(byteStr), tarVar)
			checkTestError(t, err, value.isErrNil, value.errStr)
			vars := elem.vars
			for _, varValue := range vars {
				varKey := varValue.key
				if !allTest {
					if !isInArr(varKey, varWhiteList) {
						continue
					}
				}
				keyName := decodeStrKey + ":" + byteKey + ":" + elemKey + ":" + varKey
				checkBasicResult(t, keyName, varValue.expected, varValue.got, varValue.basicType, varValue.gotPointer)
			}
		}
	}
}
