//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"io/ioutil"
	"strings"
	"testing"
)

type stringInfo struct {
	valueType int
	value     string
	line      int64
	offset    int64
}

var emptyStringInfo = stringInfo{0, "", -1, -1}
var resultKeyMap = make(map[string]stringInfo)
var currentKey string
var currentVar string
var resultVarMap = make(map[string]stringInfo)
var isVar = false
var isDic = false
var hasValue = false
var currentDicKey string
var dicMap = make(map[string]map[string]stringInfo)

func foundHandlerCall(dataKind int, l *Lexer, line int64, offset int64) (error, bool) {
	data := l.String()
	switch dataKind {
	case dataKey:
		if isDic {
			dicMap[currentDicKey][data] = emptyStringInfo
		} else {
			resultKeyMap[data] = emptyStringInfo
		}
		currentKey = data
		hasValue = false
		break
	case dataValue:
		if isVar {
			isVar = false
			resultVarMap[currentVar] = stringInfo{stringV, data, line, offset}
		} else {
			if isDic {
				dicMap[currentDicKey][currentKey] = stringInfo{l.GetValueType(), data, line, offset}
			} else {
				resultKeyMap[currentKey] = stringInfo{l.GetValueType(), data, line, offset}
			}
		}
		hasValue = true
		break
	case dataVar:
		resultVarMap[data] = emptyStringInfo
		currentVar = data
		isVar = true
		break
	case dataObjBegin:
		if hasValue {
			currentKey = ""
		}
		isDic = true
		currentVar = ""
		isVar = false
		currentDicKey = currentKey
		currentKey = ""
		if currentDicKey == "" {
			currentDicKey = "default"
		}
		if _, ok := dicMap[currentDicKey]; !ok {
			dicMap[currentDicKey] = make(map[string]stringInfo)
		}
		break
	case dataObjEnd:
		isDic = false
		currentDicKey = ""
		currentKey = ""
		currentVar = ""
		isVar = false
		break

	}
	return nil, true
}

func judgeKey(t *testing.T, key, expectedValue, expectedValueTypeStr string, expectedLine, expectedOffset int64, expectedValueType int) {
	if result, ok := resultKeyMap[key]; ok {
		if result.value != expectedValue {
			t.Fatalf("value of %s expected %s, got %s", key, expectedValue, result.value)
		}
		if result.line != expectedLine {
			t.Fatalf("line of %s expected %d, got %d", key, expectedLine, result.line)
		}
		if result.offset != expectedOffset {
			t.Fatalf("offset of %s expected %d, got %d", key, expectedOffset, result.offset)
		}
		if result.valueType != expectedValueType {
			t.Fatalf("valueType of %s expected %d, %s, got %d", key, expectedValueType, expectedValueTypeStr, result.valueType)
		}
	} else {
		t.Fatalf("%s expected %s, not got it", key, expectedValue)
	}
}

func judgeDic(t *testing.T, key1, key2, expectedValue, expectedValueTypeStr string, expectedLine, expectedOffset int64, expectedValueType int) {
	if tmpMap, ok := dicMap[key1]; ok {
		if result, tmpok := tmpMap[key2]; tmpok {
			if result.value != expectedValue {
				t.Fatalf("value of %s->%s expected %s, got %s", key1, key2, expectedValue, result.value)
			}
			if result.line != expectedLine {
				t.Fatalf("line of  %s->%s expected %d, got %d", key1, key2, expectedLine, result.line)
			}
			if result.offset != expectedOffset {
				t.Fatalf("offset of  %s->%s expected %d, got %d", key1, key2, expectedOffset, result.offset)
			}
			if result.valueType != expectedValueType {
				t.Fatalf("valueType of  %s->%s expected %d, %s, got %d", key1, key2, expectedValueType, expectedValueTypeStr, result.valueType)
			}
		} else {
			t.Fatalf("%s->%s expected  %s, not got it", key1, key2, expectedValue)
		}
	} else {
		t.Fatalf("%s expected exist, not got it", key1)
	}
}

func judgeVar(t *testing.T, key, expectedValue string, expectedLine, expectedOffset int64) {
	if version <= 100 {
		return
	}

	if result, ok := resultVarMap[key]; ok {
		if result.value != expectedValue {
			t.Fatalf("var %s expected %s, got %s", key, expectedValue, result.value)
		}
		if result.line != expectedLine {
			t.Fatalf("line of var %s expected %d, got %d", key, expectedLine, result.line)
		}
		if result.offset != expectedOffset {
			t.Fatalf("offset of var %s expected %d, got %d", key, expectedOffset, result.offset)
		}
	} else {
		t.Fatalf("var %s expected %s, not got it", key, expectedValue)
	}
}

// （样例已固定不可修改）测试Valid函数
func Test_Valid(t *testing.T) {
	// return
	tmpStr := "name: baiqiqi"
	data := []byte(tmpStr)
	if !Valid(data) {
		t.Fatalf("Valid expected return true, got false")
	}

	tmpStr =
		`text2: {
  v1: 33
  v2: testtt
  v3: 89.9
  v5: aaa bbb ccc ddd
}
: nokeyname`
	data = []byte(tmpStr)
	if Valid(data) {
		t.Fatalf("Valid expected return false, got true")
	}
}

// 特殊样例测试（样例已固定不可修改）
func Test_Scanner_Special(t *testing.T) {
	// return
	// 测试一个结尾没有空白符号的样例
	tmpStr := "name: baiqiqi"
	data := []byte(tmpStr)
	scan := newScanner()
	defer freeScanner(scan)
	scan.dataFoundHandler = foundHandlerCall
	err := work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "name", "baiqiqi", "string", 1, 13, stringV)

	// 测试一个结尾只有换行符号的样例
	scan.reset(false)
	tmpStr =
		`name: baiqiqi2
`
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "name", "baiqiqi2", "string", 1, 14, stringV)

	// 测试一个结尾只有非换行空白符号的样例
	scan.reset(false)
	tmpStr = "name: baiqiqi6   "
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "name", "baiqiqi6", "string", 1, 14, stringV)

	// 测试：先给一个数值赋值为其他类型（如FloatV），然后再变为字符串，且没有空白符
	scan.reset(false)
	tmpStr = "notnumspecial:21eE8"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "notnumspecial", "21eE8", "string", 1, 19, stringV)

	// 测试：最后一个数据为强制字符串，且最后没有空白符
	scan.reset(false)
	var tmpBuilder, resultBuilder strings.Builder
	resultBuilder.WriteString("aaa")
	stringsBuilderWriteLine(&resultBuilder, 1)
	resultBuilder.WriteString("bbb")
	stringsBuilderWriteLine(&resultBuilder, 1)
	resultBuilder.WriteString("ccc```")
	stringsBuilderWriteLine(&resultBuilder, 1)
	resultBuilder.WriteString("`")
	stringsBuilderWriteLine(&resultBuilder, 1)
	resultBuilder.WriteString("dd")
	stringsBuilderWriteLine(&resultBuilder, 1)
	resultStr := resultBuilder.String()
	tmpBuilder.WriteString("new_str12 : ``")
	tmpBuilder.WriteString(resultStr)
	tmpBuilder.WriteString("``")
	tmpStr = tmpBuilder.String()
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "new_str12", resultStr, "string", 6, 0, stringV)

	// 测试：最后一个数据为转义字符，且最后没有空白符
	scan.reset(false)
	tmpStr = "special_escape: %=%%%:"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "special_escape", "=%:", "string", 1, 22, stringV)

	// 测试：最后一个数据为}，且最后没有空白符
	scan.reset(false)
	tmpStr = "testdic : {dic1: value1 dic2: 33}"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeDic(t, "testdic", "dic1", "value1", "string", 1, 23, stringV)
	judgeDic(t, "testdic", "dic2", "33", "uint", 1, 32, uintV)

	// 测试：最后一个为var的值，且最后没有空白符
	scan.reset(false)
	tmpStr = "special_var = var_special"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeVar(t, "special_var", "var_special", 1, 25)

	// 测试：最后一个为注释，且最后没有空白符
	scan.reset(false)
	tmpStr = "test_comment: testtest#一个注释"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "test_comment", "testtest", "string", 1, 22, stringV)

	// 测试：最后一个为注释，且最后没有空白符
	scan.reset(false)
	tmpStr = "test_comment11: testtest#``一个注释``"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "test_comment11", "testtest", "string", 1, 24, stringV)

	// 测试：最后一个为注释，且最后没有空白符
	scan.reset(false)
	tmpStr = "test_forcestr1: `testtest#`"
	data = []byte(tmpStr)
	err = work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	judgeKey(t, "test_forcestr1", "testtest#", "string", 1, 26, stringV)

}

// 普通样例测试（样例已固定不可修改）
func Test_Scanner(t *testing.T) {
	// return
	tmpStr := ""
	var tmpBuilder strings.Builder
	data, _ := ioutil.ReadFile("./testdata/data.gson")
	scan := newScanner()
	defer freeScanner(scan)
	scan.dataFoundHandler = foundHandlerCall
	err := work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}
	// 测试字符串数值类型，测试键定义符紧邻键，紧邻值，不紧邻键和值，单独一行
	judgeKey(t, "name", "baiqiqi4", "string", 1, 14, stringV)
	judgeKey(t, "pass", "compile", "string", 2, 14, stringV)
	judgeKey(t, "accout", "xinrui", "string", 3, 13, stringV)
	judgeKey(t, "password", "world", "string", 6, 5, stringV)
	judgeKey(t, "database", "color", "string", 8, 16, stringV)

	// 测试数值类型，int，float，uint，小数、整型、无符号整型、带e的数值
	judgeKey(t, "height", "33", "uint", 9, 10, uintV)
	judgeKey(t, "weight", "89.77", "float", 10, 14, floatV)
	judgeKey(t, "negnum", "-338.79", "float", 11, 15, floatV)
	judgeKey(t, "negnum2", "-0.739", "float", 16, 6, floatV)
	judgeKey(t, "max", "-958", "int", 18, 10, intV)
	judgeKey(t, "min", "10e-30", "float", 20, 11, floatV)
	judgeKey(t, "num1", "32e3", "float", 21, 10, floatV)
	judgeKey(t, "num2", "75E-3", "float", 22, 11, floatV)
	judgeKey(t, "num3", "0.73e+6", "float", 23, 14, floatV)
	judgeKey(t, "num4", "-5.65E7", "float", 25, 7, floatV)
	judgeKey(t, "num5", "10.57E-68", "float", 27, 21, floatV)
	judgeKey(t, "num6", "15.7e+3", "float", 33, 7, floatV)

	// 测试带e的字符串
	judgeKey(t, "notnum", "21eE8", "string", 35, 12, stringV)

	// 测试nil、false、true
	judgeKey(t, "done", "nil", "nil", 37, 3, nilV)
	judgeKey(t, "failed", "false", "bool", 40, 5, boolV)
	judgeKey(t, "success", "true", "bool", 42, 12, boolV)
	judgeKey(t, "strtrue", "trues", "string", 44, 14, stringV)
	judgeKey(t, "strfalse", "falses", "string", 45, 16, stringV)
	judgeKey(t, "strnil", "nils", "string", 46, 12, stringV)

	// 测试带数值的字符串
	judgeKey(t, "strnum", "8.73s", "string", 47, 12, stringV)
	judgeKey(t, "strint", "999m", "string", 48, 12, stringV)
	judgeKey(t, "strpoint", "8..7s", "string", 49, 15, stringV)
	if version <= 100 {
		judgeKey(t, "pointstr", "testvar10", "string", 77, 12, stringV)
	} else {
		judgeKey(t, "pointstr", ".99", "string", 50, 13, stringV)
	}

	// 测试变量，变量定义符紧邻变量，紧邻变量值，不紧邻变量和变量值，单独一行
	judgeVar(t, "var1", "testvar1", 52, 15)
	judgeVar(t, "var2", "", -1, -1)
	judgeVar(t, "var3", "testvar3", 54, 14)
	judgeVar(t, "var4", "testvar4", 55, 13)
	judgeVar(t, "var5", "testvar5", 60, 8)
	judgeVar(t, "var6", "testvar6", 64, 8)
	judgeVar(t, "var7", "testvar7", 66, 10)
	judgeVar(t, "var8", "testvar8", 68, 17)
	judgeVar(t, "var9", "testvar9", 71, 8)
	judgeVar(t, "var10", "testvar10", 77, 12)

	// 测试注释，单独一行的注释，注释再有效字符后面，注释紧邻有效字符，注释再:后面，紧邻:，紧邻键值，几行注释
	judgeKey(t, "comment", "comment", "string", 80, 16, stringV)
	judgeKey(t, "kids", "children", "string", 81, 14, stringV)
	judgeKey(t, "kids0", "child0", "string", 83, 6, stringV)
	judgeKey(t, "kids1", "child1", "string", 85, 6, stringV)
	if version <= 100 {
		judgeKey(t, "kids2", "testvar", "string", 91, 9, stringV)
	} else {
		judgeKey(t, "kids2", "child2", "string", 88, 6, stringV)
	}
	judgeKey(t, "kids3", "child3", "string", 104, 6, stringV)

	// 测试注释紧邻变量值
	judgeVar(t, "vartest", "testvar", 91, 9)

	// 测试多行注释，注释最后紧邻有效值，注释起始位置在有效值之后，单独几行多行注释，注释起始紧邻键有效值，一行里出现多个注释将键值切割开的情况
	judgeKey(t, "comment1", "comment1", "string", 109, 19, stringV)
	judgeKey(t, "comment2", "comment2Value", "string", 114, 13, stringV)
	judgeKey(t, "comment3", "comment3Value", "string", 130, 13, stringV)
	judgeKey(t, "comment5", "comment5Value", "string", 132, 67, stringV)
	judgeKey(t, "comment6", "comment6Value", "string", 133, 21, stringV)

	// 测试强制字符串
	tmpBuilder.Reset()
	tmpBuilder.WriteString("package main")
	stringsBuilderWriteLine(&tmpBuilder, 2)
	tmpBuilder.WriteString("import (")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString(`	"github.com/bqqsrc/gson"`)
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString(")")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("func main() { ")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("	gson.Unmarshal(nil, nil)")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("}")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpStr = tmpBuilder.String()
	judgeKey(t, "code0", tmpStr, "string", 144, 0, stringV)

	// 强制字符串中有转义字符，特殊字符，强制字符串有空格
	judgeKey(t, "code1", "line=`echo $do`", "string", 145, 26, stringV)
	judgeKey(t, "code2", "`this is a string `", "string", 147, 31, stringV)
	judgeKey(t, "code3", "`a` is ``b`` and c", "string", 149, 29, stringV)

	// 强制字符串有少于和多于起始数量的`
	tmpBuilder.Reset()
	tmpBuilder.WriteString("c is `cc` d ")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("is ```dd```")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpStr = tmpBuilder.String()
	judgeKey(t, "code4", tmpStr, "string", 153, 0, stringV)

	// 强制字符串为数值字符串
	judgeKey(t, "numstr1", "98.888", "string", 154, 16, stringV)

	// 带`的字符串
	judgeKey(t, "notescape", "a`ls`", "string", 156, 17, stringV)

	// 转义字符
	judgeKey(t, "aescape", "!%@{}:$=`#", "string", 158, 30, stringV)
	judgeKey(t, "state1", "a=3", "string", 159, 12, stringV)
	judgeKey(t, "state2", "30%=0.3", "string", 160, 17, stringV)

	// 最后一个字符为%，不做转义字符处理
	judgeKey(t, "state3", "0.3=30%", "string", 161, 16, stringV)

	// 转义字符
	judgeKey(t, "state4", "0.3=30%", "string", 162, 17, stringV)
	judgeKey(t, "state5", "38.99%+$test", "string", 163, 22, stringV)

	// 键包含转义字符
	judgeKey(t, "a=3", "test=test", "string", 165, 17, stringV)
	// 键为强制字符串，且值的开头没有`，结尾的`作为普通字符处理
	judgeKey(t, "99.8%7`", "99.8%7`", "string", 167, 21, stringV)

	// 带key的对象（字典）
	judgeDic(t, "server", "listen", "80", "uint", 170, 12, uintV)
	judgeDic(t, "server", "name", "com.bqq.com", "string", 171, 19, stringV)
	judgeDic(t, "server", "location", "/", "string", 172, 13, stringV)
	judgeDic(t, "server", "timeout", "30", "uint", 173, 13, uintV)

	// 不带key的对象（字典），这里做default键处理
	judgeDic(t, "default", "dlisten", "80", "uint", 177, 13, uintV)
	judgeDic(t, "default", "dname", "com.bqq.com", "string", 178, 20, stringV)
	judgeDic(t, "default", "dlocation", "/", "string", 179, 14, stringV)
	judgeDic(t, "default", "dtimeout", "30", "uint", 180, 14, uintV)

	judgeDic(t, "default", "ddlisten", "80", "uint", 184, 14, uintV)
	judgeDic(t, "default", "ddname", "com.bqq.com", "string", 185, 21, stringV)
	judgeDic(t, "default", "ddlocation", "/", "string", 186, 15, stringV)
	judgeDic(t, "default", "ddtimeout", "30", "uint", 187, 15, uintV)

	// 键的:紧邻key和{的情况
	judgeDic(t, "server2", "listen", "80", "uint", 191, 12, uintV)
	judgeDic(t, "server2", "name", "com.bqq.com", "string", 192, 19, stringV)
	judgeDic(t, "server2", "location", "/", "string", 193, 13, stringV)
	judgeDic(t, "server2", "timeout", "30", "uint", 194, 13, uintV)

	// 强制字符串包含多于或者少于起始的`
	tmpBuilder.Reset()
	tmpBuilder.WriteString("aaa")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("bbb")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("ccc```")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("`")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("dd")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpStr = tmpBuilder.String()
	judgeKey(t, "new_str", tmpStr, "string", 202, 0, stringV)

	// 测试数值紧邻一个{的情况
	judgeKey(t, "dic_test", "dictest", "string", 204, 17, stringV)
	judgeKey(t, "dic_num_test", "999.8", "float", 208, 19, floatV)
	// fmt.Println(dicMap)
	judgeDic(t, "default", "aa", "aa", "string", 205, 8, stringV)
	judgeDic(t, "default", "bb", "bb", "string", 209, 8, stringV)

	judgeKey(t, "tmp_num1", "9", "uint", 212, 11, uintV)
}

func stringsBuilderWriteLine(builder *strings.Builder, n int) {
	for i := 0; i < n; i++ {
		builder.WriteString("\r\n")
	}
}

// 错误样例测试（样例已固定不可修改）
func Test_Scanner_Err(t *testing.T) {
	// return
	// 正确的样例
	tmpStr := "name: baiqiqi"
	data := []byte(tmpStr)
	scan := newScanner()
	defer freeScanner(scan)
	err := work(data, scan)
	if err != nil {
		t.Fatalf("Valid expected return nil, got %s", err)
	}

	// 没有键的定义直接就出现:的情况
	scan.reset(true)
	tmpStr =
		`text2: {
  v1: 33
  v2: testtt
  v3: 89.9
  v5: aaa bbb ccc ddd
}
: nokeyname`
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, "an error key-define without a key-name", byte(':'), 7, 1)

	// 没有变量的定义直接出现=的情况
	if version > 100 {
		scan.reset(true)
		tmpStr =
			`     = nokeyname
node : test`
		data = []byte(tmpStr)
		err = work(data, scan)
		checkError(t, err, "an error var-define without a var-name", byte('='), 1, 6)
	}

	// 强制字符串起始有超过3个的`
	scan.reset(true)
	tmpStr = "node : ````to many string````"
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, "too many '`' before a string", byte('`'), 1, 11)

	// 注释里的强制字符起始有超过3个的`
	scan.reset(true)
	tmpStr = "#````一个注释````node:test"
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, "too many '`' before a string", byte('`'), 1, 5)

	// 多出一个多余的{
	scan.reset(true)
	tmpStr =
		`name:bqiqiqi
dic:{
  key1: vlule
}
}`
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, "an error '}', not found its begin '{'", byte('}'), 5, 1)

	// 文件结束了也没有找到和强制字符串成对的结尾强制字符组合
	scan.reset(true)
	var tmpBuilder strings.Builder
	tmpBuilder.Reset()
	tmpBuilder.WriteString("name12345:baiqiqi")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("str : ``aaa")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("bbb")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("ccc")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("`")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("dd```")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpBuilder.WriteString("a ")
	stringsBuilderWriteLine(&tmpBuilder, 1)
	tmpStr = tmpBuilder.String()
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, `unexpected end of gson input, want a "string end character", but not found`, 0, 8, 1)

	// 测试解析层级超过最大层级
	// 这个测试样例一般不会测试，因为目前最大层级设置为10000，执行这个测试样例很耗时
	tmpBuilder.Reset()
	if maxNestingDepth < 10 {
		for i := 0; i < maxNestingDepth+2; i++ {
			tmpBuilder.WriteString("{")
		}
		tmpBuilder.WriteString("name:baiqiqi")
		for i := 0; i < maxNestingDepth+2; i++ {
			tmpBuilder.WriteString("}")
		}
		tmpStr = tmpBuilder.String()
		data = []byte(tmpStr)
		err = work(data, scan)
		checkError(t, err, "exceeded max depth", byte('{'), 1, maxNestingDepth+1)
	}

	// 文件结束了都没有找到匹配的}
	scan.reset(true)
	tmpStr =
		`name : {
	key1: value1
	key2: {
		sub1: subValue1
		subKey2: {
			kind: kind2
		}
	}
`
	data = []byte(tmpStr)
	err = work(data, scan)
	checkError(t, err, `unexpected end of gson input, want a "object end character", but not found`, 0, 9, 1)

}

func checkError(t *testing.T, err error, expectedErrorMsg string, expectedByte byte, expectedLine, expectedOffset int64) {
	if err == nil {
		t.Fatalf("err expected not nil, errormsg is %s, but got nil", expectedErrorMsg)
	} else {
		if tmpErr, ok := err.(*SyntaxError); ok {
			if tmpErr.line != expectedLine {
				t.Fatalf("err line expected %d, got %d", expectedLine, tmpErr.line)
			}
			if tmpErr.lineOffset != expectedOffset {
				t.Fatalf("err lineOffset expected %d, got %d", expectedOffset, tmpErr.lineOffset)
			}
			realExpectedMsg := expectedErrorMsg
			if expectedByte != 0 {
				realExpectedMsg = "invalid character " + quoteChar(expectedByte) + ", " + expectedErrorMsg
			}
			if tmpErr.msg != realExpectedMsg {
				t.Fatalf("err msg expected %s, got %s", realExpectedMsg, tmpErr.msg)
			}
		} else {
			t.Fatalf("err expected a *SyntaxError, errormsg is %s, but got %T", expectedErrorMsg, err)
		}
	}
}
