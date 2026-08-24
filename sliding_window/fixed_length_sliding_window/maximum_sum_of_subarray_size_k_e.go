package fixed_length_sliding_window

/*
DESCRIPTION
Given an array of integers nums and an integer k, find the maximum sum of any contiguous subarray of size k.

Example 1: Input:
nums = [2, 1, 5, 1, 3, 2]
k = 3
*/

// My answer
func MaxSum(nums []int, k int) int {
	windowMax, curMax, start := 0, 0, 0

	for i := range nums {
		windowMax += nums[i]
		if i < k {
			curMax += nums[i]
		} else {
			windowMax -= nums[start]
			curMax = max(curMax, windowMax)
			start++
		}
	}
	return curMax
}

// https://www.hellointerview.com/learn/code/sliding-window/maximum-sum-of-subarrays-of-size-k
func maxSubarraySum(nums []int, k int) int {
	maxSum := 0
	windowSum := 0
	start := 0
	for end := range nums {
		windowSum += nums[end]
		if end-start+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= nums[start]
			start++
		}
	}
	return maxSum
}

// Template
func fixedLengthSlidingWindow(nums []int, k int) int {
	// choose appropriate data structure
	// state := make(map[int]int)
	start := 0
	maxVal := 0
	windowSum := 0

	for end := 0; end < len(nums); end++ {
		// extend window
		// add nums[end] to state in O(1) time
		if end-start+1 == k {
			// INVARIANT: size of the window is k here.
			maxVal = max(maxVal, windowSum)

			// contract window
			// remove nums[start] from state in O(1) time
			start++
		}
	}

	return maxVal
}
