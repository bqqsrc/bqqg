//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"reflect"
	"strconv"
)

type DecodeError struct {
	key string
	msg string // 错误信息
}

func (e *DecodeError) Error() string {
	if e.key != "" {
		return "convert key[" + e.key + "] err: " + e.msg
	} else {
		return "decode err: " + e.msg
	}
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "gson err: SetAnyTarget(nil, true/false); can't set a nil as target"
	}
	if e.Type.Kind() != reflect.Pointer {
		return "gson err: SetAnyTarget(non-pointer " + e.Type.String() + ", true/false); can't set a non-pointer as target"
	}
	return "gson err: SetAnyTarget(nil " + e.Type.String() + ", true/false); can't set a nil " + e.Type.String() + " as target"
}

type ParseError struct {
	value     string
	valueType string
	msg       string // 错误信息
}

func (e *ParseError) Error() string {
	errMsg := `parse err: "` + e.value + `" can't parse to "` + e.valueType + `"`
	if e.msg != "" {
		errMsg = errMsg + ", errMsg: " + e.msg
	}
	return errMsg
}

type SyntaxError struct {
	msg        string // 错误信息
	line       int64  // 错误所在的行
	lineOffset int64  // 错误所在行的偏移量

}

func (e *SyntaxError) Error() string {
	return "err: line: " + strconv.FormatInt(e.line, 10) + ", offset: " + strconv.FormatInt(e.lineOffset, 10) + "; " + e.msg
}
