package binary_search

/*
https://leetcode.com/problems/binary-search/description/
Given an array of integers nums which is sorted in ascending order,
and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Constraints:
1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.
*/
/*
0  1 2 3 4 5
-1,0,3,5,9,12

搜索区间：[left, right]

初始化：
left = 0
right = n-1

循环：
left <= right

排除 mid：
right = mid-1
或者
left = mid+1


todo 总结：基础二分， [left,right] 找具体 target；left <= right；mid±1
*/
func search704(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// todo for left = 0; left <= right; left++ ❌ left 是由 mid 的比较结果决定怎么移动
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
