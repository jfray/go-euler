package euler

import "fmt"

func ProjectFive(minimum int, maximum int) int {
	/*
		Smallest multiple
		2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
		What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	*/

	// get list of unique factors between minimum and maximum, since you can remove them from the list of numbers you need to check
	var factors []int

	for x := minimum; x <= maximum; x++ {
		if maximum%x != 0 && x != maximum-1 && x != minimum {
			factors = append(factors, x)
		}
	}

	// since maximum and maximum -1  have to divide evenly, we can start at max and increment by their product
	increment_by := maximum * (maximum - 1)

	for i := increment_by; ; {
		// check the resulting number against number_list to see if all are divisible
		for pos, value := range factors {
			if i%value == 0 {
				if pos == len(factors)-1 {
					fmt.Printf("#5 %d\n", i)
					return i
				}
			} else {
				break
			}
		}
		i += increment_by
	}
	return 0
}
