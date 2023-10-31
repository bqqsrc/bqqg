//  Copyright (C) 晓白齐齐,版权所有.

package gson

// 解析器接口：所有使用gson解析的数据类型都要实现这个接口
type Unmarshaler interface {
	UnmarshalByKey(key string, w *Lexer) error
	UnmarshalByIndex(index int, w *Lexer) error
}

// 对未定义键的解析处理，根据你的需求觉得如何处理
type UnknownsUnmarshaler interface {
	UnmarshalUnknownKey(key string, in *Lexer) error
	UnmarshalUnknownIndex(index int, in *Lexer) error
}
