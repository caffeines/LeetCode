package main

import (
	"fmt"
	"math"
)

/* Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer. */

func rangeChecker(num int) int {
	up := int(math.Pow(2, 31) - 1)
	down := int(math.Pow(-2, 31))
	if num <= up && num >= down {
		return num
	}
	return 0
}

func reverse(x int) int {
	neg := false
	var res int = 0
	if x < 0 {
		x *= -1
		neg = true
	}
	for x > 0 {
		rem := x % 10
		res = res*10 + rem
		x /= 10
	}
	if neg {
		res *= -1
	}
	return rangeChecker(res)
}

func main() {
	res := reverse(-1534236469)
	fmt.Println(res)
}
