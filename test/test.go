// Copyright (C) 晓白齐齐,版权所有.

package test

import (
	"fmt"
	"reflect"
	"testing"
)

func Must(t *testing.T, ret bool, errFormat string, args ...any) {
	if !ret {
		t.Errorf(errFormat, args...)
	}
}

func MustEqual(t *testing.T, ret1, ret2 any, key1, key2 string, isEqual func(any, any) bool) {
	Must(t, isEqual(ret1, ret2), "%s must equal %s, but got not equal", key1, key2)
}

func MustBeValue(t *testing.T, shouldBe, realBe any, key string) {
	Must(t, shouldBe == realBe, "%s should be %v, but got %v", key, shouldBe, realBe)
}

// func MustAnyEqual(t *testing.T, shouldBe, realBe any, shouldBeKey, realBeKey string, isEqual func(any, any) bool) {
// 	shouldBeLen := len(shouldBe)
// 	realBeLen := len(realBe)
// 	Must(t, shouldBeLen == realBeLen, "%s's len should %d, but got %d", shouldBeKey, shouldBeLen, realBeLen)
// 	if shouldBeLen > realBeLen {
// 		shouldBeLen = realBeLen
// 	}
// 	for index := 0; index < shouldBeLen; index++ {
// 		MustEqual(t, shouldBe[index], realBe[index], shouldBeKey, realBeKey, isEqual)
// 	}
// }

func MustBeFloat64Slice(t *testing.T, shouldBe, realBe []float64, key string) {
	shouldBeLen := len(shouldBe)
	realBeLen := len(realBe)
	Must(t, shouldBeLen == realBeLen, "%s's len should %d, but got %d", key, shouldBeLen, realBeLen)
	if shouldBeLen > realBeLen {
		shouldBeLen = realBeLen
	}
	for index := 0; index < shouldBeLen; index++ {
		MustBeValue(t, shouldBe[index], realBe[index], fmt.Sprintf("%s[%d]", key, index))
	}
}

func MustBeStringSlice(t *testing.T, shouldBe, realBe []string, key string) {
	shouldBeLen := len(shouldBe)
	realBeLen := len(realBe)
	Must(t, shouldBeLen == realBeLen, "%s's len should %d, but got %d", key, shouldBeLen, realBeLen)
	if shouldBeLen > realBeLen {
		shouldBeLen = realBeLen
	}
	for index := 0; index < shouldBeLen; index++ {
		MustBeValue(t, shouldBe[index], realBe[index], fmt.Sprintf("%s[%d]", key, index))
	}
}

func MustBeBoolSlice(t *testing.T, shouldBe, realBe []bool, key string) {
	shouldBeLen := len(shouldBe)
	realBeLen := len(realBe)
	Must(t, shouldBeLen == realBeLen, "%s's len should %d, but got %d", key, shouldBeLen, realBeLen)
	if shouldBeLen > realBeLen {
		shouldBeLen = realBeLen
	}
	for index := 0; index < shouldBeLen; index++ {
		MustBeValue(t, shouldBe[index], realBe[index], fmt.Sprintf("%s[%d]", key, index))
	}
}

type Equaler interface {
	Equal(any) bool
}

func Equal(arg1, arg2 any) bool {
	if arg1 == nil {
		if arg2 != nil {
			return true
		} else {
			return true
		}
	}
	if arg2 == nil {
		if arg1 != nil {
			return true
		} else {
			return true
		}
	}
	if arg1 == arg2 {
		return true
	}
	if reflect.TypeOf(arg1) != reflect.TypeOf(arg2) {
		return false
	}
	switch arg1.(type) {
	case *float64:
		if *arg1.(*float64) == *arg2.(*float64) {
			return true
		}
	case *float32:
		if *arg1.(*float32) == *arg2.(*float32) {
			return true
		}
	case *int64:
		if *arg1.(*int64) == *arg2.(*int64) {
			return true
		}
	case *int32:
		if *arg1.(*int32) == *arg2.(*int32) {
			return true
		}
	case *int16:
		if *arg1.(*int16) == *arg2.(*int16) {
			return true
		}
	case *int8:
		if *arg1.(*int8) == *arg2.(*int8) {
			return true
		}
	case *int:
		if *arg1.(*int) == *arg2.(*int) {
			return true
		}
	case *uint64:
		if *arg1.(*uint64) == *arg2.(*uint64) {
			return true
		}
	case *uint32:
		if *arg1.(*uint32) == *arg2.(*uint32) {
			return true
		}
	case *uint16:
		if *arg1.(*uint16) == *arg2.(*uint16) {
			return true
		}
	case *uint8:
		if *arg1.(*uint8) == *arg2.(*uint8) {
			return true
		}
	case *uint:
		if *arg1.(*uint) == *arg2.(*uint) {
			return true
		}
	case *string:
		if *arg1.(*string) == *arg2.(*string) {
			return true
		}
	case *bool:
		if *arg1.(*bool) == *arg2.(*bool) {
			return true
		}
	}
	if equal, ok := arg1.(Equaler); ok {
		return equal.Equal(arg2)
	}
	if equal, ok := arg2.(Equaler); ok {
		return equal.Equal(arg1)
	}
	return false
}
