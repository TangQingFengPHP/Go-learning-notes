// 基准测试使用 go test -bench=. -benchmem 来运行测试
package main

import (
	"os"
	"strings"
	"testing"
)

func echoBench1() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echoBench2() string {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " | "
	}
	return s
}

func echoBench3() string {
	return strings.Join(os.Args[:], " , ")
}

func echoBench4() string {
	var builder strings.Builder
	for i, arg := range os.Args[:] {
		if i > 0 {
			builder.WriteString(" <> ")
		}
		builder.WriteString(arg)
	}
	return builder.String()
}

func BenchmarkEchoAll1(b *testing.B) {
	originalArgs := os.Args
	os.Args = make([]string, 100)
	for i := range os.Args {
		os.Args[i] = "arg"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = echoBench1()
	}
	os.Args = originalArgs
}

func BenchmarkEchoAll2(b *testing.B) {
	originalArgs := os.Args
	os.Args = make([]string, 100)
	for i := range os.Args {
		os.Args[i] = "arg"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = echoBench2()
	}
	os.Args = originalArgs
}

func BenchmarkEchoAll3(b *testing.B) {
	originalArgs := os.Args
	os.Args = make([]string, 100)
	for i := range os.Args {
		os.Args[i] = "arg"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = echoBench3()
	}
	os.Args = originalArgs
}

func BenchmarkEchoAll4(b *testing.B) {
	originalArgs := os.Args
	os.Args = make([]string, 100)
	for i := range os.Args {
		os.Args[i] = "arg"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = echoBench4()
	}
	os.Args = originalArgs
}
