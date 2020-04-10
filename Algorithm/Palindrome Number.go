package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.FormatInt(int64(x), 10)
	flag := true
	l := len(s)
	for i := 0; i <= l/2; i++ {
		if s[i] != s[l-i-1] {
			flag = false
			break
		}
	}
	return flag
}
func main() {
	fmt.Println(isPalindrome(1211))
}
