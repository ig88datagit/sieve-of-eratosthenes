package main

import (
	"errors"
	"fmt"
)

func ConsecutiveIntegers(n int) ([]bool, error) {
	if n < 2 {
		return nil, errors.New(fmt.Sprintf("%d is less than 2.", n))
	}

	return make([]bool, n), nil
}

func SieveOfEratosthenes(n int, nlist []bool) ([]int, error) {
	// ensure count(numlist) - 2 >= n

	var primes []int
	var p int
	for i := 0; i < n-2; i++ {
		if nlist[i] == false {
			// index is number - 2
			p = i + 2

			// we got a prime
			primes = append(primes, p)

			// mark multiples
			for j := 2; (j * p) < n; j++ {
				nlist[j*p-2] = true
			}
		}
	}

	return primes, nil
}

func main() {
	var con, _ = ConsecutiveIntegers(10)

	var numlist, _ = SieveOfEratosthenes(10, con)

	for i := 0; i < len(numlist); i++ {
		fmt.Println(numlist[i])
	}
}
