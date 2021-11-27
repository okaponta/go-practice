package main

import (
	"testing"
)

func Benchmark_mul_2exp_slice(b *testing.B) {
	var hoge int64 = 0
	numbers := make([]int64, 31)
	for i := range numbers {
		if i == 0 {
			numbers[i] = 2
		} else {
			numbers[i] = numbers[i-1] * 2
		}
	}
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		for _, j := range numbers {
			hoge = j * j
		}
	}
	if hoge == 0 {
		b.Fail()
	}
}

func Benchmark_mul_2exp_minus1_slice(b *testing.B) {
	var hoge int64 = 0
	numbers := make([]int64, 31)
	for i := range numbers {
		if i == 0 {
			numbers[i] = 2
		} else {
			numbers[i] = numbers[i-1] * 2
		}
	}
	for i := range numbers {
		numbers[i] = numbers[i] - 1
	}
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		for _, j := range numbers {
			hoge = j * j
		}
	}
	if hoge == 0 {
		b.Fail()
	}
}
