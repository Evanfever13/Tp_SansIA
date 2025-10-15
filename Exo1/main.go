package main

import "fmt"

func TwoSum(nums []int, target int) (int, int) {
	for i := 0; i < len(nums); i++ {
		// Modification : On démarre j à i+1 pour éviter de retester les mêmes paires et d’avoir i == j
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	i, j := TwoSum(nums, target)
	fmt.Println(i, j) // Résultat attendu : 0 1
}
