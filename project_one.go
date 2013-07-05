package euler

import "fmt"

func ProjectOne(limit int) int {

	/*
		Multiples of 3 and 5
		If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
		Find the sum of all the multiples of 3 or 5 below 1000.
	*/

	var mysum int

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			mysum = mysum + i
		}
	}

	fmt.Printf("#1 %d\n", mysum)
	return mysum
}
