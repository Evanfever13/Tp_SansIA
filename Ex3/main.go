package main

import "fmt"

func MaxSubArray(nums []int) int {

	cur_sum := 0
	max_sum := nums[0]

	for i := 0; i < len(nums); i++ {
		cur_sum += nums[i]
		if cur_sum < 0 {
			cur_sum = 0
		} else if max_sum < cur_sum {
			max_sum = cur_sum
		}
	}
	return max_sum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSubArray(nums)) // 6 (le sous-tableau [4, -1, 2, 1])
}
