package main

import "fmt"

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
var sun int

func main() {
	findmultiple(1, 10)
}

func findmultiple(start int, end int) {
	var num int
	for i := start; i < end; i++ {
		if (i%3 == 0) || i%5 == 0 {
			num = num + i
		}
	}
	fmt.Println(num)
}
