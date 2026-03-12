package main

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for i, val := range nums {
		sum += val
		result[i] = sum
	}
	return result
}
