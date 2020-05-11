package main

import (
	"fmt"
	"strconv"
)

type Sign struct {
	index      int
	isNegetive bool
}

func isNegetive(str string) Sign {
	numIndex, index := firstNumberIndex(str), len(str)+10
	for i := 0; i < len(str); i++ {
		if str[i] == '-' {
			index = i
		}
	}
	if numIndex > index {
		return true
	}
	return false
}

func firstNumberIndex(str string) int {
	numFound, index := false, -1

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '-' {
			continue
		}
		if numFound == false && isNum(string(str[i])) == true {
			numFound = true
			index = i
			return index
		}
	}
	return index
}

func isStartWithWord(str string) bool {
	numFound, flag, plusMinusFound := false, false, false
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		if (str[i] == '-' || str[i] == '+') && plusMinusFound == false {
			plusMinusFound = true
			continue
		}
		if numFound == false && isNum(string(str[i])) == false {
			numFound = true
			flag = true
		} else {
			numFound = true
		}
	}
	return flag
}

func numberPortion(str string) int {
	index := firstNumberIndex(str)
	for i := firstNumberIndex(str); i < len(str); i++ {
		if isNum(string(str[i])) == false {
			index = i
			break
		}
	}
	return index
}
func myAtoi(str string) int {
	var num int
	numFound := false
	if isStartWithWord(str) == true {
		return 0
	}
	for i := firstNumberIndex(str); i <= numberPortion(str); i++ {
		if isNum(string(str[i])) == true {
			numFound = true
			x, err := strconv.Atoi(string(str[i]))
			if err != nil {
				fmt.Println(err)
				return 0
			}
			num = num*10 + int(x)
		}
		if isNum(string(str[i])) == false && numFound == true {
			break
		}
	}
	if num > 2147483648 {
		num = 2147483648
	}
	if isNegetive(str) {
		num = num * -1
	}
	return num
}
func isNum(char string) bool {
	return (char >= "0" && char <= "9")
}
func main() {
	fmt.Println(myAtoi("  - 0012a42"))
}
