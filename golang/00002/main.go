package main

//A much more efficient method is the Euclidean algorithm, which uses the division algorithm in
//combination with the observation that the gcd of two numbers also divides their difference: divide 48 by 18 to get a quotient of 2 and a remainder of 12. Then divide 18 by 12 to get a
//quotient of 1 and a remainder of 6. Then divide 12 by 6 to get a remainder of 0, which means that 6 is the gcd

import "fmt"

func main() {
	i := gcd(14, 21)
	//i := factorial(40)
	fmt.Println(i)

}

func gcd(a int, b int) int {
	if b == 0 {
		return 1
	}
	return gcd(b, a%b)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
