package main

import (
	"testing"
)

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	l, _ := ConsecutiveIntegers(b.N)
	nl, _ := SieveOfEratosthenes(b.N, l)
	_ = nl
}
