// Copyright (C) 晓白齐齐,版权所有.

package test

import (	
	"testing"
	"fmt"
)


func Must(t *testing.T, ret bool, errFormat string, args ...any) {
	if !ret {
		t.Errorf(errFormat, args...)
	}
}

func MustBeValue(t *testing.T, shouldBe, realBe any, key string) {
	Must(t, shouldBe == realBe, "%s should be %v, but got %v", key, shouldBe, realBe)
}

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