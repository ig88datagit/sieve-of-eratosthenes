package main

import (
	"testing"
)

func BenchmarkComputePrimesTwoD(b *testing.B) {
	clinesize := 64

	vals, err := GenerateIntegers(100000, clinesize)
	if err != nil {
		panic(err)
	}

	_ = ComputePrimesTwoD(vals, clinesize)
}
