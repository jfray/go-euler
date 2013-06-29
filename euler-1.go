package main

/* Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

def multiples_of_3_or_5():
        for number in xrange(1000):
            if not number % 3 or not number % 5:
                yield number

    print sum(multiples_of_3_or_5())
*/

import (
	"fmt"
)

func main() {
	var number_list []int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			number_list = append(number_list, i)
		}
	}

	var sum int
	for x := 0; x < len(number_list); x++ {
		sum = sum + number_list[x]
	}
	fmt.Println(number_list)
	fmt.Println(sum)
}
