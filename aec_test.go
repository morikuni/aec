package aec_test

import (
	"fmt"
	"testing"

	"github.com/morikuni/aec"
)

var benchValues = []struct {
	n uint
}{
	{n: 0},   // 0,1 special case
	{n: 2},   // 1 digit
	{n: 20},  // 2 digit
	{n: 200}, // 3 digits
}

// Representative benchmark for all single-parameter uint CSI generators:
// Up, Down, Left, Right, NextLine, PreviousLine, Column, EraseDisplay, and EraseLine.
func BenchmarkCursor(b *testing.B) {
	for _, tc := range benchValues {
		b.Run(fmt.Sprintf("%d", tc.n), func(b *testing.B) {
			var a aec.ANSI
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				a = aec.Up(tc.n)
			}
			sinkANSI = a
		})
	}
}

// Covers the two-parameter uint CSI path used by [aec.Position].
func BenchmarkPosition(b *testing.B) {
	var tests = []struct {
		row uint
		col uint
	}{
		{row: 0, col: 0},     // 0,1 special case
		{row: 2, col: 2},     // 1 digit
		{row: 20, col: 20},   // 2 digit
		{row: 200, col: 200}, // 3 digits
		{row: 5, col: 15},
	}
	for _, tc := range tests {
		b.Run(fmt.Sprintf("%d,%d", tc.row, tc.col), func(b *testing.B) {
			var a aec.ANSI
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				a = aec.Position(tc.row, tc.col)
			}
			sinkANSI = a
		})
	}
}

// Representative benchmark for all single-parameter int CSI
// generators ([aec.ScrollUp], [aec.ScrollDown]).
func BenchmarkScrollUpDown(b *testing.B) {
	for _, tc := range benchValues {
		b.Run(fmt.Sprintf("%d", tc.n), func(b *testing.B) {
			var a aec.ANSI
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				a = aec.ScrollUp(int(tc.n))
			}
			sinkANSI = a
		})
	}
}
