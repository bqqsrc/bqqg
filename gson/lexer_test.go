//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"testing"
)

// 测试没有指定数据种类的情况
// 测试字符串
func Test_Lexer_String(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	l.Write([]byte("baiqiqi4"))
	sv := l.String()
	if sv != "baiqiqi4" {
		t.Fatalf("l.String() expected baiqiqi4, got %s", sv)
	}
}

// 测试无符号整型
// uint8取值0~255
// uint16取值0~65535
// uint32取值0~4294967296
func Test_Lexer_Uint(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	// 测试带符号
	l.Reset()
	l.Write([]byte("-1"))
	uintv, err := l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uint() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uint() expected a invalid syntax err, got nil")
	}
	uint8v, err := l.Uint8()
	if uint8v != 0 {
		t.Fatalf("l.Uint8() expected 0, got %d", uint8v)
	}
	if err == nil {
		t.Fatalf("l.Uint8() expected a invalid syntax err, got nil")
	}
	uint16v, err := l.Uint16()
	if uint16v != 0 {
		t.Fatalf("l.Uint16() expected 0, got %d", uint16v)
	}
	if err == nil {
		t.Fatalf("l.Uint16() expected a invalid syntax err, got nil")
	}
	uint32v, err := l.Uint32()
	if uint32v != 0 {
		t.Fatalf("l.Uint32() expected 0, got %d", uint32v)
	}
	if err == nil {
		t.Fatalf("l.Uint32() expected a invalid syntax err, got nil")
	}
	uint64v, err := l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("+1"))
	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uint() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uint() expected a invalid syntax err, got nil")
	}
	uint8v, err = l.Uint8()
	if uint8v != 0 {
		t.Fatalf("l.Uint8() expected 0, got %d", uint8v)
	}
	if err == nil {
		t.Fatalf("l.Uint8() expected a invalid syntax err, got nil")
	}
	uint16v, err = l.Uint16()
	if uint16v != 0 {
		t.Fatalf("l.Uint16() expected 0, got %d", uint16v)
	}
	if err == nil {
		t.Fatalf("l.Uint16() expected a invalid syntax err, got nil")
	}
	uint32v, err = l.Uint32()
	if uint32v != 0 {
		t.Fatalf("l.Uint32() expected 0, got %d", uint32v)
	}
	if err == nil {
		t.Fatalf("l.Uint32() expected a invalid syntax err, got nil")
	}
	uint64v, err = l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a invalid syntax err, got nil")
	}

	// 测试0，所有的都可以
	l.Reset()
	l.Write([]byte("0"))
	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uint() expected 0, got %d", uintv)
	}
	if err != nil {
		t.Fatalf("l.Uint() expected a nil err, got nil")
	}
	uint8v, err = l.Uint8()
	if uint8v != 0 {
		t.Fatalf("l.Uint8() expected 0, got %d", uint8v)
	}
	if err != nil {
		t.Fatalf("l.Uint8() expected a nil err, got %s", err)
	}
	uint16v, err = l.Uint16()
	if uint16v != 0 {
		t.Fatalf("l.Uint16() expected 0, got %d", uint16v)
	}
	if err != nil {
		t.Fatalf("l.Uint16() expected a nil err, got %s", err)
	}
	uint32v, err = l.Uint32()
	if uint32v != 0 {
		t.Fatalf("l.Uint32() expected 0, got %d", uint32v)
	}
	if err != nil {
		t.Fatalf("l.Uint32() expected a nil err, got %s", err)
	}
	uint64v, err = l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err != nil {
		t.Fatalf("l.Uint64() expected a nil err, got %s", err)
	}

	// 测试小于等于255
	l.Reset()
	l.Write([]byte("255"))
	uintv, err = l.Uint()
	if uintv != 255 {
		t.Fatalf("l.Uint() expected 255, got %d", uintv)
	}
	if err != nil {
		t.Fatalf("l.Uint() expected a nil err, got %s", err)
	}
	uint8v, err = l.Uint8()
	if uint8v != 255 {
		t.Fatalf("l.Uint8() expected 255, got %d", uint8v)
	}
	if err != nil {
		t.Fatalf("l.Uint8() expected a nil err, got %s", err)
	}
	uint16v, err = l.Uint16()
	if uint16v != 255 {
		t.Fatalf("l.Uint16() expected 255, got %d", uint16v)
	}
	if err != nil {
		t.Fatalf("l.Uint16() expected a nil err, got %s", err)
	}
	uint32v, err = l.Uint32()
	if uint32v != 255 {
		t.Fatalf("l.Uint32() expected 255, got %d", uint32v)
	}
	if err != nil {
		t.Fatalf("l.Uint32() expected a nil err, got %s", err)
	}
	uint64v, err = l.Uint64()
	if uint64v != 255 {
		t.Fatalf("l.Uint64() expected 255, got %d", uint64v)
	}
	if err != nil {
		t.Fatalf("l.Uint64() expected a nil err, got %s", err)
	}

	// 测试大于等于256
	l.Reset()
	l.Write([]byte("256"))
	uint8v, err = l.Uint8()
	if uint8v != 0 {
		t.Fatalf("l.Uint8() expected 0, got %d", uint8v)
	}
	if err == nil {
		t.Fatalf("l.Uint8() expected a overflow err, got nil")
	}
	uint16v, err = l.Uint16()
	if uint16v != 256 {
		t.Fatalf("l.Uint16() expected 256, got %d", uint16v)
	}
	if err != nil {
		t.Fatalf("l.Uint16() expected a nil err, got %s", err)
	}

	// 测试小于等于65535
	l.Reset()
	l.Write([]byte("65535"))
	uint16v, err = l.Uint16()
	if uint16v != 65535 {
		t.Fatalf("l.Uint16() expected 65535, got %d", uint16v)
	}
	if err != nil {
		t.Fatalf("l.Uint16() expected a nil err, got %s", err)
	}
	uint32v, err = l.Uint32()
	if uint32v != 65535 {
		t.Fatalf("l.Uint32() expected 65535, got %d", uint32v)
	}
	if err != nil {
		t.Fatalf("l.Uint32() expected a nil err, got %s", err)
	}

	// 测试大于等于65536
	l.Reset()
	l.Write([]byte("65536"))
	uint16v, err = l.Uint16()
	if uint16v != 0 {
		t.Fatalf("l.Uint16() expected 0, got %d", uint16v)
	}
	if err == nil {
		t.Fatalf("l.Uint16() expected a overflow err, got nil")
	}
	uint32v, err = l.Uint32()
	if uint32v != 65536 {
		t.Fatalf("l.Uint32() expected 65536, got %d", uint32v)
	}
	if err != nil {
		t.Fatalf("l.Uint32() expected a nil err, got %s", err)
	}

	// 测试一个小于等于4294967295
	l.Reset()
	l.Write([]byte("4294967295"))
	uint32v, err = l.Uint32()
	if uint32v != 4294967295 {
		t.Fatalf("l.Uint32() expected 4294967295, got %d", uint32v)
	}
	if err != nil {
		t.Fatalf("l.Uint32() expected a nil err, got %s", err)
	}
	uint64v, err = l.Uint64()
	if uint64v != 4294967295 {
		t.Fatalf("l.Uint64() expected 4294967295, got %d", uint64v)
	}
	if err != nil {
		t.Fatalf("l.Uint64() expected a nil err, got %s", err)
	}
	uintv, err = l.Uint()
	if uintv != 4294967295 {
		t.Fatalf("l.Uint() expected 4294967295, got %d", uintv)
	}
	if err != nil {
		t.Fatalf("l.Uint() expected a nil err, got %s", err)
	}

	// 测试一个大于等于4294967296
	l.Reset()
	l.Write([]byte("4294967296"))
	uint64v, err = l.Uint64()
	if uint64v != 4294967296 {
		t.Fatalf("l.Uint64() expected 4294967296, got %d", uint64v)
	}
	if err != nil {
		t.Fatalf("l.Uint64() expected a nil err, got %s", err)
	}
	uintv, err = l.Uint()
	if uintv != 4294967296 {
		t.Fatalf("l.Uint() expected 4294967296, got %d", uintv)
	}
	if err != nil {
		t.Fatalf("l.Uint() expected a nil err, got %s", err)
	}
	uint32v, err = l.Uint32()
	if uint32v != 0 {
		t.Fatalf("l.Uint32() expected 0, got %d", uint32v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a overflow err, got nil")
	}

	// 测试一个不是uint类型的字符串
	l.Reset()
	l.Write([]byte("1e2"))
	uint64v, err = l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("36eE7"))
	uint64v, err = l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("98.77"))
	uint64v, err = l.Uint64()
	if uint64v != 0 {
		t.Fatalf("l.Uint64() expected 0, got %d", uint64v)
	}
	if err == nil {
		t.Fatalf("l.Uint64() expected a invalid syntax err, got nil")
	}
}

// 测试有符号整型
// int8取值-128~127
// int16取值-32768~32767
// int32取值-2147483648~2147483647
func Test_Lexer_Int(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	// 测试带符号
	l.Reset()
	l.Write([]byte("-1"))
	intv, err := l.Int()
	if intv != -1 {
		t.Fatalf("l.Int() expected -1, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a invalid syntax err, got %s", err)
	}
	int8v, err := l.Int8()
	if int8v != -1 {
		t.Fatalf("l.Int8() expected -1, got %d", int8v)
	}
	if err != nil {
		t.Fatalf("l.Int8() expected a invalid syntax err, got %s", err)
	}
	int16v, err := l.Int16()
	if int16v != -1 {
		t.Fatalf("l.Int16() expected -1, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a invalid syntax err, got %s", err)
	}
	int32v, err := l.Int32()
	if int32v != -1 {
		t.Fatalf("l.Int32() expected -1, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a invalid syntax err, got %s", err)
	}
	int64v, err := l.Int64()
	if int64v != -1 {
		t.Fatalf("l.Int64() expected -1, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a invalid syntax err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("+1"))
	intv, err = l.Int()
	if intv != 1 {
		t.Fatalf("l.Int() expected 1, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int8v, err = l.Int8()
	if int8v != 1 {
		t.Fatalf("l.Int8() expected 1, got %d", int8v)
	}
	if err != nil {
		t.Fatalf("l.Int8() expected a nil err, got %s", err)
	}
	int16v, err = l.Int16()
	if int16v != 1 {
		t.Fatalf("l.Int16() expected 1, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 1 {
		t.Fatalf("l.Int32() expected 1, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != 1 {
		t.Fatalf("l.Int64() expected 1, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}

	// 测试0，所有的都可以
	l.Reset()
	l.Write([]byte("0"))
	intv, err = l.Int()
	if intv != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int8v, err = l.Int8()
	if int8v != 0 {
		t.Fatalf("l.Int8() expected 0, got %d", int8v)
	}
	if err != nil {
		t.Fatalf("l.Int8() expected a nil err, got %s", err)
	}
	int16v, err = l.Int16()
	if int16v != 0 {
		t.Fatalf("l.Int16() expected 0, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 0 {
		t.Fatalf("l.Int32() expected 0, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != 0 {
		t.Fatalf("l.Int64() expected 0, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}

	// 测试小于等于127
	l.Reset()
	l.Write([]byte("127"))
	intv, err = l.Int()
	if intv != 127 {
		t.Fatalf("l.Int() expected 127, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int8v, err = l.Int8()
	if int8v != 127 {
		t.Fatalf("l.Int8() expected 127, got %d", int8v)
	}
	if err != nil {
		t.Fatalf("l.Int8() expected a nil err, got %s", err)
	}
	int16v, err = l.Int16()
	if int16v != 127 {
		t.Fatalf("l.Int16() expected 127, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 127 {
		t.Fatalf("l.Int32() expected 127, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != 127 {
		t.Fatalf("l.Int64() expected 127, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}

	// 测试大于等于128
	l.Reset()
	l.Write([]byte("128"))
	int8v, err = l.Int8()
	if int8v != 0 {
		t.Fatalf("l.Int8() expected 0, got %d", int8v)
	}
	if err == nil {
		t.Fatalf("l.Int8() expected a overflow err, got nil")
	}
	int16v, err = l.Int16()
	if int16v != 128 {
		t.Fatalf("l.Int16() expected 128, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}

	// 测试小于等于32767
	l.Reset()
	l.Write([]byte("32767"))
	int16v, err = l.Int16()
	if int16v != 32767 {
		t.Fatalf("l.Int16() expected 32767, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 32767 {
		t.Fatalf("l.Int32() expected 32767, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}

	// 测试大于等于32768
	l.Reset()
	l.Write([]byte("32768"))
	int16v, err = l.Int16()
	if int16v != 0 {
		t.Fatalf("l.Int16() expected 0, got %d", int16v)
	}
	if err == nil {
		t.Fatalf("l.Int16() expected a overflow err, got nil")
	}
	int32v, err = l.Int32()
	if int32v != 32768 {
		t.Fatalf("l.Int32() expected 32768, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}

	// 测试一个小于等于2147483647
	l.Reset()
	l.Write([]byte("2147483647"))
	int32v, err = l.Int32()
	if int32v != 2147483647 {
		t.Fatalf("l.Int32() expected 2147483647, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != 2147483647 {
		t.Fatalf("l.Int64() expected 2147483647, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}
	intv, err = l.Int()
	if intv != 2147483647 {
		t.Fatalf("l.Int() expected 2147483647, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}

	// 测试一个大于等于2147483648
	l.Reset()
	l.Write([]byte("2147483648"))
	int64v, err = l.Int64()
	if int64v != 2147483648 {
		t.Fatalf("l.Int64() expected 2147483648, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}
	intv, err = l.Int()
	if intv != 2147483648 {
		t.Fatalf("l.Int() expected 2147483648, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 0 {
		t.Fatalf("l.Int32() expected 0, got %d", int32v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a overflow err, got nil")
	}

	// 测试大于等于-128
	l.Reset()
	l.Write([]byte("-128"))
	intv, err = l.Int()
	if intv != -128 {
		t.Fatalf("l.Int() expected -128, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int8v, err = l.Int8()
	if int8v != -128 {
		t.Fatalf("l.Int8() expected -128, got %d", int8v)
	}
	if err != nil {
		t.Fatalf("l.Int8() expected a nil err, got %s", err)
	}
	int16v, err = l.Int16()
	if int16v != -128 {
		t.Fatalf("l.Int16() expected -128, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != -128 {
		t.Fatalf("l.Int32() expected -128, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != -128 {
		t.Fatalf("l.Int64() expected -128, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}

	// 测试小于等于-129
	l.Reset()
	l.Write([]byte("-129"))
	int8v, err = l.Int8()
	if int8v != 0 {
		t.Fatalf("l.Int8() expected 0, got %d", int8v)
	}
	if err == nil {
		t.Fatalf("l.Int8() expected a overflow err, got nil")
	}
	int16v, err = l.Int16()
	if int16v != -129 {
		t.Fatalf("l.Int16() expected -128, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}

	// 测试大于等于-32768
	l.Reset()
	l.Write([]byte("-32768"))
	int16v, err = l.Int16()
	if int16v != -32768 {
		t.Fatalf("l.Int16() expected -32768, got %d", int16v)
	}
	if err != nil {
		t.Fatalf("l.Int16() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != -32768 {
		t.Fatalf("l.Int32() expected -32768, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}

	// 测试小于等于-32769
	l.Reset()
	l.Write([]byte("-32769"))
	int16v, err = l.Int16()
	if int16v != 0 {
		t.Fatalf("l.Int16() expected 0, got %d", int16v)
	}
	if err == nil {
		t.Fatalf("l.Int16() expected a overflow err, got nil")
	}
	int32v, err = l.Int32()
	if int32v != -32769 {
		t.Fatalf("l.Int32() expected -32769, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}

	// 测试一个大于等于-2147483648
	l.Reset()
	l.Write([]byte("-2147483648"))
	int32v, err = l.Int32()
	if int32v != -2147483648 {
		t.Fatalf("l.Int32() expected -2147483648, got %d", int32v)
	}
	if err != nil {
		t.Fatalf("l.Int32() expected a nil err, got %s", err)
	}
	int64v, err = l.Int64()
	if int64v != -2147483648 {
		t.Fatalf("l.Int64() expected -2147483648, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}
	intv, err = l.Int()
	if intv != -2147483648 {
		t.Fatalf("l.Int() expected -2147483648, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}

	// 测试一个小于等于-2147483649
	l.Reset()
	l.Write([]byte("-2147483649"))
	int64v, err = l.Int64()
	if int64v != -2147483649 {
		t.Fatalf("l.Int64() expected -2147483649, got %d", int64v)
	}
	if err != nil {
		t.Fatalf("l.Int64() expected a nil err, got %s", err)
	}
	intv, err = l.Int()
	if intv != -2147483649 {
		t.Fatalf("l.Int() expected -2147483649, got %d", intv)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}
	int32v, err = l.Int32()
	if int32v != 0 {
		t.Fatalf("l.Int32() expected 0, got %d", int32v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a overflow err, got nil")
	}

	// 测试一些非有符号整型字符串
	l.Reset()
	l.Write([]byte("true"))
	int64v, err = l.Int64()
	if int64v != 0 {
		t.Fatalf("l.Int64() expected 0, got %d", int64v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("89.773"))
	int64v, err = l.Int64()
	if int64v != 0 {
		t.Fatalf("l.Int64() expected 0, got %d", int64v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("13e8"))
	int64v, err = l.Int64()
	if int64v != 0 {
		t.Fatalf("l.Int64() expected 0, got %d", int64v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("13Ee8"))
	int64v, err = l.Int64()
	if int64v != 0 {
		t.Fatalf("l.Int64() expected 0, got %d", int64v)
	}
	if err == nil {
		t.Fatalf("l.Int64() expected a invalid syntax err, got nil")
	}
}

func Test_Lexer_Float(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	// 测试普通整型和非整型数值
	l.Write([]byte("9887"))
	float32v, err := l.Float32()
	if float32v != 9887 {
		t.Fatalf("l.Float32() expected 9887, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err := l.Float64()
	if float64v != 9887 {
		t.Fatalf("l.Float64() expected 9887, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("-198"))
	float32v, err = l.Float32()
	if float32v != -198 {
		t.Fatalf("l.Float32() expected -198, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != -198 {
		t.Fatalf("l.Float64() expected -198, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("+267"))
	float32v, err = l.Float32()
	if float32v != 267 {
		t.Fatalf("l.Float32() expected 267, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 267 {
		t.Fatalf("l.Float64() expected 267, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试普通正值数字
	l.Reset()
	l.Write([]byte("9.8887"))
	float32v, err = l.Float32()
	if float32v != 9.8887 {
		t.Fatalf("l.Float32() expected 9.8887, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 9.8887 {
		t.Fatalf("l.Float64() expected 9.8887, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试普通负值数字
	l.Reset()
	l.Write([]byte("-8.9967"))
	float32v, err = l.Float32()
	if float32v != -8.9967 {
		t.Fatalf("l.Float32() expected -8.9967, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != -8.9967 {
		t.Fatalf("l.Float64() expected -8.9967, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试不带整数的字符
	l.Reset()
	l.Write([]byte(".9967"))
	float32v, err = l.Float32()
	if float32v != 0.9967 {
		t.Fatalf("l.Float32() expected 0.9967, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 0.9967 {
		t.Fatalf("l.Float64() expected 0.9967, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("-.887"))
	float32v, err = l.Float32()
	if float32v != -0.887 {
		t.Fatalf("l.Float32() expected 0.887, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != -0.887 {
		t.Fatalf("l.Float64() expected 0.887, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("+.663"))
	float32v, err = l.Float32()
	if float32v != 0.663 {
		t.Fatalf("l.Float32() expected 0.663, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 0.663 {
		t.Fatalf("l.Float64() expected 0.663, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试带正号的数字
	l.Reset()
	l.Write([]byte("+8.9967"))
	float32v, err = l.Float32()
	if float32v != 8.9967 {
		t.Fatalf("l.Float32() expected 8.9967, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 8.9967 {
		t.Fatalf("l.Float64() expected 8.9967, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试带e的数字
	l.Reset()
	l.Write([]byte("3.5e-3"))
	float32v, err = l.Float32()
	if float32v != 0.0035 {
		t.Fatalf("l.Float32() expected 0.0035, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 0.0035 {
		t.Fatalf("l.Float64() expected 0.0035, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("7.5e+4"))
	float32v, err = l.Float32()
	if float32v != 75000 {
		t.Fatalf("l.Float32() expected 75000, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 75000 {
		t.Fatalf("l.Float64() expected 75000, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试溢出临界值情况3.4*10^38
	l.Reset()
	l.Write([]byte("3.4e38"))
	float32v, err = l.Float32()
	if float32v != 3.4e38 {
		t.Fatalf("l.Float32() expected 3.4e38, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 3.4e38 {
		t.Fatalf("l.Float64() expected 3.4e38, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}
	l.Reset()
	l.Write([]byte("3.4e39"))
	float32v, err = l.Float32()
	if float32v != 0 {
		t.Fatalf("l.Float32() expected 0, got %f", float32v)
	}
	if err == nil {
		t.Fatalf("l.Float32() expected a overflow err, got nil")
	}
	float64v, err = l.Float64()
	if float64v != 3.4e39 {
		t.Fatalf("l.Float64() expected 3.4e39, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}
	l.Reset()
	l.Write([]byte("3.5e38"))
	float32v, err = l.Float32()
	if float32v != 0 {
		t.Fatalf("l.Float32() expected 0, got %f", float32v)
	}
	if err == nil {
		t.Fatalf("l.Float32() expected a overflow err, got nil")
	}
	float64v, err = l.Float64()
	if float64v != 3.5e38 {
		t.Fatalf("l.Float64() expected 3.5e38, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试溢出临界值情况1.17*10^-38
	//TODO 似乎这种溢出情况没有问题？？？是被直接按照零值处理了吧
	l.Reset()
	l.Write([]byte("1.17e-38"))
	float32v, err = l.Float32()
	if float32v != 1.17e-38 {
		t.Fatalf("l.Float32() expected 1.17e-38, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 1.17e-38 {
		t.Fatalf("l.Float64() expected 1.17e-38, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}
	l.Reset()
	l.Write([]byte("1.17e-39"))
	float32v, err = l.Float32()
	if float32v != 1.17e-39 {
		t.Fatalf("l.Float32() expected 1.17e-39, got %f", float32v)
	}
	if err != nil {
		t.Fatalf("l.Float32() expected a nil err, got %s", err)
	}
	float64v, err = l.Float64()
	if float64v != 1.17e-39 {
		t.Fatalf("l.Float64() expected 1.17e-39, got %f", float64v)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	// 测试一些非float型字符串
	l.Reset()
	l.Write([]byte("false"))
	float64v, err = l.Float64()
	if float64v != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", float64v)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a invalid syntax err, got nil")
	}

	l.Reset()
	l.Write([]byte("nil"))
	float64v, err = l.Float64()
	if float64v != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", float64v)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a invalid syntax err, got nil")
	}
}

// 测试布尔型
func Test_Lexer_Bool(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	l.Write([]byte("true"))
	boolv, err := l.Bool()
	if !boolv {
		t.Fatalf("l.Bool() expected true, got false")
	}
	if err != nil {
		t.Fatalf("l.Bool() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("false"))
	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err != nil {
		t.Fatalf("l.Bool() expected a nil err, got %s", err)
	}

	l.Reset()
	l.Write([]byte("falses"))
	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}
}

// 测试Nil
func Test_Lexer_Nil(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	canNil := l.CanNil()
	if !canNil {
		t.Fatalf("l.CanNil() expected true, got false")
	}

	l.Write([]byte("nil"))
	canNil = l.CanNil()
	if !canNil {
		t.Fatalf("l.CanNil() expected true, got false")
	}

	l.Write([]byte("nill"))
	canNil = l.CanNil()
	if canNil {
		t.Fatalf("l.CanNil() expected false, got true")
	}
}

func Test_Lexer_ValueType(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	// 测试StringV类型
	l.Write([]byte("falses"))
	l.SetValueType(stringV)
	stringv := l.String()
	if stringv != "falses" {
		t.Fatalf("l.String() expected falses, got %s", stringv)
	}

	uintv, err := l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err := l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a parse err, got nil")
	}

	floatv, err := l.Float64()
	if floatv != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", floatv)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a parse err, got nil")
	}

	boolv, err := l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试NilV类型
	l.Reset()
	l.Write([]byte("nil"))
	l.SetValueType(nilV)
	stringv = l.String()
	if stringv != "nil" {
		t.Fatalf("l.String() expected nil, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err = l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a parse err, got nil")
	}

	floatv, err = l.Float64()
	if floatv != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", floatv)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a parse err, got nil")
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if !boolv {
		t.Fatalf("l.CanNil() expected true, got false")
	}

	// 测试BoolV类型
	l.Reset()
	l.Write([]byte("true"))
	l.SetValueType(boolV)
	stringv = l.String()
	if stringv != "true" {
		t.Fatalf("l.String() expected true, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err = l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a parse err, got nil")
	}

	floatv, err = l.Float64()
	if floatv != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", floatv)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a parse err, got nil")
	}

	boolv, err = l.Bool()
	if !boolv {
		t.Fatalf("l.Bool() expected true, got false")
	}
	if err != nil {
		t.Fatalf("l.Bool() expected a nil err, got %s", err)
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试FloatV类型
	l.Reset()
	l.Write([]byte("39e-3"))
	l.SetValueType(floatV)
	stringv = l.String()
	if stringv != "39e-3" {
		t.Fatalf("l.String() expected 39e-3, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err = l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a parse err, got nil")
	}

	floatv, err = l.Float64()
	if floatv != 0.039 {
		t.Fatalf("l.Float64() expected 0.039, got %f", floatv)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试FloatV类型
	l.Reset()
	l.Write([]byte("9.887"))
	l.SetValueType(floatV)
	stringv = l.String()
	if stringv != "9.887" {
		t.Fatalf("l.String() expected 9.887, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err = l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a parse err, got nil")
	}

	floatv, err = l.Float64()
	if floatv != 9.887 {
		t.Fatalf("l.Float64() expected 0.039, got %f", floatv)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试IntV类型
	l.Reset()
	l.Write([]byte("-346"))
	l.SetValueType(intV)
	stringv = l.String()
	if stringv != "-346" {
		t.Fatalf("l.String() expected -346, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a parse err, got nil")
	}

	intV, err = l.Int()
	if intV != -346 {
		t.Fatalf("l.Int() expected -346, got %d", intV)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}

	floatv, err = l.Float64()
	if floatv != -346 {
		t.Fatalf("l.Float64() expected -346, got %f", floatv)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试UintV类型
	l.Reset()
	l.Write([]byte("21378"))
	l.SetValueType(uintV)
	stringv = l.String()
	if stringv != "21378" {
		t.Fatalf("l.String() expected 21378, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 21378 {
		t.Fatalf("l.Uintv() expected 21378, got %d", uintv)
	}
	if err != nil {
		t.Fatalf("l.Uintv() expected a nil err, got %s", err)
	}

	intV, err = l.Int()
	if intV != 21378 {
		t.Fatalf("l.Int() expected 21378, got %d", intV)
	}
	if err != nil {
		t.Fatalf("l.Int() expected a nil err, got %s", err)
	}

	floatv, err = l.Float64()
	if floatv != 21378 {
		t.Fatalf("l.Float64() expected 21378, got %f", floatv)
	}
	if err != nil {
		t.Fatalf("l.Float64() expected a nil err, got %s", err)
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a parse err, got nil")
	}

	boolv = l.CanNil()
	if boolv {
		t.Fatalf("l.CanNil() expected false, got true")
	}

	// 测试无值的情况
	l.Reset()
	stringv = l.String()
	if stringv != "" {
		t.Fatalf("l.String() expected a empty_string, got %s", stringv)
	}

	uintv, err = l.Uint()
	if uintv != 0 {
		t.Fatalf("l.Uintv() expected 0, got %d", uintv)
	}
	if err == nil {
		t.Fatalf("l.Uintv() expected a invalid syntax err, got nil")
	}

	intV, err = l.Int()
	if intV != 0 {
		t.Fatalf("l.Int() expected 0, got %d", intV)
	}
	if err == nil {
		t.Fatalf("l.Int() expected a invalid syntax err, got nil")
	}

	floatv, err = l.Float64()
	if floatv != 0 {
		t.Fatalf("l.Float64() expected 0, got %f", floatv)
	}
	if err == nil {
		t.Fatalf("l.Float64() expected a invalid syntax err, got nil")
	}

	boolv, err = l.Bool()
	if boolv {
		t.Fatalf("l.Bool() expected false, got true")
	}
	if err == nil {
		t.Fatalf("l.Bool() expected a invalid syntax err, got nil")
	}

	boolv = l.CanNil()
	if !boolv {
		t.Fatalf("l.CanNil() expected true, got false")
	}
}

func Test_Lexer_Interface(t *testing.T) {
	l := NewLexer()
	defer FreeLexer(l)

	iv := l.Interface()
	if iv != nil {
		t.Fatalf("l.Interface() expected nil, got %v", iv)
	}

	l.Reset()
	l.Write([]byte("nil"))
	iv = l.Interface()
	if iv != nil {
		t.Fatalf("l.Interface() expected nil, got %v", iv)
	}

	l.Reset()
	l.Write([]byte("nil"))
	l.SetValueType(nilV)
	iv = l.Interface()
	if iv != nil {
		t.Fatalf("l.Interface() expected nil, got %v", iv)
	}

	// 测试UnknownV
	l.Reset()
	l.Write([]byte("true"))
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "true" {
		t.Fatalf("l.Interface() expected true, got %s", tmpIv)
	}

	l.Reset()
	l.Write([]byte("567"))
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "567" {
		t.Fatalf("l.Interface() expected 567, got %s", tmpIv)
	}

	l.Reset()
	l.Write([]byte("-987"))
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "-987" {
		t.Fatalf("l.Interface() expected -987, got %s", tmpIv)
	}

	l.Reset()
	l.Write([]byte("98.88"))
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "98.88" {
		t.Fatalf("l.Interface() expected 98.88, got %s", tmpIv)
	}

	l.Reset()
	l.Write([]byte("students"))
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "students" {
		t.Fatalf("l.Interface() expected students, got %s", tmpIv)
	}

	// 测试已知类型
	l.Reset()
	l.Write([]byte("true"))
	l.SetValueType(boolV)
	iv = l.Interface()
	if tmpIv, ok := iv.(bool); !ok {
		t.Fatalf("l.Interface() expected a bool-type, got %T", iv)
	} else if !tmpIv {
		t.Fatalf("l.Interface() expected true, got false")
	}

	l.Reset()
	l.Write([]byte("12567"))
	l.SetValueType(uintV)
	iv = l.Interface()
	if tmpIv, ok := iv.(int64); !ok {
		t.Fatalf("l.Interface() expected a int64-type, got %T", iv)
	} else if tmpIv != 12567 {
		t.Fatalf("l.Interface() expected 12567, got %d", tmpIv)
	}

	l.Reset()
	l.Write([]byte("-980977"))
	l.SetValueType(intV)
	iv = l.Interface()
	if tmpIv, ok := iv.(int64); !ok {
		t.Fatalf("l.Interface() expected a int64-type, got %T", iv)
	} else if tmpIv != -980977 {
		t.Fatalf("l.Interface() expected -980977, got %d", tmpIv)
	}

	l.Reset()
	l.Write([]byte("98.88e3"))
	l.SetValueType(floatV)
	iv = l.Interface()
	if tmpIv, ok := iv.(float64); !ok {
		t.Fatalf("l.Interface() expected a float64-type, got %T", iv)
	} else if tmpIv != 98880 {
		t.Fatalf("l.Interface() expected 98880, got %f", tmpIv)
	}

	l.Reset()
	l.Write([]byte("this is a students"))
	l.SetValueType(stringV)
	iv = l.Interface()
	if tmpIv, ok := iv.(string); !ok {
		t.Fatalf("l.Interface() expected a string-type, got %T", iv)
	} else if tmpIv != "this is a students" {
		t.Fatalf("l.Interface() expected this is a students, got %s", tmpIv)
	}
}
