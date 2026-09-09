package dp

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
https://leetcode.com/problems/maximum-subarray/


Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

/*
dp[i] = 以 nums[i] 结尾的最大连续子数组和
subarray 要求连续
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] = 以 nums[i] 结尾的最长递增子序列长度
	dp := make([]int, len(nums))

	// 初始化
	dp[0] = nums[0]

	result := nums[0]

	for curr := 1; curr < len(nums); curr++ {
		// 因为要求连续，只能接 curr-1
		dp[curr] = max(dp[curr], dp[curr-1]+nums[curr])
		result = max(result, dp[curr])
	}

	return result
}
