package string_test

import (
	"testing"
	"strings"
	"fmt"
)

// go test -v -run="none" -bench="."

/*
纯字符串的拼接，Builder比+要优很多
但如果有其他数据类型（如int），因为用Sprintf，两者差别不大
*/


// 字符串拼接
var base = "TeST"

func BenchmarkPlus(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < 100; i++ {
			s += base
		}
		// len(s)
	}
}

func BenchmarkBuilder(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		// s := ""
		for i := 0; i < 100; i++ {
			fmt.Fprintf(&sb, "%s", base)
		}
		sb.String()
		// len(s)
	}
}

func BenchmarkBuilder_WriteString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		// s := ""
		for i := 0; i < 100; i++ {
			sb.WriteString(base)
		}
		sb.String()
		// len(s)
	}
}

// 数字拼接

func BenchmarkPlus_num(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < 100; i++ {
			s += fmt.Sprintf("%d", i)
		}
		// len(s)
	}
}

func BenchmarkBuilder_num(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		// s := ""
		for i := 0; i < 100; i++ {
			fmt.Fprintf(&sb, "%d", i)
		}
		sb.String()
		// len(s)
	}
}

func BenchmarkBuilder_WriteString_num(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		// s := ""
		for i := 0; i < 100; i++ {
			sb.WriteString(fmt.Sprintf("%d", i))
		}
		sb.String()
		// len(s)
	}
}