package main

import "testing"

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}

//Обход в двух циклах
// BenchmarkRun-16           351927              3413 ns/op
// BenchmarkRun-16           288638              3530 ns/op
// BenchmarkRun-16           306124              3357 ns/op

//Обход в одном цикле
//BenchmarkRun-16           329523              3488 ns/op
//BenchmarkRun-16           304416              3452 ns/op
//BenchmarkRun-16           355071              3373 ns/op
