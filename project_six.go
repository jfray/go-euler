package euler

import "fmt"

func ProjectSix(maximum int) int {
	/*
		The sum of the squares of the first ten natural numbers is,
		12 + 22 + ... + 102 = 385
		The square of the sum of the first ten natural numbers is,
		(1 + 2 + ... + 10)2 = 552 = 3025
		Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025  385 = 2640.
		Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	*/
	mysum, square := 1, 1
	for i := 2; i <= maximum; i++ {
		mysum += i * i
		square += i
	}
	r := (square * square) - mysum
	fmt.Printf("#6 %d\n", r)
	return r
}
