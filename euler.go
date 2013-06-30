package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ProjectOne(limit int) int {

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

func ProjectTwo(limit int) int {

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
	var x, y int = 0, 1
	// greater than or equal..to not exceed the limit but include it
	for x <= limit {
		x, y = y, x+y
		if x%2 == 0 {
			number_list = append(number_list, x)
		}
	}
	var mysum int
	for i := 0; i < len(number_list); i++ {
		mysum = mysum + number_list[i]
	}
	return mysum
}

func ProjectThree(number int) int {
	/* The prime factors of 13195 are 5, 7, 13 and 29.

	   What is the largest prime factor of the number 600851475143 ?

	   n = 600851475143
	   i = 2
	   while i * i < n:
	     while n % i == 0:
	         n = n / i
	     i = i + 1
	   print n

	*/
	i := 2
	for i*i < number {
		for number%i == 0 {
			number = number / i
		}
		i++
	}
	return number
}

func ProjectFour() int {
	/* A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.

			Find the largest palindrome made from the product of two 3-digit numbers.

	        class Integer
	          def palindromic?
	            digits = self.to_s.split('')
	            return digits == digits.reverse
	          end
	        end

	        max = 0
	        (100..999).each do |a|
	          (a..999).each do |b|
	            product = a * b
	              if product > max and product.palindromic?
	                max = product
	              end
	          end
	        end
	        puts max
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
	return maximum
}

func main() {
	p1 := ProjectOne(1000)
	fmt.Printf("project #1: %d\n", p1)

	p2 := ProjectTwo(4000000)
	fmt.Printf("project #2: %d\n", p2)

	p3 := ProjectThree(600851475143)
	fmt.Printf("project #3: %d\n", p3)

	p4 := ProjectFour()
	fmt.Printf("project #4: %d\n", p4)

}
