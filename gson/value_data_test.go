//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"reflect"
	"testing"
)

func checkInt(t *testing.T, key string, expected, got int) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkInt8(t *testing.T, key string, expected, got int8) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkInt16(t *testing.T, key string, expected, got int16) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkInt32(t *testing.T, key string, expected, got int32) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkInt64(t *testing.T, key string, expected, got int64) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkUint(t *testing.T, key string, expected, got uint) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkUint8(t *testing.T, key string, expected, got uint8) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkUint16(t *testing.T, key string, expected, got uint16) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkUint32(t *testing.T, key string, expected, got uint32) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkUint64(t *testing.T, key string, expected, got uint64) {
	if expected != got {
		t.Fatalf("%s expected %d, got %d", key, expected, got)
	}
}
func checkFloat32(t *testing.T, key string, expected, got float32) {
	if expected != got {
		t.Fatalf("%s expected %f, got %f", key, expected, got)
	}
}
func checkFloat64(t *testing.T, key string, expected, got float64) {
	if expected != got {
		t.Fatalf("%s expected %f, got %f", key, expected, got)
	}
}
func checkBool(t *testing.T, key string, expected, got bool) {
	if expected != got {
		t.Fatalf("%s expected %t, got %t", key, expected, got)
	}
}
func checkString(t *testing.T, key string, expected, got string) {
	if expected != got {
		t.Fatalf("%s expected %s, got %s", key, expected, got)
	}
}

func checkTestError(t *testing.T, err error, isNil bool, expected string) {
	if isNil && err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if !isNil && err == nil {
		t.Fatalf("err expected %s, got nil err", expected)
	}
}

func checkBasicResult(t *testing.T, key string, expected, got any, k reflect.Kind, isGotPointer bool) {
	switch k {
	case reflect.Bool:
		expectedV := expected.(bool)
		var gotV bool
		if isGotPointer {
			gotV = *got.(*bool)
		} else {
			gotV = got.(bool)
		}
		checkBool(t, key, expectedV, gotV)
	case reflect.Int:
		expectedV := expected.(int)
		var gotV int
		if isGotPointer {
			gotV = *got.(*int)
		} else {
			gotV = got.(int)
		}
		checkInt(t, key, expectedV, gotV)
	case reflect.Int8:
		expectedV := expected.(int8)
		var gotV int8
		if isGotPointer {
			gotV = *got.(*int8)
		} else {
			gotV = got.(int8)
		}
		checkInt8(t, key, expectedV, gotV)
	case reflect.Int16:
		expectedV := expected.(int16)
		var gotV int16
		if isGotPointer {
			gotV = *got.(*int16)
		} else {
			gotV = got.(int16)
		}
		checkInt16(t, key, expectedV, gotV)
	case reflect.Int32:
		expectedV := expected.(int32)
		var gotV int32
		if isGotPointer {
			gotV = *got.(*int32)
		} else {
			gotV = got.(int32)
		}
		checkInt32(t, key, expectedV, gotV)
	case reflect.Int64:
		expectedV := expected.(int64)
		var gotV int64
		if isGotPointer {
			gotV = *got.(*int64)
		} else {
			gotV = got.(int64)
		}
		checkInt64(t, key, expectedV, gotV)
	case reflect.Uint:
		expectedV := expected.(uint)
		var gotV uint
		if isGotPointer {
			gotV = *got.(*uint)
		} else {
			gotV = got.(uint)
		}
		checkUint(t, key, expectedV, gotV)
	case reflect.Uint8:
		expectedV := expected.(uint8)
		var gotV uint8
		if isGotPointer {
			gotV = *got.(*uint8)
		} else {
			gotV = got.(uint8)
		}
		checkUint8(t, key, expectedV, gotV)
	case reflect.Uint16:
		expectedV := expected.(uint16)
		var gotV uint16
		if isGotPointer {
			gotV = *got.(*uint16)
		} else {
			gotV = got.(uint16)
		}
		checkUint16(t, key, expectedV, gotV)
	case reflect.Uint32:
		expectedV := expected.(uint32)
		var gotV uint32
		if isGotPointer {
			gotV = *got.(*uint32)
		} else {
			gotV = got.(uint32)
		}
		checkUint32(t, key, expectedV, gotV)
	case reflect.Uint64:
		expectedV := expected.(uint64)
		var gotV uint64
		if isGotPointer {
			gotV = *got.(*uint64)
		} else {
			gotV = got.(uint64)
		}
		checkUint64(t, key, expectedV, gotV)
	case reflect.Float32:
		expectedV := expected.(float32)
		var gotV float32
		if isGotPointer {
			gotV = *got.(*float32)
		} else {
			gotV = got.(float32)
		}
		checkFloat32(t, key, expectedV, gotV)
	case reflect.Float64:
		expectedV := expected.(float64)
		var gotV float64
		if isGotPointer {
			gotV = *got.(*float64)
		} else {
			gotV = got.(float64)
		}
		checkFloat64(t, key, expectedV, gotV)
	case reflect.String:
		expectedV := expected.(string)
		var gotV string
		if isGotPointer {
			gotV = *got.(*string)
		} else {
			gotV = got.(string)
		}
		checkString(t, key, expectedV, gotV)
	}
}

// 判断一个值是否存在于一组数组中
func isInArr(value string, valArr []string) bool {
	for _, k := range valArr {
		if k == value {
			return true
		}
	}
	return false
}
