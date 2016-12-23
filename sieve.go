package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

type BoolSlice []bool

func GenerateIntegers(n int, clinesize int) ([]BoolSlice, error) {
	if n < 2 {
		return nil, errors.New("the maximum integer must be larger than 2")
	}

	segcount := n / clinesize
	rval := make([]BoolSlice, segcount)

	for i := 0; i < segcount; i++ {
		rval[i] = make([]bool, clinesize)
	}

	return rval, nil
}

func (input *BoolSlice) MarkMultiples(p, base int) {
	if base <= p {
		for i := p - base; i < len(*input); i += p {
			(*input)[i] = true
		}
	} else {
		for i := base % p; i < len(*input); i += p {
			(*input)[i] = true
		}
	}
}

func (input BoolSlice) NextPrimeIndex() (int, error) {
	count := len(input)
	for i := 0; i < count; i++ {
		if input[i] {
			return i, nil
		}
	}
	return -1, errors.New("no prime found")
}

func (input *BoolSlice) ComputePrimesOneD(primes *[]int, base int) {
	// foreach known prime, mark their multiples in this BoolSlice
	for i := 0; i < len(*primes); i++ {
		input.MarkMultiples((*primes)[i], base)
	}

	for i := 0; i < len(*input); i++ {
		// find next prime
		if (*input)[i] == false {
			// add next prime
			(*primes) = append(*primes, base+i)
			// mark multiples of this prime
			input.MarkMultiples(base+i, base)
		}
	}
}

func ComputePrimesTwoD(input []BoolSlice, clinesize int) []int {
	segcount := len(input)

	n := segcount * clinesize
	primes := make([]int, 0, int(math.Sqrt(float64(n))))
	primes = append(primes, 2)

	for i := 0; i < segcount; i++ {
		input[i].ComputePrimesOneD(&primes, i*clinesize+2)
	}

	return primes
}

func main() {
	var n = flag.Int("n", 64, "max integer")
	var clinesize = flag.Int("cls", 64, "cache line size")
	var segmented = flag.Bool("segment", false, "segmented sieve, good for larger n's")
	flag.Parse()

	if !(*segmented) {
		clinesize = n
	}
	vals, err := GenerateIntegers(*n, *clinesize)
	if err != nil {
		panic(err)
	}

	primes := ComputePrimesTwoD(vals, *clinesize)
	fmt.Println(primes)
}
