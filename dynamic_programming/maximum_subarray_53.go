package dynamic_programming

/*
https://leetcode.com/problems/maximum-subarray/
Given an integer array nums, find the subarray with the largest sum, and return its sum.

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
问题：整个 nums 中，和最大的连续子数组是多少？ subarray=必须连续

第一步：怎么定义dp[]
dp[i]=以i结尾的最大连续子数组的和

第二步：要得到 dp[i]，怎么从更小的问题得到？
dp[i-1] = 以前一个元素结尾的最大连续子数组和

nums = [-2, 1, -3, 4]
假如来到了 4
① 把前面的连续子数组接过来，再加 4 dp[i-1] + nums[i]
② 前面的不要了，直接从 4 重新开始 nums[i]
dp[i]=max(num[i],dp[i-1]+nums[i])，todo 这里不是和 dp[i] 比，因为 dp[i] 还没算出来

第三步：初始化状态
dp[0] = nums[0] 可能为负数，0，正数

第四步：return 什么
return max(dp)
*/
func maxSubArray(nums []int) int {
	// todo 先校验入参
	if len(nums) == 0 {
		return 0
	}

	// 第一步定义 dp[i] = 以 nums[i] 结尾的最长递增子序列长度
	dp := make([]int, len(nums))

	// 第三步：初始化
	dp[0] = nums[0]

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		// 因为要求连续，只能接 curr-1
		// 第二步，要得到 dp[i]，怎么从更小的问题得到
		dp[i] = max(dp[i], dp[i-1]+nums[i])
		result = max(result, dp[i])
	}

	// 第四步，返回max(dp)
	return result
}
