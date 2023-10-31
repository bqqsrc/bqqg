//  Copyright (C) 晓白齐齐,版权所有.
package gson

import (
	"io/ioutil"
	"reflect"
	"testing"
)

// 本文件添加的测试样例：直接测试基本数据类型字符串的解析

var decodeFileData = []testBytesInfo{
	{decodeGsonFile, "", true, "", []testElemInfo{
		{&bool_1_1, "bool_1_1", []testVarInfo{
			{"bool_1_1", true, &bool_1_1, reflect.Bool, true},
		}},
		{&bool_1_2, "bool_1_2", []testVarInfo{
			{"bool_1_2", false, &bool_1_2, reflect.Bool, true},
		}},
		{&bool_2_1, "bool_2_1", []testVarInfo{
			{"bool_2_1", false, &bool_2_1, reflect.Bool, true},
		}},
		{&bool_2_2, "bool_2_2", []testVarInfo{
			{"bool_2_2", true, &bool_2_2, reflect.Bool, true},
		}},
		{&float32_1, "float32_1", []testVarInfo{
			{"float32_1", float32(59.999), &float32_1, reflect.Float32, true},
		}},
		{&float32_2, "float32_2", []testVarInfo{
			{"float32_2", float32(89.77), &float32_2, reflect.Float32, true},
		}},
		{&float64_1, "float64_1", []testVarInfo{
			{"float64_1", 340000.0, &float64_1, reflect.Float64, true},
		}},
		{&float64_2, "float64_2", []testVarInfo{
			{"float64_2", 98.7, &float64_2, reflect.Float64, true},
		}},
		{&int8_1, "int8_1", []testVarInfo{
			{"int8_1", int8(-33), &int8_1, reflect.Int8, true},
		}},
		{&int8_2, "int8_2", []testVarInfo{
			{"int8_2", int8(58), &int8_2, reflect.Int8, true},
		}},
		{&int16_1, "int16_1", []testVarInfo{
			{"int16_1", int16(456), &int16_1, reflect.Int16, true},
		}},
		{&int16_2, "int16_2", []testVarInfo{
			{"int16_2", int16(-567), &int16_2, reflect.Int16, true},
		}},
		{&int32_1, "int32_1", []testVarInfo{
			{"int32_1", int32(703522), &int32_1, reflect.Int32, true},
		}},
		{&int32_2, "int32_2", []testVarInfo{
			{"int32_2", int32(-89076), &int32_2, reflect.Int32, true},
		}},
		{&int64_1, "int64_1", []testVarInfo{
			{"int64_1", int64(325235298967), &int64_1, reflect.Int64, true},
		}},
		{&int64_2, "int64_2", []testVarInfo{
			{"int64_2", int64(-98799889678), &int64_2, reflect.Int64, true},
		}},
		{&int_1, "int_1", []testVarInfo{
			{"int_1", 39870935321, &int_1, reflect.Int, true},
		}},
		{&int_2, "int_2", []testVarInfo{
			{"int_2", 23458987898, &int_2, reflect.Int, true},
		}},
		{&uint8_1, "uint8_1", []testVarInfo{
			{"uint8_1", uint8(27), &uint8_1, reflect.Uint8, true},
		}},
		{&uint8_2, "uint8_2", []testVarInfo{
			{"uint8_2", uint8(35), &uint8_2, reflect.Uint8, true},
		}},
		{&uint16_1, "uint16_1", []testVarInfo{
			{"uint16_1", uint16(356), &uint16_1, reflect.Uint16, true},
		}},
		{&uint16_2, "uint16_2", []testVarInfo{
			{"uint16_2", uint16(598), &uint16_2, reflect.Uint16, true},
		}},
		{&uint32_1, "uint32_1", []testVarInfo{
			{"uint32_1", uint32(588522), &uint32_1, reflect.Uint32, true},
		}},
		{&uint32_2, "uint32_2", []testVarInfo{
			{"uint32_2", uint32(82376), &uint32_2, reflect.Uint32, true},
		}},
		{&uint64_1, "uint64_1", []testVarInfo{
			{"uint64_1", uint64(325906298967), &uint64_1, reflect.Uint64, true},
		}},
		{&uint64_2, "uint64_2", []testVarInfo{
			{"uint64_2", uint64(98757889678), &uint64_2, reflect.Uint64, true},
		}},
		{&uint_1, "uint_1", []testVarInfo{
			{"uint_1", uint(39870967321), &uint_1, reflect.Uint, true},
		}},
		{&uint_2, "uint_2", []testVarInfo{
			{"uint_2", uint(23453487898), &uint_2, reflect.Uint, true},
		}},
		{&string_1, "string_1", []testVarInfo{
			{"string_1", "baiqiqi5", &string_1, reflect.String, true},
		}},
		{&string_2, "string_2", []testVarInfo{
			{"string_2", "baiqiqi3", &string_2, reflect.String, true},
		}},
		{&testdata_1_3, "testdata_1_3", []testVarInfo{
			{"TestStr", "baiqiqi3", &testdata_1_3.TestStr, reflect.String, true},
			{"TestInt", -3599, &testdata_1_3.TestInt, reflect.Int, true},
			{"TestFloat", 99.7, &testdata_1_3.TestFloat, reflect.Float64, true},
			{"TestBool", true, &testdata_1_3.TestBool, reflect.Bool, true},
			{"TestUint", uint(235), &testdata_1_3.TestUint, reflect.Uint, true},
		}},
		{&testdata_2_3, "testdata_2_3", []testVarInfo{
			{"test_str", "baiqiqi4", &testdata_2_3.TestStr, reflect.String, true},
			{"test_int", -9088, &testdata_2_3.TestInt, reflect.Int, true},
			{"test_float", 99.53, &testdata_2_3.TestFloat, reflect.Float64, true},
			{"test_bool", false, &testdata_2_3.TestBool, reflect.Bool, true},
			{"test_uint", uint(234), &testdata_2_3.TestUint, reflect.Uint, true},
		}},
		{&testdata_2_4, "testdata_2_4", []testVarInfo{
			{"TestStr", "baiqiqi5", &testdata_2_4.TestStr, reflect.String, true},
			{"TestInt", -19088, &testdata_2_4.TestInt, reflect.Int, true},
			{"TestFloat", 991.53, &testdata_2_4.TestFloat, reflect.Float64, true},
			{"TestBool", false, &testdata_2_4.TestBool, reflect.Bool, true},
			{"TestUint", uint(2234), &testdata_2_4.TestUint, reflect.Uint, true},
		}},
		{&testdata_2_5, "testdata_2_5", []testVarInfo{
			{"TestStr", "baiqiqi6", &testdata_2_5.TestStr, reflect.String, true},
			{"test_int", -188, &testdata_2_5.TestInt, reflect.Int, true},
			{"test_float", -991.53, &testdata_2_5.TestFloat, reflect.Float64, true},
			{"TestBool", false, &testdata_2_5.TestBool, reflect.Bool, true},
			{"TestUint", uint(1134), &testdata_2_5.TestUint, reflect.Uint, true},
		}},
		{&testdata_3_3, "testdata_3_3", []testVarInfo{
			{"test_str", "baiqiqi3_3", &testdata_3_3.TestStr, reflect.String, true},
			{"test_int", -988, &testdata_3_3.TestInt, reflect.Int, true},
			{"TestFloat", 99.3, &testdata_3_3.TestFloat, reflect.Float64, true},
			{"test_bool", true, &testdata_3_3.TestBool, reflect.Bool, true},
			{"TestUint", uint(34), &testdata_3_3.TestUint, reflect.Uint, true},
		}},
		{&testdata_3_4, "testdata_3_4", []testVarInfo{
			{"TestStr", "baiqiqi3_4", &testdata_3_4.TestStr, reflect.String, true},
			{"TestInt", -198, &testdata_3_4.TestInt, reflect.Int, true},
			{"TestFloat", 91.53, &testdata_3_4.TestFloat, reflect.Float64, true},
			{"TestBool", false, &testdata_3_4.TestBool, reflect.Bool, true},
			{"TestUint", uint(224), &testdata_3_4.TestUint, reflect.Uint, true},
		}},
		{&testdata_3_5, "testdata_3_5", []testVarInfo{
			{"TestStr", "baiqiqi3_5", &testdata_3_5.TestStr, reflect.String, true},
			{"test_int", -1288, &testdata_3_5.TestInt, reflect.Int, true},
			{"TestFloat", 891.53, &testdata_3_5.TestFloat, reflect.Float64, true},
			{"TestBool", true, &testdata_3_5.TestBool, reflect.Bool, true},
			{"TestUint", uint(114), &testdata_3_5.TestUint, reflect.Uint, true},
		}},
		// 基本数据指针类型
		{&ppppppbool, "pbool", []testVarInfo{
			{"pbool", true, *****ppppppbool, reflect.Bool, true},
		}},
		{&pppppint, "pint", []testVarInfo{
			{"pint", 322334234234234, ****pppppint, reflect.Int, true},
		}},
		{&ppppint8, "pint8", []testVarInfo{
			{"pint8", int8(-122), ***ppppint8, reflect.Int8, true},
		}},
		{&pppint16, "pint16", []testVarInfo{
			{"pint16", int16(3223), **pppint16, reflect.Int16, true},
		}},
		{&ppint32, "pint32", []testVarInfo{
			{"pint32", int32(-123221312), *ppint32, reflect.Int32, true},
		}},
		{&pint64, "pint64", []testVarInfo{
			{"pint64", int64(12312214324234), pint64, reflect.Int64, true},
		}},
		{&puint, "puint", []testVarInfo{
			{"puint", uint(32315624535232), puint, reflect.Uint, true},
		}},
		{&puint8, "puint8", []testVarInfo{
			{"puint8", uint8(253), puint8, reflect.Uint8, true},
		}},
		{&puint16, "puint16", []testVarInfo{
			{"puint16", uint16(23453), puint16, reflect.Uint16, true},
		}},
		{&puint32, "puint32", []testVarInfo{
			{"puint32", uint32(2131252), puint32, reflect.Uint32, true},
		}},
		{&puint64, "puint64", []testVarInfo{
			{"puint64", uint64(12341252345345), puint64, reflect.Uint64, true},
		}},
		{&pfloat32, "pfloat32", []testVarInfo{
			{"pfloat32", float32(233.9332), pfloat32, reflect.Float32, true},
		}},
		{&pfloat64, "pfloat64", []testVarInfo{
			{"pfloat64", 11.234, pfloat64, reflect.Float64, true},
		}},
		{&pstring, "pstring", []testVarInfo{
			{"pstring", "pstringstring", pstring, reflect.String, true},
		}},
		// 数组类型
		{&int_arr, "int_arr", []testVarInfo{
			{"[0]", 3822, &int_arr[0], reflect.Int, true},
			{"[1]", 2832, &int_arr[1], reflect.Int, true},
			{"[2]", 82, &int_arr[2], reflect.Int, true},
			{"[3]", -32344, &int_arr[3], reflect.Int, true},
			{"[4]", 98323, &int_arr[4], reflect.Int, true},
			{"[5]", 21, &int_arr[5], reflect.Int, true},
			{"[6]", -23212, &int_arr[6], reflect.Int, true},
		}},
		{&float_arr, "float_arr", []testVarInfo{
			{"[0]", 93.33, &float_arr[0], reflect.Float64, true},
			{"[1]", 32.5, &float_arr[1], reflect.Float64, true},
			{"[2]", 1.0, &float_arr[2], reflect.Float64, true},
			{"[3]", 78.783, &float_arr[3], reflect.Float64, true},
			{"[4]", -322.23, &float_arr[4], reflect.Float64, true},
			{"[5]", 9.999, &float_arr[5], reflect.Float64, true},
			{"[6]", 0.0, &float_arr[6], reflect.Float64, true},
		}},
		{&string_arr, "string_arr", []testVarInfo{
			{"[0]", "baiqiqi1", &string_arr[0], reflect.String, true},
			{"[1]", "baiqiqi2", &string_arr[1], reflect.String, true},
			{"[2]", "baiqiqi3", &string_arr[2], reflect.String, true},
			{"[3]", "baiqiqi4", &string_arr[3], reflect.String, true},
			{"[4]", "baiqiqi5", &string_arr[4], reflect.String, true},
		}},
		{&bool_arr, "bool_arr", []testVarInfo{
			{"[0]", true, &bool_arr[0], reflect.Bool, true},
			{"[1]", false, &bool_arr[1], reflect.Bool, true},
			{"[2]", true, &bool_arr[2], reflect.Bool, true},
			{"[3]", true, &bool_arr[3], reflect.Bool, true},
			{"[4]", true, &bool_arr[4], reflect.Bool, true},
			{"[5]", false, &bool_arr[5], reflect.Bool, true},
			{"[6]", false, &bool_arr[6], reflect.Bool, true},
			{"[7]", true, &bool_arr[7], reflect.Bool, true},
			{"[8]", false, &bool_arr[8], reflect.Bool, true},
		}},
		{&uint_arr, "uint_arr", []testVarInfo{
			{"[0]", uint(342), &uint_arr[0], reflect.Uint, true},
			{"[1]", uint(352), &uint_arr[1], reflect.Uint, true},
			{"[2]", uint(133412), &uint_arr[2], reflect.Uint, true},
			{"[3]", uint(525322323), &uint_arr[3], reflect.Uint, true},
			{"[4]", uint(9234523), &uint_arr[4], reflect.Uint, true},
			{"[5]", uint(0), &uint_arr[5], reflect.Uint, true},
		}},
		// slice类型
		{&int_slice, "int_slice", nil},
		{&float_slice, "float_slice", nil},
		{&string_slice, "string_slice", nil},
		{&bool_slice, "bool_slice", nil},
		{&uint_slice, "uint_slice", nil},
		{&int_slice2, "int_slice2", nil},
		{&float_slice2, "float_slice2", nil},
		{&string_slice2, "string_slice2", nil},
		{&bool_slice2, "bool_slice2", nil},
		{&uint_slice2, "uint_slice2", nil},
		// 指针数组类型
		{&pint_arr, "pint_arr", nil},
		// 元素为指针的结构体
		{&point_struct, "point_struct", []testVarInfo{
			{"TestStr", "baiqiqi8", point_struct.TestStr, reflect.String, true},
			{"TestInt", -35998, point_struct.TestInt, reflect.Int, true},
			{"TestFloat", 99.77, point_struct.TestFloat, reflect.Float64, true},
			{"TestBool", true, point_struct.TestBool, reflect.Bool, true},
			{"TestUint", uint(2355), point_struct.TestUint, reflect.Uint, true},
		}},
		{&test_map_str_float, "test_map_str_float", nil},
		{&test_map_int_str, "test_map_int_str", nil},
		{&test_map_float_bool, "test_map_float_bool", nil},
		{&test_map_bool_uint, "test_map_bool_uint", nil},
		{&test_map_uint_int, "test_map_uint_int", nil},
		// 结构体嵌套结构体
		{&testdata_struct_struct, "testdata_struct_struct", []testVarInfo{
			{"test_str", "baiqiqi8", &testdata_struct_struct.TestStr, reflect.String, true},
			{"test_int", -35132, &testdata_struct_struct.TestInt, reflect.Int, true},
			{"test_float", 98.776, &testdata_struct_struct.TestFloat, reflect.Float64, true},
			{"test_bool", false, &testdata_struct_struct.TestBool, reflect.Bool, true},
			{"test_uint", uint(34156), &testdata_struct_struct.TestUint, reflect.Uint, true},
			{"test_struct.TestStr", "subbaiqiqi1", &testdata_struct_struct.TestStruct.TestStr, reflect.String, true},
			{"test_struct.TestInt", -332, &testdata_struct_struct.TestStruct.TestInt, reflect.Int, true},
			{"test_struct.TestFloat", -99.3, &testdata_struct_struct.TestStruct.TestFloat, reflect.Float64, true},
			{"test_struct.TestBool", false, &testdata_struct_struct.TestStruct.TestBool, reflect.Bool, true},
			{"test_struct.TestUint", uint(989), &testdata_struct_struct.TestStruct.TestUint, reflect.Uint, true},
		}},
		{&testdata_struct_struct_struct, "testdata_struct_struct_struct", []testVarInfo{
			{"test_str", "baiqiqi81", &testdata_struct_struct_struct.TestStr, reflect.String, true},
			{"test_int", -332, &testdata_struct_struct_struct.TestInt, reflect.Int, true},
			{"test_float", 98.7976, &testdata_struct_struct_struct.TestFloat, reflect.Float64, true},
			{"test_bool", true, &testdata_struct_struct_struct.TestBool, reflect.Bool, true},
			{"test_uint", uint(3423156), &testdata_struct_struct_struct.TestUint, reflect.Uint, true},
			{"test_struct1.TestStr", "subbaiqiqi112", &testdata_struct_struct_struct.TestStruct.TestStr, reflect.String, true},
			{"test_struct1.TestInt", 332156, &testdata_struct_struct_struct.TestStruct.TestInt, reflect.Int, true},
			{"test_struct1.TestFloat", 8999.3, &testdata_struct_struct_struct.TestStruct.TestFloat, reflect.Float64, true},
			{"test_struct1.TestBool", false, &testdata_struct_struct_struct.TestStruct.TestBool, reflect.Bool, true},
			{"test_struct1.TestUint", uint(981459), &testdata_struct_struct_struct.TestStruct.TestUint, reflect.Uint, true},

			{"TestStructNest.TestStr", "subbaiqiqi1", &testdata_struct_struct_struct.TestStructNest.TestStr, reflect.String, true},
			{"TestStructNest.TestInt", -33223, &testdata_struct_struct_struct.TestStructNest.TestInt, reflect.Int, true},
			{"TestStructNest.TestFloat", -991.13, &testdata_struct_struct_struct.TestStructNest.TestFloat, reflect.Float64, true},
			{"TestStructNest.TestBool", false, &testdata_struct_struct_struct.TestStructNest.TestBool, reflect.Bool, true},
			{"TestStructNest.TestUint", uint(989156), &testdata_struct_struct_struct.TestStructNest.TestUint, reflect.Uint, true},
			{"TestStructNest.test_struct.TestStr", "subsubbaiqiqi1", &testdata_struct_struct_struct.TestStructNest.TestStruct.TestStr, reflect.String, true},
			{"TestStructNest.test_struct.TestInt", -123332, &testdata_struct_struct_struct.TestStructNest.TestStruct.TestInt, reflect.Int, true},
			{"TestStructNest.test_struct.TestFloat", 99.3568645, &testdata_struct_struct_struct.TestStructNest.TestStruct.TestFloat, reflect.Float64, true},
			{"TestStructNest.test_struct.TestBool", true, &testdata_struct_struct_struct.TestStructNest.TestStruct.TestBool, reflect.Bool, true},
			{"TestStructNest.test_struct.TestUint", uint(981439), &testdata_struct_struct_struct.TestStructNest.TestStruct.TestUint, reflect.Uint, true},

			{"test_struct2.testFloat", float32(-99111.13), &testdata_struct_struct_struct.TestStructNest2.TestFloat, reflect.Float32, true},
			{"test_struct2.testSubStruct.TestStr", "baiqiqiqi", &testdata_struct_struct_struct.TestStructNest2.TestStruct.TestStr, reflect.String, true},
			{"test_struct2.testSubStruct.TestFloat", 99.7123, &testdata_struct_struct_struct.TestStructNest2.TestStruct.TestFloat, reflect.Float64, true},
		}},
		// 结构体嵌套map
		{&testdata_struct_map, "testdata_struct_map", nil},
		{&testdata_struct_arr, "testdata_struct_arr", nil},
		// 数组、切片嵌套结构体
		{&testdata_arr_struct, "testdata_arr_struct", []testVarInfo{
			{"[0].TestStr", "baiqiqi_arr_1", &testdata_arr_struct[0].TestStr, reflect.String, true},
			{"[0].test_int", -12878, &testdata_arr_struct[0].TestInt, reflect.Int, true},
			{"[0].TestFloat", 8973.53, &testdata_arr_struct[0].TestFloat, reflect.Float64, true},
			{"[0].TestBool", false, &testdata_arr_struct[0].TestBool, reflect.Bool, true},
			{"[0].TestUint", uint(145214), &testdata_arr_struct[0].TestUint, reflect.Uint, true},
			{"[1].test_str", "baiqiqi_arr_2", &testdata_arr_struct[1].TestStr, reflect.String, true},
			{"[1].test_int", -31288, &testdata_arr_struct[1].TestInt, reflect.Int, true},
			{"[1].test_float", 1231.533, &testdata_arr_struct[1].TestFloat, reflect.Float64, true},
			{"[1].test_bool", true, &testdata_arr_struct[1].TestBool, reflect.Bool, true},
			{"[1].test_uint", uint(13414), &testdata_arr_struct[1].TestUint, reflect.Uint, true},
			{"[2].TestStr", "baiqiqi_arr_3", &testdata_arr_struct[2].TestStr, reflect.String, true},
			{"[2].TestInt", 5623, &testdata_arr_struct[2].TestInt, reflect.Int, true},
			{"[2].TestFloat", -78491.53, &testdata_arr_struct[2].TestFloat, reflect.Float64, true},
			{"[2].TestBool", true, &testdata_arr_struct[2].TestBool, reflect.Bool, true},
			{"[2].TestUint", uint(32114), &testdata_arr_struct[2].TestUint, reflect.Uint, true},
		}},
		{&testdata_slice_struct, "testdata_slice_struct", nil},
		{&testdata_arr_map_str_str, "testdata_arr_map_str_str", nil},
		{&testdata_slice_map_float_str, "testdata_slice_map_float_str", nil},
		{&testdata_map_struct, "testdata_map_struct", nil},
		{&testdata_map_map, "testdata_map_map", nil},
		{&testdata_map_arr, "testdata_map_arr", nil},
		{&testdata_map_slice, "testdata_map_slice", nil},
	}},
}

var dataAfterDecodeFile map[string]map[string][]testVarInfo

//Tip 新增一个解析前无法获取下标时，需要添加到这里的列表
var allElemKeyAfterDecodeFile = map[string][]string{
	decodeGsonFile: {
		"int_slice",
		"float_slice",
		"string_slice",
		"bool_slice",
		"uint_slice",
		"int_slice2",
		"float_slice2",
		"string_slice2",
		"bool_slice2",
		"uint_slice2",
		"pint_arr",
		"test_map_str_float",
		"test_map_int_str",
		"test_map_float_bool",
		"test_map_bool_uint",
		"test_map_uint_int",
		"testdata_struct_map",
		"testdata_struct_arr",
		"testdata_slice_struct",
		"testdata_arr_map_str_str",
		"testdata_slice_map_float_str",
		"testdata_map_struct",
		"testdata_map_map",
		"testdata_map_arr",
		"testdata_map_slice",
	},
}

func setDataAfterDecodeFile(afterDecodeData map[string]map[string][]testVarInfo, byteKey, elemKey string) map[string]map[string][]testVarInfo {
	if afterDecodeData == nil {
		afterDecodeData = make(map[string]map[string][]testVarInfo)
	}
	if afterDecodeData[byteKey] == nil {
		afterDecodeData[byteKey] = make(map[string][]testVarInfo)
	}
	switch byteKey {
	case decodeGsonFile:
		switch elemKey {
		case "int_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", 3822, int_slice[0], reflect.Int, false},
				{"[1]", 2832, int_slice[1], reflect.Int, false},
				{"[2]", 82, int_slice[2], reflect.Int, false},
				{"[3]", -32344, int_slice[3], reflect.Int, false},
				{"[4]", 98323, int_slice[4], reflect.Int, false},
				{"[5]", 21, int_slice[5], reflect.Int, false},
				{"[6]", -23212, int_slice[6], reflect.Int, false},
				{"len(int_slice)", 7, len(int_slice), reflect.Int, false},
			}
		case "float_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", 93.33, float_slice[0], reflect.Float64, false},
				{"[1]", 32.5, float_slice[1], reflect.Float64, false},
				{"[2]", 1.0, float_slice[2], reflect.Float64, false},
				{"[3]", 78.783, float_slice[3], reflect.Float64, false},
				{"[4]", -322.23, float_slice[4], reflect.Float64, false},
				{"[5]", 9.999, float_slice[5], reflect.Float64, false},
				{"len(float_slice)", 6, len(float_slice), reflect.Int, false},
			}
		case "string_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", "baiqiqi1", string_slice[0], reflect.String, false},
				{"[1]", "baiqiqi2", string_slice[1], reflect.String, false},
				{"[2]", "baiqiqi3", string_slice[2], reflect.String, false},
				{"[3]", "baiqiqi4", string_slice[3], reflect.String, false},
				{"[4]", "baiqiqi5", string_slice[4], reflect.String, false},
				{"len(string_slice)", 5, len(string_slice), reflect.Int, false},
			}
		case "bool_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", true, bool_slice[0], reflect.Bool, false},
				{"[1]", false, bool_slice[1], reflect.Bool, false},
				{"[2]", true, bool_slice[2], reflect.Bool, false},
				{"[3]", true, bool_slice[3], reflect.Bool, false},
				{"[4]", true, bool_slice[4], reflect.Bool, false},
				{"[5]", false, bool_slice[5], reflect.Bool, false},
				{"[6]", false, bool_slice[6], reflect.Bool, false},
				{"[7]", true, bool_slice[7], reflect.Bool, false},
				{"[8]", false, bool_slice[8], reflect.Bool, false},
				{"len(bool_slice)", 9, len(bool_slice), reflect.Int, false},
			}
		case "uint_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", uint(342), uint_slice[0], reflect.Uint, false},
				{"[1]", uint(352), uint_slice[1], reflect.Uint, false},
				{"[2]", uint(133412), uint_slice[2], reflect.Uint, false},
				{"[3]", uint(525322323), uint_slice[3], reflect.Uint, false},
				{"[4]", uint(9234523), uint_slice[4], reflect.Uint, false},
				{"len(uint_slice)", 5, len(uint_slice), reflect.Int, false},
			}
		case "int_slice2":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", 3822, int_slice2[0], reflect.Int, false},
				{"[1]", 2832, int_slice2[1], reflect.Int, false},
				{"[2]", 82, int_slice2[2], reflect.Int, false},
				{"[3]", -32344, int_slice2[3], reflect.Int, false},
				{"[4]", 98323, int_slice2[4], reflect.Int, false},
				{"[5]", 21, int_slice2[5], reflect.Int, false},
				{"[6]", -23212, int_slice2[6], reflect.Int, false},
				{"len(int_slice2)", 7, len(int_slice2), reflect.Int, false},
			}
		case "float_slice2":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", 93.33, float_slice2[0], reflect.Float64, false},
				{"[1]", 32.5, float_slice2[1], reflect.Float64, false},
				{"[2]", 1.0, float_slice2[2], reflect.Float64, false},
				{"[3]", 78.783, float_slice2[3], reflect.Float64, false},
				{"[4]", -322.23, float_slice2[4], reflect.Float64, false},
				{"[5]", 9.999, float_slice2[5], reflect.Float64, false},
				{"len(float_slice2)", 6, len(float_slice2), reflect.Int, false},
			}
		case "string_slice2":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", "baiqiqi1", string_slice2[0], reflect.String, false},
				{"[1]", "baiqiqi2", string_slice2[1], reflect.String, false},
				{"[2]", "baiqiqi3", string_slice2[2], reflect.String, false},
				{"[3]", "baiqiqi4", string_slice2[3], reflect.String, false},
				{"[4]", "baiqiqi5", string_slice2[4], reflect.String, false},
				{"len(string_slice2)", 5, len(string_slice2), reflect.Int, false},
			}
		case "bool_slice2":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", true, bool_slice2[0], reflect.Bool, false},
				{"[1]", false, bool_slice2[1], reflect.Bool, false},
				{"[2]", true, bool_slice2[2], reflect.Bool, false},
				{"[3]", true, bool_slice2[3], reflect.Bool, false},
				{"[4]", true, bool_slice2[4], reflect.Bool, false},
				{"[5]", false, bool_slice2[5], reflect.Bool, false},
				{"[6]", false, bool_slice2[6], reflect.Bool, false},
				{"[7]", true, bool_slice2[7], reflect.Bool, false},
				{"[8]", false, bool_slice2[8], reflect.Bool, false},
				{"len(bool_slice2)", 9, len(bool_slice2), reflect.Int, false},
			}
		case "uint_slice2":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", uint(342), uint_slice2[0], reflect.Uint, false},
				{"[1]", uint(352), uint_slice2[1], reflect.Uint, false},
				{"[2]", uint(133412), uint_slice2[2], reflect.Uint, false},
				{"[3]", uint(525322323), uint_slice2[3], reflect.Uint, false},
				{"[4]", uint(9234523), uint_slice2[4], reflect.Uint, false},
				{"len(uint_slice2)", 5, len(uint_slice2), reflect.Int, false},
			}
		case "pint_arr":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0]", 3822, pint_arr[0], reflect.Int, true},
				{"[1]", 2832, pint_arr[1], reflect.Int, true},
				{"[2]", 82, pint_arr[2], reflect.Int, true},
				{"[3]", -32344, pint_arr[3], reflect.Int, true},
				{"[4]", 98323, pint_arr[4], reflect.Int, true},
				{"[5]", 21, pint_arr[5], reflect.Int, true},
				{"[6]", -23212, pint_arr[6], reflect.Int, true},
			}
		case "test_map_str_float":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{`["key1"]`, 998.3, test_map_str_float["key1"], reflect.Float64, false},
				{`["key2"]`, 997.6, test_map_str_float["key2"], reflect.Float64, false},
				{`["key3"]`, 99.7, test_map_str_float["key3"], reflect.Float64, false},
				{"len(test_map_str_float)", 3, len(test_map_str_float), reflect.Int, false},
			}
		case "test_map_int_str":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[1]", "baiqiqi1", test_map_int_str[1], reflect.String, false},
				{"[2]", "baiqiqi2", test_map_int_str[2], reflect.String, false},
				{"[3]", "baiqiqi3", test_map_int_str[3], reflect.String, false},
				{"[4]", "baiqiqi4", test_map_int_str[4], reflect.String, false},
				{"len(test_map_int_str)", 4, len(test_map_int_str), reflect.Int, false},
			}
		case "test_map_float_bool":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[98.787]", true, test_map_float_bool[float32(98.787)], reflect.Bool, false},
				{"[78.55]", false, test_map_float_bool[float32(78.55)], reflect.Bool, false},
				{"[-333.33]", false, test_map_float_bool[float32(-333.33)], reflect.Bool, false},
				{"[-31233.33]", true, test_map_float_bool[float32(-31233.33)], reflect.Bool, false},
				{"len(test_map_float_bool)", 4, len(test_map_float_bool), reflect.Int, false},
			}
		case "test_map_bool_uint":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[true]", uint32(3832), test_map_bool_uint[true], reflect.Uint32, false},
				{"[false]", uint32(256), test_map_bool_uint[false], reflect.Uint32, false},
				{"len(test_map_bool_uint)", 2, len(test_map_bool_uint), reflect.Int, false},
			}
		case "test_map_uint_int":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[9833]", int16(34), test_map_uint_int[uint32(9833)], reflect.Int16, false},
				{"[785]", int16(-56), test_map_uint_int[uint32(785)], reflect.Int16, false},
				{"[333]", int16(566), test_map_uint_int[uint32(333)], reflect.Int16, false},
				{"[2333]", int16(133), test_map_uint_int[uint32(2333)], reflect.Int16, false},
				{"len(test_map_uint_int)", 4, len(test_map_uint_int), reflect.Int, false},
			}
		case "testdata_struct_map":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"test_str", "baiqiqi_struct_map", testdata_struct_map.TestStr, reflect.String, false},
				{"test_int", -355632, testdata_struct_map.TestInt, reflect.Int, false},
				{"test_float", 98.1876, testdata_struct_map.TestFloat, reflect.Float64, false},
				{"test_bool", true, testdata_struct_map.TestBool, reflect.Bool, false},
				{"test_uint", uint(3156), testdata_struct_map.TestUint, reflect.Uint, false},
				{"TestMapFloatInt[983.367]", testdata_struct_map.TestMapFloatInt[float32(983.367)], int8(13), reflect.Int8, false},
				{"TestMapFloatInt[324.36]", testdata_struct_map.TestMapFloatInt[float32(324.36)], int8(-113), reflect.Int8, false},
				{"TestMapFloatInt[122.22]", testdata_struct_map.TestMapFloatInt[float32(122.22)], int8(-123), reflect.Int8, false},
				{"TestMapFloatInt[9.83]", testdata_struct_map.TestMapFloatInt[float32(9.83)], int8(127), reflect.Int8, false},
				{"len(testdata_struct_map.TestMapFloatInt)", 4, len(testdata_struct_map.TestMapFloatInt), reflect.Int, false},
				{`test_struct_map_str_uint["key1"]`, testdata_struct_map.TestMapStrUint["key1"], uint16(3523), reflect.Uint16, false},
				{`test_struct_map_str_uint["key2"]`, testdata_struct_map.TestMapStrUint["key2"], uint16(1323), reflect.Uint16, false},
				{`test_struct_map_str_uint["key3"]`, testdata_struct_map.TestMapStrUint["key3"], uint16(456), reflect.Uint16, false},
				{"len(testdata_struct_map.test_struct_map_str_uint)", 3, len(testdata_struct_map.TestMapStrUint), reflect.Int, false},
			}
		case "testdata_struct_arr":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"test_str", "baiqiqi5", testdata_struct_arr.TestStr, reflect.String, false},
				{"test_int", -98988, testdata_struct_arr.TestInt, reflect.Int, false},
				{"test_float", 98.776, testdata_struct_arr.TestFloat, reflect.Float64, false},
				{"test_bool", true, testdata_struct_arr.TestBool, reflect.Bool, false},
				{"test_uint", uint(3456), testdata_struct_arr.TestUint, reflect.Uint, false},
				{"test_str_arr[0]", "subbaiqiqi3", testdata_struct_arr.TestStrArr[0], reflect.String, false},
				{"test_str_arr[1]", "sub_arr", testdata_struct_arr.TestStrArr[1], reflect.String, false},
				{"test_str_arr[2]", "falses", testdata_struct_arr.TestStrArr[2], reflect.String, false},
				{"test_str_arr[3]", "true", testdata_struct_arr.TestStrArr[3], reflect.String, false},
				{"test_str_arr[4]", "983ee", testdata_struct_arr.TestStrArr[4], reflect.String, false},
				{"len(testdata_struct_arr.test_str_arr)", 5, len(testdata_struct_arr.TestStrArr), reflect.Int, false},
				{"test_int_arr[0]", 98, testdata_struct_arr.TestIntArr[0], reflect.Int, false},
				{"test_int_arr[1]", 143, testdata_struct_arr.TestIntArr[1], reflect.Int, false},
				{"test_int_arr[2]", -998, testdata_struct_arr.TestIntArr[2], reflect.Int, false},
				{"test_int_arr[3]", 983, testdata_struct_arr.TestIntArr[3], reflect.Int, false},
				{"len(testdata_struct_arr.test_int_arr)", 4, len(testdata_struct_arr.TestIntArr), reflect.Int, false},
				{"test_float_arr[0]", 983.77, testdata_struct_arr.TestFloatArr[0], reflect.Float64, false},
				{"test_float_arr[1]", 883.83, testdata_struct_arr.TestFloatArr[1], reflect.Float64, false},
				{"test_float_arr[2]", -93.323, testdata_struct_arr.TestFloatArr[2], reflect.Float64, false},
				{"test_float_arr[3]", 999.7, testdata_struct_arr.TestFloatArr[3], reflect.Float64, false},
				{"test_float_arr[4]", 783.3, testdata_struct_arr.TestFloatArr[4], reflect.Float64, false},
				{"len(testdata_struct_arr.test_float_arr)", 5, len(testdata_struct_arr.TestFloatArr), reflect.Int, false},
				{"test_bool_arr[0]", true, testdata_struct_arr.TestBoolArr[0], reflect.Bool, false},
				{"test_bool_arr[1]", true, testdata_struct_arr.TestBoolArr[1], reflect.Bool, false},
				{"test_bool_arr[2]", false, testdata_struct_arr.TestBoolArr[2], reflect.Bool, false},
				{"test_bool_arr[3]", true, testdata_struct_arr.TestBoolArr[3], reflect.Bool, false},
				{"test_bool_arr[4]", false, testdata_struct_arr.TestBoolArr[4], reflect.Bool, false},
				{"test_bool_arr[5]", false, testdata_struct_arr.TestBoolArr[5], reflect.Bool, false},
				{"test_bool_arr[6]", true, testdata_struct_arr.TestBoolArr[6], reflect.Bool, false},
				{"len(testdata_struct_arr.test_bool_arr)", 7, len(testdata_struct_arr.TestBoolArr), reflect.Int, false},
			}
		case "testdata_slice_struct":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0].test_str", "baiqiqi_slice_1", testdata_slice_struct[0].TestStr, reflect.String, false},
				{"[0].test_int", 25132, testdata_slice_struct[0].TestInt, reflect.Int, false},
				{"[0].test_float", -498.776, testdata_slice_struct[0].TestFloat, reflect.Float64, false},
				{"[0].test_bool", true, testdata_slice_struct[0].TestBool, reflect.Bool, false},
				{"[0].test_uint", uint(767156), testdata_slice_struct[0].TestUint, reflect.Uint, false},
				{"[0].test_struct.TestStr", "subbaiqiqi_slice_struct_1", testdata_slice_struct[0].TestStruct.TestStr, reflect.String, false},
				{"[0].test_struct.TestInt", 11332, testdata_slice_struct[0].TestStruct.TestInt, reflect.Int, false},
				{"[0].test_struct.TestFloat", 6599.38792, testdata_slice_struct[0].TestStruct.TestFloat, reflect.Float64, false},
				{"[0].test_struct.TestBool", false, testdata_slice_struct[0].TestStruct.TestBool, reflect.Bool, false},
				{"[0].test_struct.TestUint", uint(976223), testdata_slice_struct[0].TestStruct.TestUint, reflect.Uint, false},
				{"[1].test_str", "baiqiqi_slice_2", testdata_slice_struct[1].TestStr, reflect.String, false},
				{"[1].test_int", -3235132, testdata_slice_struct[1].TestInt, reflect.Int, false},
				{"[1].test_float", 9158.776, testdata_slice_struct[1].TestFloat, reflect.Float64, false},
				{"[1].test_bool", false, testdata_slice_struct[1].TestBool, reflect.Bool, false},
				{"[1].test_uint", uint(62156), testdata_slice_struct[1].TestUint, reflect.Uint, false},
				{"[1].test_struct.TestStr", "subbaiqiqi_slice_struct_2", testdata_slice_struct[1].TestStruct.TestStr, reflect.String, false},
				{"[1].test_struct.TestInt", -33472, testdata_slice_struct[1].TestStruct.TestInt, reflect.Int, false},
				{"[1].test_struct.TestFloat", -8899.3, testdata_slice_struct[1].TestStruct.TestFloat, reflect.Float64, false},
				{"[1].test_struct.TestBool", true, testdata_slice_struct[1].TestStruct.TestBool, reflect.Bool, false},
				{"[1].test_struct.TestUint", uint(176989), testdata_slice_struct[1].TestStruct.TestUint, reflect.Uint, false},
				{"len(testdata_slice_struct)", 2, len(testdata_slice_struct), reflect.Int, false},
			}
		case "testdata_arr_map_str_str":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{`[0]["key1"]`, "arr_map_1_baiqiqi1", testdata_arr_map_str_str[0]["key1"], reflect.String, false},
				{`[0]["key2"]`, "arr_map_1_baiqiqi2", testdata_arr_map_str_str[0]["key2"], reflect.String, false},
				{"len(testdata_arr_map_str_str[0])", 2, len(testdata_arr_map_str_str[0]), reflect.Int, false},
				{`[1]["key1"]`, "arr_map_2_baiqiqi1", testdata_arr_map_str_str[1]["key1"], reflect.String, false},
				{`[1]["key2"]`, "arr_map_2_baiqiqi2", testdata_arr_map_str_str[1]["key2"], reflect.String, false},
				{`[1]["key3"]`, "arr_map_2_baiqiqi3", testdata_arr_map_str_str[1]["key3"], reflect.String, false},
				{"len(testdata_arr_map_str_str[1])", 3, len(testdata_arr_map_str_str[1]), reflect.Int, false},
				{`[2]["key1"]`, "arr_map_3_baiqiqi1", testdata_arr_map_str_str[2]["key1"], reflect.String, false},
				{"len(testdata_arr_map_str_str[2])", 1, len(testdata_arr_map_str_str[2]), reflect.Int, false},
				{"len(testdata_arr_map_str_str)", 3, len(testdata_arr_map_str_str), reflect.Int, false},
			}
		case "testdata_slice_map_float_str":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[0][983.3]", "arr_map_1_983.3", testdata_slice_map_float_str[0][983.3], reflect.String, false},
				{"[0][135.55]", "arr_map_1_135.55", testdata_slice_map_float_str[0][135.55], reflect.String, false},
				{"len(testdata_slice_map_float_str[0])", 2, len(testdata_slice_map_float_str[0]), reflect.Int, false},
				{"[1][8823.32]", "arr_map_2_8823.32", testdata_slice_map_float_str[1][8823.32], reflect.String, false},
				{"len(testdata_slice_map_float_str[1])", 1, len(testdata_slice_map_float_str[1]), reflect.Int, false},
				{"[2][9183.3]", "arr_map_3_9183.3", testdata_slice_map_float_str[2][9183.3], reflect.String, false},
				{"[2][-134.32]", "arr_map_3_-134.32", testdata_slice_map_float_str[2][-134.32], reflect.String, false},
				{"[2][-1312.33]", "arr_map_3_-1312.33", testdata_slice_map_float_str[2][-1312.33], reflect.String, false},
				{"len(testdata_slice_map_float_str[2])", 3, len(testdata_slice_map_float_str[2]), reflect.Int, false},
				{"len(testdata_slice_map_float_str)", 3, len(testdata_slice_map_float_str), reflect.Int, false},
			}
		case "testdata_map_struct":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{`["key1"].TestStr`, "baiqiqi_map_struct_1", testdata_map_struct["key1"].TestStr, reflect.String, false},
				{`["key1"].TestInt`, -52323, testdata_map_struct["key1"].TestInt, reflect.Int, false},
				{`["key1"].TestFloat`, 8891.53, testdata_map_struct["key1"].TestFloat, reflect.Float64, false},
				{`["key1"].TestBool`, true, testdata_map_struct["key1"].TestBool, reflect.Bool, false},
				{`["key1"].TestUint`, uint(123), testdata_map_struct["key1"].TestUint, reflect.Uint, false},
				{`["key2"].test_str`, "baiqiqi_map_struct_2", testdata_map_struct["key2"].TestStr, reflect.String, false},
				{`["key2"].test_int`, 97455, testdata_map_struct["key2"].TestInt, reflect.Int, false},
				{`["key2"].TestFloat`, -675891.53, testdata_map_struct["key2"].TestFloat, reflect.Float64, false},
				{`["key2"].TestBool`, false, testdata_map_struct["key2"].TestBool, reflect.Bool, false},
				{`["key2"].test_uint`, uint(5623), testdata_map_struct["key2"].TestUint, reflect.Uint, false},
				{`["key3"].test_str`, "baiqiqi_map_struct_3", testdata_map_struct["key3"].TestStr, reflect.String, false},
				{`["key3"].test_int`, -1242, testdata_map_struct["key3"].TestInt, reflect.Int, false},
				{`["key3"].test_float`, 891.532342, testdata_map_struct["key3"].TestFloat, reflect.Float64, false},
				{`["key3"].test_bool`, true, testdata_map_struct["key3"].TestBool, reflect.Bool, false},
				{`["key3"].test_uint`, uint(956), testdata_map_struct["key3"].TestUint, reflect.Uint, false},
				{"len(testdata_map_struct)", 3, len(testdata_map_struct), reflect.Int, false},
			}
		case "testdata_map_map":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{`["key1"]["sub1key1"]`, -89.77, testdata_map_map["key1"]["sub1key1"], reflect.Float64, false},
				{`["key1"]["sub1key2"]`, 98.88, testdata_map_map["key1"]["sub1key2"], reflect.Float64, false},
				{`len(testdata_map_map["key1"])`, 2, len(testdata_map_map["key1"]), reflect.Int, false},
				{`["key2"]["sub2key1"]`, 983.9, testdata_map_map["key2"]["sub2key1"], reflect.Float64, false},
				{`["key2"]["sub2key2"]`, -903.33, testdata_map_map["key2"]["sub2key2"], reflect.Float64, false},
				{`len(testdata_map_map["key2"])`, 2, len(testdata_map_map["key2"]), reflect.Int, false},
				{`len(testdata_map_map)`, 2, len(testdata_map_map), reflect.Int, false},
			}
		case "testdata_map_arr":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{`["key1"][0]`, "key1_1", testdata_map_arr["key1"][0], reflect.String, false},
				{`["key1"][1]`, "key1_2", testdata_map_arr["key1"][1], reflect.String, false},
				{`["key1"][2]`, "key1_3", testdata_map_arr["key1"][2], reflect.String, false},
				{`["key2"][0]`, "key2_1", testdata_map_arr["key2"][0], reflect.String, false},
				{`["key2"][1]`, "key2_2", testdata_map_arr["key2"][1], reflect.String, false},
				{`["key2"][2]`, "key2_3", testdata_map_arr["key2"][2], reflect.String, false},
				{`["key3"][0]`, "key3_1", testdata_map_arr["key3"][0], reflect.String, false},
				{`["key3"][1]`, "key3_2", testdata_map_arr["key3"][1], reflect.String, false},
				{`["key3"][2]`, "key3_3", testdata_map_arr["key3"][2], reflect.String, false},
				{`len(testdata_map_arr)`, 3, len(testdata_map_arr), reflect.Int, false},
			}
		case "testdata_map_slice":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"[983.33][0]", 3423, testdata_map_slice[983.33][0], reflect.Int, false},
				{"[983.33][1]", -9786, testdata_map_slice[983.33][1], reflect.Int, false},
				{"[983.33][2]", 132, testdata_map_slice[983.33][2], reflect.Int, false},
				{`len(testdata_map_slice[983.33])`, 3, len(testdata_map_slice[983.33]), reflect.Int, false},
				{"[-891.323][0]", 13123, testdata_map_slice[-891.323][0], reflect.Int, false},
				{"[-891.323][1]", -98656, testdata_map_slice[-891.323][1], reflect.Int, false},
				{"[-891.323][2]", -67343, testdata_map_slice[-891.323][2], reflect.Int, false},
				{"[-891.323][3]", 56234, testdata_map_slice[-891.323][3], reflect.Int, false},
				{`len(testdata_map_slice[-891.323])`, 4, len(testdata_map_slice[-891.323]), reflect.Int, false},
				{"[13.44][0]", 980, testdata_map_slice[13.44][0], reflect.Int, false},
				{`len(testdata_map_slice[13.44])`, 1, len(testdata_map_slice[13.44]), reflect.Int, false},
				{"[41.32][0]", -897, testdata_map_slice[41.32][0], reflect.Int, false},
				{"[41.32][1]", -8755, testdata_map_slice[41.32][1], reflect.Int, false},
				{`len(testdata_map_slice[41.32])`, 2, len(testdata_map_slice[41.32]), reflect.Int, false},
				{`len(testdata_map_slice)`, 4, len(testdata_map_slice), reflect.Int, false},
			}
		}
	}
	return afterDecodeData
}

func Test_DecodeFileData(t *testing.T) {
	if decodeGsonFile == "" {
		decodeGsonFile = "decode.gson"
	}
	gsF := "./testdata/" + decodeGsonFile
	data, _ := ioutil.ReadFile(gsF)
	resetAll()
	keyFoundCallBack := func(d *Decoder, l *Lexer, isFound bool) bool {
		if isFound {
			key := l.String()
			return keyFoundCall(key, d, decodeFileKey, decodeFileData)
		}
		return true
	}
	err := DecodeData(data, keyFoundCallBack, nil, nil)
	checkTestError(t, err, true, "")
	dataAfterDecodeFile = setDataAfter(decodeFileKey, dataAfterDecodeFile, allElemKeyAfterDecodeFile, setDataAfterDecodeFile)
	testDecodeFile(t, decodeFileKey, decodeFileData, dataAfterDecodeFile)
}

func keyFoundCall(key string, d *Decoder, decodeAct string, decodeData []testBytesInfo) bool {
	byteWhiteList, ok := whiteList[decodeAct]
	if !allTest && (!ok || byteWhiteList == nil || len(byteWhiteList) == 0) {
		return true
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
			tarVar := elem.elemVar
			if elemKey == key {
				d.SetAnyTarget(tarVar, true)
				return true
			}
		}
	}
	return true
}

func setDataAfter(decodeAct string, afterDecodeData map[string]map[string][]testVarInfo, allElemKey map[string][]string, setElemMapFunc func(map[string]map[string][]testVarInfo, string, string) map[string]map[string][]testVarInfo) map[string]map[string][]testVarInfo {
	byteWhiteList, ok := whiteList[decodeAct]
	if !allTest && (!ok || byteWhiteList == nil || len(byteWhiteList) == 0) {
		return afterDecodeData
	}
	if !allTest {
		for byteKey, elemWhiteList := range byteWhiteList {
			if elemWhiteList == nil || len(elemWhiteList) == 0 {
				continue
			}
			for elemKey, varWhiteList := range elemWhiteList {
				if varWhiteList == nil || len(varWhiteList) == 0 {
					continue
				}
				if setElemMapFunc != nil {
					afterDecodeData = setElemMapFunc(afterDecodeData, byteKey, elemKey)
				}
			}
		}
	} else {
		for byteKey, elemList := range allElemKey {
			for _, elemKey := range elemList {
				if setElemMapFunc != nil {
					afterDecodeData = setElemMapFunc(afterDecodeData, byteKey, elemKey)
				}
			}
		}
	}
	return afterDecodeData
}

func testDecodeFile(t *testing.T, decodeAct string, decodeData []testBytesInfo, afterDecodeData map[string]map[string][]testVarInfo) {
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
			var tmpElems map[string][]testVarInfo
			var tmpVars []testVarInfo
			if tmpElems, ok = afterDecodeData[byteKey]; ok {
				tmpVars, ok = tmpElems[elemKey]
			}
			varsList := [][]testVarInfo{elem.vars, tmpVars}
			for _, vars := range varsList {
				if vars == nil || len(vars) == 0 {
					continue
				}
				for _, varValue := range vars {
					varKey := varValue.key
					if !allTest {
						if !isInArr(varKey, varWhiteList) {
							continue
						}
					}
					keyName := decodeAct + ".go:" + byteKey + ":" + elemKey + "." + varKey
					checkBasicResult(t, keyName, varValue.expected, varValue.got, varValue.basicType, varValue.gotPointer)
				}
			}
		}
	}
}
