package main

//In mathematics, the factorial of a positive integer n, denoted by n!, is the product of all positive
//integers less than or equal to n. For example, {\displaystyle 5!=5\times 4\times 3\times 2\times 1=120\,.} {\displaystyle 5!=5\times 4\times 3\times 2\times 1=120\,.}
//he value of 0! is 1, according to the convention for an empty product.[1]
//The factorial operation is encountered in many areas of mathematics, notably in combinatorics, algebra, and mathematical analysis. Its most basic occurrence is the fact that there
//are n! ways to arrange n distinct objects into a sequence. These arrangements are called the permutations of the set of objects.
//The definition of the factorial function can also be extended to non-integer arguments, while retaining its most important properties; this involves more advanced mathematics, notably techniques from mathematical analysis.
//we will use recursion in finding this solution

import "fmt"

func main() {
	i := factorial(5)
	fmt.Println(i)

}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
