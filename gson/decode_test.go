//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"testing"
)

// 测试InvalidUnmarshalError的测试样例
func Test_Decode_InvalidUnmarshalError(t *testing.T) {
	if !allTest {
		return
	}
	data := []byte("name : baiqiqi ")
	decoder := NewDecoder()
	defer FreeDecoder(decoder)

	// 测试样例：一个非指针的报错
	keyFoundErrCallBack1 := func(d *Decoder, l *Lexer, isFound bool) bool {
		if isFound {
			testD := TestData{}
			d.SetAnyTarget(testD, true)
		}
		return true
	}
	decoder.KeyEventCall = keyFoundErrCallBack1
	err := work(data, decoder.scan)
	if err == nil {
		t.Fatalf("err expected parse and gson err, got nil")
	}

	// 测试样例：一个普通类型报错
	keyFoundErrCallBack2 := func(d *Decoder, l *Lexer, isFound bool) bool {
		if isFound {
			var testD int
			d.SetAnyTarget(testD, true)
		}
		return true
	}
	decoder.Reset()
	decoder.KeyEventCall = keyFoundErrCallBack2
	err = work(data, decoder.scan)
	if err == nil {
		t.Fatalf("err expected parse and gson err, got nil")
	}

	// 测试样例：一个nil报错
	keyFoundErrCallBack3 := func(d *Decoder, l *Lexer, isFound bool) bool {
		if isFound {
			d.SetAnyTarget(nil, true)
		}
		return true
	}
	decoder.Reset()
	decoder.KeyEventCall = keyFoundErrCallBack3
	err = work(data, decoder.scan)
	if err == nil {
		t.Fatalf("err expected parse and gson err, got nil")
	}
}
