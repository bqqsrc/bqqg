//  Copyright (C) 晓白齐齐,版权所有.

package gson

// 本文将用于测试各种Handler回调函数
import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var allFileStr = `bool_1_1:truebool_1_2:falsebool_2_1:falsebool_2_2:truefloat32_1:59.999float32_2:89.77float64_1:34e4float64_2:9870e-2int8_1:-33int8_2:58int16_1:456int16_2:-567int32_1:703522int32_2:-89076int64_1:325235298967int64_2:-98799889678int_1:39870935321int_2:23458987898uint8_1:27uint8_2:35uint16_1:356uint16_2:598uint32_1:588522uint32_2:82376uint64_1:325906298967uint64_2:98757889678uint_1:39870967321uint_2:23453487898string_1:baiqiqi5string_2:baiqiqi3testdata_1_3:{TestStr:baiqiqi3TestInt:-3599TestFloat:99.7TestBool:trueTestUint:235}testdata_2_3:{test_str:baiqiqi4test_int:-9088test_float:99.53test_bool:falsetest_uint:234}testdata_2_4:{TestStr:baiqiqi5TestInt:-19088TestFloat:991.53TestBool:falseTestUint:2234}testdata_2_5:{TestStr:baiqiqi6test_int:-188test_float:-991.53TestBool:falseTestUint:1134}testdata_3_3:{test_str:baiqiqi3_3test_int:-988TestFloat:99.3test_bool:trueTestUint:34}testdata_3_4:{TestStr:baiqiqi3_4TestInt:-198TestFloat:91.53TestBool:falseTestUint:224}testdata_3_5:{TestStr:baiqiqi3_5test_int:-1288TestFloat:891.53TestBool:trueTestUint:114}pbool:truepint:322334234234234pint8:-122pint16:3223pint32:-123221312pint64:12312214324234puint:32315624535232puint8:253puint16:23453puint32:2131252puint64:12341252345345pfloat32:233.9332pfloat64:11234e-3pstring:pstringstringint_arr:3822283282-323449832321-23212float_arr:93.3332.51.00078783e-3-3.2223e29.999string_arr:baiqiqi1baiqiqi2baiqiqi3baiqiqi4baiqiqi5bool_arr:truefalsetruetruetruefalsefalsetruefalseuint_arr:3423521334125253223239234523int_slice:3822283282-323449832321-23212float_slice:93.3332.51.00078783e-3-3.2223e29.999string_slice:baiqiqi1baiqiqi2baiqiqi3baiqiqi4baiqiqi5bool_slice:truefalsetruetruetruefalsefalsetruefalseuint_slice:3423521334125253223239234523int_slice2:3822283282-323449832321-23212float_slice2:93.3332.51.00078783e-3-3.2223e29.999string_slice2:baiqiqi1baiqiqi2baiqiqi3baiqiqi4baiqiqi5bool_slice2:truefalsetruetruetruefalsefalsetruefalseuint_slice2:3423521334125253223239234523pint_arr:3822283282-323449832321-23212point_struct:{TestStr:baiqiqi8TestInt:-35998TestFloat:99.77TestBool:trueTestUint:2355}test_map_str_float:{key1:998.3key2:997.6key3:99.7}test_map_int_str:{1:baiqiqi12:baiqiqi23:baiqiqi34:baiqiqi4}test_map_float_bool:{98.787:true78.55:false-333.33:false-31233.33:true}test_map_bool_uint:{true:3832false:256}test_map_uint_int:{9833:34785:-56333:5662333:133}`
var callbackBuilder strings.Builder

var callbackFileData = []testBytesInfo{
	{callbackGsonFile, "", true, "", []testElemInfo{
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
		{&allFileStr, "all_callback_file_str", nil},
	}},
}

var dataAfterCallback map[string]map[string][]testVarInfo
var allElemKeyAfterCallback = map[string][]string{
	callbackGsonFile: {
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
		"all_callback_file_str",
	},
}

func setDataAfterCallback(afterDecodeData map[string]map[string][]testVarInfo, byteKey, elemKey string) map[string]map[string][]testVarInfo {
	if afterDecodeData == nil {
		afterDecodeData = make(map[string]map[string][]testVarInfo)
	}
	if afterDecodeData[byteKey] == nil {
		afterDecodeData[byteKey] = make(map[string][]testVarInfo)
	}
	switch byteKey {
	case callbackGsonFile:
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
				{"[key1]", 998.3, test_map_str_float["key1"], reflect.Float64, false},
				{"[key2]", 997.6, test_map_str_float["key2"], reflect.Float64, false},
				{"[key3]", 99.7, test_map_str_float["key3"], reflect.Float64, false},
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
		case "all_callback_file_str":
			afterDecodeData[byteKey][elemKey] = []testVarInfo{
				{"all_file", allFileStr, callbackBuilder.String(), reflect.String, false},
			}
		}
	}
	return afterDecodeData
}

var varByKey = []string{
	"testdata_1_3",
	"testdata_2_3",
	"testdata_2_4",
	"testdata_2_5",
	"testdata_3_3",
	"testdata_3_4",
	"testdata_3_5",
	"point_struct",
	"test_map_str_float",
	"test_map_int_str",
	"test_map_float_bool",
	"test_map_bool_uint",
	"test_map_uint_int",
}
var varByOneIndex = []string{
	"bool_1_1",
	"bool_1_2",
	"bool_2_1",
	"bool_2_2",
	"float32_1",
	"float32_2",
	"float64_1",
	"float64_2",
	"int8_1",
	"int8_2",
	"int16_1",
	"int16_2",
	"int32_1",
	"int32_2",
	"int64_1",
	"int64_2",
	"int_1",
	"int_2",
	"uint8_1",
	"uint8_2",
	"uint16_1",
	"uint16_2",
	"uint32_1",
	"uint32_2",
	"uint64_1",
	"uint64_2",
	"uint_1",
	"uint_2",
	"string_1",
	"string_2",
	"pbool",
	"pint8",
	"pint16",
	"pint32",
	"pint64",
	"pint",
	"puint8",
	"puint16",
	"puint32",
	"puint64",
	"puint",
	"pstring",
	"pfloat32",
	"pfloat64",
}
var varByIndex = []string{
	"int_arr",
	"float_arr",
	"string_arr",
	"bool_arr",
	"uint_arr",
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
}

var currentParseKey string
var currentTargetKey string

// 测试DecodeData接口的3个回调函数
func Test_FoundCallBackAndConvert(t *testing.T) {
	if callbackGsonFile == "" {
		callbackGsonFile = "callback.gson"
	}
	gsF := "./testdata/" + callbackGsonFile
	data, _ := ioutil.ReadFile(gsF)
	resetAll()
	callbackBuilder.Reset()

	keyFoundCallBack := func(d *Decoder, l *Lexer, isFound bool) bool {
		if l != nil && isFound {

			currentParseKey = l.String()
			callbackBuilder.WriteString(currentParseKey)
			callbackBuilder.WriteString(":")
			if currentTargetKey == "" {
				currentTargetKey = currentParseKey
			}
		} else {
			currentParseKey = ""
			key := l.String()
			if key == currentTargetKey {
				currentTargetKey = ""
			}
		}
		return true
	}
	valueCallback := func(index int, key string, d *Decoder, l *Lexer) bool {
		callbackBuilder.WriteString(l.String())
		checkString(t, "currentKey of "+currentParseKey, currentParseKey, key)

		if isInArr(key, varByOneIndex) {
			checkString(t, "currentParseKey of "+currentTargetKey, currentTargetKey, currentParseKey)
			checkInt(t, "currentIndex of "+currentTargetKey, 0, index)
		} else if isInArr(key, varByIndex) {
			checkString(t, "currentParseKey of "+currentTargetKey, currentTargetKey, currentParseKey)
		}
		err := callbackConvert(index, key, d, l, callbackFileKey, callbackFileData)
		checkTestError(t, err, true, "")
		return true
	}
	objCallback := func(key string, d *Decoder, isBegin bool) bool {
		if isBegin {
			callbackBuilder.WriteString("{")
		} else {
			callbackBuilder.WriteString("}")
		}
		return true
	}
	err := DecodeData(data, keyFoundCallBack, valueCallback, objCallback)
	checkTestError(t, err, true, "")
	dataAfterCallback = setDataAfter(callbackFileKey, dataAfterCallback, allElemKeyAfterCallback, setDataAfterCallback)
	testDecodeFile(t, callbackFileKey, callbackFileData, dataAfterCallback)
}

func callbackConvert(index int, key string, d *Decoder, l *Lexer, decodeAct string, decodeData []testBytesInfo) error {
	byteWhiteList, ok := whiteList[decodeAct]
	if !allTest && (!ok || byteWhiteList == nil || len(byteWhiteList) == 0) {
		return nil
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
			if elemKey == currentTargetKey {
				var err error
				if isInArr(currentTargetKey, varByKey) {
					err = ConvertByKey(tarVar, key, l)
				} else if isInArr(currentTargetKey, varByOneIndex) || isInArr(currentTargetKey, varByIndex) {
					err = ConvertByIndex(tarVar, index, l)
				}
				return err
			}
		}
	}
	return nil
}

var exclusiveFileStr = `bool_1_1:bool_1_2:falsebool_2_1:falsetestdata_1_3:{TestStr:baiqiqi3TestInt:-3599TestFloat:99.7TestBool:trueTestUint:235}testdata_2_3:testdata_2_4:{TestStr:baiqiqi5TestInt:-19088TestFloat:991.53TestBool:falseTestUint:2234}int_arr:3822283282-323449832321-23212float_arr:93.3332.51.00078783e-3-3.2223e29.999string_arr:test_map_str_float:test_map_int_str:{1:baiqiqi12:baiqiqi23:baiqiqi34:baiqiqi4}test_map_float_bool:{98.787:true78.55:false-333.33:false-31233.33:true}testdata_struct_struct:{test_float:98.776test_struct_tt:test_str:baiqiqi8test_int:-35132test_bool:falsetest_uint:34156}testdata_struct_struct_struct:{test_float:98.7976test_struct1:{TestStr:subbaiqiqi112TestFloat:8999.3TestInt:332156TestBool:falseTestUint:981459}test_str:baiqiqi81test_int:-332test_struct2:{testSubStruct:{TestStr:baiqiqiqiTestFloat:99.7123}testFloat:-99111.13}TestStructNest:{TestStr:subbaiqiqi1TestFloat:-991.13TestInt:-33223TestBool:falseTestUint:989156test_struct:}test_bool:truetest_uint:3423156}testdata_map_struct:`

func Test_CallbackExclusive(t *testing.T) {
	gsF := "./testdata/exclusive.gson"
	data, _ := ioutil.ReadFile(gsF)
	resetAll()
	callbackBuilder.Reset()
	keyFoundCallBack := func(d *Decoder, l *Lexer, isFound bool) bool {
		if l != nil {
			if isFound {
				key := l.String()
				callbackBuilder.WriteString(key)
				callbackBuilder.WriteString(":")
				switch key {
				case "bool_1_1":
					d.SetAnyTarget(&bool_1_1, true)
				case "bool_1_2":
					d.SetAnyTarget(&bool_1_2, false)
				case "testdata_1_3":
					d.SetAnyTarget(&testdata_1_3, false)
				case "testdata_2_3":
					d.SetAnyTarget(&testdata_2_3, true)
				case "float_arr":
					d.SetAnyTarget(&float_arr, false)
				case "string_arr":
					d.SetAnyTarget(&string_arr, true)
				case "test_map_str_float":
					d.SetAnyTarget(&test_map_str_float, true)
				case "test_map_float_bool":
					d.SetAnyTarget(&test_map_float_bool, false)
				case "test_struct_tt":
					d.SetAnyTarget(&testdata_2_4, true)
				case "test_struct1":
					d.SetAnyTarget(&testdata_2_5, false)
				case "test_struct":
					d.SetAnyTarget(&testdata_3_3, true)
				case "testdata_map_struct":
					d.SetAnyTarget(&testdata_map_struct, true)
				}
			}
		}
		return true
	}
	valueCallback := func(index int, key string, d *Decoder, l *Lexer) bool {
		callbackBuilder.WriteString(l.String())
		return true
	}
	objCallback := func(key string, d *Decoder, isBegin bool) bool {
		if isBegin {
			callbackBuilder.WriteString("{")
		} else {
			callbackBuilder.WriteString("}")
		}
		return true
	}
	err := DecodeData(data, keyFoundCallBack, valueCallback, objCallback)
	checkTestError(t, err, true, "")
	checkString(t, "exclusiveFileStr", exclusiveFileStr, callbackBuilder.String())
}
