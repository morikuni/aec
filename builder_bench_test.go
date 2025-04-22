package aec_test

import (
	"testing"

	"github.com/morikuni/aec"
)

var sinkString string

func BenchmarkBuilderFluent(b *testing.B) {
	b.ReportAllocs()

	var s string
	for i := 0; i < b.N; i++ {
		s = aec.EmptyBuilder.
			Right(2).
			RGB8BitF(128, 255, 64).
			RedB().
			Bold().
			ANSI.
			Apply("Hello World")
	}
	sinkString = s
}

func BenchmarkBuilderWith(b *testing.B) {
	b.ReportAllocs()

	var ansi aec.ANSI
	for i := 0; i < b.N; i++ {
		ansi = aec.NewBuilder(aec.Up(2)).
			With(aec.Column(10), aec.Bold, aec.RedF).
			ANSI
	}
	sinkANSI = ansi
}
