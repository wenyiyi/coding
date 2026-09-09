package dp

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

/*
dp[i] = 以 nums[i] 结尾的最长递增子序列长度
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] = 以 nums[i] 结尾的最长递增子序列长度
	dp := make([]int, len(nums))

	// 每个元素自己至少可以构成长度为 1 的子序列
	// Go 没有 Arrays.fill，所以手动初始化
	for i := range dp {
		dp[i] = 1
	}

	result := 0

	// curr = 当前
	// prev = 前一个候选位置
	for curr := 0; curr < len(nums); curr++ {
		for prev := 0; prev < curr; prev++ {
			// 对当前元素 curr，检查前面所有 prev
			if nums[prev] < nums[curr] {
				dp[curr] = max(dp[curr], dp[prev]+1)
			}
		}
		result = max(result, dp[curr])
	}

	return result
}
