package fixed_length_sliding_window

import "math"

/*
https://www.hellointerview.com/learn/code/sliding-window/maximum-sum-of-distinct-subarrays-with-length-k

Max Sum of Distinct Subarrays Length k
DESCRIPTION (inspired by Leetcode.com)
Given an integer array nums and an integer k, write a function to identify the highest possible sum of a subarray within nums,
where the subarray meets the following criteria: its length is k, and all of its elements are unique. If no such subarray exists, return 0.

Example 1: Input:

nums = [3, 2, 2, 3, 4, 6, 7, 7, -1]
k = 4
Output: 20
Explanation: The subarrays of nums with length 4 are:

[3, 2, 2, 3] # elements 3 and 2 are repeated.
[2, 2, 3, 4] # element 2 is repeated.
[2, 3, 4, 6] # meets the requirements and has a sum of 15.
[3, 4, 6, 7] # meets the requirements and has a sum of 20.
[4, 6, 7, 7] # element 7 is repeated.
[6, 7, 7, -1] # element 7 is repeated.
We return 20 because it is the maximum subarray sum of all the subarrays that meet the conditions.

Example 2:
Input: nums = [5, 5, 5, 5, 5] k = 3
Output:  0
Explanation: Every subarray of length 3 contains duplicate elements, so no valid subarray exists. Return 0.
*/

func maxSum(nums []int, k int) int64 {
	// Initialize maxSum with the smallest possible value to handle negative sums : [-5,-5] k=2, if maxSum=0,max(maxSum, windowSum)=0
	maxSum := int64(math.MinInt64)
	var windowSum int64
	// Use a map to detect duplicates
	countMap := make(map[int]int)

	start := 0

	for end := range nums {
		countMap[nums[end]]++

		windowSum += int64(nums[end])

		if end-start+1 == k {
			// If the map contains k keys, all numbers are distinct.
			if len(countMap) == k {
				maxSum = max(maxSum, windowSum)
			}

			// Decrease the frequency of the leftmost number.
			countMap[nums[start]]--
			if countMap[nums[start]] == 0 {
				delete(countMap, nums[start])
			}
			windowSum -= int64(nums[start])
			start++
		}
	}

	// Return 0 if no valid window was found.
	if maxSum == math.MinInt64 {
		return 0
	}

	return maxSum
}
