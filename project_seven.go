package euler

import (
	"fmt"
	"math"
)

func isprime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func ProjectSeven(target int) int {
	/*
		By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
		What is the 10,001st prime number?
	*/
	count := 3
	index := 1

	for {
		if isprime(count) {
			index++
		}
		if index == target {
			fmt.Printf("#7 %d\n", count)
			return count
		}
		count += 2
	}
}
