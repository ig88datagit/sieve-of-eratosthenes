package main_test

import (
	"fmt"
	"testing"
)

func TestConsecutiveIntegers(t *testing.T) {
	testCases := []int{1, 32, 64, 128, 256, 512, 2024, 4028, 8056}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc), func(T *testing.T) {

		})
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	l, _ := ConsecutiveIntegers(b.N)
	nl, _ := SieveOfEratosthenes(b.N, l)
	_ = nl
}
