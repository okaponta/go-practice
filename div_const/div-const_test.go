package main

import (
	"testing"
)

func Benchmark_div_2(b *testing.B) {
	var ans uint16 = 0
	//var j uint16 = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			ans = div2(uint16(j))
		}
	}
	if ans == 0 {
		b.Fail()
	}
}

func Benchmark_div_3(b *testing.B) {
	var ans uint16 = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			ans = div3(uint16(j))
		}
	}
	if ans == 0 {
		b.Fail()
	}
}

func Benchmark_div_3_original(b *testing.B) {
	var ans uint16 = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			ans = mul43691(uint16(j))
		}
	}
	if ans == 0 {
		b.Fail()
	}
}
