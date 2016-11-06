package main

import (
	"errors"
	"fmt"
	"math"
)

func ConsecutiveIntegers(n int) ([][]bool, error) {
	if n < 2 {
		return nil, errors.New(fmt.Sprintf("%d is less than 2.", n))
	}

	if n < 64 {
		r := make([][]bool, 1)
		r[0] = make([]bool, n-2)
		return r, nil
	}

	segsize := int(math.Sqrt(float64(n - 2)))
	segcount := int(math.Ceil(float64(n-2) / float64(segsize)))

	r := make([][]bool, segcount)
	for i := 0; i < segcount; i++ {
		r[i] = make([]bool, segsize)
	}

	return r, nil
}

func SegmentedSieveOfEratosthenes(n int, nlist [][]bool) ([]int, error) {
	var primes []int
	var p int

	for i := 0; i < n-2; i++ {

		for j := 0; j < len(nlist); j++ {
			for k := 0; k < len(nlist[j]); k++ {
				if nlist[j][k] == false {
					p = i + 2

					primes = append(primes, p)

					for l := 2; (l * p) < len(nlist)*(j-1)*len(nlist[j-1])+len(nlist[j]); l++ {
						nlist[j][l*p-2] = true
					}
				}
			}
		}
	}

	return primes, nil
}

func main() {
	var con, _ = ConsecutiveIntegers(10)
	var nlist, _ = SegmentedSieveOfEratosthenes(10, con)

	for i := 0; i < len(nlist); i++ {
		fmt.Println(nlist[i])
	}
	fmt.Printf("%v", con)
}
