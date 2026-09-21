package dynamic_programming

/*
https://leetcode.com/problems/longest-increasing-subsequence/
Given an integer array nums, return the length of  longest strictly increasing subsequence.

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4
Explanation: [0,1,2,3] 不要求连续

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

/*
问题：给点一个数组，返回最长的递增子序列，subsequence = 不要求 contiguous连续

第一步：怎么定义dp[]
dp[i] = 以 nums[i] 结尾的最长递增子序列长度

第二步：要得到 dp[i]，怎么从更小的问题得到？
dp[i] = dp[i-1]+1 ❌，LIS 是 subsequence，不要求连续，所以不能只看前一个位置，前面的都要看

比如：nums = [2, 5, 8, 7]
从 2 接 7：dp[0] + 1 = 2
从 5 接 7：dp[1] + 1 = 3
从 8 接不了 7，8 > 7，不符合递增

只有 nums[pre] < nums[i] 时，nums[i] 才能接在 nums[pre] 后面，因为要递增
dp[i] = max(dp[i],dp[pre]+1), pre 从 0 到 i-1

第三步：初始化状态
dp[0] = 1
dp[1] = 1
每个元素自己至少可以构成长度为 1 的子序列

第四步：return 什么
return dp[len(nums)-1] ❌
return max(dp)
因为 dp[i] 表示“以 nums[i] 结尾”的 LIS，所以整个数组的 LIS 可能结束在任意位置
最终答案往往需要在所有 dp[i] 中再取一次最优值，而不一定是最后一个状态
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 第一步：定义 dp[i] = 以 nums[i] 结尾的最长递增子序列长度
	dp := make([]int, len(nums))

	// 第三步：初始化 每个元素自己至少可以构成长度为 1 的子序列
	for i := range dp {
		dp[i] = 1
	}

	result := 0
	// i = 当前
	// prev = 前一个候选位置
	for i := 0; i < len(nums); i++ {
		// 不要求连续，需要检查前面所有的字符
		for prev := 0; prev < i; prev++ {
			// 对当前元素，检查前面所有 prev
			if nums[prev] < nums[i] {
				// 第二步：要得到 dp[i]，怎么从更小的问题得到
				dp[i] = max(dp[i], dp[prev]+1)
			}
		}
		result = max(result, dp[i])
	}
	// 第四步，返回max(dp)
	return result
}
