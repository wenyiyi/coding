package search

/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target,
return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1
*/

/*
	找 target 在不在数据组里，在就输出 index
	普通二分面对 [1,2,3,4,5,6,7] 很容易决定往左还是往右；
	现在变成 [4,5,6,7,0,1,2] 后，怎么决定下一步搜左边还是右边
	关键：每次用 mid 切开后，左右两边至少有一边一定是正常升序的。

	index:  0  1  2  3  4  5  6
	nums:   4  5  6  7  0  1  2
        	L        M        R
	左边 4567正常生序
	右边 7012 不是有序的
	需要判断旋转断点在哪里
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 每一轮都重新算 mid
		mid := left + (right-left)/2
		// 1 先判断中间
		if nums[mid] == target {
			return mid
		}
		// 问题1：哪一边有序？
		// 问题2：target 在不在有序的那一边？
		if nums[left] <= nums[mid] { // 左边有序
			// target 在左边
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
