//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"reflect"
)

// 一个字符串的测试样例数据，一个字符串可能包含多个元素
type testBytesInfo struct {
	key      string         // 该字符串的标志键
	byteStr  string         // 待解析的字符串，如："true"
	isErrNil bool           // 解析后的err是否期望为nil，true期望为nil，false期望不为nil
	errStr   string         // 解析后如果期望err不为nil，期望在日志中显示的错误提示字符串，，如："overflow error"
	elems    []testElemInfo // 该字符串中要解析出的元素
}

// 一个元素的测试样例数据，一个元素可能包含多个要检测的变量
type testElemInfo struct {
	elemVar any           // 待赋值的目标变量，必须为指针，如：&boolv
	key     string        // 该变量对应的键，如："boolv"
	vars    []testVarInfo // 该元素要检测的所有变量值
}

// 一个变量的测试样例
type testVarInfo struct {
	key        string       // 该变量的标志键
	expected   any          // 该变量的期望值
	got        any          // 该变量获取的值
	basicType  reflect.Kind // 该变量的类型
	gotPointer bool         // 传入的该变量的值是否是它的指针
}
