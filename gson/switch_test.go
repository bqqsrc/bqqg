//  Copyright (C) 晓白齐齐,版权所有.
package gson

// 本文件是定义了一些测试开关，有时候调试时只需要测试其中某个测试文件的某些测试样例，可以通过这里修改来实现，同时可以修改读取文件的测试

//====================以下是和Decode相关的一些开关==========================//
const allTest = true

// decode_file_test.go测试从哪个文件读取数据，为空时默认为decode.gson
var decodeGsonFile = "decode.gson" //cpy.gson"  //option.gson"

// callback_file_test.go测试从哪个文件读取数据，为空时默认为callback.gson
var callbackGsonFile = "callback.gson"

// 测试白名单（allDecodeTest为false时生效）
// 键：下面的常量有说明，值：想要测试的样例的键
// 测试数据类型->待解析的字符标志->待解析的键->测试值的键
var whiteList = map[string]map[string]map[string][]string{
	// decode_str_test.go里面的测试样例
	decodeStrKey: {
		// "decode_str.stringgv": {
		// 	"decode_str.stringgv": {
		// 		"decode_str.stringgv",
		// 	},
		// },
		// "decode_str.uint64v": {
		// 	"decode_str.uint64v": {
		// 		"decode_str.uint64v",
		// 	},
		// },
	},
	decodeFileKey: {
		"decode.cpy.gson": {
			"testdata_map_arr": {
				`["key1"][0]`,
			},
		},
		"decode.option.gson": {
			"bool_1_1": {
				"bool_1_1",
			},
			"testdata_1_3": {
				"TestStr",
			},
			"pbool": {
				"pbool",
			},
			"int_arr": {
				"[0]",
			},
			"string_slice": {
				"[1]",
			},
			"pint_arr": {
				"[3]",
			},
			"point_struct": {
				"TestFloat",
			},
			"test_map_str_float": {
				`["key3"]`,
			},
			"testdata_struct_struct": {
				"test_float",
				"test_struct.TestStr",
			},
			"testdata_struct_struct_struct": {
				"test_str",
				"test_struct2.testFloat",
			},
			"testdata_struct_map": {
				"test_float",
				"TestMapFloatInt[983.367]",
			},
			"testdata_struct_arr": {
				"test_float",
				"test_str_arr[0]",
				"test_bool_arr[2]",
			},
			"testdata_arr_struct": {
				"[0].TestStr",
				"[1].test_bool",
			},
			"testdata_slice_struct": {
				"[1].test_float",
				"[0].test_struct.TestUint",
			},
			"testdata_arr_map_str_str": {
				`[0]["key1"]`,
			},
			"testdata_map_struct": {
				`["key1"].TestStr`,
				`["key2"].test_uint`,
			},
			"testdata_map_map": {
				`["key1"]["sub1key1"]`,
			},
		},
	},
	callbackFileKey: {
		callbackGsonFile: {
			// "int64_1": {
			// 	"int64_1",
			// },
			// "all_callback_file_str": {
			// 	"all_file",
			// },
		},
	},
}

// 一些白名单的key取值常量定义
const decodeStrKey = "decode_str_test"       // 基本数据类型字符串，decode_str_test.go里的样例
const decodeFileKey = "decode_file_test"     // 从文件中读取的数据解析，decode_file_test.go里的样例
const callbackFileKey = "callback_file_test" //从文件读取数据回调，handler_file_test.go里的测试样例
