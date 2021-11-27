package main

import (
	"math/rand"
	"testing"
	"time"
)

func get2exp() int64 {
	rand.Seed(time.Now().UnixNano())
	exp := 30 + rand.Intn(2)
	var ans int64 = 1
	for i := 0; i < exp; i++ {
		ans = ans * 2
	}
	return ans
}

func Benchmark_mul_2exp(b *testing.B) {
	var base int64 = get2exp() // 2^30 or 2^31
	var hoge int64 = get2exp() - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hoge = mul64(hoge, base)
		}
	}
	if hoge == 1 {
		b.Fail()
	}
}

func Benchmark_mul_2exp_minus1(b *testing.B) {
	var base int64 = get2exp() - 1 // 2^30 or 2^31
	var hoge int64 = get2exp() - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hoge = mul64(hoge, base)
		}
	}
	if hoge == 1 {
		b.Fail()
	}
}

func Benchmark_mul_2exp_plus1(b *testing.B) {
	var base int64 = get2exp() + 1 // 2^30 or 2^31
	var hoge int64 = get2exp() - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hoge = mul64(hoge, base)
		}
	}
	if hoge == 1 {
		b.Fail()
	}
}

func Benchmark_plus_2exp(b *testing.B) {
	var base int64 = get2exp() // 2^30 or 2^31
	var hoge int64 = get2exp() - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hoge = plus(hoge, base)
		}
	}
	if hoge == 1 {
		b.Fail()
	}
}
