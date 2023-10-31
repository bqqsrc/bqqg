//  Copyright (C) 晓白齐齐,版权所有.
package gson

// 本文件定义了一些基本数据类型的变量，它们用于基本数据字符串的解析测试样例（decode_str_basic_test.go）、从文件中解析出基本数据的测试样例

// 基本数据定义，这些基本数据会在decode_str_basic_test.go中使用
var bool_1_1, bool_1_2, bool_2_1, bool_2_2 bool
var float32_1, float32_2 float32
var float64_1, float64_2 float64
var int8_1, int8_2 int8
var int16_1, int16_2 int16
var int32_1, int32_2 int32
var int64_1, int64_2 int64
var int_1, int_2 int
var uint8_1, uint8_2 uint8
var uint16_1, uint16_2 uint16
var uint32_1, uint32_2 uint32
var uint64_1, uint64_2 uint64
var uint_1, uint_2 uint
var string_1, string_2 string

// 重置这些数据
func resetBasicData() {
	bool_1_1 = false
	bool_1_2 = false
	bool_2_1 = false
	bool_2_2 = false
	int_1 = 0
	int_2 = 0
	int8_1 = 0
	int8_2 = 0
	int16_1 = 0
	int16_2 = 0
	int32_1 = 0
	int32_2 = 0
	int64_1 = 0
	int64_2 = 0
	uint_1 = 0
	uint_2 = 0
	uint8_1 = 0
	uint8_2 = 0
	uint16_1 = 0
	uint16_2 = 0
	uint32_1 = 0
	uint32_2 = 0
	uint64_1 = 0
	uint64_2 = 0
	float32_1 = 0
	float32_2 = 0
	float64_1 = 0
	float64_2 = 0
	string_1 = ""
	string_2 = ""
}

// 定义一个普通的结构体
type TestData struct {
	TestStr   string
	TestInt   int
	TestFloat float64
	TestBool  bool
	TestUint  uint
}

// 定义一个所有字段都有标签的结构体
type TestDataAllWithTag struct {
	TestStr   string  `gson:"test_str"`
	TestInt   int     `gson:"test_int"`
	TestFloat float64 `gson:"test_float"`
	TestBool  bool    `gson:"test_bool"`
	TestUint  uint    `gson:"test_uint"`
}

// 定义一个只有部分字段有标签的结构体
type TestDataSomeWithTag struct {
	TestStr   string `gson:"test_str"`
	TestInt   int    `gson:"test_int"`
	TestFloat float64
	TestBool  bool `gson:"test_bool"`
	TestUint  uint
}

// 普通结构体样例数据
var testdata_1_3 TestData
var testdata_2_3, testdata_2_4, testdata_2_5 TestDataAllWithTag
var testdata_3_3, testdata_3_4, testdata_3_5 TestDataSomeWithTag

func resetStructData() {
	testdata_1_3.TestStr = ""
	testdata_1_3.TestInt = 0
	testdata_1_3.TestFloat = 0
	testdata_1_3.TestBool = false
	testdata_1_3.TestUint = 0
	testdata_2_3.TestStr = ""
	testdata_2_3.TestInt = 0
	testdata_2_3.TestFloat = 0
	testdata_2_3.TestBool = false
	testdata_2_3.TestUint = 0
	testdata_2_4.TestStr = ""
	testdata_2_4.TestInt = 0
	testdata_2_4.TestFloat = 0
	testdata_2_4.TestBool = false
	testdata_2_4.TestUint = 0
	testdata_2_5.TestStr = ""
	testdata_2_5.TestInt = 0
	testdata_2_5.TestFloat = 0
	testdata_2_5.TestBool = false
	testdata_2_5.TestUint = 0
	testdata_3_3.TestStr = ""
	testdata_3_3.TestInt = 0
	testdata_3_3.TestFloat = 0
	testdata_3_3.TestBool = false
	testdata_3_3.TestUint = 0
	testdata_3_4.TestStr = ""
	testdata_3_4.TestInt = 0
	testdata_3_4.TestFloat = 0
	testdata_3_4.TestBool = false
	testdata_3_4.TestUint = 0
	testdata_3_5.TestStr = ""
	testdata_3_5.TestInt = 0
	testdata_3_5.TestFloat = 0
	testdata_3_5.TestBool = false
	testdata_3_5.TestUint = 0
}

// 指针样例数据
var vbool bool
var pbool *bool = &vbool
var ppbool **bool = &pbool
var pppbool ***bool = &ppbool
var ppppbool ****bool = &pppbool
var pppppbool *****bool = &ppppbool
var ppppppbool ******bool = &pppppbool
var vint int
var pint *int = &vint
var ppint **int = &pint
var pppint ***int = &ppint
var ppppint ****int = &pppint
var pppppint *****int = &ppppint
var vint8 int8
var pint8 *int8 = &vint8
var ppint8 **int8 = &pint8
var pppint8 ***int8 = &ppint8
var ppppint8 ****int8 = &pppint8
var vint16 int16
var pint16 *int16 = &vint16
var ppint16 **int16 = &pint16
var pppint16 ***int16 = &ppint16
var vint32 int32
var pint32 *int32 = &vint32
var ppint32 **int32 = &pint32
var vint64 int64
var pint64 *int64 = &vint64
var vuint uint
var puint *uint = &vuint
var vuint8 uint8
var puint8 *uint8 = &vuint8
var vuint16 uint16
var puint16 *uint16 = &vuint16
var vuint32 uint32
var puint32 *uint32 = &vuint32
var vuint64 uint64
var puint64 *uint64 = &vuint64
var vfloat32 float32
var pfloat32 *float32 = &vfloat32
var vfloat64 float64
var pfloat64 *float64 = &vfloat64
var vstring string
var pstring *string = &vstring

func resetPointerData() {
	vbool = false
	vint = 0
	vint8 = 0
	vint16 = 0
	vint32 = 0
	vint64 = 0
	vuint = 0
	vuint8 = 0
	vuint16 = 0
	vuint32 = 0
	vuint64 = 0
	vfloat32 = 0
	vfloat64 = 0
	vstring = ""
}

// 数组样例数据
var int_arr [7]int
var float_arr [7]float64
var string_arr [5]string
var bool_arr [9]bool
var uint_arr [6]uint

func resetArrData() {
	for i := 0; i < len(int_arr); i++ {
		int_arr[i] = 0
	}
	for i := 0; i < len(float_arr); i++ {
		float_arr[i] = 0
	}
	for i := 0; i < len(string_arr); i++ {
		string_arr[i] = ""
	}
	for i := 0; i < len(bool_arr); i++ {
		bool_arr[i] = false
	}
	for i := 0; i < len(uint_arr); i++ {
		uint_arr[i] = 0
	}
}

// slice样例数据
var int_slice = make([]int, 0, 7)
var float_slice = make([]float64, 0, 5)
var string_slice = make([]string, 0, 2)
var bool_slice = make([]bool, 0, 4)
var uint_slice []uint
var int_slice2 = make([]int, 7)
var float_slice2 = make([]float64, 8)
var string_slice2 = make([]string, 2)
var bool_slice2 = make([]bool, 4)
var uint_slice2 = make([]uint, 9)

func resetSliceData() {
	int_slice = make([]int, 0, 7)
	float_slice = make([]float64, 0, 5)
	string_slice = make([]string, 0, 2)
	bool_slice = make([]bool, 0, 4)
	uint_slice = make([]uint, 0)
	int_slice2 = make([]int, 7)
	float_slice2 = make([]float64, 8)
	string_slice2 = make([]string, 2)
	bool_slice2 = make([]bool, 4)
	uint_slice2 = make([]uint, 9)
}

// 指针数组
var pint_earr0, pint_earr1, pint_earr2, pint_earr3, pint_earr4, pint_earr5, pint_earr6 int
var pint_arr [7]*int

func resetPointerArrData() {
	pint_earr0 = 0
	pint_earr1 = 0
	pint_earr2 = 0
	pint_earr3 = 0
	pint_earr4 = 0
	pint_earr5 = 0
	pint_earr6 = 0
	pint_arr[0] = &pint_earr0
	pint_arr[1] = &pint_earr1
	pint_arr[2] = &pint_earr2
	pint_arr[3] = &pint_earr3
	pint_arr[4] = &pint_earr4
	pint_arr[5] = &pint_earr5
	pint_arr[6] = &pint_earr6
}

// 定义一个元素为指针的结构体
type TestPointData struct {
	TestStr   *string
	TestInt   *int
	TestFloat *float64
	TestBool  *bool
	TestUint  *uint
}

// 元素为指针的结构体
var pointStructStr string
var pointStructInt int
var pointStructFloat float64
var pointStructBool bool
var pointStructUint uint
var point_struct = TestPointData{&pointStructStr, &pointStructInt, &pointStructFloat, &pointStructBool, &pointStructUint}

func resetPointerStructData() {
	*point_struct.TestStr = ""
	*point_struct.TestInt = 0
	*point_struct.TestFloat = 0
	*point_struct.TestBool = false
	*point_struct.TestUint = 0
}

// map数据类型
var test_map_str_float map[string]float64
var test_map_int_str map[int]string
var test_map_float_bool map[float32]bool
var test_map_bool_uint map[bool]uint32
var test_map_uint_int map[uint32]int16

func resetMapData() {
	test_map_str_float = make(map[string]float64)
	test_map_int_str = make(map[int]string)
	test_map_float_bool = make(map[float32]bool)
	test_map_bool_uint = make(map[bool]uint32)
	test_map_uint_int = make(map[uint32]int16)
}

// 结构体嵌套了结构体
type TestStructNestStructData struct {
	TestStr    string   `gson:"test_str"`
	TestInt    int      `gson:"test_int"`
	TestFloat  float64  `gson:"test_float"`
	TestBool   bool     `gson:"test_bool"`
	TestUint   uint     `gson:"test_uint"`
	TestStruct TestData `gson:"test_struct"`
}

// 结构体嵌套结构体再嵌套结构体
type TestStructNestTwoStructData struct {
	TestStr         string   `gson:"test_str"`
	TestInt         int      `gson:"test_int"`
	TestFloat       float64  `gson:"test_float"`
	TestBool        bool     `gson:"test_bool"`
	TestUint        uint     `gson:"test_uint"`
	TestStruct      TestData `gson:"test_struct1"`
	TestStructNest  TestStructNestStructData
	TestStructNest2 struct {
		TestFloat  float32 `gson:"testFloat"`
		TestStruct struct {
			TestStr   string
			TestFloat float64
		} `gson:"testSubStruct"`
	} `gson:"test_struct2"`
}

var testdata_struct_struct TestStructNestStructData
var testdata_struct_struct_struct TestStructNestTwoStructData

func resetStructNestStructData() {
	testdata_struct_struct.TestStr = ""
	testdata_struct_struct.TestInt = 0
	testdata_struct_struct.TestFloat = 0
	testdata_struct_struct.TestBool = false
	testdata_struct_struct.TestUint = 0
	testdata_struct_struct.TestStruct.TestStr = ""
	testdata_struct_struct.TestStruct.TestInt = 0
	testdata_struct_struct.TestStruct.TestFloat = 0
	testdata_struct_struct.TestStruct.TestBool = false
	testdata_struct_struct.TestStruct.TestUint = 0

	testdata_struct_struct_struct.TestStr = ""
	testdata_struct_struct_struct.TestInt = 0
	testdata_struct_struct_struct.TestFloat = 0
	testdata_struct_struct_struct.TestBool = false
	testdata_struct_struct_struct.TestUint = 0
	testdata_struct_struct_struct.TestStruct.TestStr = ""
	testdata_struct_struct_struct.TestStruct.TestInt = 0
	testdata_struct_struct_struct.TestStruct.TestFloat = 0
	testdata_struct_struct_struct.TestStruct.TestBool = false
	testdata_struct_struct_struct.TestStruct.TestUint = 0
	testdata_struct_struct_struct.TestStructNest.TestStr = ""
	testdata_struct_struct_struct.TestStructNest.TestInt = 0
	testdata_struct_struct_struct.TestStructNest.TestFloat = 0
	testdata_struct_struct_struct.TestStructNest.TestBool = false
	testdata_struct_struct_struct.TestStructNest.TestUint = 0
	testdata_struct_struct_struct.TestStructNest.TestStruct.TestStr = ""
	testdata_struct_struct_struct.TestStructNest.TestStruct.TestInt = 0
	testdata_struct_struct_struct.TestStructNest.TestStruct.TestFloat = 0
	testdata_struct_struct_struct.TestStructNest.TestStruct.TestBool = false
	testdata_struct_struct_struct.TestStructNest.TestStruct.TestUint = 0

	testdata_struct_struct_struct.TestStructNest2.TestFloat = 0
	testdata_struct_struct_struct.TestStructNest2.TestStruct.TestStr = ""
	testdata_struct_struct_struct.TestStructNest2.TestStruct.TestFloat = 0
}

// 结构体嵌套map
type TestStructNestMapData struct {
	TestStr         string  `gson:"test_str"`
	TestInt         int     `gson:"test_int"`
	TestFloat       float64 `gson:"test_float"`
	TestBool        bool    `gson:"test_bool"`
	TestMapFloatInt map[float32]int8
	TestUint        uint              `gson:"test_uint"`
	TestMapStrUint  map[string]uint16 `gson:"test_struct_map_str_uint"`
}

var testdata_struct_map TestStructNestMapData

func resetStructNestMapData() {
	testdata_struct_map.TestStr = ""
	testdata_struct_map.TestInt = 0
	testdata_struct_map.TestFloat = 0
	testdata_struct_map.TestBool = false
	testdata_struct_map.TestUint = 0
	testdata_struct_map.TestMapFloatInt = make(map[float32]int8)
	testdata_struct_map.TestMapStrUint = make(map[string]uint16)
}

// 结构体嵌套数组
type TestStructNestArrData struct {
	TestStr      string    `gson:"test_str"`
	TestInt      int       `gson:"test_int"`
	TestFloat    float64   `gson:"test_float"`
	TestStrArr   [5]string `gson:"test_str_arr"`
	TestBool     bool      `gson:"test_bool"`
	TestIntArr   [4]int    `gson:"test_int_arr"`
	TestUint     uint      `gson:"test_uint"`
	TestFloatArr []float64 `gson:"test_float_arr"`
	TestBoolArr  []bool    `gson:"test_bool_arr"`
}

var testdata_struct_arr TestStructNestArrData

func resetStructNestArrAndSliceData() {
	testdata_struct_arr.TestStr = ""
	testdata_struct_arr.TestInt = 0
	testdata_struct_arr.TestFloat = 0
	testdata_struct_arr.TestBool = false
	testdata_struct_arr.TestUint = 0
	testdata_struct_arr.TestStrArr[0] = ""
	testdata_struct_arr.TestStrArr[1] = ""
	testdata_struct_arr.TestStrArr[2] = ""
	testdata_struct_arr.TestStrArr[3] = ""
	testdata_struct_arr.TestStrArr[4] = ""
	testdata_struct_arr.TestIntArr[0] = 0
	testdata_struct_arr.TestIntArr[1] = 0
	testdata_struct_arr.TestIntArr[2] = 0
	testdata_struct_arr.TestIntArr[3] = 0
	testdata_struct_arr.TestFloatArr = make([]float64, 0)
	testdata_struct_arr.TestBoolArr = make([]bool, 0)

}

// 数组嵌套结构体
var testdata_arr_struct [3]TestDataAllWithTag

func resetArrNestStructData() {
	for _, v := range testdata_arr_struct {
		v.TestStr = ""
		v.TestInt = 0
		v.TestFloat = 0
		v.TestBool = false
		v.TestUint = 0
	}
}

// slice嵌套结构体
var testdata_slice_struct []TestStructNestStructData

func resetSliceNestStructData() {
	testdata_slice_struct = make([]TestStructNestStructData, 0)
}

// 数组嵌套map
var testdata_arr_map_str_str [3]map[string]string

func resetArrNestMapData() {
	for i := 0; i < len(testdata_arr_map_str_str); i++ {
		testdata_arr_map_str_str[i] = make(map[string]string)
	}
}

// 切片嵌套map
var testdata_slice_map_float_str []map[float64]string

func resetSliceNestMapData() {
	testdata_slice_map_float_str = make([]map[float64]string, 0)
}

// 字典嵌套结构体
var testdata_map_struct map[string]TestDataAllWithTag

func resetMapNestStructData() {
	testdata_map_struct = make(map[string]TestDataAllWithTag)
}

// 字典嵌套字典
var testdata_map_map map[string]map[string]float64

func resetMapNestMapData() {
	testdata_map_map = make(map[string]map[string]float64)
}

// 字典嵌套数组或Slice
var testdata_map_arr map[string][3]string
var testdata_map_slice map[float64][]int

func resetMapNestArrOrSliceData() {
	testdata_map_arr = make(map[string][3]string)
	testdata_map_slice = make(map[float64][]int)
}

func resetAll() {
	resetBasicData()
	resetStructData()
	resetPointerData()
	resetArrData()
	resetPointerArrData()
	resetPointerStructData()
	resetMapData()
	resetStructNestStructData()
	resetStructNestMapData()
	resetStructNestArrAndSliceData()
	resetArrNestStructData()
	resetSliceNestStructData()
	resetArrNestMapData()
	resetSliceNestMapData()
	resetMapNestStructData()
	resetMapNestMapData()
	resetMapNestArrOrSliceData()
}
