package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	nums = []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		for i2, v2 := range nums {
			summ := v + v2
			if summ == target && i != i2 {
				return []int{i, i2}
			}

		}

	}

	return nil
}
