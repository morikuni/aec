package aec_test

import (
	"testing"

	"github.com/morikuni/aec"
)

var sinkANSI aec.ANSI // used in benchmarks to avoid compiler optimizations

func BenchmarkColor3BitF(b *testing.B) {
	b.ReportAllocs()

	var a aec.ANSI
	for i := 0; i < b.N; i++ {
		a = aec.Color3BitF(3)
	}
	sinkANSI = a
}

func BenchmarkColor8BitF(b *testing.B) {
	b.ReportAllocs()

	var a aec.ANSI
	for i := 0; i < b.N; i++ {
		a = aec.Color8BitF(128)
	}
	sinkANSI = a
}

func BenchmarkFullColorF(b *testing.B) {
	b.ReportAllocs()

	var a aec.ANSI
	for i := 0; i < b.N; i++ {
		a = aec.FullColorF(255, 128, 0)
	}
	sinkANSI = a
}
