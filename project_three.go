package euler

import "fmt"

func ProjectThree(number int) int {
	/*
		The prime factors of 13195 are 5, 7, 13 and 29.
		What is the largest prime factor of the number 600851475143?
	*/
	i := 2
	for i*i < number {
		for number%i == 0 {
			number = number / i
		}
		i++
	}
	fmt.Printf("#3 %d\n", number)

	return number
}
