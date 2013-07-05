package euler

import (
	"fmt"
	"strconv"
	"strings"
)

func ProjectFour() int {
	/*
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
		Find the largest palindrome made from the product of two 3-digit numbers.
	*/

	maximum := 0
	// counting down instead of up was nearly 9x faster according to benchmark
	for a := 999; a > 100; a-- {
		for b := a; b > 100; b-- {
			product := a * b
			if product > maximum {
				s := strings.Split(strconv.Itoa(product), "")
				front := fmt.Sprintf("%s%s%s", s[0], s[1], s[2])
				back := fmt.Sprintf("%s%s%s", s[len(s)-1], s[len(s)-2], s[len(s)-3])
				if front == back {
					maximum = product
				}
			}
		}
	}
	fmt.Printf("#4 %d\n", maximum)

	return maximum
}
