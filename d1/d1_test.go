package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if twoSum([]int{10, 15, 3, 7}, 17) != true {
		t.Error("Expected true, got false")
	}

	defaultAssertionErrorMessage := "Expected %v, got %v"

	type TestCase struct {
		nums []int
		k    int
		want bool
	}

	testCases := []TestCase{
		{[]int{10, 15, 3, 7}, 17, true},
		{[]int{10, 15, 3, 7}, 25, true},
		{[]int{10, 15, 3, 7}, 26, false},
		{[]int{10, 15, 3, 7}, 10, true},
		{[]int{10, 15, 3, 7}, 11, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("twoSum(%v, %d) should return %v)", tc.nums, tc.k, tc.want), func(t *testing.T) {
			if got := twoSum(tc.nums, tc.k); got != tc.want {
				t.Errorf(defaultAssertionErrorMessage, tc.want, got)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{10, 15, 3, 7}, 17)
	}
}
