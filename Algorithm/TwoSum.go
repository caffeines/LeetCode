package main

import "fmt"

type Num struct {
	index int
	exist int
}

func twoSum(nums []int, target int) []int {
	flag := make(map[int]Num)
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		flag[curr] = Num{i, 1}
	}
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		rest := target - curr
		if flag[rest].exist == 1 && flag[rest].index != i {
			fmt.Println(rest)
			arr = append(arr, i, flag[rest].index)
			break
		}
	}
	return arr
}
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
