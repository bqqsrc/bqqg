//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"bytes"
	"strconv"
	"sync"
)

// 扫描后可能的数据类型
const (
	unknownV = iota // 未知
	stringV         // 字符串
	nilV            // 空
	boolV           // 布尔型
	floatV          // 浮点型
	intV            // 有符号整型
	uintV           // 无符号整型
)

type Lexer struct {
	valueType int          // 数据的可能类型，当dataKind为dataValue时，表示最合适的数据类型
	buffer    bytes.Buffer // 当前数据的字节数组
}

var lexerPool = sync.Pool{
	New: func() any {
		return &Lexer{}
	},
}

func NewLexer() *Lexer {
	lexer := lexerPool.Get().(*Lexer)
	lexer.Reset()
	return lexer
}

func FreeLexer(lexer *Lexer) {
	lexer.buffer.Reset()
	lexerPool.Put(lexer)
}

func (l *Lexer) Reset() {
	l.valueType = unknownV
	l.buffer.Reset()
}

func (l *Lexer) SetValueType(valueType int) {
	l.valueType = valueType
}

func (l *Lexer) GetValueType() int {
	return l.valueType
}

func (l *Lexer) Write(b []byte) {
	l.buffer.Write(b)
}

func (l *Lexer) Len() int {
	return l.buffer.Len()
}

func (l *Lexer) String() string {
	return l.buffer.String()
}

func (l *Lexer) Bytes() []byte {
	return l.buffer.Bytes()
}

func (l *Lexer) Uint8() (uint8, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == unknownV {
		n, err := strconv.ParseUint(sv, 10, 8)
		if err != nil {
			return 0, &ParseError{sv, "uint8", err.Error()}
		} else {
			return uint8(n), nil
		}
	} else {
		return 0, &ParseError{sv, "uint8", ""}
	}
}

func (l *Lexer) Uint16() (uint16, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == unknownV {
		n, err := strconv.ParseUint(sv, 10, 16)
		if err != nil {
			return 0, &ParseError{sv, "uint16", err.Error()}
		} else {
			return uint16(n), nil
		}
	} else {
		return 0, &ParseError{sv, "uint16", ""}
	}
}

func (l *Lexer) Uint32() (uint32, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == unknownV {
		n, err := strconv.ParseUint(sv, 10, 32)
		if err != nil {
			return 0, &ParseError{sv, "uint32", err.Error()}
		} else {
			return uint32(n), nil
		}
	} else {
		return 0, &ParseError{sv, "uint32", ""}
	}
}

func (l *Lexer) Uint64() (uint64, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == unknownV {
		n, err := strconv.ParseUint(sv, 10, 64)
		if err != nil {
			return 0, &ParseError{sv, "uint64", err.Error()}
		} else {
			return n, nil
		}
	} else {
		return 0, &ParseError{sv, "uint64", ""}
	}
}

func (l *Lexer) Uint() (uint, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == unknownV {
		n, err := strconv.ParseUint(sv, 10, 64)
		if err != nil {
			return 0, &ParseError{sv, "uint", err.Error()}
		} else {
			return uint(n), nil
		}
	} else {
		return 0, &ParseError{sv, "uint", ""}
	}
}

func (l *Lexer) Int8() (int8, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseInt(sv, 10, 8)
		if err != nil {
			return 0, &ParseError{sv, "int8", err.Error()}
		} else {
			return int8(n), nil
		}
	} else {
		return 0, &ParseError{sv, "int8", ""}
	}
}

func (l *Lexer) Int16() (int16, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseInt(sv, 10, 16)
		if err != nil {
			return 0, &ParseError{sv, "int16", err.Error()}
		} else {
			return int16(n), nil
		}
	} else {
		return 0, &ParseError{sv, "int16", ""}
	}
}

func (l *Lexer) Int32() (int32, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseInt(sv, 10, 32)
		if err != nil {
			return 0, &ParseError{sv, "int32", err.Error()}
		} else {
			return int32(n), nil
		}
	} else {
		return 0, &ParseError{sv, "int32", ""}
	}
}

func (l *Lexer) Int64() (int64, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseInt(sv, 10, 64)
		if err != nil {
			return 0, &ParseError{sv, "int64", err.Error()}
		} else {
			return n, nil
		}
	} else {
		return 0, &ParseError{sv, "int64", ""}
	}
}

func (l *Lexer) Int() (int, error) {
	sv := l.String()
	if l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseInt(sv, 10, 64)
		if err != nil {
			return 0, &ParseError{sv, "int", err.Error()}
		} else {
			return int(n), nil
		}
	} else {
		return 0, &ParseError{sv, "int", ""}
	}
}

func (l *Lexer) Float32() (float32, error) {
	sv := l.String()
	if l.valueType == floatV || l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseFloat(sv, 32)
		if err != nil {
			return 0, &ParseError{sv, "float32", err.Error()}
		} else {
			return float32(n), nil
		}
	} else {
		return 0, &ParseError{sv, "float32", ""}
	}
}

func (l *Lexer) Float64() (float64, error) {
	sv := l.String()
	if l.valueType == floatV || l.valueType == uintV || l.valueType == intV || l.valueType == unknownV {
		n, err := strconv.ParseFloat(sv, 64)
		if err != nil {
			return 0, &ParseError{sv, "float64", err.Error()}
		} else {
			return n, nil
		}
	} else {
		return 0, &ParseError{sv, "float64", ""}
	}
}

func (l *Lexer) Bool() (bool, error) {
	sv := l.String()
	if sv == "false" {
		return false, nil
	}
	if sv == "true" {
		return true, nil
	}
	return false, &ParseError{sv, "bool", ""}
}

func (l *Lexer) CanNil() bool {
	if l.valueType == nilV || l.Len() == 0 {
		return true
	} else if l.valueType == unknownV {
		sv := l.String()
		if sv == "nil" {
			return true
		}
	}
	return false
}

// 添加返回interface的接口
func (l *Lexer) Interface() any {
	if l.CanNil() {
		return nil
	}
	switch l.valueType {
	case intV, uintV:
		ret, _ := l.Int64()
		return ret
	case floatV:
		ret, _ := l.Float64()
		return ret
	case boolV:
		ret, _ := l.Bool()
		return ret
	}
	return l.String()
}
