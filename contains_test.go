package main

import (
	"testing"
)

var sa []string = []string{"q", "w", "e", "r", "t", "a", "b", "c", "d", "f", "g", "h", "i", "j", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7"}
var needle string = "5"

func BenchmarkContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Contains(sa, needle)
	}
}

func BenchmarkStringsContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringsContains(sa, needle)
	}
}

func BenchmarkWgoContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		WgoContains(sa, needle)
	}
}

func BenchmarkWgoStringsContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		WgoStringsContains(sa, needle)
	}
}

func BenchmarkWgSwitchContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		WgoSwitchContains(sa, needle)
	}
}

func BenchmarkSliceContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SliceContains(sa, needle)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("For %d + %d, expected %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})

	t.Run("mix", func(t *testing.T) {
		if Mul(-2, Mul(2, -3)) != 12 {
			t.Fatal("fail")
		}
	})
}
