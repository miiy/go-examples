package main

import "fmt"

func main()  {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}