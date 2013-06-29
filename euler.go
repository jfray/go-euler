package main

import (
	"fmt"
)

func needFloat(x float64) float64 {
	return x * 0.1
}

func project_one(limit int) int {

	/* Multiples of 3 and 5

	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

	Find the sum of all the multiples of 3 or 5 below 1000.

	def multiples_of_3_or_5():
	  for number in xrange(1000):
	    if not number % 3 or not number % 5:
	      yield number

	print sum(multiples_of_3_or_5())
	*/

	var number_list []int

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			number_list = append(number_list, i)
		}
	}

	var mysum int
	for x := 0; x < len(number_list); x++ {
		mysum = mysum + number_list[x]
	}
	return mysum
}

func project_two(limit int) int {

	/* Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

	   1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	   By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
	   def fib():
	     x,y = 0,1
	     while True:
	       yield x
	       x,y = y, x+y

	   def even(seq):
	     for number in seq:
	       if not number % 2:
	         yield number

	   def under_a_million(seq):
	     for number in seq:
	       if number > 1000000:
	         break
	       yield number

	   print sum(even(under_a_million(fib())))

	*/

	var number_list []int
	return

}

func main() {
	p1 := project_one(1000)
	fmt.Printf("project #1: %d\n", p1)

	p2 := project_two(4000000)
	fmt.Printf("project #2: %d\n", p2)
}