package euler

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ProjectFour(digits int) int {
	/*
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
		Find the largest palindrome made from the product of two 3-digit numbers.
	*/

	high_num := int(math.Pow(10, float64(digits))) - 1
	low_num := int(math.Pow(10, float64(digits-1)))
	maximum := 0

	// counting down instead of up was nearly 9x faster according to benchmark
	for a := high_num; a > low_num; a-- {
		for b := a; b > low_num; b-- {
			product := a * b
			if product > maximum {
				s := strings.Split(strconv.Itoa(product), "")

				var front string
				var back string

				// turns out the number of digits you put in will always equal len(product) / 2 evenly so just use digits
				for i := 0; i < digits; i++ {
					front = front + fmt.Sprintf("%s", s[i])
					back = back + fmt.Sprintf("%s", s[len(s)-(i+1)])
				}
				if front == back {
					maximum = product
				}
			}
		}
	}
	fmt.Printf("#4 %d\n", maximum)

	return maximum
}
