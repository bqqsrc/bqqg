//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"io/ioutil"
	"testing"
)

func Benchmark_Scanner(b *testing.B) {
	data, _ := ioutil.ReadFile("./testdata/data.gson")
	for i := 0; i < b.N; i++ {
		Valid(data)
	}
}
