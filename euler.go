package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

	return mysum
}

func ProjectTwo(limit int) int {

	/*
		Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
		1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
		By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
	*/

	var mysum int

	var x, y int = 0, 1
	// greater than or equal..to not exceed the limit but include it
	for x <= limit {
		x, y = y, x+y
		if x%2 == 0 {
			mysum = mysum + x
		}
	}

	return mysum
}

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
	return number
}

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
	return maximum
}

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
	return (square * square) - mysum
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
			return count
		}
		count += 2
	}
}

func ProjectEight() int {
	/*
		Find the greatest product of five consecutive digits in the 1000-digit number.

		73167176531330624919225119674426574742355349194934
		96983520312774506326239578318016984801869478851843
		85861560789112949495459501737958331952853208805511
		12540698747158523863050715693290963295227443043557
		66896648950445244523161731856403098711121722383113
		62229893423380308135336276614282806444486645238749
		30358907296290491560440772390713810515859307960866
		70172427121883998797908792274921901699720888093776
		65727333001053367881220235421809751254540594752243
		52584907711670556013604839586446706324415722155397
		53697817977846174064955149290862569321978468622482
		83972241375657056057490261407972968652414535100474
		82166370484403199890008895243450658541227588666881
		16427171479924442928230863465674813919123162824586
		17866458359124566529476545682848912883142607690042
		24219022671055626321111109370544217506941658960408
		07198403850962455444362981230987879927244284909188
		84580156166097919133875499200524063689912560717606
		05886116467109405077541002256983155200055935729725
		71636269561882670428252483600823257530420752963450
	*/

	return 0
}

func isprime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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

	p5 := ProjectFive(1, 20)
	fmt.Printf("project #5: %d\n", p5)

	p6 := ProjectSix(100)
	fmt.Printf("project #6: %d\n", p6)

	p7 := ProjectSeven(10001)
	fmt.Printf("project #7: %d\n", p7)

	p8 := ProjectEight()
	fmt.Printf("project #8: %d\n", p8)
}
