package main

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?
func twoSum(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return nums[i]+nums[j] == k
		}
		m[k-v] = i
	}
	return false
}
