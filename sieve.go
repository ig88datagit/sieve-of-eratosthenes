package main

import (
	"errors"
	"flag"
	"fmt"
)

func GenerateIntegers(n int, clinesize int) ([][]bool, error) {
	if n < 2 {
		return nil, errors.New("the maximum integer must be larger than 2")
	}

	segcount := n / clinesize
	rval := make([][]bool, segcount)

	for i := 0; i < segcount; i++ {
		rval[i] = make([]bool, clinesize)
	}

	return rval, nil
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

	err = vals.SimpleCalc()
	if err != nil {
		panic(err)
	}

	fmt.Print(vals)
}
